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
// +groupName=sfn.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/sfn/v1alpha1"
)

type LoggingConfigurationObservation struct {
}

type LoggingConfigurationParameters struct {
	IncludeExecutionData *bool `json:"includeExecutionData,omitempty" tf:"include_execution_data"`

	Level *string `json:"level,omitempty" tf:"level"`

	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination"`
}

type SfnStateMachineObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreationDate string `json:"creationDate" tf:"creation_date"`

	Status string `json:"status" tf:"status"`
}

type SfnStateMachineParameters struct {
	Definition string `json:"definition" tf:"definition"`

	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration"`

	Name string `json:"name" tf:"name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TracingConfiguration []TracingConfigurationParameters `json:"tracingConfiguration,omitempty" tf:"tracing_configuration"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type TracingConfigurationObservation struct {
}

type TracingConfigurationParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

// SfnStateMachineSpec defines the desired state of SfnStateMachine
type SfnStateMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SfnStateMachineParameters `json:"forProvider"`
}

// SfnStateMachineStatus defines the observed state of SfnStateMachine.
type SfnStateMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SfnStateMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SfnStateMachine is the Schema for the SfnStateMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SfnStateMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SfnStateMachineSpec   `json:"spec"`
	Status            SfnStateMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SfnStateMachineList contains a list of SfnStateMachines
type SfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SfnStateMachine `json:"items"`
}

// Repository type metadata.
var (
	SfnStateMachineKind             = "SfnStateMachine"
	SfnStateMachineGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SfnStateMachineKind}.String()
	SfnStateMachineKindAPIVersion   = SfnStateMachineKind + "." + v1alpha1.GroupVersion.String()
	SfnStateMachineGroupVersionKind = v1alpha1.GroupVersion.WithKind(SfnStateMachineKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SfnStateMachine{}, &SfnStateMachineList{})
}
