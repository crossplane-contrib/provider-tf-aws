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
// +groupName=glue.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/glue/v1alpha1"
)

type GlueUserDefinedFunctionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateTime string `json:"createTime" tf:"create_time"`
}

type GlueUserDefinedFunctionParameters struct {
	CatalogId *string `json:"catalogId,omitempty" tf:"catalog_id"`

	ClassName string `json:"className" tf:"class_name"`

	DatabaseName string `json:"databaseName" tf:"database_name"`

	Name string `json:"name" tf:"name"`

	OwnerName string `json:"ownerName" tf:"owner_name"`

	OwnerType string `json:"ownerType" tf:"owner_type"`

	ResourceUris []ResourceUrisParameters `json:"resourceUris,omitempty" tf:"resource_uris"`
}

type ResourceUrisObservation struct {
}

type ResourceUrisParameters struct {
	ResourceType string `json:"resourceType" tf:"resource_type"`

	Uri string `json:"uri" tf:"uri"`
}

// GlueUserDefinedFunctionSpec defines the desired state of GlueUserDefinedFunction
type GlueUserDefinedFunctionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueUserDefinedFunctionParameters `json:"forProvider"`
}

// GlueUserDefinedFunctionStatus defines the observed state of GlueUserDefinedFunction.
type GlueUserDefinedFunctionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueUserDefinedFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueUserDefinedFunction is the Schema for the GlueUserDefinedFunctions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlueUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueUserDefinedFunctionSpec   `json:"spec"`
	Status            GlueUserDefinedFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueUserDefinedFunctionList contains a list of GlueUserDefinedFunctions
type GlueUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueUserDefinedFunction `json:"items"`
}

// Repository type metadata.
var (
	GlueUserDefinedFunctionKind             = "GlueUserDefinedFunction"
	GlueUserDefinedFunctionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: GlueUserDefinedFunctionKind}.String()
	GlueUserDefinedFunctionKindAPIVersion   = GlueUserDefinedFunctionKind + "." + v1alpha1.GroupVersion.String()
	GlueUserDefinedFunctionGroupVersionKind = v1alpha1.GroupVersion.WithKind(GlueUserDefinedFunctionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&GlueUserDefinedFunction{}, &GlueUserDefinedFunctionList{})
}
