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

type ReservedIPBlockObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	AddressFamily *float64 `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	CidrNotation *string `json:"cidrNotation,omitempty" tf:"cidr_notation,omitempty"`

	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Manageable *bool `json:"manageable,omitempty" tf:"manageable,omitempty"`

	Management *bool `json:"management,omitempty" tf:"management,omitempty"`

	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	Public *bool `json:"public,omitempty" tf:"public,omitempty"`
}

type ReservedIPBlockParameters struct {

	// the size of the network to reserve from an existing vrf ip_range. `cidr` can only be specified with `vrf_id`. Minimum range is 22-29, with 30-31 supported and necessary for virtual-circuits
	// +kubebuilder:validation:Optional
	Cidr *float64 `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may be helpful for self-managed IPAM. The object must be valid JSON.
	// +kubebuilder:validation:Optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data,omitempty"`

	// Arbitrary description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with metro
	// +kubebuilder:validation:Optional
	Facility *string `json:"facility,omitempty" tf:"facility,omitempty"`

	// Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with facility
	// +kubebuilder:validation:Optional
	Metro *string `json:"metro,omitempty" tf:"metro,omitempty"`

	// an unreserved network address from an existing vrf ip_range. `network` can only be specified with vrf_id
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The metal project ID where to allocate the address block
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// The number of allocated /32 addresses, a power of 2
	// +kubebuilder:validation:Optional
	Quantity *float64 `json:"quantity,omitempty" tf:"quantity,omitempty"`

	// Tags attached to the reserved block
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Either global_ipv4, public_ipv4, or vrf. Defaults to public_ipv4.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// VRF ID for type=vrf reservations
	// +crossplane:generate:reference:type=Vrf
	// +kubebuilder:validation:Optional
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`

	// Wait for the IP reservation block to reach a desired state on resource creation. One of: `pending`, `created`. The `created` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the `denied` state is encountered.
	// +kubebuilder:validation:Optional
	WaitForState *string `json:"waitForState,omitempty" tf:"wait_for_state,omitempty"`
}

// ReservedIPBlockSpec defines the desired state of ReservedIPBlock
type ReservedIPBlockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservedIPBlockParameters `json:"forProvider"`
}

// ReservedIPBlockStatus defines the observed state of ReservedIPBlock.
type ReservedIPBlockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservedIPBlockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReservedIPBlock is the Schema for the ReservedIPBlocks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type ReservedIPBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReservedIPBlockSpec   `json:"spec"`
	Status            ReservedIPBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservedIPBlockList contains a list of ReservedIPBlocks
type ReservedIPBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReservedIPBlock `json:"items"`
}

// Repository type metadata.
var (
	ReservedIPBlock_Kind             = "ReservedIPBlock"
	ReservedIPBlock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReservedIPBlock_Kind}.String()
	ReservedIPBlock_KindAPIVersion   = ReservedIPBlock_Kind + "." + CRDGroupVersion.String()
	ReservedIPBlock_GroupVersionKind = CRDGroupVersion.WithKind(ReservedIPBlock_Kind)
)

func init() {
	SchemeBuilder.Register(&ReservedIPBlock{}, &ReservedIPBlockList{})
}
