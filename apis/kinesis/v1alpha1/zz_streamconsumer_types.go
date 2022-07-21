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

type StreamConsumerObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StreamConsumerParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StreamArn *string `json:"streamArn" tf:"stream_arn,omitempty"`
}

// StreamConsumerSpec defines the desired state of StreamConsumer
type StreamConsumerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamConsumerParameters `json:"forProvider"`
}

// StreamConsumerStatus defines the observed state of StreamConsumer.
type StreamConsumerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamConsumerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamConsumer is the Schema for the StreamConsumers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type StreamConsumer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamConsumerSpec   `json:"spec"`
	Status            StreamConsumerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamConsumerList contains a list of StreamConsumers
type StreamConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamConsumer `json:"items"`
}

// Repository type metadata.
var (
	StreamConsumer_Kind             = "StreamConsumer"
	StreamConsumer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamConsumer_Kind}.String()
	StreamConsumer_KindAPIVersion   = StreamConsumer_Kind + "." + CRDGroupVersion.String()
	StreamConsumer_GroupVersionKind = CRDGroupVersion.WithKind(StreamConsumer_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamConsumer{}, &StreamConsumerList{})
}
