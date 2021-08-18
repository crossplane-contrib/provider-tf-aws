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
// +groupName=apigatewayv2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/apigatewayv2/v1alpha1"
)

type Apigatewayv2DeploymentObservation struct {
	AutoDeployed bool `json:"autoDeployed" tf:"auto_deployed"`
}

type Apigatewayv2DeploymentParameters struct {
	ApiId string `json:"apiId" tf:"api_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Triggers map[string]string `json:"triggers,omitempty" tf:"triggers"`
}

// Apigatewayv2DeploymentSpec defines the desired state of Apigatewayv2Deployment
type Apigatewayv2DeploymentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2DeploymentParameters `json:"forProvider"`
}

// Apigatewayv2DeploymentStatus defines the observed state of Apigatewayv2Deployment.
type Apigatewayv2DeploymentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Deployment is the Schema for the Apigatewayv2Deployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Apigatewayv2Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2DeploymentSpec   `json:"spec"`
	Status            Apigatewayv2DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2DeploymentList contains a list of Apigatewayv2Deployments
type Apigatewayv2DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Deployment `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2DeploymentKind             = "Apigatewayv2Deployment"
	Apigatewayv2DeploymentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Apigatewayv2DeploymentKind}.String()
	Apigatewayv2DeploymentKindAPIVersion   = Apigatewayv2DeploymentKind + "." + v1alpha1.GroupVersion.String()
	Apigatewayv2DeploymentGroupVersionKind = v1alpha1.GroupVersion.WithKind(Apigatewayv2DeploymentKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Apigatewayv2Deployment{}, &Apigatewayv2DeploymentList{})
}
