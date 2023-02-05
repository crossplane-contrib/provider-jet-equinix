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

type IPAttachmentObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Address family as integer (4 or 6)
	AddressFamily *float64 `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// Length of CIDR prefix of the block as integer
	Cidr *float64 `json:"cidr,omitempty" tf:"cidr,omitempty"`

	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Flag indicating whether IP block is global, i.e. assignable in any location
	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Manageable *bool `json:"manageable,omitempty" tf:"manageable,omitempty"`

	Management *bool `json:"management,omitempty" tf:"management,omitempty"`

	// Mask in decimal notation, e.g. 255.255.255.0
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// Network IP address portion of the block specification
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Flag indicating whether IP block is addressable from the Internet
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type IPAttachmentParameters struct {

	// +kubebuilder:validation:Required
	CidrNotation *string `json:"cidrNotation" tf:"cidr_notation,omitempty"`

	// +crossplane:generate:reference:type=Device
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Reference to a Device to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDRef *v1.Reference `json:"deviceIdRef,omitempty" tf:"-"`

	// Selector for a Device to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDSelector *v1.Selector `json:"deviceIdSelector,omitempty" tf:"-"`
}

// IPAttachmentSpec defines the desired state of IPAttachment
type IPAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPAttachmentParameters `json:"forProvider"`
}

// IPAttachmentStatus defines the observed state of IPAttachment.
type IPAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPAttachment is the Schema for the IPAttachments API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type IPAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPAttachmentSpec   `json:"spec"`
	Status            IPAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPAttachmentList contains a list of IPAttachments
type IPAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPAttachment `json:"items"`
}

// Repository type metadata.
var (
	IPAttachment_Kind             = "IPAttachment"
	IPAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPAttachment_Kind}.String()
	IPAttachment_KindAPIVersion   = IPAttachment_Kind + "." + CRDGroupVersion.String()
	IPAttachment_GroupVersionKind = CRDGroupVersion.WithKind(IPAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&IPAttachment{}, &IPAttachmentList{})
}
