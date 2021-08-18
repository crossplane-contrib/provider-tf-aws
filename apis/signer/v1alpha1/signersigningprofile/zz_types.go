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
// +groupName=signer.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/signer/v1alpha1"
)

type RevocationRecordObservation struct {
	RevocationEffectiveFrom string `json:"revocationEffectiveFrom" tf:"revocation_effective_from"`

	RevokedAt string `json:"revokedAt" tf:"revoked_at"`

	RevokedBy string `json:"revokedBy" tf:"revoked_by"`
}

type RevocationRecordParameters struct {
}

type SignatureValidityPeriodObservation struct {
}

type SignatureValidityPeriodParameters struct {
	Type string `json:"type" tf:"type"`

	Value int64 `json:"value" tf:"value"`
}

type SignerSigningProfileObservation struct {
	Arn string `json:"arn" tf:"arn"`

	PlatformDisplayName string `json:"platformDisplayName" tf:"platform_display_name"`

	RevocationRecord []RevocationRecordObservation `json:"revocationRecord" tf:"revocation_record"`

	Status string `json:"status" tf:"status"`

	Version string `json:"version" tf:"version"`

	VersionArn string `json:"versionArn" tf:"version_arn"`
}

type SignerSigningProfileParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	PlatformId string `json:"platformId" tf:"platform_id"`

	SignatureValidityPeriod []SignatureValidityPeriodParameters `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SignerSigningProfileSpec defines the desired state of SignerSigningProfile
type SignerSigningProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SignerSigningProfileParameters `json:"forProvider"`
}

// SignerSigningProfileStatus defines the observed state of SignerSigningProfile.
type SignerSigningProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SignerSigningProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SignerSigningProfile is the Schema for the SignerSigningProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SignerSigningProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignerSigningProfileSpec   `json:"spec"`
	Status            SignerSigningProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SignerSigningProfileList contains a list of SignerSigningProfiles
type SignerSigningProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignerSigningProfile `json:"items"`
}

// Repository type metadata.
var (
	SignerSigningProfileKind             = "SignerSigningProfile"
	SignerSigningProfileGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SignerSigningProfileKind}.String()
	SignerSigningProfileKindAPIVersion   = SignerSigningProfileKind + "." + v1alpha1.GroupVersion.String()
	SignerSigningProfileGroupVersionKind = v1alpha1.GroupVersion.WithKind(SignerSigningProfileKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SignerSigningProfile{}, &SignerSigningProfileList{})
}
