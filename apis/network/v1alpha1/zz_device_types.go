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

type ClusterDetailsObservation struct {

	// The id of the cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// An object that has node0 details
	// +kubebuilder:validation:Required
	Node0 []Node0Observation `json:"node0,omitempty" tf:"node0,omitempty"`

	// An object that has node1 details
	// +kubebuilder:validation:Required
	Node1 []Node1Observation `json:"node1,omitempty" tf:"node1,omitempty"`

	// The number of nodes in the cluster
	NumOfNodes *float64 `json:"numOfNodes,omitempty" tf:"num_of_nodes,omitempty"`
}

type ClusterDetailsParameters struct {

	// The name of the cluster device
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// An object that has node0 details
	// +kubebuilder:validation:Required
	Node0 []Node0Parameters `json:"node0" tf:"node0,omitempty"`

	// An object that has node1 details
	// +kubebuilder:validation:Required
	Node1 []Node1Parameters `json:"node1" tf:"node1,omitempty"`
}

type DeviceObservation struct {

	// Autonomous system number
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// An object that has the cluster details
	// +kubebuilder:validation:Optional
	ClusterDetails []ClusterDetailsObservation `json:"clusterDetails,omitempty" tf:"cluster_details,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Device location Equinix Business Exchange name
	Ibx *string `json:"ibx,omitempty" tf:"ibx,omitempty"`

	// List of device interfaces
	Interface []InterfaceObservation `json:"interface,omitempty" tf:"interface,omitempty"`

	// Unique identifier of applied license file
	LicenseFileID *string `json:"licenseFileId,omitempty" tf:"license_file_id,omitempty"`

	// Device license registration status
	LicenseStatus *string `json:"licenseStatus,omitempty" tf:"license_status,omitempty"`

	// Device redundancy type applicable for HA devices, either primary or secondary
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type,omitempty"`

	// Unique identifier for a redundant device, applicable for HA device
	RedundantID *string `json:"redundantId,omitempty" tf:"redundant_id,omitempty"`

	// Device location region
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// IP address of SSH enabled interface on the device
	SSHIPAddress *string `json:"sshIpAddress,omitempty" tf:"ssh_ip_address,omitempty"`

	// FQDN of SSH enabled interface on the device
	SSHIPFqdn *string `json:"sshIpFqdn,omitempty" tf:"ssh_ip_fqdn,omitempty"`

	// Definition of secondary device applicable for HA setup
	// +kubebuilder:validation:Optional
	SecondaryDevice []SecondaryDeviceObservation `json:"secondaryDevice,omitempty" tf:"secondary_device,omitempty"`

	// Device provisioning status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Device unique identifier
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Device location zone code
	ZoneCode *string `json:"zoneCode,omitempty" tf:"zone_code,omitempty"`
}

type DeviceParameters struct {

	// Unique identifier of applied ACL template
	// +kubebuilder:validation:Optional
	ACLTemplateID *string `json:"aclTemplateId,omitempty" tf:"acl_template_id,omitempty"`

	// Device billing account number
	// +kubebuilder:validation:Required
	AccountNumber *string `json:"accountNumber" tf:"account_number,omitempty"`

	// Additional Internet bandwidth, in Mbps, that will be allocated to the device
	// +kubebuilder:validation:Optional
	AdditionalBandwidth *float64 `json:"additionalBandwidth,omitempty" tf:"additional_bandwidth,omitempty"`

	// Boolean value that determines device licensing mode: bring your own license or subscription (default)
	// +kubebuilder:validation:Optional
	Byol *bool `json:"byol,omitempty" tf:"byol,omitempty"`

	// An object that has the cluster details
	// +kubebuilder:validation:Optional
	ClusterDetails []ClusterDetailsParameters `json:"clusterDetails,omitempty" tf:"cluster_details,omitempty"`

	// Number of CPU cores used by device
	// +kubebuilder:validation:Required
	CoreCount *float64 `json:"coreCount" tf:"core_count,omitempty"`

	// Device hostname prefix
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Number of network interfaces on a device. If not specified, default number for a given device type will be used
	// +kubebuilder:validation:Optional
	InterfaceCount *float64 `json:"interfaceCount,omitempty" tf:"interface_count,omitempty"`

	// Path to the license file that will be uploaded and applied on a device, applicable for some device types in BYOL licensing mode
	// +kubebuilder:validation:Optional
	LicenseFile *string `json:"licenseFile,omitempty" tf:"license_file,omitempty"`

	// License Token applicable for some device types in BYOL licensing mode
	// +kubebuilder:validation:Optional
	LicenseToken *string `json:"licenseToken,omitempty" tf:"license_token,omitempty"`

	// Device location metro code
	// +kubebuilder:validation:Required
	MetroCode *string `json:"metroCode" tf:"metro_code,omitempty"`

	// Unique identifier of applied MGMT ACL template
	// +kubebuilder:validation:Optional
	MgmtACLTemplateUUID *string `json:"mgmtAclTemplateUuid,omitempty" tf:"mgmt_acl_template_uuid,omitempty"`

	// Device name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// List of email addresses that will receive device status notifications
	// +kubebuilder:validation:Required
	Notifications []*string `json:"notifications" tf:"notifications,omitempty"`

	// Name/number used to identify device order on the invoice
	// +kubebuilder:validation:Optional
	OrderReference *string `json:"orderReference,omitempty" tf:"order_reference,omitempty"`

	// Device software package code
	// +kubebuilder:validation:Required
	PackageCode *string `json:"packageCode" tf:"package_code,omitempty"`

	// Purchase order number associated with a device order
	// +kubebuilder:validation:Optional
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty" tf:"purchase_order_number,omitempty"`

	// Definition of SSH key that will be provisioned on a device
	// +kubebuilder:validation:Optional
	SSHKey []DeviceSSHKeyParameters `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`

	// Definition of secondary device applicable for HA setup
	// +kubebuilder:validation:Optional
	SecondaryDevice []SecondaryDeviceParameters `json:"secondaryDevice,omitempty" tf:"secondary_device,omitempty"`

	// Boolean value that determines device management mode: self-managed or subscription (default)
	// +kubebuilder:validation:Optional
	SelfManaged *bool `json:"selfManaged,omitempty" tf:"self_managed,omitempty"`

	// Device term length
	// +kubebuilder:validation:Required
	TermLength *float64 `json:"termLength" tf:"term_length,omitempty"`

	// Device license throughput
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Device license throughput unit (Mbps or Gbps)
	// +kubebuilder:validation:Optional
	ThroughputUnit *string `json:"throughputUnit,omitempty" tf:"throughput_unit,omitempty"`

	// Device type code
	// +kubebuilder:validation:Required
	TypeCode *string `json:"typeCode" tf:"type_code,omitempty"`

	// Map of vendor specific configuration parameters for a device (controller1, activationKey, managementType, siteId, systemIpAddress)
	// +kubebuilder:validation:Optional
	VendorConfiguration map[string]*string `json:"vendorConfiguration,omitempty" tf:"vendor_configuration,omitempty"`

	// Device software software version
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`

	// device interface id picked for WAN
	// +kubebuilder:validation:Optional
	WanInterfaceID *string `json:"wanInterfaceId,omitempty" tf:"wan_interface_id,omitempty"`
}

type DeviceSSHKeyObservation struct {
}

type DeviceSSHKeyParameters struct {

	// Reference by name to previously provisioned public SSH key
	// +kubebuilder:validation:Required
	KeyName *string `json:"keyName" tf:"key_name,omitempty"`

	// Username associated with given key
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type InterfaceObservation struct {
	AssignedType *string `json:"assignedType,omitempty" tf:"assigned_type,omitempty"`

	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OperationalStatus *string `json:"operationalStatus,omitempty" tf:"operational_status,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InterfaceParameters struct {
}

type Node0Observation struct {

	// The name of the node
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The unique id of the node
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type Node0Parameters struct {

	// License file id. This is necessary for Fortinet and Juniper clusters
	// +kubebuilder:validation:Optional
	LicenseFileIDSecretRef *v1.SecretKeySelector `json:"licenseFileIdSecretRef,omitempty" tf:"-"`

	// License token. This is necessary for Palo Alto clusters
	// +kubebuilder:validation:Optional
	LicenseTokenSecretRef *v1.SecretKeySelector `json:"licenseTokenSecretRef,omitempty" tf:"-"`

	// An object that has fields relevant to the vendor of the cluster device
	// +kubebuilder:validation:Optional
	VendorConfiguration []VendorConfigurationParameters `json:"vendorConfiguration,omitempty" tf:"vendor_configuration,omitempty"`
}

type Node1Observation struct {

	// The name of the node
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The unique id of the node
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type Node1Parameters struct {

	// License file id. This is necessary for Fortinet and Juniper clusters
	// +kubebuilder:validation:Optional
	LicenseFileIDSecretRef *v1.SecretKeySelector `json:"licenseFileIdSecretRef,omitempty" tf:"-"`

	// License token. This is necessary for Palo Alto clusters
	// +kubebuilder:validation:Optional
	LicenseTokenSecretRef *v1.SecretKeySelector `json:"licenseTokenSecretRef,omitempty" tf:"-"`

	// An object that has fields relevant to the vendor of the cluster device
	// +kubebuilder:validation:Optional
	VendorConfiguration []Node1VendorConfigurationParameters `json:"vendorConfiguration,omitempty" tf:"vendor_configuration,omitempty"`
}

type Node1VendorConfigurationObservation struct {
}

type Node1VendorConfigurationParameters struct {

	// Activation key. This is required for Velocloud clusters
	// +kubebuilder:validation:Optional
	ActivationKeySecretRef *v1.SecretKeySelector `json:"activationKeySecretRef,omitempty" tf:"-"`

	// The administrative password of the device. You can use it to log in to the console. This field is not available for all device types
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// System IP Address. Mandatory for the Fortinet SDWAN cluster device
	// +kubebuilder:validation:Optional
	Controller1 *string `json:"controller1,omitempty" tf:"controller1,omitempty"`

	// Controller fqdn. This is required for Velocloud clusters
	// +kubebuilder:validation:Optional
	ControllerFqdn *string `json:"controllerFqdn,omitempty" tf:"controller_fqdn,omitempty"`

	// Hostname. This is necessary for Palo Alto, Juniper, and Fortinet clusters
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The CLI password of the device. This field is relevant only for the Velocloud SDWAN cluster
	// +kubebuilder:validation:Optional
	RootPasswordSecretRef *v1.SecretKeySelector `json:"rootPasswordSecretRef,omitempty" tf:"-"`
}

type SSHKeyObservation struct {
}

type SSHKeyParameters struct {

	// Reference by name to previously provisioned public SSH key
	// +kubebuilder:validation:Required
	KeyName *string `json:"keyName" tf:"key_name,omitempty"`

	// Username associated with given key
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type SecondaryDeviceInterfaceObservation struct {
	AssignedType *string `json:"assignedType,omitempty" tf:"assigned_type,omitempty"`

	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OperationalStatus *string `json:"operationalStatus,omitempty" tf:"operational_status,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SecondaryDeviceInterfaceParameters struct {
}

type SecondaryDeviceObservation struct {

	// Autonomous system number
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// Device location Equinix Business Exchange name
	Ibx *string `json:"ibx,omitempty" tf:"ibx,omitempty"`

	// List of device interfaces
	Interface []SecondaryDeviceInterfaceObservation `json:"interface,omitempty" tf:"interface,omitempty"`

	// Unique identifier of applied license file
	LicenseFileID *string `json:"licenseFileId,omitempty" tf:"license_file_id,omitempty"`

	// Device license registration status
	LicenseStatus *string `json:"licenseStatus,omitempty" tf:"license_status,omitempty"`

	// Device redundancy type applicable for HA devices, either primary or secondary
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type,omitempty"`

	// Unique identifier for a redundant device, applicable for HA device
	RedundantID *string `json:"redundantId,omitempty" tf:"redundant_id,omitempty"`

	// Device location region
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// IP address of SSH enabled interface on the device
	SSHIPAddress *string `json:"sshIpAddress,omitempty" tf:"ssh_ip_address,omitempty"`

	// FQDN of SSH enabled interface on the device
	SSHIPFqdn *string `json:"sshIpFqdn,omitempty" tf:"ssh_ip_fqdn,omitempty"`

	// Device provisioning status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Device unique identifier
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Device location zone code
	ZoneCode *string `json:"zoneCode,omitempty" tf:"zone_code,omitempty"`
}

type SecondaryDeviceParameters struct {

	// Unique identifier of applied ACL template
	// +kubebuilder:validation:Optional
	ACLTemplateID *string `json:"aclTemplateId,omitempty" tf:"acl_template_id,omitempty"`

	// Device billing account number
	// +kubebuilder:validation:Required
	AccountNumber *string `json:"accountNumber" tf:"account_number,omitempty"`

	// Additional Internet bandwidth, in Mbps, that will be allocated to the device
	// +kubebuilder:validation:Optional
	AdditionalBandwidth *float64 `json:"additionalBandwidth,omitempty" tf:"additional_bandwidth,omitempty"`

	// Device hostname prefix
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Path to the license file that will be uploaded and applied on a device, applicable for some device types in BYOL licensing mode
	// +kubebuilder:validation:Optional
	LicenseFile *string `json:"licenseFile,omitempty" tf:"license_file,omitempty"`

	// License Token applicable for some device types in BYOL licensing mode
	// +kubebuilder:validation:Optional
	LicenseToken *string `json:"licenseToken,omitempty" tf:"license_token,omitempty"`

	// Device location metro code
	// +kubebuilder:validation:Required
	MetroCode *string `json:"metroCode" tf:"metro_code,omitempty"`

	// Unique identifier of applied MGMT ACL template
	// +kubebuilder:validation:Optional
	MgmtACLTemplateUUID *string `json:"mgmtAclTemplateUuid,omitempty" tf:"mgmt_acl_template_uuid,omitempty"`

	// Device name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// List of email addresses that will receive device status notifications
	// +kubebuilder:validation:Required
	Notifications []*string `json:"notifications" tf:"notifications,omitempty"`

	// Definition of SSH key that will be provisioned on a device
	// +kubebuilder:validation:Optional
	SSHKey []SSHKeyParameters `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`

	// Map of vendor specific configuration parameters for a device (controller1, activationKey, managementType, siteId, systemIpAddress)
	// +kubebuilder:validation:Optional
	VendorConfiguration map[string]*string `json:"vendorConfiguration,omitempty" tf:"vendor_configuration,omitempty"`

	// device interface id picked for WAN
	// +kubebuilder:validation:Optional
	WanInterfaceID *string `json:"wanInterfaceId,omitempty" tf:"wan_interface_id,omitempty"`
}

type VendorConfigurationObservation struct {
}

type VendorConfigurationParameters struct {

	// Activation key. This is required for Velocloud clusters
	// +kubebuilder:validation:Optional
	ActivationKeySecretRef *v1.SecretKeySelector `json:"activationKeySecretRef,omitempty" tf:"-"`

	// The administrative password of the device. You can use it to log in to the console. This field is not available for all device types
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// System IP Address. Mandatory for the Fortinet SDWAN cluster device
	// +kubebuilder:validation:Optional
	Controller1 *string `json:"controller1,omitempty" tf:"controller1,omitempty"`

	// Controller fqdn. This is required for Velocloud clusters
	// +kubebuilder:validation:Optional
	ControllerFqdn *string `json:"controllerFqdn,omitempty" tf:"controller_fqdn,omitempty"`

	// Hostname. This is necessary for Palo Alto, Juniper, and Fortinet clusters
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The CLI password of the device. This field is relevant only for the Velocloud SDWAN cluster
	// +kubebuilder:validation:Optional
	RootPasswordSecretRef *v1.SecretKeySelector `json:"rootPasswordSecretRef,omitempty" tf:"-"`
}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceParameters `json:"forProvider"`
}

// DeviceStatus defines the observed state of Device.
type DeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Device is the Schema for the Devices API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSpec   `json:"spec"`
	Status            DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}

// Repository type metadata.
var (
	Device_Kind             = "Device"
	Device_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Device_Kind}.String()
	Device_KindAPIVersion   = Device_Kind + "." + CRDGroupVersion.String()
	Device_GroupVersionKind = CRDGroupVersion.WithKind(Device_Kind)
)

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}
