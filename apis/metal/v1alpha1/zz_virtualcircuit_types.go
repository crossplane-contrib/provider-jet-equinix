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

type VirtualCircuitInitParameters struct {

	// (String) UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	// UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// (String) The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// (String) The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	// The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	CustomerIPv6 *string `json:"customerIpv6,omitempty" tf:"customer_ipv6,omitempty"`

	// (String) Description of the Virtual Circuit resource
	// Description of the Virtual Circuit resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String, Sensitive) The password that can be set for the VRF BGP peer
	// The password that can be set for the VRF BGP peer
	Md5SecretRef *v1.SecretKeySelector `json:"md5SecretRef,omitempty" tf:"-"`

	// (String) The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// (String) The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	// The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	MetalIPv6 *string `json:"metalIpv6,omitempty" tf:"metal_ipv6,omitempty"`

	// (String) Name of the Virtual Circuit resource
	// Name of the Virtual Circuit resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// to-network VLAN ID
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// (Number) The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// (String) UUID of the Connection Port where the VC is scoped to
	// UUID of the Connection Port where the VC is scoped to
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// (String) UUID of the Project where the VC is scoped to
	// UUID of the Project where the VC is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// (String) A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (String) A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	SubnetIPv6 *string `json:"subnetIpv6,omitempty" tf:"subnet_ipv6,omitempty"`

	// (List of String) Tags attached to the virtual circuit
	// Tags attached to the virtual circuit
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	// UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	VirtualCircuitID *string `json:"virtualCircuitId,omitempty" tf:"virtual_circuit_id,omitempty"`

	// (String) UUID of the VLAN to associate
	// UUID of the VLAN to associate
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (String) UUID of the VRF to associate
	// UUID of the VRF to associate
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type VirtualCircuitObservation struct {

	// (String) UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	// UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// (String) The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// (String) The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	// The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	CustomerIPv6 *string `json:"customerIpv6,omitempty" tf:"customer_ipv6,omitempty"`

	// (String) Description of the Virtual Circuit resource
	// Description of the Virtual Circuit resource
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// (String) The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	// The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	MetalIPv6 *string `json:"metalIpv6,omitempty" tf:"metal_ipv6,omitempty"`

	// (String) Name of the Virtual Circuit resource
	// Name of the Virtual Circuit resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// to-network VLAN ID
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// (Number) Nni VLAN ID parameter, see https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/
	// Nni VLAN ID parameter, see https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/
	NniVnid *float64 `json:"nniVnid,omitempty" tf:"nni_vnid,omitempty"`

	// (Number) The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// (String) UUID of the Connection Port where the VC is scoped to
	// UUID of the Connection Port where the VC is scoped to
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// (String) UUID of the Project where the VC is scoped to
	// UUID of the Project where the VC is scoped to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// (String) Status of the virtual circuit resource
	// Status of the virtual circuit resource
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (String) A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	SubnetIPv6 *string `json:"subnetIpv6,omitempty" tf:"subnet_ipv6,omitempty"`

	// (List of String) Tags attached to the virtual circuit
	// Tags attached to the virtual circuit
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	// UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	VirtualCircuitID *string `json:"virtualCircuitId,omitempty" tf:"virtual_circuit_id,omitempty"`

	// (String) UUID of the VLAN to associate
	// UUID of the VLAN to associate
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (Number) VNID VLAN parameter, see https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/
	// VNID VLAN parameter, see https://deploy.equinix.com/developers/docs/metal/interconnections/introduction/
	Vnid *float64 `json:"vnid,omitempty" tf:"vnid,omitempty"`

	// (String) UUID of the VRF to associate
	// UUID of the VRF to associate
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type VirtualCircuitParameters struct {

	// (String) UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	// UUID of Connection where the VC is scoped to.  Only used for dedicated connections
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// (String) The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	// +kubebuilder:validation:Optional
	CustomerIP *string `json:"customerIp,omitempty" tf:"customer_ip,omitempty"`

	// (String) The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	// The Customer IPv6 address which the CSR switch will peer with. Will default to the other usable IP in the IPv6 subnet.
	// +kubebuilder:validation:Optional
	CustomerIPv6 *string `json:"customerIpv6,omitempty" tf:"customer_ipv6,omitempty"`

	// (String) Description of the Virtual Circuit resource
	// Description of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String, Sensitive) The password that can be set for the VRF BGP peer
	// The password that can be set for the VRF BGP peer
	// +kubebuilder:validation:Optional
	Md5SecretRef *v1.SecretKeySelector `json:"md5SecretRef,omitempty" tf:"-"`

	// (String) The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	// +kubebuilder:validation:Optional
	MetalIP *string `json:"metalIp,omitempty" tf:"metal_ip,omitempty"`

	// (String) The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	// The Metal IPv6 address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the IPv6 subnet.
	// +kubebuilder:validation:Optional
	MetalIPv6 *string `json:"metalIpv6,omitempty" tf:"metal_ipv6,omitempty"`

	// (String) Name of the Virtual Circuit resource
	// Name of the Virtual Circuit resource
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// to-network VLAN ID
	// Equinix Metal network-to-network VLAN ID (optional when the connection has mode=tunnel)
	// +kubebuilder:validation:Optional
	NniVlan *float64 `json:"nniVlan,omitempty" tf:"nni_vlan,omitempty"`

	// (Number) The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
	// +kubebuilder:validation:Optional
	PeerAsn *float64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// (String) UUID of the Connection Port where the VC is scoped to
	// UUID of the Connection Port where the VC is scoped to
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// (String) UUID of the Project where the VC is scoped to
	// UUID of the Project where the VC is scoped to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// Description of the Virtual Circuit speed. This is for information purposes and is computed when the connection type is shared.
	// +kubebuilder:validation:Optional
	Speed *string `json:"speed,omitempty" tf:"speed,omitempty"`

	// (String) A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// (String) A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// A subnet from one of the IPv6 blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /126 or /127.
	// * For a /127 block, it will only have two IP addresses, which will be used for the metal_ip and customer_ip.
	// * For a /126 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	// +kubebuilder:validation:Optional
	SubnetIPv6 *string `json:"subnetIpv6,omitempty" tf:"subnet_ipv6,omitempty"`

	// (List of String) Tags attached to the virtual circuit
	// Tags attached to the virtual circuit
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (String) UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	// UUID of an existing VC to configure. Used in the case of shared interconnections where the VC has already been created.
	// +kubebuilder:validation:Optional
	VirtualCircuitID *string `json:"virtualCircuitId,omitempty" tf:"virtual_circuit_id,omitempty"`

	// (String) UUID of the VLAN to associate
	// UUID of the VLAN to associate
	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (String) UUID of the VRF to associate
	// UUID of the VRF to associate
	// +kubebuilder:validation:Optional
	VrfID *string `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

// VirtualCircuitSpec defines the desired state of VirtualCircuit
type VirtualCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualCircuitParameters `json:"forProvider"`
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
	InitProvider VirtualCircuitInitParameters `json:"initProvider,omitempty"`
}

// VirtualCircuitStatus defines the observed state of VirtualCircuit.
type VirtualCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VirtualCircuit is the Schema for the VirtualCircuits API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type VirtualCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.portId) || (has(self.initProvider) && has(self.initProvider.portId))",message="spec.forProvider.portId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   VirtualCircuitSpec   `json:"spec"`
	Status VirtualCircuitStatus `json:"status,omitempty"`
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
