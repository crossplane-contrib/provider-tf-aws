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
// +groupName=route.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route/v1alpha1"
)

type RouteTableAssociationObservation struct {
}

type RouteTableAssociationParameters struct {
	GatewayId *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	RouteTableId string `json:"routeTableId" tf:"route_table_id"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

// RouteTableAssociationSpec defines the desired state of RouteTableAssociation
type RouteTableAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteTableAssociationParameters `json:"forProvider"`
}

// RouteTableAssociationStatus defines the observed state of RouteTableAssociation.
type RouteTableAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociation is the Schema for the RouteTableAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableAssociationSpec   `json:"spec"`
	Status            RouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableAssociationList contains a list of RouteTableAssociations
type RouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	RouteTableAssociationKind             = "RouteTableAssociation"
	RouteTableAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: RouteTableAssociationKind}.String()
	RouteTableAssociationKindAPIVersion   = RouteTableAssociationKind + "." + v1alpha1.GroupVersion.String()
	RouteTableAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(RouteTableAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&RouteTableAssociation{}, &RouteTableAssociationList{})
}
