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

type WebsiteCertificateAuthorityAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	WebsiteCAID *string `json:"websiteCaId,omitempty" tf:"website_ca_id,omitempty"`
}

type WebsiteCertificateAuthorityAssociationParameters struct {

	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	FleetArn *string `json:"fleetArn" tf:"fleet_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// WebsiteCertificateAuthorityAssociationSpec defines the desired state of WebsiteCertificateAuthorityAssociation
type WebsiteCertificateAuthorityAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebsiteCertificateAuthorityAssociationParameters `json:"forProvider"`
}

// WebsiteCertificateAuthorityAssociationStatus defines the observed state of WebsiteCertificateAuthorityAssociation.
type WebsiteCertificateAuthorityAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebsiteCertificateAuthorityAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebsiteCertificateAuthorityAssociation is the Schema for the WebsiteCertificateAuthorityAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type WebsiteCertificateAuthorityAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebsiteCertificateAuthorityAssociationSpec   `json:"spec"`
	Status            WebsiteCertificateAuthorityAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebsiteCertificateAuthorityAssociationList contains a list of WebsiteCertificateAuthorityAssociations
type WebsiteCertificateAuthorityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebsiteCertificateAuthorityAssociation `json:"items"`
}

// Repository type metadata.
var (
	WebsiteCertificateAuthorityAssociation_Kind             = "WebsiteCertificateAuthorityAssociation"
	WebsiteCertificateAuthorityAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebsiteCertificateAuthorityAssociation_Kind}.String()
	WebsiteCertificateAuthorityAssociation_KindAPIVersion   = WebsiteCertificateAuthorityAssociation_Kind + "." + CRDGroupVersion.String()
	WebsiteCertificateAuthorityAssociation_GroupVersionKind = CRDGroupVersion.WithKind(WebsiteCertificateAuthorityAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&WebsiteCertificateAuthorityAssociation{}, &WebsiteCertificateAuthorityAssociationList{})
}
