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
// +groupName=cognito.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cognito/v1alpha1"
)

type CognitoIdentityPoolObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CognitoIdentityPoolParameters struct {
	AllowClassicFlow *bool `json:"allowClassicFlow,omitempty" tf:"allow_classic_flow"`

	AllowUnauthenticatedIdentities *bool `json:"allowUnauthenticatedIdentities,omitempty" tf:"allow_unauthenticated_identities"`

	CognitoIdentityProviders []CognitoIdentityProvidersParameters `json:"cognitoIdentityProviders,omitempty" tf:"cognito_identity_providers"`

	DeveloperProviderName *string `json:"developerProviderName,omitempty" tf:"developer_provider_name"`

	IdentityPoolName string `json:"identityPoolName" tf:"identity_pool_name"`

	OpenidConnectProviderArns []string `json:"openidConnectProviderArns,omitempty" tf:"openid_connect_provider_arns"`

	SamlProviderArns []string `json:"samlProviderArns,omitempty" tf:"saml_provider_arns"`

	SupportedLoginProviders map[string]string `json:"supportedLoginProviders,omitempty" tf:"supported_login_providers"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type CognitoIdentityProvidersObservation struct {
}

type CognitoIdentityProvidersParameters struct {
	ClientId *string `json:"clientId,omitempty" tf:"client_id"`

	ProviderName *string `json:"providerName,omitempty" tf:"provider_name"`

	ServerSideTokenCheck *bool `json:"serverSideTokenCheck,omitempty" tf:"server_side_token_check"`
}

// CognitoIdentityPoolSpec defines the desired state of CognitoIdentityPool
type CognitoIdentityPoolSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoIdentityPoolParameters `json:"forProvider"`
}

// CognitoIdentityPoolStatus defines the observed state of CognitoIdentityPool.
type CognitoIdentityPoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoIdentityPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPool is the Schema for the CognitoIdentityPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoIdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityPoolSpec   `json:"spec"`
	Status            CognitoIdentityPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityPoolList contains a list of CognitoIdentityPools
type CognitoIdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityPool `json:"items"`
}

// Repository type metadata.
var (
	CognitoIdentityPoolKind             = "CognitoIdentityPool"
	CognitoIdentityPoolGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CognitoIdentityPoolKind}.String()
	CognitoIdentityPoolKindAPIVersion   = CognitoIdentityPoolKind + "." + v1alpha1.GroupVersion.String()
	CognitoIdentityPoolGroupVersionKind = v1alpha1.GroupVersion.WithKind(CognitoIdentityPoolKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CognitoIdentityPool{}, &CognitoIdentityPoolList{})
}
