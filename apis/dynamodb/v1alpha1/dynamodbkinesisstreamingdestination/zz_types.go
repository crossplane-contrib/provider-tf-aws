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
// +groupName=dynamodb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dynamodb/v1alpha1"
)

type DynamodbKinesisStreamingDestinationObservation struct {
}

type DynamodbKinesisStreamingDestinationParameters struct {
	StreamArn string `json:"streamArn" tf:"stream_arn"`

	TableName string `json:"tableName" tf:"table_name"`
}

// DynamodbKinesisStreamingDestinationSpec defines the desired state of DynamodbKinesisStreamingDestination
type DynamodbKinesisStreamingDestinationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DynamodbKinesisStreamingDestinationParameters `json:"forProvider"`
}

// DynamodbKinesisStreamingDestinationStatus defines the observed state of DynamodbKinesisStreamingDestination.
type DynamodbKinesisStreamingDestinationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DynamodbKinesisStreamingDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbKinesisStreamingDestination is the Schema for the DynamodbKinesisStreamingDestinations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DynamodbKinesisStreamingDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbKinesisStreamingDestinationSpec   `json:"spec"`
	Status            DynamodbKinesisStreamingDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbKinesisStreamingDestinationList contains a list of DynamodbKinesisStreamingDestinations
type DynamodbKinesisStreamingDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamodbKinesisStreamingDestination `json:"items"`
}

// Repository type metadata.
var (
	DynamodbKinesisStreamingDestinationKind             = "DynamodbKinesisStreamingDestination"
	DynamodbKinesisStreamingDestinationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DynamodbKinesisStreamingDestinationKind}.String()
	DynamodbKinesisStreamingDestinationKindAPIVersion   = DynamodbKinesisStreamingDestinationKind + "." + v1alpha1.GroupVersion.String()
	DynamodbKinesisStreamingDestinationGroupVersionKind = v1alpha1.GroupVersion.WithKind(DynamodbKinesisStreamingDestinationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DynamodbKinesisStreamingDestination{}, &DynamodbKinesisStreamingDestinationList{})
}
