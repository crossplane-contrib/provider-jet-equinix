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
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/terrajet/pkg/terraform"

	"github.com/crossplane-contrib/provider-jet-equinix/apis/v1alpha1"
)

const (
	fmtEnvVar = "%s=%s"

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

	// Equinix credentials environment variable names
	envClientID     = "EQUINIX_API_CLIENTID"
	envClientSecret = "EQUINIX_API_CLIENTSECRET"
	envAuthToken    = "METAL_AUTH_TOKEN"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
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

		// set environment variables for sensitive provider configuration
		// Deprecated: In shared gRPC mode we do not support injecting
		// credentials via the environment variables. You should specify
		// credentials via the Terraform main.tf.json instead.
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envClientID, equinixCreds[keyClientID]),
			fmt.Sprintf(fmtEnvVar, envClientSecret, equinixCreds[keyClientSecret]),
			fmt.Sprintf(fmtEnvVar, envAuthToken, equinixCreds[keyAuthToken]),
		}
		// set credentials in Terraform provider configuration
		ps.Configuration = map[string]interface{}{}
		for _, key := range []string{
			keyEndpoint,
			keyRequestTimeout,
			keyResponseMaxPageSize,
		} {
			if equinixCreds[key] != "" {
				ps.Configuration[key] = equinixCreds[key]
			}
		}
		return ps, nil
	}
}
