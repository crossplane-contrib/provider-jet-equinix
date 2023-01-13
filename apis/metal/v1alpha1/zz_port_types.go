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

type PortObservation struct {
	BondID *string `json:"bondId,omitempty" tf:"bond_id,omitempty"`

	BondName *string `json:"bondName,omitempty" tf:"bond_name,omitempty"`

	DisbondSupported *bool `json:"disbondSupported,omitempty" tf:"disbond_supported,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PortParameters struct {

	// Flag indicating whether the port should be bonded
	// +kubebuilder:validation:Required
	Bonded *bool `json:"bonded" tf:"bonded,omitempty"`

	// Flag indicating whether the port is in layer2 (or layer3) mode. The `layer2` flag can be set only for bond ports.
	// +kubebuilder:validation:Optional
	Layer2 *bool `json:"layer2,omitempty" tf:"layer2,omitempty"`

	// UUID of native VLAN of the port
	// +crossplane:generate:reference:type=Vlan
	// +kubebuilder:validation:Optional
	NativeVlanID *string `json:"nativeVlanId,omitempty" tf:"native_vlan_id,omitempty"`

	// +kubebuilder:validation:Optional
	NativeVlanIDRef *v1.Reference `json:"nativeVlanIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NativeVlanIDSelector *v1.Selector `json:"nativeVlanIdSelector,omitempty" tf:"-"`

	// UUID of the port to lookup
	// +kubebuilder:validation:Required
	PortID *string `json:"portId" tf:"port_id,omitempty"`

	// Behavioral setting to reset the port to default settings (layer3 bonded mode without any vlan attached) before delete/destroy
	// +kubebuilder:validation:Optional
	ResetOnDelete *bool `json:"resetOnDelete,omitempty" tf:"reset_on_delete,omitempty"`

	// UUIDs VLANs to attach. To avoid jitter, use the UUID and not the VXLAN
	// +kubebuilder:validation:Optional
	VlanIds []*string `json:"vlanIds,omitempty" tf:"vlan_ids,omitempty"`

	// VLAN VXLAN ids to attach (example: [1000])
	// +kubebuilder:validation:Optional
	VxlanIds []*float64 `json:"vxlanIds,omitempty" tf:"vxlan_ids,omitempty"`
}

// PortSpec defines the desired state of Port
type PortSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortParameters `json:"forProvider"`
}

// PortStatus defines the observed state of Port.
type PortStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Port is the Schema for the Ports API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type Port struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortSpec   `json:"spec"`
	Status            PortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortList contains a list of Ports
type PortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Port `json:"items"`
}

// Repository type metadata.
var (
	Port_Kind             = "Port"
	Port_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Port_Kind}.String()
	Port_KindAPIVersion   = Port_Kind + "." + CRDGroupVersion.String()
	Port_GroupVersionKind = CRDGroupVersion.WithKind(Port_Kind)
)

func init() {
	SchemeBuilder.Register(&Port{}, &PortList{})
}
