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

type DeploymentStrategyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeploymentStrategyParameters struct {

	// +kubebuilder:validation:Required
	DeploymentDurationInMinutes *float64 `json:"deploymentDurationInMinutes" tf:"deployment_duration_in_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FinalBakeTimeInMinutes *float64 `json:"finalBakeTimeInMinutes,omitempty" tf:"final_bake_time_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	GrowthFactor *float64 `json:"growthFactor" tf:"growth_factor,omitempty"`

	// +kubebuilder:validation:Optional
	GrowthType *string `json:"growthType,omitempty" tf:"growth_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ReplicateTo *string `json:"replicateTo" tf:"replicate_to,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DeploymentStrategySpec defines the desired state of DeploymentStrategy
type DeploymentStrategySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentStrategyParameters `json:"forProvider"`
}

// DeploymentStrategyStatus defines the observed state of DeploymentStrategy.
type DeploymentStrategyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentStrategyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentStrategy is the Schema for the DeploymentStrategys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DeploymentStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentStrategySpec   `json:"spec"`
	Status            DeploymentStrategyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentStrategyList contains a list of DeploymentStrategys
type DeploymentStrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentStrategy `json:"items"`
}

// Repository type metadata.
var (
	DeploymentStrategy_Kind             = "DeploymentStrategy"
	DeploymentStrategy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentStrategy_Kind}.String()
	DeploymentStrategy_KindAPIVersion   = DeploymentStrategy_Kind + "." + CRDGroupVersion.String()
	DeploymentStrategy_GroupVersionKind = CRDGroupVersion.WithKind(DeploymentStrategy_Kind)
)

func init() {
	SchemeBuilder.Register(&DeploymentStrategy{}, &DeploymentStrategyList{})
}
