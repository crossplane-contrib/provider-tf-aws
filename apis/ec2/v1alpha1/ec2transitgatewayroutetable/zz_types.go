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

type Ec2TransitGatewayRouteTableObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DefaultAssociationRouteTable bool `json:"defaultAssociationRouteTable" tf:"default_association_route_table"`

	DefaultPropagationRouteTable bool `json:"defaultPropagationRouteTable" tf:"default_propagation_route_table"`
}

type Ec2TransitGatewayRouteTableParameters struct {
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TransitGatewayId string `json:"transitGatewayId" tf:"transit_gateway_id"`
}

// Ec2TransitGatewayRouteTableSpec defines the desired state of Ec2TransitGatewayRouteTable
type Ec2TransitGatewayRouteTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayRouteTableParameters `json:"forProvider"`
}

// Ec2TransitGatewayRouteTableStatus defines the observed state of Ec2TransitGatewayRouteTable.
type Ec2TransitGatewayRouteTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTable is the Schema for the Ec2TransitGatewayRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteTableSpec   `json:"spec"`
	Status            Ec2TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableList contains a list of Ec2TransitGatewayRouteTables
type Ec2TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRouteTable `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayRouteTableKind             = "Ec2TransitGatewayRouteTable"
	Ec2TransitGatewayRouteTableGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2TransitGatewayRouteTableKind}.String()
	Ec2TransitGatewayRouteTableKindAPIVersion   = Ec2TransitGatewayRouteTableKind + "." + v1alpha1.GroupVersion.String()
	Ec2TransitGatewayRouteTableGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2TransitGatewayRouteTableKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2TransitGatewayRouteTable{}, &Ec2TransitGatewayRouteTableList{})
}
