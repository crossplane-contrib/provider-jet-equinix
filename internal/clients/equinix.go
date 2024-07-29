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

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	equinixprovider "github.com/equinix/terraform-provider-equinix/equinix/provider"

	"github.com/equinix/terraform-provider-equinix/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	terraformsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-jet-equinix/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal equinix credentials as JSON"

	keyClientID            = "client_id"
	keyClientSecret        = "client_secret"
	keyAuthToken           = "auth_token"
	keyEndpoint            = "endpoint"
	keyRequestTimeout      = "request_timeout"
	keyResponseMaxPageSize = "response_max_page_size"
	keyMaxRetries          = "max_retries"
	keyMaxRetryWaitSeconds = "max_retry_wait_seconds"
)

type SetupConfig struct {
	ProviderSource     *string
	ProviderVersion    *string
	TerraformVersion   *string
	TerraformProvider  *schema.Provider
	NativeProviderPath *string
	DefaultScheduler   terraform.ProviderScheduler
}

func prepareTerraformProviderConfiguration(creds map[string]string, pc v1beta1.ProviderConfiguration) map[string]any {
	config := map[string]any{}
	config[keyMaxRetries] = pc.MaxRetries
	config[keyMaxRetryWaitSeconds] = pc.MaxRetryWaitSeconds
	config[keyRequestTimeout] = pc.RequestTimeout
	config[keyResponseMaxPageSize] = pc.ResponseMaxPageSize
	config[keyEndpoint] = pc.Endpoint

	// TODO: should we deprecate config field setting via credentials
	for _, key := range []string{
		keyEndpoint,
		keyRequestTimeout,
		keyResponseMaxPageSize,
		keyAuthToken,
		keyClientID,
		keyClientSecret,
	} {
		if creds[key] != "" {
			config[key] = creds[key]
		}
	}

	return config
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(setupCfg SetupConfig) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: *setupCfg.TerraformVersion,
			Requirement: terraform.ProviderRequirement{
				Source:  *setupCfg.ProviderSource,
				Version: *setupCfg.ProviderVersion,
			},
			Scheduler: setupCfg.DefaultScheduler,
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		equinixCreds := map[string]string{}
		if err := json.Unmarshal(data, &equinixCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = prepareTerraformProviderConfiguration(equinixCreds, pc.Spec.Configuration)
		return ps, errors.Wrap(configureTerraformPluginSDKEquinixClient(ctx, &ps, *setupCfg.TerraformProvider), "failed to configure the Terraform Plugin SDK Equinix client")
	}
}

func configureTerraformPluginSDKEquinixClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &terraformsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}

	fwProvider := equinixprovider.CreateFrameworkProvider(version.ProviderVersion)
	ps.FrameworkProvider = fwProvider
	return nil
}
