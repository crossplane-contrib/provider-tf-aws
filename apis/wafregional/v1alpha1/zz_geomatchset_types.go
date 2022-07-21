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

type GeoMatchConstraintObservation struct {
}

type GeoMatchConstraintParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type GeoMatchSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GeoMatchSetParameters struct {

	// +kubebuilder:validation:Optional
	GeoMatchConstraint []GeoMatchConstraintParameters `json:"geoMatchConstraint,omitempty" tf:"geo_match_constraint,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GeoMatchSetSpec defines the desired state of GeoMatchSet
type GeoMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GeoMatchSetParameters `json:"forProvider"`
}

// GeoMatchSetStatus defines the observed state of GeoMatchSet.
type GeoMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GeoMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GeoMatchSet is the Schema for the GeoMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GeoMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GeoMatchSetSpec   `json:"spec"`
	Status            GeoMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GeoMatchSetList contains a list of GeoMatchSets
type GeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeoMatchSet `json:"items"`
}

// Repository type metadata.
var (
	GeoMatchSet_Kind             = "GeoMatchSet"
	GeoMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GeoMatchSet_Kind}.String()
	GeoMatchSet_KindAPIVersion   = GeoMatchSet_Kind + "." + CRDGroupVersion.String()
	GeoMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(GeoMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&GeoMatchSet{}, &GeoMatchSetList{})
}
