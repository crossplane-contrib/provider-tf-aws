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
// +groupName=dx.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dx/v1alpha1"
)

type DxTransitVirtualInterfaceObservation struct {
	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`

	Arn string `json:"arn" tf:"arn"`

	AwsDevice string `json:"awsDevice" tf:"aws_device"`

	JumboFrameCapable bool `json:"jumboFrameCapable" tf:"jumbo_frame_capable"`
}

type DxTransitVirtualInterfaceParameters struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`

	BgpAsn int64 `json:"bgpAsn" tf:"bgp_asn"`

	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`

	ConnectionId string `json:"connectionId" tf:"connection_id"`

	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`

	DxGatewayId string `json:"dxGatewayId" tf:"dx_gateway_id"`

	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Vlan int64 `json:"vlan" tf:"vlan"`
}

// DxTransitVirtualInterfaceSpec defines the desired state of DxTransitVirtualInterface
type DxTransitVirtualInterfaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxTransitVirtualInterfaceParameters `json:"forProvider"`
}

// DxTransitVirtualInterfaceStatus defines the observed state of DxTransitVirtualInterface.
type DxTransitVirtualInterfaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxTransitVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxTransitVirtualInterface is the Schema for the DxTransitVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxTransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxTransitVirtualInterfaceSpec   `json:"spec"`
	Status            DxTransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxTransitVirtualInterfaceList contains a list of DxTransitVirtualInterfaces
type DxTransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxTransitVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	DxTransitVirtualInterfaceKind             = "DxTransitVirtualInterface"
	DxTransitVirtualInterfaceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxTransitVirtualInterfaceKind}.String()
	DxTransitVirtualInterfaceKindAPIVersion   = DxTransitVirtualInterfaceKind + "." + v1alpha1.GroupVersion.String()
	DxTransitVirtualInterfaceGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxTransitVirtualInterfaceKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxTransitVirtualInterface{}, &DxTransitVirtualInterfaceList{})
}
