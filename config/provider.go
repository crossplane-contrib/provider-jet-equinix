/*
Copyright 2021 The Crossplane Authors.

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

	upconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-equinix/config/ecx/l2connection"
	"github.com/crossplane-contrib/provider-jet-equinix/config/metal/device"
)

const (
	resourcePrefix = "equinix"
	modulePath     = "github.com/crossplane-contrib/provider-jet-equinix"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *upconfig.Provider {
	pc := upconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		upconfig.WithShortName("equinix"),
		upconfig.WithRootGroup("equinix.jet.crossplane.io"),
		// upconfig.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector("github.com/crossplane-contrib/provider-jet-equinix")}),
		upconfig.WithDefaultResourceOptions(
			KnownReferencers(),
			IdentifierAssignedByEquinix(),
		),
		upconfig.WithIncludeList([]string{
			".*",
		}),
		upconfig.WithBasePackages(upconfig.BasePackages{
			APIVersion: []string{
				// Default package for ProviderConfig APIs
				"apis/v1alpha1",
			},
			Controller: []string{
				// Default package for ProviderConfig controllers
				"internal/controller/providerconfig",
			},
		}),
	)

	for _, configure := range []func(provider *upconfig.Provider){
		// add custom config functions
		device.Configure,
		l2connection.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
