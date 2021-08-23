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
// +groupName=route53.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/route53/v1alpha1"
)

type Route53ResolverFirewallRuleObservation struct {
}

type Route53ResolverFirewallRuleParameters struct {
	Action string `json:"action" tf:"action"`

	BlockOverrideDnsType *string `json:"blockOverrideDnsType,omitempty" tf:"block_override_dns_type"`

	BlockOverrideDomain *string `json:"blockOverrideDomain,omitempty" tf:"block_override_domain"`

	BlockOverrideTtl *int64 `json:"blockOverrideTtl,omitempty" tf:"block_override_ttl"`

	BlockResponse *string `json:"blockResponse,omitempty" tf:"block_response"`

	FirewallDomainListId string `json:"firewallDomainListId" tf:"firewall_domain_list_id"`

	FirewallRuleGroupId string `json:"firewallRuleGroupId" tf:"firewall_rule_group_id"`

	Name string `json:"name" tf:"name"`

	Priority int64 `json:"priority" tf:"priority"`
}

// Route53ResolverFirewallRuleSpec defines the desired state of Route53ResolverFirewallRule
type Route53ResolverFirewallRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Route53ResolverFirewallRuleParameters `json:"forProvider"`
}

// Route53ResolverFirewallRuleStatus defines the observed state of Route53ResolverFirewallRule.
type Route53ResolverFirewallRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Route53ResolverFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallRule is the Schema for the Route53ResolverFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Route53ResolverFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverFirewallRuleSpec   `json:"spec"`
	Status            Route53ResolverFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverFirewallRuleList contains a list of Route53ResolverFirewallRules
type Route53ResolverFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverFirewallRuleKind             = "Route53ResolverFirewallRule"
	Route53ResolverFirewallRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Route53ResolverFirewallRuleKind}.String()
	Route53ResolverFirewallRuleKindAPIVersion   = Route53ResolverFirewallRuleKind + "." + v1alpha1.GroupVersion.String()
	Route53ResolverFirewallRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(Route53ResolverFirewallRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Route53ResolverFirewallRule{}, &Route53ResolverFirewallRuleList{})
}
