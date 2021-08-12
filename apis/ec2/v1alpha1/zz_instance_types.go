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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)



type CapacityReservationSpecificationParameters struct {


CapacityReservationPreference string `json:"capacityReservationPreference" tf:"capacity_reservation_preference"`

CapacityReservationTarget []CapacityReservationTargetParameters `json:"capacityReservationTarget" tf:"capacity_reservation_target"`
}

type CapacityReservationTargetParameters struct {


CapacityReservationId string `json:"capacityReservationId" tf:"capacity_reservation_id"`
}

type CreditSpecificationParameters struct {


CpuCredits string `json:"cpuCredits" tf:"cpu_credits"`
}

type EbsBlockDeviceObservation struct {


VolumeId string `json:"volumeId" tf:"volume_id"`
}

type EbsBlockDeviceParameters struct {


DeleteOnTermination bool `json:"deleteOnTermination" tf:"delete_on_termination"`

DeviceName string `json:"deviceName" tf:"device_name"`

Encrypted bool `json:"encrypted" tf:"encrypted"`

Iops int64 `json:"iops" tf:"iops"`

KmsKeyId string `json:"kmsKeyId" tf:"kms_key_id"`

SnapshotId string `json:"snapshotId" tf:"snapshot_id"`

Tags map[string]string `json:"tags" tf:"tags"`

Throughput int64 `json:"throughput" tf:"throughput"`

VolumeSize int64 `json:"volumeSize" tf:"volume_size"`

VolumeType string `json:"volumeType" tf:"volume_type"`
}

type EnclaveOptionsParameters struct {


Enabled bool `json:"enabled" tf:"enabled"`
}

type EphemeralBlockDeviceParameters struct {


DeviceName string `json:"deviceName" tf:"device_name"`

NoDevice bool `json:"noDevice" tf:"no_device"`

VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type InstanceObservation struct {


Arn string `json:"arn" tf:"arn"`

InstanceState string `json:"instanceState" tf:"instance_state"`

OutpostArn string `json:"outpostArn" tf:"outpost_arn"`

PasswordData string `json:"passwordData" tf:"password_data"`

PrimaryNetworkInterfaceId string `json:"primaryNetworkInterfaceId" tf:"primary_network_interface_id"`

PrivateDns string `json:"privateDns" tf:"private_dns"`

PublicDns string `json:"publicDns" tf:"public_dns"`

PublicIp string `json:"publicIp" tf:"public_ip"`
}

type InstanceParameters struct {


Ami string `json:"ami" tf:"ami"`

AssociatePublicIpAddress bool `json:"associatePublicIpAddress" tf:"associate_public_ip_address"`

AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

CapacityReservationSpecification []CapacityReservationSpecificationParameters `json:"capacityReservationSpecification" tf:"capacity_reservation_specification"`

CpuCoreCount int64 `json:"cpuCoreCount" tf:"cpu_core_count"`

CpuThreadsPerCore int64 `json:"cpuThreadsPerCore" tf:"cpu_threads_per_core"`

CreditSpecification []CreditSpecificationParameters `json:"creditSpecification" tf:"credit_specification"`

DisableApiTermination bool `json:"disableApiTermination" tf:"disable_api_termination"`

EbsBlockDevice []EbsBlockDeviceParameters `json:"ebsBlockDevice" tf:"ebs_block_device"`

EbsOptimized bool `json:"ebsOptimized" tf:"ebs_optimized"`

EnclaveOptions []EnclaveOptionsParameters `json:"enclaveOptions" tf:"enclave_options"`

EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice" tf:"ephemeral_block_device"`

GetPasswordData bool `json:"getPasswordData" tf:"get_password_data"`

Hibernation bool `json:"hibernation" tf:"hibernation"`

HostId string `json:"hostId" tf:"host_id"`

IamInstanceProfile string `json:"iamInstanceProfile" tf:"iam_instance_profile"`

InstanceInitiatedShutdownBehavior string `json:"instanceInitiatedShutdownBehavior" tf:"instance_initiated_shutdown_behavior"`

InstanceType string `json:"instanceType" tf:"instance_type"`

Ipv6AddressCount int64 `json:"ipv6AddressCount" tf:"ipv6_address_count"`

Ipv6Addresses []string `json:"ipv6Addresses" tf:"ipv6_addresses"`

KeyName string `json:"keyName" tf:"key_name"`

LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate" tf:"launch_template"`

MetadataOptions []MetadataOptionsParameters `json:"metadataOptions" tf:"metadata_options"`

Monitoring bool `json:"monitoring" tf:"monitoring"`

NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface"`

PlacementGroup string `json:"placementGroup" tf:"placement_group"`

PrivateIp string `json:"privateIp" tf:"private_ip"`

RootBlockDevice []RootBlockDeviceParameters `json:"rootBlockDevice" tf:"root_block_device"`

SecondaryPrivateIps []string `json:"secondaryPrivateIps" tf:"secondary_private_ips"`

SecurityGroups []string `json:"securityGroups" tf:"security_groups"`

SourceDestCheck bool `json:"sourceDestCheck" tf:"source_dest_check"`

SubnetId string `json:"subnetId" tf:"subnet_id"`

Tags map[string]string `json:"tags" tf:"tags"`

TagsAll map[string]string `json:"tagsAll" tf:"tags_all"`

Tenancy string `json:"tenancy" tf:"tenancy"`

UserData string `json:"userData" tf:"user_data"`

UserDataBase64 string `json:"userDataBase64" tf:"user_data_base64"`

VolumeTags map[string]string `json:"volumeTags" tf:"volume_tags"`

VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds" tf:"vpc_security_group_ids"`
}

type LaunchTemplateParameters struct {


Id string `json:"id" tf:"id"`

Name string `json:"name" tf:"name"`

Version string `json:"version" tf:"version"`
}

type MetadataOptionsParameters struct {


HttpEndpoint string `json:"httpEndpoint" tf:"http_endpoint"`

HttpPutResponseHopLimit int64 `json:"httpPutResponseHopLimit" tf:"http_put_response_hop_limit"`

HttpTokens string `json:"httpTokens" tf:"http_tokens"`
}

type NetworkInterfaceParameters struct {


DeleteOnTermination bool `json:"deleteOnTermination" tf:"delete_on_termination"`

DeviceIndex int64 `json:"deviceIndex" tf:"device_index"`

NetworkInterfaceId string `json:"networkInterfaceId" tf:"network_interface_id"`
}

type RootBlockDeviceObservation struct {


DeviceName string `json:"deviceName" tf:"device_name"`

VolumeId string `json:"volumeId" tf:"volume_id"`
}

type RootBlockDeviceParameters struct {


DeleteOnTermination bool `json:"deleteOnTermination" tf:"delete_on_termination"`

Encrypted bool `json:"encrypted" tf:"encrypted"`

Iops int64 `json:"iops" tf:"iops"`

KmsKeyId string `json:"kmsKeyId" tf:"kms_key_id"`

Tags map[string]string `json:"tags" tf:"tags"`

Throughput int64 `json:"throughput" tf:"throughput"`

VolumeSize int64 `json:"volumeSize" tf:"volume_size"`

VolumeType string `json:"volumeType" tf:"volume_type"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	InstanceKind             = "Instance"
	InstanceGroupKind        = schema.GroupKind{Group: Group, Kind: InstanceKind}.String()
	InstanceKindAPIVersion   = InstanceKind + "." + GroupVersion.String()
	InstanceGroupVersionKind = GroupVersion.WithKind(InstanceKind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
