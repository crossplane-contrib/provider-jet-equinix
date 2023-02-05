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

type BGPSessionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the session - up or down
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BGPSessionParameters struct {

	// ipv4 or ipv6
	// +kubebuilder:validation:Required
	AddressFamily *string `json:"addressFamily" tf:"address_family,omitempty"`

	// Boolean flag to set the default route policy. False by default
	// +kubebuilder:validation:Optional
	DefaultRoute *bool `json:"defaultRoute,omitempty" tf:"default_route,omitempty"`

	// ID of device
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

// BGPSessionSpec defines the desired state of BGPSession
type BGPSessionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BGPSessionParameters `json:"forProvider"`
}

// BGPSessionStatus defines the observed state of BGPSession.
type BGPSessionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BGPSessionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BGPSession is the Schema for the BGPSessions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type BGPSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BGPSessionSpec   `json:"spec"`
	Status            BGPSessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BGPSessionList contains a list of BGPSessions
type BGPSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPSession `json:"items"`
}

// Repository type metadata.
var (
	BGPSession_Kind             = "BGPSession"
	BGPSession_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BGPSession_Kind}.String()
	BGPSession_KindAPIVersion   = BGPSession_Kind + "." + CRDGroupVersion.String()
	BGPSession_GroupVersionKind = CRDGroupVersion.WithKind(BGPSession_Kind)
)

func init() {
	SchemeBuilder.Register(&BGPSession{}, &BGPSessionList{})
}
