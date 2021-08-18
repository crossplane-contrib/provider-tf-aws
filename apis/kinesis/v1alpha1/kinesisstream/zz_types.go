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
// +groupName=kinesis.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kinesis/v1alpha1"
)

type KinesisStreamObservation struct {
}

type KinesisStreamParameters struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`

	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type"`

	EnforceConsumerDeletion *bool `json:"enforceConsumerDeletion,omitempty" tf:"enforce_consumer_deletion"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Name string `json:"name" tf:"name"`

	RetentionPeriod *int64 `json:"retentionPeriod,omitempty" tf:"retention_period"`

	ShardCount int64 `json:"shardCount" tf:"shard_count"`

	ShardLevelMetrics []string `json:"shardLevelMetrics,omitempty" tf:"shard_level_metrics"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// KinesisStreamSpec defines the desired state of KinesisStream
type KinesisStreamSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KinesisStreamParameters `json:"forProvider"`
}

// KinesisStreamStatus defines the observed state of KinesisStream.
type KinesisStreamStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KinesisStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisStream is the Schema for the KinesisStreams API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KinesisStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisStreamSpec   `json:"spec"`
	Status            KinesisStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisStreamList contains a list of KinesisStreams
type KinesisStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisStream `json:"items"`
}

// Repository type metadata.
var (
	KinesisStreamKind             = "KinesisStream"
	KinesisStreamGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: KinesisStreamKind}.String()
	KinesisStreamKindAPIVersion   = KinesisStreamKind + "." + v1alpha1.GroupVersion.String()
	KinesisStreamGroupVersionKind = v1alpha1.GroupVersion.WithKind(KinesisStreamKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&KinesisStream{}, &KinesisStreamList{})
}
