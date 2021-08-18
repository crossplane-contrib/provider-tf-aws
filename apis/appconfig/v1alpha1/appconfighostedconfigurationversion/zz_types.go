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
// +groupName=appconfig.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/appconfig/v1alpha1"
)

type AppconfigHostedConfigurationVersionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	VersionNumber int64 `json:"versionNumber" tf:"version_number"`
}

type AppconfigHostedConfigurationVersionParameters struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	ConfigurationProfileId string `json:"configurationProfileId" tf:"configuration_profile_id"`

	Content string `json:"content" tf:"content"`

	ContentType string `json:"contentType" tf:"content_type"`

	Description *string `json:"description,omitempty" tf:"description"`
}

// AppconfigHostedConfigurationVersionSpec defines the desired state of AppconfigHostedConfigurationVersion
type AppconfigHostedConfigurationVersionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppconfigHostedConfigurationVersionParameters `json:"forProvider"`
}

// AppconfigHostedConfigurationVersionStatus defines the observed state of AppconfigHostedConfigurationVersion.
type AppconfigHostedConfigurationVersionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppconfigHostedConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigHostedConfigurationVersion is the Schema for the AppconfigHostedConfigurationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AppconfigHostedConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppconfigHostedConfigurationVersionSpec   `json:"spec"`
	Status            AppconfigHostedConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppconfigHostedConfigurationVersionList contains a list of AppconfigHostedConfigurationVersions
type AppconfigHostedConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppconfigHostedConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	AppconfigHostedConfigurationVersionKind             = "AppconfigHostedConfigurationVersion"
	AppconfigHostedConfigurationVersionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: AppconfigHostedConfigurationVersionKind}.String()
	AppconfigHostedConfigurationVersionKindAPIVersion   = AppconfigHostedConfigurationVersionKind + "." + v1alpha1.GroupVersion.String()
	AppconfigHostedConfigurationVersionGroupVersionKind = v1alpha1.GroupVersion.WithKind(AppconfigHostedConfigurationVersionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&AppconfigHostedConfigurationVersion{}, &AppconfigHostedConfigurationVersionList{})
}
