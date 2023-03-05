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

type AddressObservation struct {
}

type AddressParameters struct {

	// Postal address.
	// Postal address
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// City name.
	// City name
	// +kubebuilder:validation:Required
	City *string `json:"city" tf:"city,omitempty"`

	// Two letter country code (ISO 3166-1 alpha-2), e.g. US.
	// Two letter country code (ISO 3166-1 alpha-2), e.g. US
	// +kubebuilder:validation:Required
	Country *string `json:"country" tf:"country,omitempty"`

	// State name.
	// State name
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Zip Code.
	// Zip Code
	// +kubebuilder:validation:Required
	ZipCode *string `json:"zipCode" tf:"zip_code,omitempty"`
}

type OrganizationObservation struct {

	// The timestamp for when the organization was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The unique ID of the organization.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp for the last time the organization was updated.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type OrganizationParameters struct {

	// An object that has the address information. See Address
	// below for more details.
	// Address information block
	// +kubebuilder:validation:Required
	Address []AddressParameters `json:"address" tf:"address,omitempty"`

	// Description string.
	// Description string
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Logo URL.
	// Logo URL
	// +kubebuilder:validation:Optional
	Logo *string `json:"logo,omitempty" tf:"logo,omitempty"`

	// The name of the Organization.
	// The name of the Organization
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Twitter handle.
	// Twitter handle
	// +kubebuilder:validation:Optional
	Twitter *string `json:"twitter,omitempty" tf:"twitter,omitempty"`

	// Website link.
	// Website link
	// +kubebuilder:validation:Optional
	Website *string `json:"website,omitempty" tf:"website,omitempty"`
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Organization is the Schema for the Organizations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSpec   `json:"spec"`
	Status            OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
