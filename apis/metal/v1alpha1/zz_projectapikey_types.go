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

type ProjectAPIKeyInitParameters struct {

	// Description string for the Project API Key resource.
	// Description string for the API key
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// UUID of the project where the API key is scoped to.
	// UUID of project which the new API key is scoped to
	// +crossplane:generate:reference:type=Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type ProjectAPIKeyObservation struct {

	// Description string for the Project API Key resource.
	// Description string for the API key
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// UUID of the project where the API key is scoped to.
	// UUID of project which the new API key is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type ProjectAPIKeyParameters struct {

	// Description string for the Project API Key resource.
	// Description string for the API key
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// UUID of the project where the API key is scoped to.
	// UUID of project which the new API key is scoped to
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Flag indicating whether the API key shoud be read-only
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

// ProjectAPIKeySpec defines the desired state of ProjectAPIKey
type ProjectAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectAPIKeyParameters `json:"forProvider"`
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
	InitProvider ProjectAPIKeyInitParameters `json:"initProvider,omitempty"`
}

// ProjectAPIKeyStatus defines the observed state of ProjectAPIKey.
type ProjectAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectAPIKey is the Schema for the ProjectAPIKeys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type ProjectAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.readOnly) || (has(self.initProvider) && has(self.initProvider.readOnly))",message="spec.forProvider.readOnly is a required parameter"
	Spec   ProjectAPIKeySpec   `json:"spec"`
	Status ProjectAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectAPIKeyList contains a list of ProjectAPIKeys
type ProjectAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectAPIKey `json:"items"`
}

// Repository type metadata.
var (
	ProjectAPIKey_Kind             = "ProjectAPIKey"
	ProjectAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectAPIKey_Kind}.String()
	ProjectAPIKey_KindAPIVersion   = ProjectAPIKey_Kind + "." + CRDGroupVersion.String()
	ProjectAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(ProjectAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectAPIKey{}, &ProjectAPIKeyList{})
}
