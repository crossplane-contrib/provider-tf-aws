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
// +groupName=cloudfront.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudfront/v1alpha1"
)

type CloudfrontOriginRequestPolicyObservation struct {
}

type CloudfrontOriginRequestPolicyParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	CookiesConfig []CookiesConfigParameters `json:"cookiesConfig" tf:"cookies_config"`

	Etag *string `json:"etag,omitempty" tf:"etag"`

	HeadersConfig []HeadersConfigParameters `json:"headersConfig" tf:"headers_config"`

	Name string `json:"name" tf:"name"`

	QueryStringsConfig []QueryStringsConfigParameters `json:"queryStringsConfig" tf:"query_strings_config"`
}

type CookiesConfigObservation struct {
}

type CookiesConfigParameters struct {
	CookieBehavior string `json:"cookieBehavior" tf:"cookie_behavior"`

	Cookies []CookiesParameters `json:"cookies,omitempty" tf:"cookies"`
}

type CookiesObservation struct {
}

type CookiesParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

type HeadersConfigObservation struct {
}

type HeadersConfigParameters struct {
	HeaderBehavior *string `json:"headerBehavior,omitempty" tf:"header_behavior"`

	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers"`
}

type HeadersObservation struct {
}

type HeadersParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

type QueryStringsConfigObservation struct {
}

type QueryStringsConfigParameters struct {
	QueryStringBehavior string `json:"queryStringBehavior" tf:"query_string_behavior"`

	QueryStrings []QueryStringsParameters `json:"queryStrings,omitempty" tf:"query_strings"`
}

type QueryStringsObservation struct {
}

type QueryStringsParameters struct {
	Items []string `json:"items,omitempty" tf:"items"`
}

// CloudfrontOriginRequestPolicySpec defines the desired state of CloudfrontOriginRequestPolicy
type CloudfrontOriginRequestPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudfrontOriginRequestPolicyParameters `json:"forProvider"`
}

// CloudfrontOriginRequestPolicyStatus defines the observed state of CloudfrontOriginRequestPolicy.
type CloudfrontOriginRequestPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudfrontOriginRequestPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontOriginRequestPolicy is the Schema for the CloudfrontOriginRequestPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudfrontOriginRequestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontOriginRequestPolicySpec   `json:"spec"`
	Status            CloudfrontOriginRequestPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontOriginRequestPolicyList contains a list of CloudfrontOriginRequestPolicys
type CloudfrontOriginRequestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontOriginRequestPolicy `json:"items"`
}

// Repository type metadata.
var (
	CloudfrontOriginRequestPolicyKind             = "CloudfrontOriginRequestPolicy"
	CloudfrontOriginRequestPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudfrontOriginRequestPolicyKind}.String()
	CloudfrontOriginRequestPolicyKindAPIVersion   = CloudfrontOriginRequestPolicyKind + "." + v1alpha1.GroupVersion.String()
	CloudfrontOriginRequestPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudfrontOriginRequestPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudfrontOriginRequestPolicy{}, &CloudfrontOriginRequestPolicyList{})
}
