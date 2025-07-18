// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logical

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mitchellh/copystructure"
)

// RequestWrapInfo is a struct that stores information about desired response
// and seal wrapping behavior
type RequestWrapInfo struct {
	// Setting to non-zero specifies that the response should be wrapped.
	// Specifies the desired TTL of the wrapping token.
	TTL time.Duration `json:"ttl" structs:"ttl" mapstructure:"ttl" sentinel:""`

	// The format to use for the wrapped response; if not specified it's a bare
	// token
	Format string `json:"format" structs:"format" mapstructure:"format" sentinel:""`

	// A flag to conforming backends that data for a given request should be
	// seal wrapped
	SealWrap bool `json:"seal_wrap" structs:"seal_wrap" mapstructure:"seal_wrap" sentinel:""`
}

func (r *RequestWrapInfo) SentinelGet(key string) (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	switch key {
	case "ttl":
		return r.TTL, nil
	case "ttl_seconds":
		return int64(r.TTL.Seconds()), nil
	}

	return nil, nil
}

func (r *RequestWrapInfo) SentinelKeys() []string {
	return []string{
		"ttl",
		"ttl_seconds",
	}
}

type ClientTokenSource uint32

const (
	NoClientToken ClientTokenSource = iota
	ClientTokenFromVaultHeader
	ClientTokenFromAuthzHeader
	ClientTokenFromInlineAuth
)

// Request is a struct that stores the parameters and context of a request
// being made to Vault. It is used to abstract the details of the higher level
// request protocol from the handlers.
//
// Note: Many of these have Sentinel disabled because they are values populated
// by the router after policy checks; the token namespace would be the right
// place to access them via Sentinel
type Request struct {
	// Id is the uuid associated with each request
	ID string `json:"id" structs:"id" mapstructure:"id" sentinel:""`

	// If set, the name given to the replication secondary where this request
	// originated
	ReplicationCluster string `json:"replication_cluster" structs:"replication_cluster" mapstructure:"replication_cluster" sentinel:""`

	// Operation is the requested operation type
	Operation Operation `json:"operation" structs:"operation" mapstructure:"operation"`

	// Path is the full path of the request
	Path string `json:"path" structs:"path" mapstructure:"path" sentinel:""`

	// Request data is an opaque map that must have string keys.
	Data map[string]interface{} `json:"map" structs:"data" mapstructure:"data"`

	// Storage can be used to durably store and retrieve state.
	Storage Storage `json:"-" sentinel:""`

	// OriginalStorage can hold an original storage instance during TX operations for restoring Storage.
	OriginalStorage Storage `json:"-" sentinel:""`

	// Secret will be non-nil only for Revoke and Renew operations
	// to represent the secret that was returned prior.
	Secret *Secret `json:"secret" structs:"secret" mapstructure:"secret" sentinel:""`

	// Auth will be non-nil only for Renew operations
	// to represent the auth that was returned prior.
	Auth *Auth `json:"auth" structs:"auth" mapstructure:"auth" sentinel:""`

	// Headers will contain the http headers from the request. This value will
	// be used in the audit broker to ensure we are auditing only the allowed
	// headers.
	Headers map[string][]string `json:"headers" structs:"headers" mapstructure:"headers" sentinel:""`

	// Connection will be non-nil only for credential providers to
	// inspect the connection information and potentially use it for
	// authentication/protection.
	Connection *Connection `json:"connection" structs:"connection" mapstructure:"connection"`

	// ClientToken is provided to the core so that the identity
	// can be verified and ACLs applied. This value is passed
	// through to the logical backends but after being salted and
	// hashed.
	ClientToken string `json:"client_token" structs:"client_token" mapstructure:"client_token" sentinel:""`

	// ClientTokenAccessor is provided to the core so that the it can get
	// logged as part of request audit logging.
	ClientTokenAccessor string `json:"client_token_accessor" structs:"client_token_accessor" mapstructure:"client_token_accessor" sentinel:""`

	// DisplayName is provided to the logical backend to help associate
	// dynamic secrets with the source entity. This is not a sensitive
	// name, but is useful for operators.
	DisplayName string `json:"display_name" structs:"display_name" mapstructure:"display_name" sentinel:""`

	// MountPoint is provided so that a logical backend can generate
	// paths relative to itself. The `Path` is effectively the client
	// request path with the MountPoint trimmed off.
	MountPoint string `json:"mount_point" structs:"mount_point" mapstructure:"mount_point" sentinel:""`

	// MountType is provided so that a logical backend can make decisions
	// based on the specific mount type (e.g., if a mount type has different
	// aliases, generating different defaults depending on the alias)
	MountType string `json:"mount_type" structs:"mount_type" mapstructure:"mount_type" sentinel:""`

	// MountAccessor is provided so that identities returned by the authentication
	// backends can be tied to the mount it belongs to.
	MountAccessor string `json:"mount_accessor" structs:"mount_accessor" mapstructure:"mount_accessor" sentinel:""`

	// mountRunningVersion is used internally to propagate the semantic version
	// of the mounted plugin as reported by its vault.MountEntry to audit logging
	mountRunningVersion string

	// mountRunningSha256 is used internally to propagate the encoded sha256
	// of the mounted plugin as reported its vault.MountEntry to audit logging
	mountRunningSha256 string

	// mountIsExternalPlugin is used internally to propagate whether
	// the backend of the mounted plugin is running externally (i.e., over GRPC)
	// to audit logging
	mountIsExternalPlugin bool

	// mountClass is used internally to propagate the mount class of the mounted plugin to audit logging
	mountClass string

	// WrapInfo contains requested response wrapping parameters
	WrapInfo *RequestWrapInfo `json:"wrap_info" structs:"wrap_info" mapstructure:"wrap_info" sentinel:""`

	// ClientTokenRemainingUses represents the allowed number of uses left on the
	// token supplied
	ClientTokenRemainingUses int `json:"client_token_remaining_uses" structs:"client_token_remaining_uses" mapstructure:"client_token_remaining_uses"`

	// EntityID is the identity of the caller extracted out of the token used
	// to make this request
	EntityID string `json:"entity_id" structs:"entity_id" mapstructure:"entity_id" sentinel:""`

	// PolicyOverride indicates that the requestor wishes to override
	// soft-mandatory Sentinel policies
	PolicyOverride bool `json:"policy_override" structs:"policy_override" mapstructure:"policy_override"`

	// Whether the request is unauthenticated, as in, had no client token
	// attached. Useful in some situations where the client token is not made
	// accessible.
	Unauthenticated bool `json:"unauthenticated" structs:"unauthenticated" mapstructure:"unauthenticated"`

	// MFACreds holds the parsed MFA information supplied over the API as part of
	// X-Vault-MFA header
	MFACreds MFACreds `json:"mfa_creds" structs:"mfa_creds" mapstructure:"mfa_creds" sentinel:""`

	// Cached token entry. This avoids another lookup in request handling when
	// we've already looked it up at http handling time. Note that this token
	// has not been "used", as in it will not properly take into account use
	// count limitations. As a result this field should only ever be used for
	// transport to a function that would otherwise do a lookup and then
	// properly use the token.
	tokenEntry *TokenEntry

	// For replication, contains the last WAL on the remote side after handling
	// the request, used for best-effort avoidance of stale read-after-write
	lastRemoteWAL uint64

	// ClientTokenSource tells us where the client token was sourced from, so
	// we can delete it before sending off to plugins
	ClientTokenSource ClientTokenSource

	// HTTPRequest, if set, can be used to access fields from the HTTP request
	// that generated this logical.Request object, such as the request body.
	HTTPRequest *http.Request `json:"-" sentinel:""`

	// ResponseWriter if set can be used to stream a response value to the http
	// request that generated this logical.Request object.
	ResponseWriter *HTTPResponseWriter `json:"-" sentinel:""`

	// ClientID is the identity of the caller. If the token is associated with an
	// entity, it will be the same as the EntityID . If the token has no entity,
	// this will be the sha256(sorted policies + namespace) associated with the
	// client token.
	ClientID string `json:"client_id" structs:"client_id" mapstructure:"client_id" sentinel:""`

	// InboundSSCToken is the token that arrives on an inbound request, supplied
	// by the vault user.
	InboundSSCToken string

	// When a request has been forwarded, contains information of the host the request was forwarded 'from'
	ForwardedFrom string `json:"forwarded_from,omitempty"`

	// Whether we are an inline authentication request; in this case, we
	// should not persist the generated token to storage.
	IsInlineAuth bool `json:"is_inline_auth,omitempty" sentinel:""`

	// Results from inline authentication to bypass storage-based token
	// resolution. Only valid when HasInlineAuth=true.
	HasInlineAuth bool  `json:"has_inline_auth,omitempty" sentinel:""`
	InlineAuth    *Auth `json:"-" sentinel:""`
}

// Clone returns a deep copy of the request by using copystructure
func (r *Request) Clone() (*Request, error) {
	cpy, err := copystructure.Copy(r)
	if err != nil {
		return nil, err
	}
	return cpy.(*Request), nil
}

// Get returns a data field and guards for nil Data
func (r *Request) Get(key string) interface{} {
	if r.Data == nil {
		return nil
	}
	return r.Data[key]
}

// GetString returns a data field as a string
func (r *Request) GetString(key string) string {
	raw := r.Get(key)
	s, _ := raw.(string)
	return s
}

func (r *Request) GoString() string {
	return fmt.Sprintf("*%#v", *r)
}

func (r *Request) SentinelGet(key string) (interface{}, error) {
	switch key {
	case "path":
		// Sanitize it here so that it's consistent in policies
		return strings.TrimPrefix(r.Path, "/"), nil

	case "wrapping", "wrap_info":
		// If the pointer is nil accessing the wrap info is considered
		// "undefined" so this allows us to instead discover a TTL of zero
		if r.WrapInfo == nil {
			return &RequestWrapInfo{}, nil
		}
		return r.WrapInfo, nil
	}

	return nil, nil
}

func (r *Request) SentinelKeys() []string {
	return []string{
		"path",
		"wrapping",
		"wrap_info",
	}
}

func (r *Request) MountRunningVersion() string {
	return r.mountRunningVersion
}

func (r *Request) SetMountRunningVersion(mountRunningVersion string) {
	r.mountRunningVersion = mountRunningVersion
}

func (r *Request) MountRunningSha256() string {
	return r.mountRunningSha256
}

func (r *Request) SetMountRunningSha256(mountRunningSha256 string) {
	r.mountRunningSha256 = mountRunningSha256
}

func (r *Request) MountIsExternalPlugin() bool {
	return r.mountIsExternalPlugin
}

func (r *Request) SetMountIsExternalPlugin(mountIsExternalPlugin bool) {
	r.mountIsExternalPlugin = mountIsExternalPlugin
}

func (r *Request) MountClass() string {
	return r.mountClass
}

func (r *Request) SetMountClass(mountClass string) {
	r.mountClass = mountClass
}

func (r *Request) TokenEntry() *TokenEntry {
	return r.tokenEntry
}

func (r *Request) SetTokenEntry(te *TokenEntry) {
	r.tokenEntry = te
}

// RenewRequest creates the structure of the renew request.
func RenewRequest(path string, secret *Secret, data map[string]interface{}) *Request {
	return &Request{
		Operation: RenewOperation,
		Path:      path,
		Data:      data,
		Secret:    secret,
	}
}

// RenewAuthRequest creates the structure of the renew request for an auth.
func RenewAuthRequest(path string, auth *Auth, data map[string]interface{}) *Request {
	return &Request{
		Operation: RenewOperation,
		Path:      path,
		Data:      data,
		Auth:      auth,
	}
}

// RevokeRequest creates the structure of the revoke request.
func RevokeRequest(path string, secret *Secret, data map[string]interface{}) *Request {
	return &Request{
		Operation: RevokeOperation,
		Path:      path,
		Data:      data,
		Secret:    secret,
	}
}

// RollbackRequest creates the structure of the revoke request.
func RollbackRequest(path string) *Request {
	return &Request{
		Operation: RollbackOperation,
		Path:      path,
		Data:      make(map[string]interface{}),
	}
}

// Operation is an enum that is used to specify the type
// of request being made
type Operation string

const (
	// The operations below are called per path
	CreateOperation         Operation = "create"
	ReadOperation           Operation = "read"
	UpdateOperation         Operation = "update"
	PatchOperation          Operation = "patch"
	DeleteOperation         Operation = "delete"
	ListOperation           Operation = "list"
	ScanOperation           Operation = "scan"
	HelpOperation           Operation = "help"
	AliasLookaheadOperation Operation = "alias-lookahead"
	ResolveRoleOperation    Operation = "resolve-role"
	HeaderOperation         Operation = "header"

	// The operations below are called globally, the path is less relevant.
	RevokeOperation   Operation = "revoke"
	RenewOperation    Operation = "renew"
	RollbackOperation Operation = "rollback"
)

type MFACreds map[string][]string

// InitializationRequest stores the parameters and context of an Initialize()
// call being made to a logical.Backend.
type InitializationRequest struct {
	// Storage can be used to durably store and retrieve state.
	Storage Storage
}

type CustomHeader struct {
	Name  string
	Value string
}

type CtxKeyInFlightRequestID struct{}

func (c CtxKeyInFlightRequestID) String() string {
	return "in-flight-request-ID"
}

type CtxKeyRequestRole struct{}

func (c CtxKeyRequestRole) String() string {
	return "request-role"
}

type ctxKeyOriginalBody struct{}

func ContextOriginalBodyValue(ctx context.Context) (io.ReadCloser, bool) {
	value, ok := ctx.Value(ctxKeyOriginalBody{}).(io.ReadCloser)
	return value, ok
}

func CreateContextOriginalBody(parent context.Context, body io.ReadCloser) context.Context {
	return context.WithValue(parent, ctxKeyOriginalBody{}, body)
}
