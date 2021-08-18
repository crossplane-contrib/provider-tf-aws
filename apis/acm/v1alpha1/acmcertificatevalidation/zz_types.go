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
// +groupName=acm.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/acm/v1alpha1"
)

type AcmCertificateValidationObservation struct {
}

type AcmCertificateValidationParameters struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`

	ValidationRecordFqdns []string `json:"validationRecordFqdns,omitempty" tf:"validation_record_fqdns"`
}

// AcmCertificateValidationSpec defines the desired state of AcmCertificateValidation
type AcmCertificateValidationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AcmCertificateValidationParameters `json:"forProvider"`
}

// AcmCertificateValidationStatus defines the observed state of AcmCertificateValidation.
type AcmCertificateValidationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AcmCertificateValidationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AcmCertificateValidation is the Schema for the AcmCertificateValidations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AcmCertificateValidation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmCertificateValidationSpec   `json:"spec"`
	Status            AcmCertificateValidationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmCertificateValidationList contains a list of AcmCertificateValidations
type AcmCertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmCertificateValidation `json:"items"`
}

// Repository type metadata.
var (
	AcmCertificateValidationKind             = "AcmCertificateValidation"
	AcmCertificateValidationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AcmCertificateValidationKind}.String()
	AcmCertificateValidationKindAPIVersion   = AcmCertificateValidationKind + "." + v1alpha1.GroupVersion.String()
	AcmCertificateValidationGroupVersionKind = v1alpha1.GroupVersion.WithKind(AcmCertificateValidationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&AcmCertificateValidation{}, &AcmCertificateValidationList{})
}
