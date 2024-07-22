/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"strings"

	upconfig "github.com/crossplane/upjet/pkg/config"
)

// IdentifierAssignedByEquinix will work for all Equinix types because even if
// the ID is assigned by user, we'll see it in the TF State ID. The
// resource-specific configurations should override this whenever possible.
// See https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#case-2-identifier-from-provider for more details.
func IdentifierAssignedByEquinix() upconfig.ResourceOption {
	return func(r *upconfig.Resource) {
		r.ExternalName = upconfig.IdentifierFromProvider
	}
}

var knownReferencerTFResource = map[string]map[string]string{
	"metal": {
		"project_id":          "equinix_metal_project",
		"project_ids":         "equinix_metal_project",
		"organization_id":     "equinix_metal_organization",
		"connection_id":       "equinix_metal_connection",
		"device_id":           "equinix_metal_device",
		"vlan_id":             "equinix_metal_vlan",
		"vlan_ids":            "equinix_metal_vlan",
		"vrf_id":              "equinix_metal_vrf",
		"ip_reservation_id":   "equinix_metal_reserved_ip_block",
		"virtual_circuit_id":  "equinix_metal_virtual_circuit",
		"gateway_id":          "equinix_metal_gateway",
		"project_ssh_key_ids": "equinix_metal_project_ssh_key",
		"user_ssh_key_ids":    "equinix_metal_ssh_key",
		// "ssh_key_ids" // These would be ambiguously matched to the above at random, so leave them out.
	},
	"network": {
		// TODO: need a way to disambiguate id from redundant_uuid. the user may want to reference both from the same equinix_network_device.
		// "device_ids": "equinix_network_device",
	},
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
// See https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#cross-resource-referencing for more details.
func KnownReferencers() upconfig.ResourceOption {
	return func(r *upconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}

			// Loop over knownReferencerFields and add references
			for suffix, resource := range knownReferencerTFResource[r.ShortGroup] {
				if !strings.HasSuffix(k, suffix) {
					continue
				}

				r.References[k] = upconfig.Reference{
					TerraformName: resource,
				}
				break // TODO: if there are multiple suffix matches, only the first processed will be handled. Go map order is random.
			}
		}
	}
}

// SkipOptCompLateIntialization generalize the LateInitializer rule above to apply to allow fields that are Optional + Computed + ConflictsWith another Computed + Optional field
// See https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#further-details-on-late-initialization for details on this mutually-exclusive scenario
// See https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#overriding-terraform-resource-schema for default behavior
func SkipOptCompLateInitialization() upconfig.ResourceOption {
	return func(r *upconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			if s.Computed && s.Optional {
				for _, conflict := range s.ConflictsWith {
					if cs := r.TerraformResource.Schema[conflict]; cs.Computed && cs.Optional {
						r.LateInitializer.AddIgnoredCanonicalFields(conflict)
						r.LateInitializer.AddIgnoredCanonicalFields(k)
					}
				}
			}
		}
	}
}
