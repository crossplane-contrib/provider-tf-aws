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
// +groupName=alb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/alb/v1alpha1"
)

type AccessLogsObservation struct {
}

type AccessLogsParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type AlbObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ArnSuffix string `json:"arnSuffix" tf:"arn_suffix"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	VpcId string `json:"vpcId" tf:"vpc_id"`

	ZoneId string `json:"zoneId" tf:"zone_id"`
}

type AlbParameters struct {
	AccessLogs []AccessLogsParameters `json:"accessLogs,omitempty" tf:"access_logs"`

	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`

	DropInvalidHeaderFields *bool `json:"dropInvalidHeaderFields,omitempty" tf:"drop_invalid_header_fields"`

	EnableCrossZoneLoadBalancing *bool `json:"enableCrossZoneLoadBalancing,omitempty" tf:"enable_cross_zone_load_balancing"`

	EnableDeletionProtection *bool `json:"enableDeletionProtection,omitempty" tf:"enable_deletion_protection"`

	EnableHttp2 *bool `json:"enableHttp2,omitempty" tf:"enable_http2"`

	IdleTimeout *int64 `json:"idleTimeout,omitempty" tf:"idle_timeout"`

	Internal *bool `json:"internal,omitempty" tf:"internal"`

	IpAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type"`

	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SubnetMapping []SubnetMappingParameters `json:"subnetMapping,omitempty" tf:"subnet_mapping"`

	Subnets []string `json:"subnets,omitempty" tf:"subnets"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type SubnetMappingObservation struct {
	OutpostId string `json:"outpostId" tf:"outpost_id"`
}

type SubnetMappingParameters struct {
	AllocationId *string `json:"allocationId,omitempty" tf:"allocation_id"`

	Ipv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address"`

	PrivateIpv4Address *string `json:"privateIpv4Address,omitempty" tf:"private_ipv4_address"`

	SubnetId string `json:"subnetId" tf:"subnet_id"`
}

// AlbSpec defines the desired state of Alb
type AlbSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AlbParameters `json:"forProvider"`
}

// AlbStatus defines the observed state of Alb.
type AlbStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AlbObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alb is the Schema for the Albs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Alb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbSpec   `json:"spec"`
	Status            AlbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbList contains a list of Albs
type AlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alb `json:"items"`
}

// Repository type metadata.
var (
	AlbKind             = "Alb"
	AlbGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AlbKind}.String()
	AlbKindAPIVersion   = AlbKind + "." + v1alpha1.GroupVersion.String()
	AlbGroupVersionKind = v1alpha1.GroupVersion.WithKind(AlbKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Alb{}, &AlbList{})
}
