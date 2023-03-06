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

type FileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// File upload status.
	// File upload status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Unique identifier of file resource.
	// Unique identifier of file resource
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type FileParameters struct {

	// Boolean value that determines device licensing mode, i.e.,
	// bring your own license or subscription.
	// Boolean value that determines device licensing mode: bring your own license or subscription
	// +kubebuilder:validation:Required
	Byol *bool `json:"byol" tf:"byol,omitempty"`

	// Uploaded file content, expected to be a UTF-8 encoded string.
	// Uploaded file content, expected to be a UTF-8 encoded string
	// +kubebuilder:validation:Required
	ContentSecretRef v1.SecretKeySelector `json:"contentSecretRef" tf:"-"`

	// Device type code.
	// Device type code
	// +kubebuilder:validation:Required
	DeviceTypeCode *string `json:"deviceTypeCode" tf:"device_type_code,omitempty"`

	// File name.
	// File name
	// +kubebuilder:validation:Required
	FileName *string `json:"fileName" tf:"file_name,omitempty"`

	// File upload location metro code. It should match the device location metro code.
	// File upload location metro code
	// +kubebuilder:validation:Required
	MetroCode *string `json:"metroCode" tf:"metro_code,omitempty"`

	// File process type (LICENSE or CLOUD_INIT).
	// File process type (LICENSE or CLOUD_INIT)
	// +kubebuilder:validation:Required
	ProcessType *string `json:"processType" tf:"process_type,omitempty"`

	// Boolean value that determines device management mode, i.e.,
	// self-managed or Equinix-managed.
	// Boolean value that determines device management mode: self-managed or equinix-managed
	// +kubebuilder:validation:Required
	SelfManaged *bool `json:"selfManaged" tf:"self_managed,omitempty"`
}

// FileSpec defines the desired state of File
type FileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileParameters `json:"forProvider"`
}

// FileStatus defines the observed state of File.
type FileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// File is the Schema for the Files API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSpec   `json:"spec"`
	Status            FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of Files
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// Repository type metadata.
var (
	File_Kind             = "File"
	File_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: File_Kind}.String()
	File_KindAPIVersion   = File_Kind + "." + CRDGroupVersion.String()
	File_GroupVersionKind = CRDGroupVersion.WithKind(File_Kind)
)

func init() {
	SchemeBuilder.Register(&File{}, &FileList{})
}
