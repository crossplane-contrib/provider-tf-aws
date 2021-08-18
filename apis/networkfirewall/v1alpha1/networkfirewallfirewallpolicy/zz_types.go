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
// +groupName=networkfirewall.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/networkfirewall/v1alpha1"
)

type ActionDefinitionObservation struct {
}

type ActionDefinitionParameters struct {
	PublishMetricAction []PublishMetricActionParameters `json:"publishMetricAction" tf:"publish_metric_action"`
}

type DimensionObservation struct {
}

type DimensionParameters struct {
	Value string `json:"value" tf:"value"`
}

type FirewallPolicyObservation struct {
}

type FirewallPolicyParameters struct {
	StatefulRuleGroupReference []StatefulRuleGroupReferenceParameters `json:"statefulRuleGroupReference,omitempty" tf:"stateful_rule_group_reference"`

	StatelessCustomAction []StatelessCustomActionParameters `json:"statelessCustomAction,omitempty" tf:"stateless_custom_action"`

	StatelessDefaultActions []string `json:"statelessDefaultActions" tf:"stateless_default_actions"`

	StatelessFragmentDefaultActions []string `json:"statelessFragmentDefaultActions" tf:"stateless_fragment_default_actions"`

	StatelessRuleGroupReference []StatelessRuleGroupReferenceParameters `json:"statelessRuleGroupReference,omitempty" tf:"stateless_rule_group_reference"`
}

type NetworkfirewallFirewallPolicyObservation struct {
	Arn string `json:"arn" tf:"arn"`

	UpdateToken string `json:"updateToken" tf:"update_token"`
}

type NetworkfirewallFirewallPolicyParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	FirewallPolicy []FirewallPolicyParameters `json:"firewallPolicy" tf:"firewall_policy"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type PublishMetricActionObservation struct {
}

type PublishMetricActionParameters struct {
	Dimension []DimensionParameters `json:"dimension" tf:"dimension"`
}

type StatefulRuleGroupReferenceObservation struct {
}

type StatefulRuleGroupReferenceParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type StatelessCustomActionObservation struct {
}

type StatelessCustomActionParameters struct {
	ActionDefinition []ActionDefinitionParameters `json:"actionDefinition" tf:"action_definition"`

	ActionName string `json:"actionName" tf:"action_name"`
}

type StatelessRuleGroupReferenceObservation struct {
}

type StatelessRuleGroupReferenceParameters struct {
	Priority int64 `json:"priority" tf:"priority"`

	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

// NetworkfirewallFirewallPolicySpec defines the desired state of NetworkfirewallFirewallPolicy
type NetworkfirewallFirewallPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkfirewallFirewallPolicyParameters `json:"forProvider"`
}

// NetworkfirewallFirewallPolicyStatus defines the observed state of NetworkfirewallFirewallPolicy.
type NetworkfirewallFirewallPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkfirewallFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkfirewallFirewallPolicy is the Schema for the NetworkfirewallFirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkfirewallFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkfirewallFirewallPolicySpec   `json:"spec"`
	Status            NetworkfirewallFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkfirewallFirewallPolicyList contains a list of NetworkfirewallFirewallPolicys
type NetworkfirewallFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkfirewallFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	NetworkfirewallFirewallPolicyKind             = "NetworkfirewallFirewallPolicy"
	NetworkfirewallFirewallPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NetworkfirewallFirewallPolicyKind}.String()
	NetworkfirewallFirewallPolicyKindAPIVersion   = NetworkfirewallFirewallPolicyKind + "." + v1alpha1.GroupVersion.String()
	NetworkfirewallFirewallPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(NetworkfirewallFirewallPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&NetworkfirewallFirewallPolicy{}, &NetworkfirewallFirewallPolicyList{})
}
