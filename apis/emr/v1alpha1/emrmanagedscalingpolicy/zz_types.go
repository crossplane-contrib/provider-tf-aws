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
// +groupName=emr.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/emr/v1alpha1"
)

type ComputeLimitsObservation struct {
}

type ComputeLimitsParameters struct {
	MaximumCapacityUnits int64 `json:"maximumCapacityUnits" tf:"maximum_capacity_units"`

	MaximumCoreCapacityUnits *int64 `json:"maximumCoreCapacityUnits,omitempty" tf:"maximum_core_capacity_units"`

	MaximumOndemandCapacityUnits *int64 `json:"maximumOndemandCapacityUnits,omitempty" tf:"maximum_ondemand_capacity_units"`

	MinimumCapacityUnits int64 `json:"minimumCapacityUnits" tf:"minimum_capacity_units"`

	UnitType string `json:"unitType" tf:"unit_type"`
}

type EmrManagedScalingPolicyObservation struct {
}

type EmrManagedScalingPolicyParameters struct {
	ClusterId string `json:"clusterId" tf:"cluster_id"`

	ComputeLimits []ComputeLimitsParameters `json:"computeLimits" tf:"compute_limits"`
}

// EmrManagedScalingPolicySpec defines the desired state of EmrManagedScalingPolicy
type EmrManagedScalingPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmrManagedScalingPolicyParameters `json:"forProvider"`
}

// EmrManagedScalingPolicyStatus defines the observed state of EmrManagedScalingPolicy.
type EmrManagedScalingPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmrManagedScalingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmrManagedScalingPolicy is the Schema for the EmrManagedScalingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmrManagedScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrManagedScalingPolicySpec   `json:"spec"`
	Status            EmrManagedScalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrManagedScalingPolicyList contains a list of EmrManagedScalingPolicys
type EmrManagedScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrManagedScalingPolicy `json:"items"`
}

// Repository type metadata.
var (
	EmrManagedScalingPolicyKind             = "EmrManagedScalingPolicy"
	EmrManagedScalingPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EmrManagedScalingPolicyKind}.String()
	EmrManagedScalingPolicyKindAPIVersion   = EmrManagedScalingPolicyKind + "." + v1alpha1.GroupVersion.String()
	EmrManagedScalingPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(EmrManagedScalingPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EmrManagedScalingPolicy{}, &EmrManagedScalingPolicyList{})
}
