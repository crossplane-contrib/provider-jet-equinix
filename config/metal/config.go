// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0
package metal

import (
	"context"
	"errors"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the metal group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("equinix_metal_port_vlan_attachment", func(r *config.Resource) {
		// ID assigned in Terraform is PortID + ":" + VlanID
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.IdentifierFields = []string{"port_id", "vlan_id"}
		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, cfg map[string]interface{}) (string, error) {
			portID, ok := parameters["port_id"]
			if !ok || portID.(string) == "" {
				return "", errors.New("port_id cannot be empty")
			}
			vlanID, ok := parameters["vlan_id"]
			if !ok || vlanID.(string) == "" {
				return "", errors.New("vlan_id cannot be empty")
			}
			return fmt.Sprintf("%s:%s", portID.(string), vlanID.(string)), nil
		}
	})
}
