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
// +groupName=lightsail.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lightsail/v1alpha1"
)

type LightsailKeyPairObservation struct {
	Arn string `json:"arn" tf:"arn"`

	EncryptedFingerprint string `json:"encryptedFingerprint" tf:"encrypted_fingerprint"`

	EncryptedPrivateKey string `json:"encryptedPrivateKey" tf:"encrypted_private_key"`

	Fingerprint string `json:"fingerprint" tf:"fingerprint"`

	PrivateKey string `json:"privateKey" tf:"private_key"`
}

type LightsailKeyPairParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key"`

	PublicKey *string `json:"publicKey,omitempty" tf:"public_key"`
}

// LightsailKeyPairSpec defines the desired state of LightsailKeyPair
type LightsailKeyPairSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LightsailKeyPairParameters `json:"forProvider"`
}

// LightsailKeyPairStatus defines the observed state of LightsailKeyPair.
type LightsailKeyPairStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LightsailKeyPairObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailKeyPair is the Schema for the LightsailKeyPairs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailKeyPairSpec   `json:"spec"`
	Status            LightsailKeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailKeyPairList contains a list of LightsailKeyPairs
type LightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailKeyPair `json:"items"`
}

// Repository type metadata.
var (
	LightsailKeyPairKind             = "LightsailKeyPair"
	LightsailKeyPairGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LightsailKeyPairKind}.String()
	LightsailKeyPairKindAPIVersion   = LightsailKeyPairKind + "." + v1alpha1.GroupVersion.String()
	LightsailKeyPairGroupVersionKind = v1alpha1.GroupVersion.WithKind(LightsailKeyPairKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LightsailKeyPair{}, &LightsailKeyPairList{})
}
