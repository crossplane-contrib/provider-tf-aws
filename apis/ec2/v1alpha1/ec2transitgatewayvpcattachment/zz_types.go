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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type Ec2TransitGatewayVpcAttachmentObservation struct {
	VpcOwnerId string `json:"vpcOwnerId" tf:"vpc_owner_id"`
}

type Ec2TransitGatewayVpcAttachmentParameters struct {
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support"`

	DnsSupport *string `json:"dnsSupport,omitempty" tf:"dns_support"`

	Ipv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association"`

	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation"`

	TransitGatewayId string `json:"transitGatewayId" tf:"transit_gateway_id"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// Ec2TransitGatewayVpcAttachmentSpec defines the desired state of Ec2TransitGatewayVpcAttachment
type Ec2TransitGatewayVpcAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayVpcAttachmentParameters `json:"forProvider"`
}

// Ec2TransitGatewayVpcAttachmentStatus defines the observed state of Ec2TransitGatewayVpcAttachment.
type Ec2TransitGatewayVpcAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayVpcAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachment is the Schema for the Ec2TransitGatewayVpcAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2TransitGatewayVpcAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayVpcAttachmentSpec   `json:"spec"`
	Status            Ec2TransitGatewayVpcAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachmentList contains a list of Ec2TransitGatewayVpcAttachments
type Ec2TransitGatewayVpcAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayVpcAttachment `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayVpcAttachmentKind             = "Ec2TransitGatewayVpcAttachment"
	Ec2TransitGatewayVpcAttachmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2TransitGatewayVpcAttachmentKind}.String()
	Ec2TransitGatewayVpcAttachmentKindAPIVersion   = Ec2TransitGatewayVpcAttachmentKind + "." + v1alpha1.GroupVersion.String()
	Ec2TransitGatewayVpcAttachmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2TransitGatewayVpcAttachmentKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2TransitGatewayVpcAttachment{}, &Ec2TransitGatewayVpcAttachmentList{})
}
