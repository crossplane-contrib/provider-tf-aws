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
// +groupName=wafregional.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/wafregional/v1alpha1"
)

type ActionObservation struct {
}

type ActionParameters struct {
	Type string `json:"type" tf:"type"`
}

type DefaultActionObservation struct {
}

type DefaultActionParameters struct {
	Type string `json:"type" tf:"type"`
}

type FieldToMatchObservation struct {
}

type FieldToMatchParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	Type string `json:"type" tf:"type"`
}

type LoggingConfigurationObservation struct {
}

type LoggingConfigurationParameters struct {
	LogDestination string `json:"logDestination" tf:"log_destination"`

	RedactedFields []RedactedFieldsParameters `json:"redactedFields,omitempty" tf:"redacted_fields"`
}

type OverrideActionObservation struct {
}

type OverrideActionParameters struct {
	Type string `json:"type" tf:"type"`
}

type RedactedFieldsObservation struct {
}

type RedactedFieldsParameters struct {
	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`
}

type RuleObservation struct {
}

type RuleParameters struct {
	Action []ActionParameters `json:"action,omitempty" tf:"action"`

	OverrideAction []OverrideActionParameters `json:"overrideAction,omitempty" tf:"override_action"`

	Priority int64 `json:"priority" tf:"priority"`

	RuleId string `json:"ruleId" tf:"rule_id"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type WafregionalWebAclObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafregionalWebAclParameters struct {
	DefaultAction []DefaultActionParameters `json:"defaultAction" tf:"default_action"`

	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration"`

	MetricName string `json:"metricName" tf:"metric_name"`

	Name string `json:"name" tf:"name"`

	Rule []RuleParameters `json:"rule,omitempty" tf:"rule"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// WafregionalWebAclSpec defines the desired state of WafregionalWebAcl
type WafregionalWebAclSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalWebAclParameters `json:"forProvider"`
}

// WafregionalWebAclStatus defines the observed state of WafregionalWebAcl.
type WafregionalWebAclStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalWebAclObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalWebAcl is the Schema for the WafregionalWebAcls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type WafregionalWebAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalWebAclSpec   `json:"spec"`
	Status            WafregionalWebAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalWebAclList contains a list of WafregionalWebAcls
type WafregionalWebAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalWebAcl `json:"items"`
}

// Repository type metadata.
var (
	WafregionalWebAclKind             = "WafregionalWebAcl"
	WafregionalWebAclGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: WafregionalWebAclKind}.String()
	WafregionalWebAclKindAPIVersion   = WafregionalWebAclKind + "." + v1alpha1.GroupVersion.String()
	WafregionalWebAclGroupVersionKind = v1alpha1.GroupVersion.WithKind(WafregionalWebAclKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&WafregionalWebAcl{}, &WafregionalWebAclList{})
}
