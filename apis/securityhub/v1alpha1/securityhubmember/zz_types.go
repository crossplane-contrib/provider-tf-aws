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
// +groupName=securityhub.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/securityhub/v1alpha1"
)

type SecurityhubMemberObservation struct {
	MasterId string `json:"masterId" tf:"master_id"`

	MemberStatus string `json:"memberStatus" tf:"member_status"`
}

type SecurityhubMemberParameters struct {
	AccountId string `json:"accountId" tf:"account_id"`

	Email string `json:"email" tf:"email"`

	Invite *bool `json:"invite,omitempty" tf:"invite"`
}

// SecurityhubMemberSpec defines the desired state of SecurityhubMember
type SecurityhubMemberSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecurityhubMemberParameters `json:"forProvider"`
}

// SecurityhubMemberStatus defines the observed state of SecurityhubMember.
type SecurityhubMemberStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecurityhubMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubMember is the Schema for the SecurityhubMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityhubMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityhubMemberSpec   `json:"spec"`
	Status            SecurityhubMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubMemberList contains a list of SecurityhubMembers
type SecurityhubMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityhubMember `json:"items"`
}

// Repository type metadata.
var (
	SecurityhubMemberKind             = "SecurityhubMember"
	SecurityhubMemberGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SecurityhubMemberKind}.String()
	SecurityhubMemberKindAPIVersion   = SecurityhubMemberKind + "." + v1alpha1.GroupVersion.String()
	SecurityhubMemberGroupVersionKind = v1alpha1.GroupVersion.WithKind(SecurityhubMemberKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SecurityhubMember{}, &SecurityhubMemberList{})
}
