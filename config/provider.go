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
	"context"
	_ "embed"

	upconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	"github.com/equinix/terraform-provider-equinix/equinix"
	framework "github.com/equinix/terraform-provider-equinix/equinix/provider"
	"github.com/equinix/terraform-provider-equinix/version"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const (
	resourcePrefix = "equinix"
	modulePath     = "github.com/crossplane-contrib/provider-jet-equinix"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*upconfig.Provider, error) {
	var p *schema.Provider
	var err error

	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = equinix.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	fwProvider := framework.CreateFrameworkProvider(version.ProviderVersion)

	pc := upconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		upconfig.WithShortName("equinix"),
		upconfig.WithRootGroup("equinix.jet.crossplane.io"),
		// upconfig.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector("github.com/crossplane-contrib/provider-jet-equinix")}),
		upconfig.WithDefaultResourceOptions(
			KnownReferencers(),
			IdentifierAssignedByEquinix(),
			SkipOptCompLateInitialization(),
		),
		upconfig.WithFeaturesPackage("internal/features"),
		upconfig.WithTerraformProvider(p),
		upconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		upconfig.WithIncludeList([]string{
			".*",
		}),
		upconfig.WithSkipList([]string{
			// ".*", // helpful when debugging to minimize the number of resources
		}),
		// config.WithTerraformPluginSDKIncludeList(resourceList(terraformSDKIncludeList)),
		// config.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
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
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, err
}
