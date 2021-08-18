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

type SesReceiptRuleSetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SesReceiptRuleSetParameters struct {
	RuleSetName string `json:"ruleSetName" tf:"rule_set_name"`
}

// SesReceiptRuleSetSpec defines the desired state of SesReceiptRuleSet
type SesReceiptRuleSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesReceiptRuleSetParameters `json:"forProvider"`
}

// SesReceiptRuleSetStatus defines the observed state of SesReceiptRuleSet.
type SesReceiptRuleSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesReceiptRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRuleSet is the Schema for the SesReceiptRuleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SesReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesReceiptRuleSetSpec   `json:"spec"`
	Status            SesReceiptRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesReceiptRuleSetList contains a list of SesReceiptRuleSets
type SesReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesReceiptRuleSet `json:"items"`
}

// Repository type metadata.
var (
	SesReceiptRuleSetKind             = "SesReceiptRuleSet"
	SesReceiptRuleSetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SesReceiptRuleSetKind}.String()
	SesReceiptRuleSetKindAPIVersion   = SesReceiptRuleSetKind + "." + v1alpha1.GroupVersion.String()
	SesReceiptRuleSetGroupVersionKind = v1alpha1.GroupVersion.WithKind(SesReceiptRuleSetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SesReceiptRuleSet{}, &SesReceiptRuleSetList{})
}
