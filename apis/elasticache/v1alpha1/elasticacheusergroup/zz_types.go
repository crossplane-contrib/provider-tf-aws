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
// +groupName=elasticache.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/elasticache/v1alpha1"
)

type ElasticacheUserGroupObservation struct {
}

type ElasticacheUserGroupParameters struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`

	Engine string `json:"engine" tf:"engine"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UserGroupId string `json:"userGroupId" tf:"user_group_id"`

	UserIds []string `json:"userIds,omitempty" tf:"user_ids"`
}

// ElasticacheUserGroupSpec defines the desired state of ElasticacheUserGroup
type ElasticacheUserGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticacheUserGroupParameters `json:"forProvider"`
}

// ElasticacheUserGroupStatus defines the observed state of ElasticacheUserGroup.
type ElasticacheUserGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticacheUserGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheUserGroup is the Schema for the ElasticacheUserGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ElasticacheUserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheUserGroupSpec   `json:"spec"`
	Status            ElasticacheUserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheUserGroupList contains a list of ElasticacheUserGroups
type ElasticacheUserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheUserGroup `json:"items"`
}

// Repository type metadata.
var (
	ElasticacheUserGroupKind             = "ElasticacheUserGroup"
	ElasticacheUserGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ElasticacheUserGroupKind}.String()
	ElasticacheUserGroupKindAPIVersion   = ElasticacheUserGroupKind + "." + v1alpha1.GroupVersion.String()
	ElasticacheUserGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(ElasticacheUserGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ElasticacheUserGroup{}, &ElasticacheUserGroupList{})
}
