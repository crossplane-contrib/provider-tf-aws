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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SignatureValidityPeriodObservation struct {
}

type SignatureValidityPeriodParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type SigningProfileObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PlatformDisplayName *string `json:"platformDisplayName,omitempty" tf:"platform_display_name,omitempty"`

	RevocationRecord []SigningProfileRevocationRecordObservation `json:"revocationRecord,omitempty" tf:"revocation_record,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	VersionArn *string `json:"versionArn,omitempty" tf:"version_arn,omitempty"`
}

type SigningProfileParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PlatformID *string `json:"platformId" tf:"platform_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SignatureValidityPeriod []SignatureValidityPeriodParameters `json:"signatureValidityPeriod,omitempty" tf:"signature_validity_period,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SigningProfileRevocationRecordObservation struct {
	RevocationEffectiveFrom *string `json:"revocationEffectiveFrom,omitempty" tf:"revocation_effective_from,omitempty"`

	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	RevokedBy *string `json:"revokedBy,omitempty" tf:"revoked_by,omitempty"`
}

type SigningProfileRevocationRecordParameters struct {
}

// SigningProfileSpec defines the desired state of SigningProfile
type SigningProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SigningProfileParameters `json:"forProvider"`
}

// SigningProfileStatus defines the observed state of SigningProfile.
type SigningProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SigningProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfile is the Schema for the SigningProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SigningProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SigningProfileSpec   `json:"spec"`
	Status            SigningProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SigningProfileList contains a list of SigningProfiles
type SigningProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningProfile `json:"items"`
}

// Repository type metadata.
var (
	SigningProfile_Kind             = "SigningProfile"
	SigningProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SigningProfile_Kind}.String()
	SigningProfile_KindAPIVersion   = SigningProfile_Kind + "." + CRDGroupVersion.String()
	SigningProfile_GroupVersionKind = CRDGroupVersion.WithKind(SigningProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&SigningProfile{}, &SigningProfileList{})
}
