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
// +groupName=api.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/api/v1alpha1"
)

type ApiGatewayRestApiPolicyObservation struct {
}

type ApiGatewayRestApiPolicyParameters struct {
	Policy string `json:"policy" tf:"policy"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`
}

// ApiGatewayRestApiPolicySpec defines the desired state of ApiGatewayRestApiPolicy
type ApiGatewayRestApiPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayRestApiPolicyParameters `json:"forProvider"`
}

// ApiGatewayRestApiPolicyStatus defines the observed state of ApiGatewayRestApiPolicy.
type ApiGatewayRestApiPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayRestApiPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayRestApiPolicy is the Schema for the ApiGatewayRestApiPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApiGatewayRestApiPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayRestApiPolicySpec   `json:"spec"`
	Status            ApiGatewayRestApiPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayRestApiPolicyList contains a list of ApiGatewayRestApiPolicys
type ApiGatewayRestApiPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayRestApiPolicy `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayRestApiPolicyKind             = "ApiGatewayRestApiPolicy"
	ApiGatewayRestApiPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ApiGatewayRestApiPolicyKind}.String()
	ApiGatewayRestApiPolicyKindAPIVersion   = ApiGatewayRestApiPolicyKind + "." + v1alpha1.GroupVersion.String()
	ApiGatewayRestApiPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(ApiGatewayRestApiPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ApiGatewayRestApiPolicy{}, &ApiGatewayRestApiPolicyList{})
}
