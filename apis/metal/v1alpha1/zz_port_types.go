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

type PortInitParameters struct {

	// Flag indicating whether the port should be bonded
	Bonded *bool `json:"bonded,omitempty" tf:"bonded,omitempty"`

	// Flag indicating whether the port is in layer2 (or layer3) mode. The `layer2` flag can be set only for bond ports.
	Layer2 *bool `json:"layer2,omitempty" tf:"layer2,omitempty"`

	// UUID of native VLAN of the port
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-equinix/apis/metal/v1alpha1.Vlan
	NativeVlanID *string `json:"nativeVlanId,omitempty" tf:"native_vlan_id,omitempty"`

	// Reference to a Vlan in metal to populate nativeVlanId.
	// +kubebuilder:validation:Optional
	NativeVlanIDRef *v1.Reference `json:"nativeVlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan in metal to populate nativeVlanId.
	// +kubebuilder:validation:Optional
	NativeVlanIDSelector *v1.Selector `json:"nativeVlanIdSelector,omitempty" tf:"-"`

	// UUID of the port to lookup
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Behavioral setting to reset the port to default settings (layer3 bonded mode without any vlan attached) before delete/destroy
	ResetOnDelete *bool `json:"resetOnDelete,omitempty" tf:"reset_on_delete,omitempty"`

	// UUIDs VLANs to attach. To avoid jitter, use the UUID and not the VXLAN
	// +listType=set
	VlanIds []*string `json:"vlanIds,omitempty" tf:"vlan_ids,omitempty"`

	// VLAN VXLAN ids to attach (example: [1000])
	// +listType=set
	VxlanIds []*float64 `json:"vxlanIds,omitempty" tf:"vxlan_ids,omitempty"`
}

type PortObservation struct {

	// UUID of the bond port
	BondID *string `json:"bondId,omitempty" tf:"bond_id,omitempty"`

	// Name of the bond port
	BondName *string `json:"bondName,omitempty" tf:"bond_name,omitempty"`

	// Flag indicating whether the port should be bonded
	Bonded *bool `json:"bonded,omitempty" tf:"bonded,omitempty"`

	// Flag indicating whether the port can be removed from a bond
	DisbondSupported *bool `json:"disbondSupported,omitempty" tf:"disbond_supported,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Flag indicating whether the port is in layer2 (or layer3) mode. The `layer2` flag can be set only for bond ports.
	Layer2 *bool `json:"layer2,omitempty" tf:"layer2,omitempty"`

	// MAC address of the port
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Name of the port to look up, e.g. bond0, eth1
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// UUID of native VLAN of the port
	NativeVlanID *string `json:"nativeVlanId,omitempty" tf:"native_vlan_id,omitempty"`

	// One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// UUID of the port to lookup
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Behavioral setting to reset the port to default settings (layer3 bonded mode without any vlan attached) before delete/destroy
	ResetOnDelete *bool `json:"resetOnDelete,omitempty" tf:"reset_on_delete,omitempty"`

	// Port type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// UUIDs VLANs to attach. To avoid jitter, use the UUID and not the VXLAN
	// +listType=set
	VlanIds []*string `json:"vlanIds,omitempty" tf:"vlan_ids,omitempty"`

	// VLAN VXLAN ids to attach (example: [1000])
	// +listType=set
	VxlanIds []*float64 `json:"vxlanIds,omitempty" tf:"vxlan_ids,omitempty"`
}

type PortParameters struct {

	// Flag indicating whether the port should be bonded
	// +kubebuilder:validation:Optional
	Bonded *bool `json:"bonded,omitempty" tf:"bonded,omitempty"`

	// Flag indicating whether the port is in layer2 (or layer3) mode. The `layer2` flag can be set only for bond ports.
	// +kubebuilder:validation:Optional
	Layer2 *bool `json:"layer2,omitempty" tf:"layer2,omitempty"`

	// UUID of native VLAN of the port
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-equinix/apis/metal/v1alpha1.Vlan
	// +kubebuilder:validation:Optional
	NativeVlanID *string `json:"nativeVlanId,omitempty" tf:"native_vlan_id,omitempty"`

	// Reference to a Vlan in metal to populate nativeVlanId.
	// +kubebuilder:validation:Optional
	NativeVlanIDRef *v1.Reference `json:"nativeVlanIdRef,omitempty" tf:"-"`

	// Selector for a Vlan in metal to populate nativeVlanId.
	// +kubebuilder:validation:Optional
	NativeVlanIDSelector *v1.Selector `json:"nativeVlanIdSelector,omitempty" tf:"-"`

	// UUID of the port to lookup
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Behavioral setting to reset the port to default settings (layer3 bonded mode without any vlan attached) before delete/destroy
	// +kubebuilder:validation:Optional
	ResetOnDelete *bool `json:"resetOnDelete,omitempty" tf:"reset_on_delete,omitempty"`

	// UUIDs VLANs to attach. To avoid jitter, use the UUID and not the VXLAN
	// +kubebuilder:validation:Optional
	// +listType=set
	VlanIds []*string `json:"vlanIds,omitempty" tf:"vlan_ids,omitempty"`

	// VLAN VXLAN ids to attach (example: [1000])
	// +kubebuilder:validation:Optional
	// +listType=set
	VxlanIds []*float64 `json:"vxlanIds,omitempty" tf:"vxlan_ids,omitempty"`
}

// PortSpec defines the desired state of Port
type PortSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PortParameters `json:"forProvider"`
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
	InitProvider PortInitParameters `json:"initProvider,omitempty"`
}

// PortStatus defines the observed state of Port.
type PortStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Port is the Schema for the Ports API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Port struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bonded) || (has(self.initProvider) && has(self.initProvider.bonded))",message="spec.forProvider.bonded is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portId) || (has(self.initProvider) && has(self.initProvider.portId))",message="spec.forProvider.portId is a required parameter"
	Spec   PortSpec   `json:"spec"`
	Status PortStatus `json:"status,omitempty"`
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
