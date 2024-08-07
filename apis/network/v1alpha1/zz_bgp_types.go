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

type BGPInitParameters struct {

	// shared key used for BGP peer authentication.
	// Shared key used for BGP peer authentication
	AuthenticationKeySecretRef *v1.SecretKeySelector `json:"authenticationKeySecretRef,omitempty" tf:"-"`

	// identifier of a connection established between. network device and remote service provider that will be used for peering.
	// Identifier of a connection established between network device and remote service provider that will be used for peering
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Local ASN number.
	// Local ASN number
	LocalAsn *float64 `json:"localAsn,omitempty" tf:"local_asn,omitempty"`

	// IP address in CIDR format of a local device.
	// IP address in CIDR format of a local device
	LocalIPAddress *string `json:"localIpAddress,omitempty" tf:"local_ip_address,omitempty"`

	// Remote ASN number.
	// Remote ASN number
	RemoteAsn *float64 `json:"remoteAsn,omitempty" tf:"remote_asn,omitempty"`

	// IP address of remote peer.
	// IP address of remote peer
	RemoteIPAddress *string `json:"remoteIpAddress,omitempty" tf:"remote_ip_address,omitempty"`
}

type BGPObservation struct {

	// identifier of a connection established between. network device and remote service provider that will be used for peering.
	// Identifier of a connection established between network device and remote service provider that will be used for peering
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// unique identifier of a network device that is a local peer in a given BGP peering configuration.
	// Unique identifier of a network device that is a local peer in a given BGP peering configuration
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Local ASN number.
	// Local ASN number
	LocalAsn *float64 `json:"localAsn,omitempty" tf:"local_asn,omitempty"`

	// IP address in CIDR format of a local device.
	// IP address in CIDR format of a local device
	LocalIPAddress *string `json:"localIpAddress,omitempty" tf:"local_ip_address,omitempty"`

	// BGP peering configuration provisioning status, one of PROVISIONING, PENDING_UPDATE, PROVISIONED, FAILED.
	// BGP peering configuration provisioning status
	ProvisioningStatus *string `json:"provisioningStatus,omitempty" tf:"provisioning_status,omitempty"`

	// Remote ASN number.
	// Remote ASN number
	RemoteAsn *float64 `json:"remoteAsn,omitempty" tf:"remote_asn,omitempty"`

	// IP address of remote peer.
	// IP address of remote peer
	RemoteIPAddress *string `json:"remoteIpAddress,omitempty" tf:"remote_ip_address,omitempty"`

	// BGP peer state, one of Idle, Connect, Active, OpenSent, OpenConfirm, Established.
	// BGP peer state
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// BGP peering configuration unique identifier.
	// BGP peering configuration unique identifier
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type BGPParameters struct {

	// shared key used for BGP peer authentication.
	// Shared key used for BGP peer authentication
	// +kubebuilder:validation:Optional
	AuthenticationKeySecretRef *v1.SecretKeySelector `json:"authenticationKeySecretRef,omitempty" tf:"-"`

	// identifier of a connection established between. network device and remote service provider that will be used for peering.
	// Identifier of a connection established between network device and remote service provider that will be used for peering
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Local ASN number.
	// Local ASN number
	// +kubebuilder:validation:Optional
	LocalAsn *float64 `json:"localAsn,omitempty" tf:"local_asn,omitempty"`

	// IP address in CIDR format of a local device.
	// IP address in CIDR format of a local device
	// +kubebuilder:validation:Optional
	LocalIPAddress *string `json:"localIpAddress,omitempty" tf:"local_ip_address,omitempty"`

	// Remote ASN number.
	// Remote ASN number
	// +kubebuilder:validation:Optional
	RemoteAsn *float64 `json:"remoteAsn,omitempty" tf:"remote_asn,omitempty"`

	// IP address of remote peer.
	// IP address of remote peer
	// +kubebuilder:validation:Optional
	RemoteIPAddress *string `json:"remoteIpAddress,omitempty" tf:"remote_ip_address,omitempty"`
}

// BGPSpec defines the desired state of BGP
type BGPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BGPParameters `json:"forProvider"`
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
	InitProvider BGPInitParameters `json:"initProvider,omitempty"`
}

// BGPStatus defines the observed state of BGP.
type BGPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BGPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BGP is the Schema for the BGPs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type BGP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionId) || (has(self.initProvider) && has(self.initProvider.connectionId))",message="spec.forProvider.connectionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.localAsn) || (has(self.initProvider) && has(self.initProvider.localAsn))",message="spec.forProvider.localAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.localIpAddress) || (has(self.initProvider) && has(self.initProvider.localIpAddress))",message="spec.forProvider.localIpAddress is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.remoteAsn) || (has(self.initProvider) && has(self.initProvider.remoteAsn))",message="spec.forProvider.remoteAsn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.remoteIpAddress) || (has(self.initProvider) && has(self.initProvider.remoteIpAddress))",message="spec.forProvider.remoteIpAddress is a required parameter"
	Spec   BGPSpec   `json:"spec"`
	Status BGPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BGPList contains a list of BGPs
type BGPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGP `json:"items"`
}

// Repository type metadata.
var (
	BGP_Kind             = "BGP"
	BGP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BGP_Kind}.String()
	BGP_KindAPIVersion   = BGP_Kind + "." + CRDGroupVersion.String()
	BGP_GroupVersionKind = CRDGroupVersion.WithKind(BGP_Kind)
)

func init() {
	SchemeBuilder.Register(&BGP{}, &BGPList{})
}
