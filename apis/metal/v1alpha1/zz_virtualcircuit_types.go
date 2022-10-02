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

type VirtualCircuitObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NniVnid *float64 `json:"nniVnid,omitempty" tf:"nni_vnid,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Vnid *float64 `json:"vnid,omitempty" tf:"vnid,omitempty"`
}

type VirtualCircuitParameters struct {

	// UUID of Connection where the VC is scoped to
	// +crossplane:generate:reference:type=Connection
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// +kubebuilder:validation:Optional
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// Description of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The password that can be set for the VRF BGP peer
	// +kubebuilder:validation:Optional
	Md5SecretRef *v1.SecretKeySelector `json:"md5SecretRef,omitempty" tf:"-"`

	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// +kubebuilder:validation:Optional
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// Name of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	// +kubebuilder:validation:Optional
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// UUID of the Connection Port where the VC is scoped to
	// +kubebuilder:validation:Required
	PortID *string `json:"portId" tf:"port_id,omitempty"`

	// UUID of the Project where the VC is scoped to
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// +kubebuilder:validation:Optional
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// Tags attached to the virtual circuit
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// UUID of the VLAN to associate
	// +crossplane:generate:reference:type=Vlan
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// +kubebuilder:validation:Optional
	VlanIDRef *v1.Reference `json:"vlanIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VlanIDSelector *v1.Selector `json:"vlanIdSelector,omitempty" tf:"-"`

	// UUID of the VRF to associate
	// +crossplane:generate:reference:type=Vrf
	// +kubebuilder:validation:Optional
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`

	// +kubebuilder:validation:Optional
	VrfIDRef *v1.Reference `json:"vrfIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VrfIDSelector *v1.Selector `json:"vrfIdSelector,omitempty" tf:"-"`
}

// VirtualCircuitSpec defines the desired state of VirtualCircuit
type VirtualCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualCircuitParameters `json:"forProvider"`
}

// VirtualCircuitStatus defines the observed state of VirtualCircuit.
type VirtualCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualCircuit is the Schema for the VirtualCircuits API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type VirtualCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualCircuitSpec   `json:"spec"`
	Status            VirtualCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualCircuitList contains a list of VirtualCircuits
type VirtualCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCircuit `json:"items"`
}

// Repository type metadata.
var (
	VirtualCircuit_Kind             = "VirtualCircuit"
	VirtualCircuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualCircuit_Kind}.String()
	VirtualCircuit_KindAPIVersion   = VirtualCircuit_Kind + "." + CRDGroupVersion.String()
	VirtualCircuit_GroupVersionKind = CRDGroupVersion.WithKind(VirtualCircuit_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualCircuit{}, &VirtualCircuitList{})
}
