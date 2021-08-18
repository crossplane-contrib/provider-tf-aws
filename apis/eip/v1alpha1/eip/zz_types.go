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
// +groupName=eip.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/eip/v1alpha1"
)

type EipObservation struct {
	AllocationId string `json:"allocationId" tf:"allocation_id"`

	AssociationId string `json:"associationId" tf:"association_id"`

	CarrierIp string `json:"carrierIp" tf:"carrier_ip"`

	CustomerOwnedIp string `json:"customerOwnedIp" tf:"customer_owned_ip"`

	Domain string `json:"domain" tf:"domain"`

	PrivateDns string `json:"privateDns" tf:"private_dns"`

	PrivateIp string `json:"privateIp" tf:"private_ip"`

	PublicDns string `json:"publicDns" tf:"public_dns"`

	PublicIp string `json:"publicIp" tf:"public_ip"`
}

type EipParameters struct {
	Address *string `json:"address,omitempty" tf:"address"`

	AssociateWithPrivateIp *string `json:"associateWithPrivateIp,omitempty" tf:"associate_with_private_ip"`

	CustomerOwnedIpv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`

	Instance *string `json:"instance,omitempty" tf:"instance"`

	NetworkBorderGroup *string `json:"networkBorderGroup,omitempty" tf:"network_border_group"`

	NetworkInterface *string `json:"networkInterface,omitempty" tf:"network_interface"`

	PublicIpv4Pool *string `json:"publicIpv4Pool,omitempty" tf:"public_ipv4_pool"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Vpc *bool `json:"vpc,omitempty" tf:"vpc"`
}

// EipSpec defines the desired state of Eip
type EipSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EipParameters `json:"forProvider"`
}

// EipStatus defines the observed state of Eip.
type EipStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Eip is the Schema for the Eips API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Eip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EipSpec   `json:"spec"`
	Status            EipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EipList contains a list of Eips
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eip `json:"items"`
}

// Repository type metadata.
var (
	EipKind             = "Eip"
	EipGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EipKind}.String()
	EipKindAPIVersion   = EipKind + "." + v1alpha1.GroupVersion.String()
	EipGroupVersionKind = v1alpha1.GroupVersion.WithKind(EipKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Eip{}, &EipList{})
}
