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

type DxHostedTransitVirtualInterfaceAccepterObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DxHostedTransitVirtualInterfaceAccepterParameters struct {
	DxGatewayId string `json:"dxGatewayId" tf:"dx_gateway_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VirtualInterfaceId string `json:"virtualInterfaceId" tf:"virtual_interface_id"`
}

// DxHostedTransitVirtualInterfaceAccepterSpec defines the desired state of DxHostedTransitVirtualInterfaceAccepter
type DxHostedTransitVirtualInterfaceAccepterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxHostedTransitVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// DxHostedTransitVirtualInterfaceAccepterStatus defines the observed state of DxHostedTransitVirtualInterfaceAccepter.
type DxHostedTransitVirtualInterfaceAccepterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxHostedTransitVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterfaceAccepter is the Schema for the DxHostedTransitVirtualInterfaceAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxHostedTransitVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxHostedTransitVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            DxHostedTransitVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterfaceAccepterList contains a list of DxHostedTransitVirtualInterfaceAccepters
type DxHostedTransitVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedTransitVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	DxHostedTransitVirtualInterfaceAccepterKind             = "DxHostedTransitVirtualInterfaceAccepter"
	DxHostedTransitVirtualInterfaceAccepterGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxHostedTransitVirtualInterfaceAccepterKind}.String()
	DxHostedTransitVirtualInterfaceAccepterKindAPIVersion   = DxHostedTransitVirtualInterfaceAccepterKind + "." + v1alpha1.GroupVersion.String()
	DxHostedTransitVirtualInterfaceAccepterGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxHostedTransitVirtualInterfaceAccepterKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxHostedTransitVirtualInterfaceAccepter{}, &DxHostedTransitVirtualInterfaceAccepterList{})
}
