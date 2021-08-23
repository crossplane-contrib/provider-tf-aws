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
// +groupName=shield.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/shield/v1alpha1"
)

type ShieldProtectionGroupObservation struct {
	ProtectionGroupArn string `json:"protectionGroupArn" tf:"protection_group_arn"`
}

type ShieldProtectionGroupParameters struct {
	Aggregation string `json:"aggregation" tf:"aggregation"`

	Members []string `json:"members,omitempty" tf:"members"`

	Pattern string `json:"pattern" tf:"pattern"`

	ProtectionGroupId string `json:"protectionGroupId" tf:"protection_group_id"`

	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ShieldProtectionGroupSpec defines the desired state of ShieldProtectionGroup
type ShieldProtectionGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ShieldProtectionGroupParameters `json:"forProvider"`
}

// ShieldProtectionGroupStatus defines the observed state of ShieldProtectionGroup.
type ShieldProtectionGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ShieldProtectionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ShieldProtectionGroup is the Schema for the ShieldProtectionGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ShieldProtectionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShieldProtectionGroupSpec   `json:"spec"`
	Status            ShieldProtectionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShieldProtectionGroupList contains a list of ShieldProtectionGroups
type ShieldProtectionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShieldProtectionGroup `json:"items"`
}

// Repository type metadata.
var (
	ShieldProtectionGroupKind             = "ShieldProtectionGroup"
	ShieldProtectionGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ShieldProtectionGroupKind}.String()
	ShieldProtectionGroupKindAPIVersion   = ShieldProtectionGroupKind + "." + v1alpha1.GroupVersion.String()
	ShieldProtectionGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(ShieldProtectionGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ShieldProtectionGroup{}, &ShieldProtectionGroupList{})
}
