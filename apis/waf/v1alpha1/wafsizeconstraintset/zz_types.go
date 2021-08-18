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
// +groupName=waf.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/waf/v1alpha1"
)

type FieldToMatchObservation struct {
}

type FieldToMatchParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	Type string `json:"type" tf:"type"`
}

type SizeConstraintsObservation struct {
}

type SizeConstraintsParameters struct {
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`

	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	Size int64 `json:"size" tf:"size"`

	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

type WafSizeConstraintSetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafSizeConstraintSetParameters struct {
	Name string `json:"name" tf:"name"`

	SizeConstraints []SizeConstraintsParameters `json:"sizeConstraints,omitempty" tf:"size_constraints"`
}

// WafSizeConstraintSetSpec defines the desired state of WafSizeConstraintSet
type WafSizeConstraintSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafSizeConstraintSetParameters `json:"forProvider"`
}

// WafSizeConstraintSetStatus defines the observed state of WafSizeConstraintSet.
type WafSizeConstraintSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafSizeConstraintSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafSizeConstraintSet is the Schema for the WafSizeConstraintSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WafSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafSizeConstraintSetSpec   `json:"spec"`
	Status            WafSizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafSizeConstraintSetList contains a list of WafSizeConstraintSets
type WafSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafSizeConstraintSet `json:"items"`
}

// Repository type metadata.
var (
	WafSizeConstraintSetKind             = "WafSizeConstraintSet"
	WafSizeConstraintSetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WafSizeConstraintSetKind}.String()
	WafSizeConstraintSetKindAPIVersion   = WafSizeConstraintSetKind + "." + v1alpha1.GroupVersion.String()
	WafSizeConstraintSetGroupVersionKind = v1alpha1.GroupVersion.WithKind(WafSizeConstraintSetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&WafSizeConstraintSet{}, &WafSizeConstraintSetList{})
}
