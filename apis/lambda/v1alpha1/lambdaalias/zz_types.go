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
// +groupName=lambda.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lambda/v1alpha1"
)

type LambdaAliasObservation struct {
	Arn string `json:"arn" tf:"arn"`

	InvokeArn string `json:"invokeArn" tf:"invoke_arn"`
}

type LambdaAliasParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	FunctionName string `json:"functionName" tf:"function_name"`

	FunctionVersion string `json:"functionVersion" tf:"function_version"`

	Name string `json:"name" tf:"name"`

	RoutingConfig []RoutingConfigParameters `json:"routingConfig,omitempty" tf:"routing_config"`
}

type RoutingConfigObservation struct {
}

type RoutingConfigParameters struct {
	AdditionalVersionWeights map[string]float64 `json:"additionalVersionWeights,omitempty" tf:"additional_version_weights"`
}

// LambdaAliasSpec defines the desired state of LambdaAlias
type LambdaAliasSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaAliasParameters `json:"forProvider"`
}

// LambdaAliasStatus defines the observed state of LambdaAlias.
type LambdaAliasStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaAlias is the Schema for the LambdaAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LambdaAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaAliasSpec   `json:"spec"`
	Status            LambdaAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaAliasList contains a list of LambdaAliass
type LambdaAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaAlias `json:"items"`
}

// Repository type metadata.
var (
	LambdaAliasKind             = "LambdaAlias"
	LambdaAliasGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LambdaAliasKind}.String()
	LambdaAliasKindAPIVersion   = LambdaAliasKind + "." + v1alpha1.GroupVersion.String()
	LambdaAliasGroupVersionKind = v1alpha1.GroupVersion.WithKind(LambdaAliasKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LambdaAlias{}, &LambdaAliasList{})
}
