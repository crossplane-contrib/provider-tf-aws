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

type DxGatewayObservation struct {
	OwnerAccountId string `json:"ownerAccountId" tf:"owner_account_id"`
}

type DxGatewayParameters struct {
	AmazonSideAsn string `json:"amazonSideAsn" tf:"amazon_side_asn"`

	Name string `json:"name" tf:"name"`
}

// DxGatewaySpec defines the desired state of DxGateway
type DxGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxGatewayParameters `json:"forProvider"`
}

// DxGatewayStatus defines the observed state of DxGateway.
type DxGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxGateway is the Schema for the DxGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewaySpec   `json:"spec"`
	Status            DxGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxGatewayList contains a list of DxGateways
type DxGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxGateway `json:"items"`
}

// Repository type metadata.
var (
	DxGatewayKind             = "DxGateway"
	DxGatewayGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxGatewayKind}.String()
	DxGatewayKindAPIVersion   = DxGatewayKind + "." + v1alpha1.GroupVersion.String()
	DxGatewayGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxGatewayKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxGateway{}, &DxGatewayList{})
}
