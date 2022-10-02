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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectAPIKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProjectAPIKeyParameters struct {

	// Description string for the API key
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// UUID of project which the new API key is scoped to
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Flag indicating whether the API key shoud be read-only
	// +kubebuilder:validation:Required
	ReadOnly *bool `json:"readOnly" tf:"read_only,omitempty"`
}

// ProjectAPIKeySpec defines the desired state of ProjectAPIKey
type ProjectAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectAPIKeyParameters `json:"forProvider"`
}

// ProjectAPIKeyStatus defines the observed state of ProjectAPIKey.
type ProjectAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectAPIKey is the Schema for the ProjectAPIKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type ProjectAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectAPIKeySpec   `json:"spec"`
	Status            ProjectAPIKeyStatus `json:"status,omitempty"`
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
