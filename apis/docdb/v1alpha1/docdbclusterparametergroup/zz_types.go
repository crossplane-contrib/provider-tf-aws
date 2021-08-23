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
// +groupName=docdb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/docdb/v1alpha1"
)

type DocdbClusterParameterGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DocdbClusterParameterGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Family string `json:"family" tf:"family"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method"`

	Name string `json:"name" tf:"name"`

	Value string `json:"value" tf:"value"`
}

// DocdbClusterParameterGroupSpec defines the desired state of DocdbClusterParameterGroup
type DocdbClusterParameterGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DocdbClusterParameterGroupParameters `json:"forProvider"`
}

// DocdbClusterParameterGroupStatus defines the observed state of DocdbClusterParameterGroup.
type DocdbClusterParameterGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DocdbClusterParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterParameterGroup is the Schema for the DocdbClusterParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocdbClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterParameterGroupSpec   `json:"spec"`
	Status            DocdbClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterParameterGroupList contains a list of DocdbClusterParameterGroups
type DocdbClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbClusterParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	DocdbClusterParameterGroupKind             = "DocdbClusterParameterGroup"
	DocdbClusterParameterGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DocdbClusterParameterGroupKind}.String()
	DocdbClusterParameterGroupKindAPIVersion   = DocdbClusterParameterGroupKind + "." + v1alpha1.GroupVersion.String()
	DocdbClusterParameterGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(DocdbClusterParameterGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DocdbClusterParameterGroup{}, &DocdbClusterParameterGroupList{})
}
