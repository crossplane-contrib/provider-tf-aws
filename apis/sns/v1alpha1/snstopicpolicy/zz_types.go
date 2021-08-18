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
// +groupName=sns.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/sns/v1alpha1"
)

type SnsTopicPolicyObservation struct {
	Owner string `json:"owner" tf:"owner"`
}

type SnsTopicPolicyParameters struct {
	Arn string `json:"arn" tf:"arn"`

	Policy string `json:"policy" tf:"policy"`
}

// SnsTopicPolicySpec defines the desired state of SnsTopicPolicy
type SnsTopicPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SnsTopicPolicyParameters `json:"forProvider"`
}

// SnsTopicPolicyStatus defines the observed state of SnsTopicPolicy.
type SnsTopicPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SnsTopicPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnsTopicPolicy is the Schema for the SnsTopicPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnsTopicPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicPolicySpec   `json:"spec"`
	Status            SnsTopicPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnsTopicPolicyList contains a list of SnsTopicPolicys
type SnsTopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnsTopicPolicy `json:"items"`
}

// Repository type metadata.
var (
	SnsTopicPolicyKind             = "SnsTopicPolicy"
	SnsTopicPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SnsTopicPolicyKind}.String()
	SnsTopicPolicyKindAPIVersion   = SnsTopicPolicyKind + "." + v1alpha1.GroupVersion.String()
	SnsTopicPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(SnsTopicPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SnsTopicPolicy{}, &SnsTopicPolicyList{})
}
