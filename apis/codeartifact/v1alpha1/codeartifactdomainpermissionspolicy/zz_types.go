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
// +groupName=codeartifact.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/codeartifact/v1alpha1"
)

type CodeartifactDomainPermissionsPolicyObservation struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type CodeartifactDomainPermissionsPolicyParameters struct {
	Domain string `json:"domain" tf:"domain"`

	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner"`

	PolicyDocument string `json:"policyDocument" tf:"policy_document"`

	PolicyRevision *string `json:"policyRevision,omitempty" tf:"policy_revision"`
}

// CodeartifactDomainPermissionsPolicySpec defines the desired state of CodeartifactDomainPermissionsPolicy
type CodeartifactDomainPermissionsPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodeartifactDomainPermissionsPolicyParameters `json:"forProvider"`
}

// CodeartifactDomainPermissionsPolicyStatus defines the observed state of CodeartifactDomainPermissionsPolicy.
type CodeartifactDomainPermissionsPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodeartifactDomainPermissionsPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactDomainPermissionsPolicy is the Schema for the CodeartifactDomainPermissionsPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodeartifactDomainPermissionsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodeartifactDomainPermissionsPolicySpec   `json:"spec"`
	Status            CodeartifactDomainPermissionsPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactDomainPermissionsPolicyList contains a list of CodeartifactDomainPermissionsPolicys
type CodeartifactDomainPermissionsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeartifactDomainPermissionsPolicy `json:"items"`
}

// Repository type metadata.
var (
	CodeartifactDomainPermissionsPolicyKind             = "CodeartifactDomainPermissionsPolicy"
	CodeartifactDomainPermissionsPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CodeartifactDomainPermissionsPolicyKind}.String()
	CodeartifactDomainPermissionsPolicyKindAPIVersion   = CodeartifactDomainPermissionsPolicyKind + "." + v1alpha1.GroupVersion.String()
	CodeartifactDomainPermissionsPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(CodeartifactDomainPermissionsPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CodeartifactDomainPermissionsPolicy{}, &CodeartifactDomainPermissionsPolicyList{})
}
