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
// +groupName=ses.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ses/v1alpha1"
)

type CloudwatchDestinationObservation struct {
}

type CloudwatchDestinationParameters struct {
	DefaultValue string `json:"defaultValue" tf:"default_value"`

	DimensionName string `json:"dimensionName" tf:"dimension_name"`

	ValueSource string `json:"valueSource" tf:"value_source"`
}

type KinesisDestinationObservation struct {
}

type KinesisDestinationParameters struct {
	RoleArn string `json:"roleArn" tf:"role_arn"`

	StreamArn string `json:"streamArn" tf:"stream_arn"`
}

type SesEventDestinationObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SesEventDestinationParameters struct {
	CloudwatchDestination []CloudwatchDestinationParameters `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination"`

	ConfigurationSetName string `json:"configurationSetName" tf:"configuration_set_name"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	KinesisDestination []KinesisDestinationParameters `json:"kinesisDestination,omitempty" tf:"kinesis_destination"`

	MatchingTypes []string `json:"matchingTypes" tf:"matching_types"`

	Name string `json:"name" tf:"name"`

	SnsDestination []SnsDestinationParameters `json:"snsDestination,omitempty" tf:"sns_destination"`
}

type SnsDestinationObservation struct {
}

type SnsDestinationParameters struct {
	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

// SesEventDestinationSpec defines the desired state of SesEventDestination
type SesEventDestinationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesEventDestinationParameters `json:"forProvider"`
}

// SesEventDestinationStatus defines the observed state of SesEventDestination.
type SesEventDestinationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesEventDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesEventDestination is the Schema for the SesEventDestinations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SesEventDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesEventDestinationSpec   `json:"spec"`
	Status            SesEventDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesEventDestinationList contains a list of SesEventDestinations
type SesEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesEventDestination `json:"items"`
}

// Repository type metadata.
var (
	SesEventDestinationKind             = "SesEventDestination"
	SesEventDestinationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SesEventDestinationKind}.String()
	SesEventDestinationKindAPIVersion   = SesEventDestinationKind + "." + v1alpha1.GroupVersion.String()
	SesEventDestinationGroupVersionKind = v1alpha1.GroupVersion.WithKind(SesEventDestinationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SesEventDestination{}, &SesEventDestinationList{})
}
