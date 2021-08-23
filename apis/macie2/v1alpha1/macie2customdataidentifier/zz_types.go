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
// +groupName=macie2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/macie2/v1alpha1"
)

type Macie2CustomDataIdentifierObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedAt string `json:"createdAt" tf:"created_at"`
}

type Macie2CustomDataIdentifierParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	IgnoreWords []string `json:"ignoreWords,omitempty" tf:"ignore_words"`

	Keywords []string `json:"keywords,omitempty" tf:"keywords"`

	MaximumMatchDistance *int64 `json:"maximumMatchDistance,omitempty" tf:"maximum_match_distance"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Regex *string `json:"regex,omitempty" tf:"regex"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Macie2CustomDataIdentifierSpec defines the desired state of Macie2CustomDataIdentifier
type Macie2CustomDataIdentifierSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2CustomDataIdentifierParameters `json:"forProvider"`
}

// Macie2CustomDataIdentifierStatus defines the observed state of Macie2CustomDataIdentifier.
type Macie2CustomDataIdentifierStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2CustomDataIdentifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2CustomDataIdentifier is the Schema for the Macie2CustomDataIdentifiers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Macie2CustomDataIdentifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2CustomDataIdentifierSpec   `json:"spec"`
	Status            Macie2CustomDataIdentifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2CustomDataIdentifierList contains a list of Macie2CustomDataIdentifiers
type Macie2CustomDataIdentifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2CustomDataIdentifier `json:"items"`
}

// Repository type metadata.
var (
	Macie2CustomDataIdentifierKind             = "Macie2CustomDataIdentifier"
	Macie2CustomDataIdentifierGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Macie2CustomDataIdentifierKind}.String()
	Macie2CustomDataIdentifierKindAPIVersion   = Macie2CustomDataIdentifierKind + "." + v1alpha1.GroupVersion.String()
	Macie2CustomDataIdentifierGroupVersionKind = v1alpha1.GroupVersion.WithKind(Macie2CustomDataIdentifierKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Macie2CustomDataIdentifier{}, &Macie2CustomDataIdentifierList{})
}
