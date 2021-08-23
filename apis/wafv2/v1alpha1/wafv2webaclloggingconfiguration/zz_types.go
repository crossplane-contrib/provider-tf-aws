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
// +groupName=wafv2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/wafv2/v1alpha1"
)

type ActionConditionObservation struct {
}

type ActionConditionParameters struct {
	Action string `json:"action" tf:"action"`
}

type AllQueryArgumentsObservation struct {
}

type AllQueryArgumentsParameters struct {
}

type BodyObservation struct {
}

type BodyParameters struct {
}

type ConditionObservation struct {
}

type ConditionParameters struct {
	ActionCondition []ActionConditionParameters `json:"actionCondition,omitempty" tf:"action_condition"`

	LabelNameCondition []LabelNameConditionParameters `json:"labelNameCondition,omitempty" tf:"label_name_condition"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	Behavior string `json:"behavior" tf:"behavior"`

	Condition []ConditionParameters `json:"condition" tf:"condition"`

	Requirement string `json:"requirement" tf:"requirement"`
}

type LabelNameConditionObservation struct {
}

type LabelNameConditionParameters struct {
	LabelName string `json:"labelName" tf:"label_name"`
}

type LoggingFilterObservation struct {
}

type LoggingFilterParameters struct {
	DefaultBehavior string `json:"defaultBehavior" tf:"default_behavior"`

	Filter []FilterParameters `json:"filter" tf:"filter"`
}

type MethodObservation struct {
}

type MethodParameters struct {
}

type QueryStringObservation struct {
}

type QueryStringParameters struct {
}

type RedactedFieldsObservation struct {
}

type RedactedFieldsParameters struct {
	AllQueryArguments []AllQueryArgumentsParameters `json:"allQueryArguments,omitempty" tf:"all_query_arguments"`

	Body []BodyParameters `json:"body,omitempty" tf:"body"`

	Method []MethodParameters `json:"method,omitempty" tf:"method"`

	QueryString []QueryStringParameters `json:"queryString,omitempty" tf:"query_string"`

	SingleHeader []SingleHeaderParameters `json:"singleHeader,omitempty" tf:"single_header"`

	SingleQueryArgument []SingleQueryArgumentParameters `json:"singleQueryArgument,omitempty" tf:"single_query_argument"`

	UriPath []UriPathParameters `json:"uriPath,omitempty" tf:"uri_path"`
}

type SingleHeaderObservation struct {
}

type SingleHeaderParameters struct {
	Name string `json:"name" tf:"name"`
}

type SingleQueryArgumentObservation struct {
}

type SingleQueryArgumentParameters struct {
	Name string `json:"name" tf:"name"`
}

type UriPathObservation struct {
}

type UriPathParameters struct {
}

type Wafv2WebAclLoggingConfigurationObservation struct {
}

type Wafv2WebAclLoggingConfigurationParameters struct {
	LogDestinationConfigs []string `json:"logDestinationConfigs" tf:"log_destination_configs"`

	LoggingFilter []LoggingFilterParameters `json:"loggingFilter,omitempty" tf:"logging_filter"`

	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields"`

	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

// Wafv2WebAclLoggingConfigurationSpec defines the desired state of Wafv2WebAclLoggingConfiguration
type Wafv2WebAclLoggingConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Wafv2WebAclLoggingConfigurationParameters `json:"forProvider"`
}

// Wafv2WebAclLoggingConfigurationStatus defines the observed state of Wafv2WebAclLoggingConfiguration.
type Wafv2WebAclLoggingConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Wafv2WebAclLoggingConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclLoggingConfiguration is the Schema for the Wafv2WebAclLoggingConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Wafv2WebAclLoggingConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Wafv2WebAclLoggingConfigurationSpec   `json:"spec"`
	Status            Wafv2WebAclLoggingConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclLoggingConfigurationList contains a list of Wafv2WebAclLoggingConfigurations
type Wafv2WebAclLoggingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2WebAclLoggingConfiguration `json:"items"`
}

// Repository type metadata.
var (
	Wafv2WebAclLoggingConfigurationKind             = "Wafv2WebAclLoggingConfiguration"
	Wafv2WebAclLoggingConfigurationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Wafv2WebAclLoggingConfigurationKind}.String()
	Wafv2WebAclLoggingConfigurationKindAPIVersion   = Wafv2WebAclLoggingConfigurationKind + "." + v1alpha1.GroupVersion.String()
	Wafv2WebAclLoggingConfigurationGroupVersionKind = v1alpha1.GroupVersion.WithKind(Wafv2WebAclLoggingConfigurationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Wafv2WebAclLoggingConfiguration{}, &Wafv2WebAclLoggingConfigurationList{})
}
