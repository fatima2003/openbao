// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vault

import (
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/openbao/openbao/helper/namespace"
)

var rawPolicy = strings.TrimSpace(`
# Developer policy
name = "dev"
# Deny all paths by default
path "*" {
	policy = "deny"
}
# Allow full access to staging
path "stage/*" {
	policy = "sudo"
}
# Limited read privilege to production
path "prod/version" {
	policy = "read"
	comment = "this comment is stored but not parsed"
}
# Read access to foobar
# Also tests stripping of leading slash and parsing of min/max as string and
# integer
path "/foo/bar" {
	policy = "read"
	min_wrapping_ttl = 300
	max_wrapping_ttl = "1h"
}
# Add capabilities for creation and sudo to foobar
# This will be separate; they are combined when compiled into an ACL
# Also tests reverse string/int handling to the above
path "foo/bar" {
	capabilities = ["create", "sudo"]
	min_wrapping_ttl = "300s"
	max_wrapping_ttl = 3600
}
# Check that only allowed_parameters are being added to foobar
path "foo/bar" {
	capabilities = ["create", "sudo"]
	allowed_parameters = {
	  "zip" = []
	  "zap" = []
	}
}
# Check that only denied_parameters are being added to bazbar
path "baz/bar" {
	capabilities = ["create", "sudo"]
	denied_parameters = {
	  "zip" = []
	  "zap" = []
	}
}
# Check that both allowed and denied parameters are being added to bizbar
path "biz/bar" {
	capabilities = ["create", "sudo"]
	allowed_parameters = {
	  "zim" = []
	  "zam" = []
	}
	denied_parameters = {
	  "zip" = []
	  "zap" = []
	}
}
path "test/types" {
	capabilities = ["create", "sudo"]
	allowed_parameters = {
		"map" = [{"good" = "one"}]
		"int" = [1, 2]
	}
	denied_parameters = {
		"string" = ["test"]
		"bool" = [false]
	}
}
path "test/req" {
	capabilities = ["create", "sudo"]
	required_parameters = ["foo"]
}
path "test/patch" {
	capabilities = ["patch"]
}
path "test/scan" {
	capabilities = ["scan"]
}
path "test/mfa" {
	capabilities = ["create", "sudo"]
	mfa_methods = ["my_totp", "my_totp2"]
}
path "test/+/segment" {
	capabilities = ["create", "sudo"]
}
path "test/segment/at/end/+" {
	capabilities = ["create", "sudo"]
}
path "test/segment/at/end/v2/+/" {
	capabilities = ["create", "sudo"]
}
path "test/+/wildcard/+/*" {
	capabilities = ["create", "sudo"]
}
path "test/+/wildcard/+/end*" {
	capabilities = ["create", "sudo"]
}
path "paginated-kv/metadata" {
	capabilities = ["list"]
	pagination_limit = 12345
}
path "unpaginated-kv/metadata" {
	capabilities = ["list"]
}
path "some-expired-path" {
	capabilities = ["list"]
	expiration = "2006-01-02T15:04:05Z"
}
path "test/metadata/*" {
	capabilities = ["list"]
	list_scan_response_keys_filter_path = "test/metadata/{{ .key }}"
}
`)

var rawPolicyJSON = strings.TrimSpace(`
{
  "name": "dev",
  "path": {
    "*": {
        "policy": "deny"
    },
    "stage/*": {
        "policy": "sudo"
    },
    "prod/version": {
        "policy": "read",
        "comment": "this comment is stored but not parsed"
    },
    "/foo/bar": {
      "policy": "read",
      "min_wrapping_ttl": 300,
      "max_wrapping_ttl": "1h"
    },
    "foo/bar": {
      "capabilities": ["create", "sudo"],
      "min_wrapping_ttl": "300s",
      "max_wrapping_ttl": 3600
    },
    "foo/bar": {
      "capabilities": ["create", "sudo"],
      "allowed_parameters": {
        "zip": [],
        "zap": []
      }
    },
    "baz/bar": {
      "capabilities": ["create", "sudo"],
      "denied_parameters": {
        "zip": [],
        "zap": []
      }
    },
    "biz/bar": {
      "capabilities": ["create", "sudo"],
      "allowed_parameters": {
        "zim": [],
        "zam": []
      },
      "denied_parameters": {
        "zip": [],
        "zap": []
      }
    },
    "test/types": {
      "capabilities": ["create", "sudo"],
      "allowed_parameters": {
        "map": [{"good": "one"}],
        "int": [1, 2]
      },
      "denied_parameters": {
        "string": ["test"],
        "bool": [false]
      }
    },
    "test/req": {
      "capabilities": ["create", "sudo"],
      "required_parameters": ["foo"]
    },
    "test/patch": {
      "capabilities": ["patch"]
    },
    "test/scan": {
      "capabilities": ["scan"]
    },
    "test/mfa": {
      "capabilities": ["create", "sudo"],
      "mfa_methods": ["my_totp", "my_totp2"]
    },
    "test/+/segment": {
      "capabilities": ["create", "sudo"]
    },
    "test/segment/at/end/+": {
      "capabilities": ["create", "sudo"]
    },
    "test/segment/at/end/v2/+/": {
      "capabilities": ["create", "sudo"]
    },
    "test/+/wildcard/+/*": {
      "capabilities": ["create", "sudo"]
    },
    "test/+/wildcard/+/end*": {
      "capabilities": ["create", "sudo"]
    },
    "paginated-kv/metadata": {
      "capabilities": ["list"],
      "pagination_limit": 12345
    },
    "unpaginated-kv/metadata": {
      "capabilities": ["list"]
    },
    "test/metadata/*": {
      "capabilities": ["list"],
      "list_scan_response_keys_filter_path": "test/metadata/{{ .key }}"
    }
  }
}
`)

func TestPolicy_Parse(t *testing.T) {
	t.Run("HCL", func(t *testing.T) {
		pHcl, err := ParseACLPolicy(namespace.RootNamespace, rawPolicy)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		validatePolicy(t, pHcl)
	})

	/*
		TODO(ascheel): When https://github.com/hashicorp/hcl/pull/741 merges, we'll
		want to update and uncomment this test.

		t.Run("JSON", func(t *testing.T) {
			var parsed map[string]interface{}
			err := json.Unmarshal([]byte(rawPolicyJSON), &parsed)
			if err != nil {
				t.Fatalf("failed to parse JSON: %v", err)
			}

			pJson, err := ParseACLPolicy(namespace.RootNamespace, rawPolicyJSON)
			if err != nil {
				t.Fatalf("err: %v", err)
			}

			t.Logf("value: %#v", pJson.Paths[8])

			validatePolicy(t, pJson)
		})
	*/
}

func validatePolicy(t *testing.T, p *Policy) {
	if p.Name != "dev" {
		t.Fatalf("bad name: %q", p.Name)
	}

	expect := []*PathRules{
		{
			Path:   "",
			Policy: "deny",
			Capabilities: []string{
				"deny",
			},
			Permissions: &ACLPermissions{CapabilitiesBitmap: DenyCapabilityInt},
			IsPrefix:    true,
		},
		{
			Path:   "stage/",
			Policy: "sudo",
			Capabilities: []string{
				"create",
				"read",
				"update",
				"delete",
				"list",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | ReadCapabilityInt | UpdateCapabilityInt | DeleteCapabilityInt | ListCapabilityInt | SudoCapabilityInt),
			},
			IsPrefix: true,
		},
		{
			Path:   "prod/version",
			Policy: "read",
			Capabilities: []string{
				"read",
				"list",
			},
			Permissions: &ACLPermissions{CapabilitiesBitmap: (ReadCapabilityInt | ListCapabilityInt)},
		},
		{
			Path:   "foo/bar",
			Policy: "read",
			Capabilities: []string{
				"read",
				"list",
			},
			MinWrappingTTLHCL: 300,
			MaxWrappingTTLHCL: "1h",
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (ReadCapabilityInt | ListCapabilityInt),
				MinWrappingTTL:     300 * time.Second,
				MaxWrappingTTL:     3600 * time.Second,
			},
		},
		{
			Path: "foo/bar",
			Capabilities: []string{
				"create",
				"sudo",
			},
			MinWrappingTTLHCL: "300s",
			MaxWrappingTTLHCL: 3600,
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				MinWrappingTTL:     300 * time.Second,
				MaxWrappingTTL:     3600 * time.Second,
			},
		},
		{
			Path: "foo/bar",
			Capabilities: []string{
				"create",
				"sudo",
			},
			AllowedParametersHCL: map[string][]interface{}{"zip": {}, "zap": {}},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				AllowedParameters:  map[string][]interface{}{"zip": {}, "zap": {}},
			},
		},
		{
			Path: "baz/bar",
			Capabilities: []string{
				"create",
				"sudo",
			},
			DeniedParametersHCL: map[string][]interface{}{"zip": {}, "zap": {}},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				DeniedParameters:   map[string][]interface{}{"zip": {}, "zap": {}},
			},
		},
		{
			Path: "biz/bar",
			Capabilities: []string{
				"create",
				"sudo",
			},
			AllowedParametersHCL: map[string][]interface{}{"zim": {}, "zam": {}},
			DeniedParametersHCL:  map[string][]interface{}{"zip": {}, "zap": {}},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				AllowedParameters:  map[string][]interface{}{"zim": {}, "zam": {}},
				DeniedParameters:   map[string][]interface{}{"zip": {}, "zap": {}},
			},
		},
		{
			Path:   "test/types",
			Policy: "",
			Capabilities: []string{
				"create",
				"sudo",
			},
			AllowedParametersHCL: map[string][]interface{}{"map": {map[string]interface{}{"good": "one"}}, "int": {1, 2}},
			DeniedParametersHCL:  map[string][]interface{}{"string": {"test"}, "bool": {false}},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				AllowedParameters:  map[string][]interface{}{"map": {map[string]interface{}{"good": "one"}}, "int": {1, 2}},
				DeniedParameters:   map[string][]interface{}{"string": {"test"}, "bool": {false}},
			},
			IsPrefix: false,
		},
		{
			Path: "test/req",
			Capabilities: []string{
				"create",
				"sudo",
			},
			RequiredParametersHCL: []string{"foo"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				RequiredParameters: []string{"foo"},
			},
		},
		{
			Path:         "test/patch",
			Capabilities: []string{"patch"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (PatchCapabilityInt),
			},
		},
		{
			Path:         "test/scan",
			Capabilities: []string{"scan"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (ScanCapabilityInt),
			},
		},
		{
			Path: "test/mfa",
			Capabilities: []string{
				"create",
				"sudo",
			},
			MFAMethodsHCL: []string{
				"my_totp",
				"my_totp2",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
				MFAMethods: []string{
					"my_totp",
					"my_totp2",
				},
			},
		},
		{
			Path: "test/+/segment",
			Capabilities: []string{
				"create",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
			},
			HasSegmentWildcards: true,
		},
		{
			Path: "test/segment/at/end/+",
			Capabilities: []string{
				"create",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
			},
			HasSegmentWildcards: true,
		},
		{
			Path: "test/segment/at/end/v2/+/",
			Capabilities: []string{
				"create",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
			},
			HasSegmentWildcards: true,
		},
		{
			Path: "test/+/wildcard/+/*",
			Capabilities: []string{
				"create",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
			},
			HasSegmentWildcards: true,
		},
		{
			Path: "test/+/wildcard/+/end*",
			Capabilities: []string{
				"create",
				"sudo",
			},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: (CreateCapabilityInt | SudoCapabilityInt),
			},
			HasSegmentWildcards: true,
		},
		{
			Path:         "paginated-kv/metadata",
			Capabilities: []string{"list"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: ListCapabilityInt,
				PaginationLimit:    12345,
			},
			PaginationLimitHCL: 12345,
		},
		{
			Path:         "unpaginated-kv/metadata",
			Capabilities: []string{"list"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap: ListCapabilityInt,
			},
		},
		{
			Path:         "test/metadata/",
			Capabilities: []string{"list"},
			Permissions: &ACLPermissions{
				CapabilitiesBitmap:     ListCapabilityInt,
				ResponseKeysFilterPath: "test/metadata/{{ .key }}",
			},
			ResponseKeysFilterPathHCL: "test/metadata/{{ .key }}",
			IsPrefix:                  true,
		},
	}

	if diff := deep.Equal(p.Paths, expect); diff != nil {
		t.Error(diff)
	}
}

func TestPolicy_ParseBadRoot(t *testing.T) {
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
name = "test"
bad  = "foo"
nope = "yes"
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `invalid key "bad" on line 2`) {
		t.Errorf("bad error: %q", err)
	}

	if !strings.Contains(err.Error(), `invalid key "nope" on line 3`) {
		t.Errorf("bad error: %q", err)
	}
}

func TestPolicy_ParseBadPath(t *testing.T) {
	// The wrong spelling is intended here
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
path "/" {
	capabilities = ["read"]
	capabilites  = ["read"]
}
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `invalid key "capabilites" on line 3`) {
		t.Errorf("bad error: %s", err)
	}
}

func TestPolicy_ParseBadPolicy(t *testing.T) {
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
path "/" {
	policy = "banana"
}
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `path "/": invalid policy "banana"`) {
		t.Errorf("bad error: %s", err)
	}
}

func TestPolicy_ParseBadWrapping(t *testing.T) {
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
path "/" {
	policy = "read"
	min_wrapping_ttl = 400
	max_wrapping_ttl = 200
}
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `max_wrapping_ttl cannot be less than min_wrapping_ttl`) {
		t.Errorf("bad error: %s", err)
	}
}

func TestPolicy_ParseBadCapabilities(t *testing.T) {
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
path "/" {
	capabilities = ["read", "banana"]
}
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `path "/": invalid capability "banana"`) {
		t.Errorf("bad error: %s", err)
	}
}

func TestPolicy_ParseBadSegmentWildcard(t *testing.T) {
	_, err := ParseACLPolicy(namespace.RootNamespace, strings.TrimSpace(`
path "foo/+*" {
	capabilities = ["read"]
}
`))
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `path "foo/+*": invalid use of wildcards ('+*' is forbidden)`) {
		t.Errorf("bad error: %s", err)
	}
}
