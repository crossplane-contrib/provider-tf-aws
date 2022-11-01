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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LogDestinationConfigObservation struct {
}

type LogDestinationConfigParameters struct {

	// +kubebuilder:validation:Required
	LogDestination map[string]*string `json:"logDestination" tf:"log_destination,omitempty"`

	// +kubebuilder:validation:Required
	LogDestinationType *string `json:"logDestinationType" tf:"log_destination_type,omitempty"`

	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

type LoggingConfigurationLoggingConfigurationObservation struct {
}

type LoggingConfigurationLoggingConfigurationParameters struct {

	// +kubebuilder:validation:Required
	LogDestinationConfig []LogDestinationConfigParameters `json:"logDestinationConfig" tf:"log_destination_config,omitempty"`
}

type LoggingConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoggingConfigurationParameters struct {

	// +crossplane:generate:reference:type=Firewall
	// +kubebuilder:validation:Optional
	FirewallArn *string `json:"firewallArn,omitempty" tf:"firewall_arn,omitempty"`

	// +kubebuilder:validation:Optional
	FirewallArnRef *v1.Reference `json:"firewallArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FirewallArnSelector *v1.Selector `json:"firewallArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	LoggingConfiguration []LoggingConfigurationLoggingConfigurationParameters `json:"loggingConfiguration" tf:"logging_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LoggingConfigurationSpec defines the desired state of LoggingConfiguration
type LoggingConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoggingConfigurationParameters `json:"forProvider"`
}

// LoggingConfigurationStatus defines the observed state of LoggingConfiguration.
type LoggingConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoggingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingConfiguration is the Schema for the LoggingConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingConfigurationSpec   `json:"spec"`
	Status            LoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingConfigurationList contains a list of LoggingConfigurations
type LoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LoggingConfiguration_Kind             = "LoggingConfiguration"
	LoggingConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoggingConfiguration_Kind}.String()
	LoggingConfiguration_KindAPIVersion   = LoggingConfiguration_Kind + "." + CRDGroupVersion.String()
	LoggingConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LoggingConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LoggingConfiguration{}, &LoggingConfigurationList{})
}
