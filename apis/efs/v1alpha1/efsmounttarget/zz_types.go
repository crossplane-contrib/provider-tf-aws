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

// +kubebuilder:object:generate=true
// +groupName=efs.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/efs/v1alpha1"
)

type EfsMountTargetObservation struct {
	AvailabilityZoneId string `json:"availabilityZoneId" tf:"availability_zone_id"`

	AvailabilityZoneName string `json:"availabilityZoneName" tf:"availability_zone_name"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	FileSystemArn string `json:"fileSystemArn" tf:"file_system_arn"`

	MountTargetDnsName string `json:"mountTargetDnsName" tf:"mount_target_dns_name"`

	NetworkInterfaceId string `json:"networkInterfaceId" tf:"network_interface_id"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type EfsMountTargetParameters struct {
	FileSystemId string `json:"fileSystemId" tf:"file_system_id"`

	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// EfsMountTargetSpec defines the desired state of EfsMountTarget
type EfsMountTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EfsMountTargetParameters `json:"forProvider"`
}

// EfsMountTargetStatus defines the observed state of EfsMountTarget.
type EfsMountTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EfsMountTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EfsMountTarget is the Schema for the EfsMountTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EfsMountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsMountTargetSpec   `json:"spec"`
	Status            EfsMountTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsMountTargetList contains a list of EfsMountTargets
type EfsMountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsMountTarget `json:"items"`
}

// Repository type metadata.
var (
	EfsMountTargetKind             = "EfsMountTarget"
	EfsMountTargetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EfsMountTargetKind}.String()
	EfsMountTargetKindAPIVersion   = EfsMountTargetKind + "." + v1alpha1.GroupVersion.String()
	EfsMountTargetGroupVersionKind = v1alpha1.GroupVersion.WithKind(EfsMountTargetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EfsMountTarget{}, &EfsMountTargetList{})
}
