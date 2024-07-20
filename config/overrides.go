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
func IdentifierAssignedByEquinix() upconfig.ResourceOption {
	return func(r *upconfig.Resource) {
		r.ExternalName = upconfig.IdentifierFromProvider
	}
}

var knownReferencerFields = map[string]map[string]string{
	"metal": {
		"project_id":         "Project",
		"organization_id":    "Organization",
		"connection_id":      "Connection",
		"device_id":          "Device",
		"vlan_id":            "Vlan",
		"vrf_id":             "Vrf",
		"ip_reservation_id":  "ReservedIPBlock",
		"virtual_circuit_id": "VirtualCircuit",
		"gateway_id":         "Gateway",
	},
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() upconfig.ResourceOption {
	return func(r *upconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}

			// Loop over knownReferencerFields and add references
			for suffix, resource := range knownReferencerFields[r.ShortGroup] {
				if !strings.HasSuffix(k, suffix) {
					continue
				}
				r.References[k] = upconfig.Reference{
					Type: resource,
				}
			}
		}
	}
}

// SkipOptCompLateIntialization generalize the LateInitializer rule above to apply to allow fields that are Optional + Computed + ConflictsWith another Computed + Optional field
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
