// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: helper/identity/mfa/types.proto

package mfa

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config represents the configuration information used *along with* the MFA
// secret tied to caller's identity, to verify the MFA credentials supplied.
// Configuration information differs by type. Handler of each type should know
// what to expect from the Config field.
type Config struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	ID string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	MountAccessor string `protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	UsernameFormat string `protobuf:"bytes,5,opt,name=username_format,json=usernameFormat,proto3" json:"username_format,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	//
	// Types that are valid to be assigned to Config:
	//
	//	*Config_TOTPConfig
	//	*Config_OktaConfig
	//	*Config_DuoConfig
	//	*Config_PingIDConfig
	Config isConfig_Config `protobuf_oneof:"config" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	NamespaceID   string `protobuf:"bytes,10,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty" sentinel:"-"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Config) GetMountAccessor() string {
	if x != nil {
		return x.MountAccessor
	}
	return ""
}

func (x *Config) GetUsernameFormat() string {
	if x != nil {
		return x.UsernameFormat
	}
	return ""
}

func (x *Config) GetConfig() isConfig_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Config) GetTOTPConfig() *TOTPConfig {
	if x != nil {
		if x, ok := x.Config.(*Config_TOTPConfig); ok {
			return x.TOTPConfig
		}
	}
	return nil
}

func (x *Config) GetOktaConfig() *OktaConfig {
	if x != nil {
		if x, ok := x.Config.(*Config_OktaConfig); ok {
			return x.OktaConfig
		}
	}
	return nil
}

func (x *Config) GetDuoConfig() *DuoConfig {
	if x != nil {
		if x, ok := x.Config.(*Config_DuoConfig); ok {
			return x.DuoConfig
		}
	}
	return nil
}

func (x *Config) GetPingIDConfig() *PingIDConfig {
	if x != nil {
		if x, ok := x.Config.(*Config_PingIDConfig); ok {
			return x.PingIDConfig
		}
	}
	return nil
}

func (x *Config) GetNamespaceID() string {
	if x != nil {
		return x.NamespaceID
	}
	return ""
}

type isConfig_Config interface {
	isConfig_Config()
}

type Config_TOTPConfig struct {
	TOTPConfig *TOTPConfig `protobuf:"bytes,6,opt,name=totp_config,json=totpConfig,proto3,oneof"`
}

type Config_OktaConfig struct {
	OktaConfig *OktaConfig `protobuf:"bytes,7,opt,name=okta_config,json=oktaConfig,proto3,oneof"`
}

type Config_DuoConfig struct {
	DuoConfig *DuoConfig `protobuf:"bytes,8,opt,name=duo_config,json=duoConfig,proto3,oneof"`
}

type Config_PingIDConfig struct {
	PingIDConfig *PingIDConfig `protobuf:"bytes,9,opt,name=pingid_config,json=pingidConfig,proto3,oneof"`
}

func (*Config_TOTPConfig) isConfig_Config() {}

func (*Config_OktaConfig) isConfig_Config() {}

func (*Config_DuoConfig) isConfig_Config() {}

func (*Config_PingIDConfig) isConfig_Config() {}

// TOTPConfig represents the configuration information required to generate
// a TOTP key. The generated key will be stored in the entity along with these
// options. Validation of credentials supplied over the API will be validated
// by the information stored in the entity and not from the values in the
// configuration.
type TOTPConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Period uint32 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Algorithm int32 `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Digits int32 `protobuf:"varint,4,opt,name=digits,proto3" json:"digits,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Skew uint32 `protobuf:"varint,5,opt,name=skew,proto3" json:"skew,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	KeySize uint32 `protobuf:"varint,6,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	QRSize int32 `protobuf:"varint,7,opt,name=qr_size,json=qrSize,proto3" json:"qr_size,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	MaxValidationAttempts uint32 `protobuf:"varint,8,opt,name=max_validation_attempts,json=maxValidationAttempts,proto3" json:"max_validation_attempts,omitempty" sentinel:"-"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TOTPConfig) Reset() {
	*x = TOTPConfig{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TOTPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TOTPConfig) ProtoMessage() {}

func (x *TOTPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TOTPConfig.ProtoReflect.Descriptor instead.
func (*TOTPConfig) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{1}
}

func (x *TOTPConfig) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *TOTPConfig) GetPeriod() uint32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *TOTPConfig) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *TOTPConfig) GetDigits() int32 {
	if x != nil {
		return x.Digits
	}
	return 0
}

func (x *TOTPConfig) GetSkew() uint32 {
	if x != nil {
		return x.Skew
	}
	return 0
}

func (x *TOTPConfig) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *TOTPConfig) GetQRSize() int32 {
	if x != nil {
		return x.QRSize
	}
	return 0
}

func (x *TOTPConfig) GetMaxValidationAttempts() uint32 {
	if x != nil {
		return x.MaxValidationAttempts
	}
	return 0
}

// DuoConfig represents the configuration information required to perform
// Duo authentication.
type DuoConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	IntegrationKey string `protobuf:"bytes,1,opt,name=integration_key,json=integrationKey,proto3" json:"integration_key,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	APIHostname string `protobuf:"bytes,3,opt,name=api_hostname,json=apiHostname,proto3" json:"api_hostname,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	PushInfo string `protobuf:"bytes,4,opt,name=push_info,json=pushInfo,proto3" json:"push_info,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	UsePasscode   bool `protobuf:"varint,5,opt,name=use_passcode,json=usePasscode,proto3" json:"use_passcode,omitempty" sentinel:"-"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DuoConfig) Reset() {
	*x = DuoConfig{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DuoConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuoConfig) ProtoMessage() {}

func (x *DuoConfig) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuoConfig.ProtoReflect.Descriptor instead.
func (*DuoConfig) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{2}
}

func (x *DuoConfig) GetIntegrationKey() string {
	if x != nil {
		return x.IntegrationKey
	}
	return ""
}

func (x *DuoConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *DuoConfig) GetAPIHostname() string {
	if x != nil {
		return x.APIHostname
	}
	return ""
}

func (x *DuoConfig) GetPushInfo() string {
	if x != nil {
		return x.PushInfo
	}
	return ""
}

func (x *DuoConfig) GetUsePasscode() bool {
	if x != nil {
		return x.UsePasscode
	}
	return false
}

// OktaConfig contains Okta configuration parameters required to perform Okta
// authentication.
type OktaConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	OrgName string `protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	APIToken string `protobuf:"bytes,2,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Production bool `protobuf:"varint,3,opt,name=production,proto3" json:"production,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	BaseURL string `protobuf:"bytes,4,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	PrimaryEmail  bool `protobuf:"varint,5,opt,name=primary_email,json=primaryEmail,proto3" json:"primary_email,omitempty" sentinel:"-"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OktaConfig) Reset() {
	*x = OktaConfig{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OktaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OktaConfig) ProtoMessage() {}

func (x *OktaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OktaConfig.ProtoReflect.Descriptor instead.
func (*OktaConfig) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{3}
}

func (x *OktaConfig) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *OktaConfig) GetAPIToken() string {
	if x != nil {
		return x.APIToken
	}
	return ""
}

func (x *OktaConfig) GetProduction() bool {
	if x != nil {
		return x.Production
	}
	return false
}

func (x *OktaConfig) GetBaseURL() string {
	if x != nil {
		return x.BaseURL
	}
	return ""
}

func (x *OktaConfig) GetPrimaryEmail() bool {
	if x != nil {
		return x.PrimaryEmail
	}
	return false
}

// PingIDConfig contains PingID configuration information
type PingIDConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	UseBase64Key string `protobuf:"bytes,1,opt,name=use_base64_key,json=useBase64Key,proto3" json:"use_base64_key,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	UseSignature bool `protobuf:"varint,2,opt,name=use_signature,json=useSignature,proto3" json:"use_signature,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	IDPURL string `protobuf:"bytes,4,opt,name=idp_url,json=idpUrl,proto3" json:"idp_url,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	OrgAlias string `protobuf:"bytes,5,opt,name=org_alias,json=orgAlias,proto3" json:"org_alias,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	AdminURL string `protobuf:"bytes,6,opt,name=admin_url,json=adminUrl,proto3" json:"admin_url,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	AuthenticatorURL string `protobuf:"bytes,7,opt,name=authenticator_url,json=authenticatorUrl,proto3" json:"authenticator_url,omitempty" sentinel:"-"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PingIDConfig) Reset() {
	*x = PingIDConfig{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingIDConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingIDConfig) ProtoMessage() {}

func (x *PingIDConfig) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingIDConfig.ProtoReflect.Descriptor instead.
func (*PingIDConfig) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{4}
}

func (x *PingIDConfig) GetUseBase64Key() string {
	if x != nil {
		return x.UseBase64Key
	}
	return ""
}

func (x *PingIDConfig) GetUseSignature() bool {
	if x != nil {
		return x.UseSignature
	}
	return false
}

func (x *PingIDConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PingIDConfig) GetIDPURL() string {
	if x != nil {
		return x.IDPURL
	}
	return ""
}

func (x *PingIDConfig) GetOrgAlias() string {
	if x != nil {
		return x.OrgAlias
	}
	return ""
}

func (x *PingIDConfig) GetAdminURL() string {
	if x != nil {
		return x.AdminURL
	}
	return ""
}

func (x *PingIDConfig) GetAuthenticatorURL() string {
	if x != nil {
		return x.AuthenticatorURL
	}
	return ""
}

// Secret represents all the types of secrets which the entity can hold.
// Each MFA type should add a secret type to the oneof block in this message.
type Secret struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty" sentinel:"-"`
	// Types that are valid to be assigned to Value:
	//
	//	*Secret_TOTPSecret
	Value         isSecret_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{5}
}

func (x *Secret) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *Secret) GetValue() isSecret_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Secret) GetTOTPSecret() *TOTPSecret {
	if x != nil {
		if x, ok := x.Value.(*Secret_TOTPSecret); ok {
			return x.TOTPSecret
		}
	}
	return nil
}

type isSecret_Value interface {
	isSecret_Value()
}

type Secret_TOTPSecret struct {
	// @inject_tag: sentinel:"-"
	TOTPSecret *TOTPSecret `protobuf:"bytes,2,opt,name=totp_secret,json=totpSecret,proto3,oneof" sentinel:"-"`
}

func (*Secret_TOTPSecret) isSecret_Value() {}

// TOTPSecret represents the secret that gets stored in the entity about a
// particular MFA method. This information is used to validate the MFA
// credential supplied over the API during request time.
type TOTPSecret struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: sentinel:"-"
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Period uint32 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Algorithm int32 `protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Digits int32 `protobuf:"varint,4,opt,name=digits,proto3" json:"digits,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Skew uint32 `protobuf:"varint,5,opt,name=skew,proto3" json:"skew,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	KeySize uint32 `protobuf:"varint,6,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty" sentinel:"-"`
	// reserving 7 here just to keep parity with the config message above
	// @inject_tag: sentinel:"-"
	AccountName string `protobuf:"bytes,8,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty" sentinel:"-"`
	// @inject_tag: sentinel:"-"
	Key           string `protobuf:"bytes,9,opt,name=key,proto3" json:"key,omitempty" sentinel:"-"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TOTPSecret) Reset() {
	*x = TOTPSecret{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TOTPSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TOTPSecret) ProtoMessage() {}

func (x *TOTPSecret) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TOTPSecret.ProtoReflect.Descriptor instead.
func (*TOTPSecret) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{6}
}

func (x *TOTPSecret) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *TOTPSecret) GetPeriod() uint32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *TOTPSecret) GetAlgorithm() int32 {
	if x != nil {
		return x.Algorithm
	}
	return 0
}

func (x *TOTPSecret) GetDigits() int32 {
	if x != nil {
		return x.Digits
	}
	return 0
}

func (x *TOTPSecret) GetSkew() uint32 {
	if x != nil {
		return x.Skew
	}
	return 0
}

func (x *TOTPSecret) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *TOTPSecret) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *TOTPSecret) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// MFAEnforcementConfig is what the user provides to the
// mfa/login_enforcement endpoint.
type MFAEnforcementConfig struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Name                string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NamespaceID         string                 `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceID,proto3" json:"namespace_id,omitempty"`
	MFAMethodIDs        []string               `protobuf:"bytes,3,rep,name=mfa_method_ids,json=mfaMethodIds,proto3" json:"mfa_method_ids,omitempty"`
	AuthMethodAccessors []string               `protobuf:"bytes,4,rep,name=auth_method_accessors,json=authMethodAccessors,proto3" json:"auth_method_accessors,omitempty"`
	AuthMethodTypes     []string               `protobuf:"bytes,5,rep,name=auth_method_types,json=authMethodTypes,proto3" json:"auth_method_types,omitempty"`
	IdentityGroupIds    []string               `protobuf:"bytes,6,rep,name=identity_group_ids,json=identityGroupIds,proto3" json:"identity_group_ids,omitempty"`
	IdentityEntityIDs   []string               `protobuf:"bytes,7,rep,name=identity_entity_ids,json=identityEntityIds,proto3" json:"identity_entity_ids,omitempty"`
	ID                  string                 `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MFAEnforcementConfig) Reset() {
	*x = MFAEnforcementConfig{}
	mi := &file_helper_identity_mfa_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MFAEnforcementConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MFAEnforcementConfig) ProtoMessage() {}

func (x *MFAEnforcementConfig) ProtoReflect() protoreflect.Message {
	mi := &file_helper_identity_mfa_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MFAEnforcementConfig.ProtoReflect.Descriptor instead.
func (*MFAEnforcementConfig) Descriptor() ([]byte, []int) {
	return file_helper_identity_mfa_types_proto_rawDescGZIP(), []int{7}
}

func (x *MFAEnforcementConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MFAEnforcementConfig) GetNamespaceID() string {
	if x != nil {
		return x.NamespaceID
	}
	return ""
}

func (x *MFAEnforcementConfig) GetMFAMethodIDs() []string {
	if x != nil {
		return x.MFAMethodIDs
	}
	return nil
}

func (x *MFAEnforcementConfig) GetAuthMethodAccessors() []string {
	if x != nil {
		return x.AuthMethodAccessors
	}
	return nil
}

func (x *MFAEnforcementConfig) GetAuthMethodTypes() []string {
	if x != nil {
		return x.AuthMethodTypes
	}
	return nil
}

func (x *MFAEnforcementConfig) GetIdentityGroupIds() []string {
	if x != nil {
		return x.IdentityGroupIds
	}
	return nil
}

func (x *MFAEnforcementConfig) GetIdentityEntityIDs() []string {
	if x != nil {
		return x.IdentityEntityIDs
	}
	return nil
}

func (x *MFAEnforcementConfig) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_helper_identity_mfa_types_proto protoreflect.FileDescriptor

var file_helper_identity_mfa_types_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x6d, 0x66, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6d, 0x66, 0x61, 0x22, 0x90, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x66, 0x61, 0x2e, 0x54, 0x4f, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a,
	0x0b, 0x6f, 0x6b, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x66, 0x61, 0x2e, 0x4f, 0x6b, 0x74, 0x61, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6b, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2f, 0x0a, 0x0a, 0x64, 0x75, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x66, 0x61, 0x2e, 0x44, 0x75, 0x6f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x64, 0x75, 0x6f, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x38, 0x0a, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x66, 0x61, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c,
	0x70, 0x69, 0x6e, 0x67, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xf2, 0x01, 0x0a, 0x0a, 0x54, 0x4f,
	0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6b, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6b,
	0x65, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x71, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x71, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x22, 0xb6,
	0x01, 0x0a, 0x09, 0x44, 0x75, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x4f, 0x6b, 0x74, 0x61,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xef,
	0x01, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x36, 0x34, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x64, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x67,
	0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72,
	0x67, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x72, 0x6c,
	0x22, 0x66, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x66, 0x61, 0x2e, 0x54, 0x4f, 0x54, 0x50, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x54, 0x4f, 0x54,
	0x50, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6b, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6b, 0x65,
	0x77, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0xc1, 0x02, 0x0a, 0x14, 0x4d, 0x46, 0x41, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x66, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x66, 0x61, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x62, 0x61, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x6d, 0x66, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_helper_identity_mfa_types_proto_rawDescOnce sync.Once
	file_helper_identity_mfa_types_proto_rawDescData []byte
)

func file_helper_identity_mfa_types_proto_rawDescGZIP() []byte {
	file_helper_identity_mfa_types_proto_rawDescOnce.Do(func() {
		file_helper_identity_mfa_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_helper_identity_mfa_types_proto_rawDesc), len(file_helper_identity_mfa_types_proto_rawDesc)))
	})
	return file_helper_identity_mfa_types_proto_rawDescData
}

var file_helper_identity_mfa_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_helper_identity_mfa_types_proto_goTypes = []any{
	(*Config)(nil),               // 0: mfa.Config
	(*TOTPConfig)(nil),           // 1: mfa.TOTPConfig
	(*DuoConfig)(nil),            // 2: mfa.DuoConfig
	(*OktaConfig)(nil),           // 3: mfa.OktaConfig
	(*PingIDConfig)(nil),         // 4: mfa.PingIDConfig
	(*Secret)(nil),               // 5: mfa.Secret
	(*TOTPSecret)(nil),           // 6: mfa.TOTPSecret
	(*MFAEnforcementConfig)(nil), // 7: mfa.MFAEnforcementConfig
}
var file_helper_identity_mfa_types_proto_depIDxs = []int32{
	1, // 0: mfa.Config.totp_config:type_name -> mfa.TOTPConfig
	3, // 1: mfa.Config.okta_config:type_name -> mfa.OktaConfig
	2, // 2: mfa.Config.duo_config:type_name -> mfa.DuoConfig
	4, // 3: mfa.Config.pingid_config:type_name -> mfa.PingIDConfig
	6, // 4: mfa.Secret.totp_secret:type_name -> mfa.TOTPSecret
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_helper_identity_mfa_types_proto_init() }
func file_helper_identity_mfa_types_proto_init() {
	if File_helper_identity_mfa_types_proto != nil {
		return
	}
	file_helper_identity_mfa_types_proto_msgTypes[0].OneofWrappers = []any{
		(*Config_TOTPConfig)(nil),
		(*Config_OktaConfig)(nil),
		(*Config_DuoConfig)(nil),
		(*Config_PingIDConfig)(nil),
	}
	file_helper_identity_mfa_types_proto_msgTypes[5].OneofWrappers = []any{
		(*Secret_TOTPSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_helper_identity_mfa_types_proto_rawDesc), len(file_helper_identity_mfa_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helper_identity_mfa_types_proto_goTypes,
		DependencyIndexes: file_helper_identity_mfa_types_proto_depIDxs,
		MessageInfos:      file_helper_identity_mfa_types_proto_msgTypes,
	}.Build()
	File_helper_identity_mfa_types_proto = out.File
	file_helper_identity_mfa_types_proto_goTypes = nil
	file_helper_identity_mfa_types_proto_depIDxs = nil
}
