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
// +groupName=cloudwatch.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudwatch/v1alpha1"
)

type CloudwatchLogGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudwatchLogGroupParameters struct {
	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// CloudwatchLogGroupSpec defines the desired state of CloudwatchLogGroup
type CloudwatchLogGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchLogGroupParameters `json:"forProvider"`
}

// CloudwatchLogGroupStatus defines the observed state of CloudwatchLogGroup.
type CloudwatchLogGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchLogGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogGroup is the Schema for the CloudwatchLogGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudwatchLogGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogGroupSpec   `json:"spec"`
	Status            CloudwatchLogGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogGroupList contains a list of CloudwatchLogGroups
type CloudwatchLogGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogGroup `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchLogGroupKind             = "CloudwatchLogGroup"
	CloudwatchLogGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudwatchLogGroupKind}.String()
	CloudwatchLogGroupKindAPIVersion   = CloudwatchLogGroupKind + "." + v1alpha1.GroupVersion.String()
	CloudwatchLogGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudwatchLogGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudwatchLogGroup{}, &CloudwatchLogGroupList{})
}
