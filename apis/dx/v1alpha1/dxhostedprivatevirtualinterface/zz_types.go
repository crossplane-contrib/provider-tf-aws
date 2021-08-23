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

type DxHostedPrivateVirtualInterfaceObservation struct {
	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`

	Arn string `json:"arn" tf:"arn"`

	AwsDevice string `json:"awsDevice" tf:"aws_device"`

	JumboFrameCapable bool `json:"jumboFrameCapable" tf:"jumbo_frame_capable"`
}

type DxHostedPrivateVirtualInterfaceParameters struct {
	AddressFamily string `json:"addressFamily" tf:"address_family"`

	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address"`

	BgpAsn int64 `json:"bgpAsn" tf:"bgp_asn"`

	BgpAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key"`

	ConnectionId string `json:"connectionId" tf:"connection_id"`

	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address"`

	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`

	Name string `json:"name" tf:"name"`

	OwnerAccountId string `json:"ownerAccountId" tf:"owner_account_id"`

	Vlan int64 `json:"vlan" tf:"vlan"`
}

// DxHostedPrivateVirtualInterfaceSpec defines the desired state of DxHostedPrivateVirtualInterface
type DxHostedPrivateVirtualInterfaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxHostedPrivateVirtualInterfaceParameters `json:"forProvider"`
}

// DxHostedPrivateVirtualInterfaceStatus defines the observed state of DxHostedPrivateVirtualInterface.
type DxHostedPrivateVirtualInterfaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxHostedPrivateVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterface is the Schema for the DxHostedPrivateVirtualInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedPrivateVirtualInterfaceSpec   `json:"spec"`
	Status            DxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterfaceList contains a list of DxHostedPrivateVirtualInterfaces
type DxHostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPrivateVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	DxHostedPrivateVirtualInterfaceKind             = "DxHostedPrivateVirtualInterface"
	DxHostedPrivateVirtualInterfaceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxHostedPrivateVirtualInterfaceKind}.String()
	DxHostedPrivateVirtualInterfaceKindAPIVersion   = DxHostedPrivateVirtualInterfaceKind + "." + v1alpha1.GroupVersion.String()
	DxHostedPrivateVirtualInterfaceGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxHostedPrivateVirtualInterfaceKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxHostedPrivateVirtualInterface{}, &DxHostedPrivateVirtualInterfaceList{})
}
