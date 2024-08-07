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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SSHKeyInitParameters_2 struct {

	// The name of SSH key used for identification.
	// The name of SSH key used for identification
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique Identifier for the project resource where the SSH key is scoped to.If you leave it out, the ssh key will be created under the default project id of your organization.
	// The unique identifier of Project Resource to which ssh key is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The SSH public key. If this is a file, it can be read using the file interpolation function.
	// The SSH public key. If this is a file, it can be read using the file interpolation function
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The type of SSH key: RSA (default) or DSA.
	// The type of SSH key: RSA (default) or DSA
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SSHKeyObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of SSH key used for identification.
	// The name of SSH key used for identification
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique Identifier for the project resource where the SSH key is scoped to.If you leave it out, the ssh key will be created under the default project id of your organization.
	// The unique identifier of Project Resource to which ssh key is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The SSH public key. If this is a file, it can be read using the file interpolation function.
	// The SSH public key. If this is a file, it can be read using the file interpolation function
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The type of SSH key: RSA (default) or DSA.
	// The type of SSH key: RSA (default) or DSA
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The unique identifier of the key
	// The unique identifier of the key
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type SSHKeyParameters_2 struct {

	// The name of SSH key used for identification.
	// The name of SSH key used for identification
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique Identifier for the project resource where the SSH key is scoped to.If you leave it out, the ssh key will be created under the default project id of your organization.
	// The unique identifier of Project Resource to which ssh key is scoped to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The SSH public key. If this is a file, it can be read using the file interpolation function.
	// The SSH public key. If this is a file, it can be read using the file interpolation function
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The type of SSH key: RSA (default) or DSA.
	// The type of SSH key: RSA (default) or DSA
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHKeyParameters_2 `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SSHKeyInitParameters_2 `json:"initProvider,omitempty"`
}

// SSHKeyStatus defines the observed state of SSHKey.
type SSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHKeyObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSHKey is the Schema for the SSHKeys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   SSHKeySpec   `json:"spec"`
	Status SSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKeyList contains a list of SSHKeys
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}

// Repository type metadata.
var (
	SSHKey_Kind             = "SSHKey"
	SSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHKey_Kind}.String()
	SSHKey_KindAPIVersion   = SSHKey_Kind + "." + CRDGroupVersion.String()
	SSHKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
