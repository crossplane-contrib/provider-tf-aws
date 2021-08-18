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
// +groupName=kms.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1"
)

type KmsAliasObservation struct {
	Arn string `json:"arn" tf:"arn"`

	TargetKeyArn string `json:"targetKeyArn" tf:"target_key_arn"`
}

type KmsAliasParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	TargetKeyId string `json:"targetKeyId" tf:"target_key_id"`
}

// KmsAliasSpec defines the desired state of KmsAlias
type KmsAliasSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KmsAliasParameters `json:"forProvider"`
}

// KmsAliasStatus defines the observed state of KmsAlias.
type KmsAliasStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KmsAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KmsAlias is the Schema for the KmsAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KmsAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsAliasSpec   `json:"spec"`
	Status            KmsAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsAliasList contains a list of KmsAliass
type KmsAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsAlias `json:"items"`
}

// Repository type metadata.
var (
	KmsAliasKind             = "KmsAlias"
	KmsAliasGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: KmsAliasKind}.String()
	KmsAliasKindAPIVersion   = KmsAliasKind + "." + v1alpha1.GroupVersion.String()
	KmsAliasGroupVersionKind = v1alpha1.GroupVersion.WithKind(KmsAliasKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&KmsAlias{}, &KmsAliasList{})
}
