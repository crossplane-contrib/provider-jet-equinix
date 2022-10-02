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

	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// IdentifierAssignedByEquinix will work for all Equinix types because even if
// the ID is assigned by user, we'll see it in the TF State ID. The
// resource-specific configurations should override this whenever possible.
func IdentifierAssignedByEquinix() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	}
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() tjconfig.ResourceOption { //nolint:gocyclo
	return func(r *tjconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			if r.ShortGroup == "metal" {
				switch {
				case strings.HasSuffix(k, "project_id"):
					r.References[k] = tjconfig.Reference{
						// github.com/crossplane-contrib/provider-jet-equinix/apis/metal/v1alpha1.Project
						Type: "Project",
					}
				case strings.HasSuffix(k, "organization_id"):
					r.References[k] = tjconfig.Reference{
						Type: "Organization",
					}
				case strings.HasSuffix(k, "connection_id"):
					r.References[k] = tjconfig.Reference{
						Type: "Connection",
					}
				case strings.HasSuffix(k, "device_id"):
					r.References[k] = tjconfig.Reference{
						Type: "Device",
					}
				case strings.HasSuffix(k, "vlan_id"):
					// vlan_vnid is ignored because it is an int type
					r.References[k] = tjconfig.Reference{
						Type: "Vlan",
					}
				case strings.HasSuffix(k, "vrf_id"):
					r.References[k] = tjconfig.Reference{
						Type: "Vrf",
					}
				case strings.HasSuffix(k, "ip_reservation_id"):
					r.References[k] = tjconfig.Reference{
						Type: "ReservedIPBlock",
					}
				}
			}
		}
	}
}
