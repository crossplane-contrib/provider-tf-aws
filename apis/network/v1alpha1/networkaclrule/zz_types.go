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
// +groupName=network.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/network/v1alpha1"
)

type NetworkAclRuleObservation struct {
}

type NetworkAclRuleParameters struct {
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	Egress *bool `json:"egress,omitempty" tf:"egress"`

	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port"`

	IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code"`

	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type"`

	Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	NetworkAclId string `json:"networkAclId" tf:"network_acl_id"`

	Protocol string `json:"protocol" tf:"protocol"`

	RuleAction string `json:"ruleAction" tf:"rule_action"`

	RuleNumber int64 `json:"ruleNumber" tf:"rule_number"`

	ToPort *int64 `json:"toPort,omitempty" tf:"to_port"`
}

// NetworkAclRuleSpec defines the desired state of NetworkAclRule
type NetworkAclRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       NetworkAclRuleParameters `json:"forProvider"`
}

// NetworkAclRuleStatus defines the observed state of NetworkAclRule.
type NetworkAclRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          NetworkAclRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAclRule is the Schema for the NetworkAclRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkAclRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkAclRuleSpec   `json:"spec"`
	Status            NetworkAclRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAclRuleList contains a list of NetworkAclRules
type NetworkAclRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkAclRule `json:"items"`
}

// Repository type metadata.
var (
	NetworkAclRuleKind             = "NetworkAclRule"
	NetworkAclRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: NetworkAclRuleKind}.String()
	NetworkAclRuleKindAPIVersion   = NetworkAclRuleKind + "." + v1alpha1.GroupVersion.String()
	NetworkAclRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(NetworkAclRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&NetworkAclRule{}, &NetworkAclRuleList{})
}
