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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LocalGatewayRouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LocalGatewayRouteParameters struct {

	// +kubebuilder:validation:Required
	DestinationCidrBlock *string `json:"destinationCidrBlock" tf:"destination_cidr_block,omitempty"`

	// +kubebuilder:validation:Required
	LocalGatewayRouteTableID *string `json:"localGatewayRouteTableId" tf:"local_gateway_route_table_id,omitempty"`

	// +kubebuilder:validation:Required
	LocalGatewayVirtualInterfaceGroupID *string `json:"localGatewayVirtualInterfaceGroupId" tf:"local_gateway_virtual_interface_group_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LocalGatewayRouteSpec defines the desired state of LocalGatewayRoute
type LocalGatewayRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalGatewayRouteParameters `json:"forProvider"`
}

// LocalGatewayRouteStatus defines the observed state of LocalGatewayRoute.
type LocalGatewayRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocalGatewayRoute is the Schema for the LocalGatewayRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocalGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalGatewayRouteSpec   `json:"spec"`
	Status            LocalGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalGatewayRouteList contains a list of LocalGatewayRoutes
type LocalGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	LocalGatewayRoute_Kind             = "LocalGatewayRoute"
	LocalGatewayRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocalGatewayRoute_Kind}.String()
	LocalGatewayRoute_KindAPIVersion   = LocalGatewayRoute_Kind + "." + CRDGroupVersion.String()
	LocalGatewayRoute_GroupVersionKind = CRDGroupVersion.WithKind(LocalGatewayRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&LocalGatewayRoute{}, &LocalGatewayRouteList{})
}
