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

type BehaviorObservation struct {
}

type BehaviorParameters struct {

	// List of attributes that are allowed to change without recreating the instance. Supported attributes: custom_data, user_data"
	// List of attributes that are allowed to change without recreating the instance. Supported attributes: `custom_data`, `user_data`
	// +kubebuilder:validation:Optional
	AllowChanges []*string `json:"allowChanges,omitempty" tf:"allow_changes,omitempty"`
}

type DeviceObservation struct {

	// The ipv4 private IP assigned to the device.
	// The ipv4 private IP assigned to the device
	AccessPrivateIPv4 *string `json:"accessPrivateIpv4,omitempty" tf:"access_private_ipv4,omitempty"`

	// The ipv4 maintenance IP assigned to the device.
	// The ipv4 maintenance IP assigned to the device
	AccessPublicIPv4 *string `json:"accessPublicIpv4,omitempty" tf:"access_public_ipv4,omitempty"`

	// The ipv6 maintenance IP assigned to the device.
	// The ipv6 maintenance IP assigned to the device
	AccessPublicIPv6 *string `json:"accessPublicIpv6,omitempty" tf:"access_public_ipv6,omitempty"`

	// The timestamp for when the device was created.
	// The timestamp for when the device was created
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The facility where the device is deployed.
	// The facility where the device is deployed
	DeployedFacility *string `json:"deployedFacility,omitempty" tf:"deployed_facility,omitempty"`

	// ID of hardware reservation where this device was deployed.
	// It is useful when using the next-available hardware reservation.
	// ID of hardware reservation where this device was deployed. It is useful when using the next-available hardware reservation
	DeployedHardwareReservationID *string `json:"deployedHardwareReservationId,omitempty" tf:"deployed_hardware_reservation_id,omitempty"`

	// The ID of the device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the device is locked.
	// Whether the device is locked
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// The device's private and public IP (v4 and v6) network details. See
	// Network Attribute below for more details.
	// The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 addresses: public ipv4, private ipv4 and ipv6
	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// (Deprecated) Network type of a device, used in
	// Layer 2 networking. Since this
	// attribute is deprecated you should handle Network Type with one of
	// equinix_metal_port,
	// equinix_metal_device_network_type resources or
	// equinix_metal_port datasource.
	// See network_types guide for more info.
	// Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// List of ports assigned to the device. See Ports Attribute below for
	// more details.
	// Ports assigned to the device
	Ports []DevicePortsObservation `json:"ports,omitempty" tf:"ports,omitempty"`

	// List of IDs of SSH keys deployed in the device, can be both user and project SSH keys.
	// List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
	SSHKeyIds []*string `json:"sshKeyIds,omitempty" tf:"ssh_key_ids,omitempty"`

	// The status of the device.
	// The status of the device
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The timestamp for the last time the device was updated.
	// The timestamp for the last time the device was updated
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type DeviceParameters struct {

	// If true, a device with OS custom_ipxe will continue to boot via iPXE
	// on reboots.
	// If true, a device with OS custom_ipxe will
	// +kubebuilder:validation:Optional
	AlwaysPxe *bool `json:"alwaysPxe,omitempty" tf:"always_pxe,omitempty"`

	// Behavioral overrides that change how the resource handles certain attribute updates. See Behavior below for more details.
	// +kubebuilder:validation:Optional
	Behavior []BehaviorParameters `json:"behavior,omitempty" tf:"behavior,omitempty"`

	// monthly or hourly
	// monthly or hourly
	// +kubebuilder:validation:Optional
	BillingCycle *string `json:"billingCycle,omitempty" tf:"billing_cycle,omitempty"`

	// A string of the desired Custom Data for the device.  By default, changing this attribute will cause the provider to destroy and recreate your device.  If reinstall is specified or behavior.allow_changes includes "custom_data", the device will be updated in-place instead of recreated.
	// A string of the desired Custom Data for the device.  By default, changing this attribute will cause the provider to destroy and recreate your device.  If `reinstall` is specified or `behavior.allow_changes` includes `"custom_data"`, the device will be updated in-place instead of recreated.
	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// The device description.
	// Description string for the device
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of facility codes with deployment preferences. Equinix Metal API will go
	// through the list and will deploy your device to first facility with free capacity. List items must
	// be facility codes or any (a wildcard). To find the facility code, visit
	// Facilities API docs, set your API auth
	// token in the top of the page and see JSON from the API response. Conflicts with metro.
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or any (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with metro
	// +kubebuilder:validation:Optional
	Facilities []*string `json:"facilities,omitempty" tf:"facilities,omitempty"`

	// Delete device even if it has volumes attached. Only applies
	// for destroy action.
	// Delete device even if it has volumes attached. Only applies for destroy action
	// +kubebuilder:validation:Optional
	ForceDetachVolumes *bool `json:"forceDetachVolumes,omitempty" tf:"force_detach_volumes,omitempty"`

	// The UUID of the hardware reservation where you want this
	// device deployed, or next-available if you want to pick your next available reservation
	// automatically. Changing this from a reservation UUID to next-available will re-create the device
	// in another reservation. Please be careful when using hardware reservation UUID and next-available
	// together for the same pool of reservations. It might happen that the reservation which Equinix
	// Metal API will pick as next-available is the reservation which you refer with UUID in another
	// equinix_metal_device resource. If that happens, and the equinix_metal_device with the UUID is
	// created later, resource creation will fail because the reservation is already in use (by the
	// resource created with next-available). To workaround this, have the next-available resource
	// explicitly depend_on
	// the resource with hardware reservation UUID, so that the latter is created first. For more details,
	// see issue #176.
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your next available reservation automatically
	// +kubebuilder:validation:Optional
	HardwareReservationID *string `json:"hardwareReservationId,omitempty" tf:"hardware_reservation_id,omitempty"`

	// The device hostname used in deployments taking advantage of Layer3 DHCP
	// or metadata service configuration.
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// A list of IP address types for the device. See
	// IP address below for more details.
	// A list of IP address types for the device (structure is documented below)
	// +kubebuilder:validation:Optional
	IPAddress []IPAddressParameters `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// URL pointing to a hosted iPXE script. More information is in the
	// Custom iPXE doc.
	// URL pointing to a hosted iPXE script. More
	// +kubebuilder:validation:Optional
	IpxeScriptURL *string `json:"ipxeScriptUrl,omitempty" tf:"ipxe_script_url,omitempty"`

	// Metro area for the new device. Conflicts with facilities.
	// Metro area for the new device. Conflicts with facilities
	// +kubebuilder:validation:Optional
	Metro *string `json:"metro,omitempty" tf:"metro,omitempty"`

	// The operating system slug. To find the slug, or visit
	// Operating Systems API docs, set your
	// API auth token in the top of the page and see JSON from the API response.
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.  By default, changing this attribute will cause your device to be deleted and recreated.  If `reinstall` is enabled, the device will be updated in-place instead of recreated.
	// +kubebuilder:validation:Required
	OperatingSystem *string `json:"operatingSystem" tf:"operating_system,omitempty"`

	// The device plan slug. To find the plan slug, visit
	// Device plans API docs, set your auth token in the
	// top of the page and see JSON from the API response.
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response
	// +kubebuilder:validation:Required
	Plan *string `json:"plan" tf:"plan,omitempty"`

	// The ID of the project in which to create the device
	// The ID of the project in which to create the device
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Array of IDs of the project SSH keys which should be added to the device.
	// If you omit this, SSH keys of all the members of the parent project will be added to the device. If
	// you specify this array, only the listed project SSH keys will be added. Project SSH keys can be
	// created with the equinix_metal_project_ssh_key resource.
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys (and any user_ssh_key_ids) will be added. Project SSH keys can be created with the [equinix_metal_project_ssh_key](equinix_metal_project_ssh_key.md) resource
	// +kubebuilder:validation:Optional
	ProjectSSHKeyIds []*string `json:"projectSshKeyIds,omitempty" tf:"project_ssh_key_ids,omitempty"`

	// Whether the device should be reinstalled instead of destroyed when
	// modifying user_data, custom_data, or operating system. See Reinstall below for more
	// details.
	// +kubebuilder:validation:Optional
	Reinstall []ReinstallParameters `json:"reinstall,omitempty" tf:"reinstall,omitempty"`

	// JSON for custom partitioning. Only usable on reserved hardware. More
	// information in in the
	// Custom Partitioning and RAID
	// doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can
	// be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc
	// +kubebuilder:validation:Optional
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// Tags attached to the device.
	// Tags attached to the device
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Timestamp for device termination. For example 2021-09-03T16:32:00+03:00.
	// If you don't supply timezone info, timestamp is assumed to be in UTC.
	// Timestamp for device termination. For example "2021-09-03T16:32:00+03:00". If you don't supply timezone info, timestamp is assumed to be in UTC.
	// +kubebuilder:validation:Optional
	TerminationTime *string `json:"terminationTime,omitempty" tf:"termination_time,omitempty"`

	// A string of the desired User Data for the device.  By default, changing this attribute will cause the provider to destroy and recreate your device.  If reinstall is specified or behavior.allow_changes includes "user_data", the device will be updated in-place instead of recreated.
	// A string of the desired User Data for the device.  By default, changing this attribute will cause the provider to destroy and recreate your device.  If `reinstall` is specified or `behavior.allow_changes` includes `"user_data"`, the device will be updated in-place instead of recreated.
	// +kubebuilder:validation:Optional
	UserDataSecretRef *v1.SecretKeySelector `json:"userDataSecretRef,omitempty" tf:"-"`

	// Array of IDs of the user SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed user SSH keys (and any project_ssh_key_ids) will be added. User SSH keys can be created with the equinix_metal_ssh_key resource
	// Array of IDs of the user SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed user SSH keys (and any project_ssh_key_ids) will be added. User SSH keys can be created with the [equinix_metal_ssh_key](equinix_metal_ssh_key.md) resource
	// +kubebuilder:validation:Optional
	UserSSHKeyIds []*string `json:"userSshKeyIds,omitempty" tf:"user_ssh_key_ids,omitempty"`

	// Only used for devices in reserved hardware. If
	// set, the deletion of this device will block until the hardware reservation is marked provisionable
	// (about 4 minutes in August 2019).
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019)
	// +kubebuilder:validation:Optional
	WaitForReservationDeprovision *bool `json:"waitForReservationDeprovision,omitempty" tf:"wait_for_reservation_deprovision,omitempty"`
}

type DevicePortsObservation struct {

	// Whether this port is part of a bond in bonded network setup.
	Bonded *bool `json:"bonded,omitempty" tf:"bonded,omitempty"`

	// The ID of the device.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// MAC address assigned to the port.
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Name of the port (e.g. eth0, or bond0).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the port (e.g. NetworkPort or NetworkBondPort).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DevicePortsParameters struct {
}

type IPAddressObservation struct {
}

type IPAddressParameters struct {

	// CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
	// CIDR suffix for IP block assigned to this device
	// +kubebuilder:validation:Optional
	Cidr *float64 `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// List of UUIDs of IP block reservations
	// from which the public IPv4 address should be taken.
	// IDs of reservations to pick the blocks from
	// +kubebuilder:validation:Optional
	ReservationIds []*string `json:"reservationIds,omitempty" tf:"reservation_ids,omitempty"`

	// One of private_ipv4, public_ipv4, public_ipv6.
	// one of public_ipv4,private_ipv4,public_ipv6
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkObservation struct {

	// IPv4 or IPv6 address string.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
	Cidr *float64 `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// IP version. One of 4, 6.
	Family *float64 `json:"family,omitempty" tf:"family,omitempty"`

	// Address of router.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Whether the address is routable from the Internet.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`
}

type NetworkParameters struct {
}

type ReinstallObservation struct {
}

type ReinstallParameters struct {

	// Whether the OS disk should be filled with 00h bytes before reinstall.
	// Defaults to false.
	// Whether the OS disk should be filled with `00h` bytes before reinstall
	// +kubebuilder:validation:Optional
	DeprovisionFast *bool `json:"deprovisionFast,omitempty" tf:"deprovision_fast,omitempty"`

	// Whether the provider should favour reinstall over destroy and create. Defaults to
	// false.
	// Whether the device should be reinstalled instead of destroyed
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the non-OS disks should be kept or wiped during reinstall.
	// Defaults to false.
	// Whether the non-OS disks should be kept or wiped during reinstall
	// +kubebuilder:validation:Optional
	PreserveData *bool `json:"preserveData,omitempty" tf:"preserve_data,omitempty"`
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

// Device is the Schema for the Devices API.
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
