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
// +groupName=organizations.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/organizations/v1alpha1"
)

type AccountsObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Email string `json:"email" tf:"email"`

	Id string `json:"id" tf:"id"`

	Name string `json:"name" tf:"name"`
}

type AccountsParameters struct {
}

type OrganizationsOrganizationalUnitObservation struct {
	Accounts []AccountsObservation `json:"accounts" tf:"accounts"`

	Arn string `json:"arn" tf:"arn"`
}

type OrganizationsOrganizationalUnitParameters struct {
	Name string `json:"name" tf:"name"`

	ParentId string `json:"parentId" tf:"parent_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// OrganizationsOrganizationalUnitSpec defines the desired state of OrganizationsOrganizationalUnit
type OrganizationsOrganizationalUnitSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OrganizationsOrganizationalUnitParameters `json:"forProvider"`
}

// OrganizationsOrganizationalUnitStatus defines the observed state of OrganizationsOrganizationalUnit.
type OrganizationsOrganizationalUnitStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OrganizationsOrganizationalUnitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganizationalUnit is the Schema for the OrganizationsOrganizationalUnits API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type OrganizationsOrganizationalUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsOrganizationalUnitSpec   `json:"spec"`
	Status            OrganizationsOrganizationalUnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganizationalUnitList contains a list of OrganizationsOrganizationalUnits
type OrganizationsOrganizationalUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsOrganizationalUnit `json:"items"`
}

// Repository type metadata.
var (
	OrganizationsOrganizationalUnitKind             = "OrganizationsOrganizationalUnit"
	OrganizationsOrganizationalUnitGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: OrganizationsOrganizationalUnitKind}.String()
	OrganizationsOrganizationalUnitKindAPIVersion   = OrganizationsOrganizationalUnitKind + "." + v1alpha1.GroupVersion.String()
	OrganizationsOrganizationalUnitGroupVersionKind = v1alpha1.GroupVersion.WithKind(OrganizationsOrganizationalUnitKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&OrganizationsOrganizationalUnit{}, &OrganizationsOrganizationalUnitList{})
}
