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
// +groupName=ec2.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ec2/v1alpha1"
)

type CapacityRebalanceObservation struct {
}

type CapacityRebalanceParameters struct {
	ReplacementStrategy *string `json:"replacementStrategy,omitempty" tf:"replacement_strategy"`
}

type Ec2FleetObservation struct {
}

type Ec2FleetParameters struct {
	ExcessCapacityTerminationPolicy *string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy"`

	LaunchTemplateConfig []LaunchTemplateConfigParameters `json:"launchTemplateConfig" tf:"launch_template_config"`

	OnDemandOptions []OnDemandOptionsParameters `json:"onDemandOptions,omitempty" tf:"on_demand_options"`

	ReplaceUnhealthyInstances *bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances"`

	SpotOptions []SpotOptionsParameters `json:"spotOptions,omitempty" tf:"spot_options"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetCapacitySpecification []TargetCapacitySpecificationParameters `json:"targetCapacitySpecification" tf:"target_capacity_specification"`

	TerminateInstances *bool `json:"terminateInstances,omitempty" tf:"terminate_instances"`

	TerminateInstancesWithExpiration *bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type LaunchTemplateConfigObservation struct {
}

type LaunchTemplateConfigParameters struct {
	LaunchTemplateSpecification []LaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification"`

	Override []OverrideParameters `json:"override,omitempty" tf:"override"`
}

type LaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateSpecificationParameters struct {
	LaunchTemplateId *string `json:"launchTemplateId,omitempty" tf:"launch_template_id"`

	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name"`

	Version string `json:"version" tf:"version"`
}

type MaintenanceStrategiesObservation struct {
}

type MaintenanceStrategiesParameters struct {
	CapacityRebalance []CapacityRebalanceParameters `json:"capacityRebalance,omitempty" tf:"capacity_rebalance"`
}

type OnDemandOptionsObservation struct {
}

type OnDemandOptionsParameters struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy"`
}

type OverrideObservation struct {
}

type OverrideParameters struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price"`

	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type SpotOptionsObservation struct {
}

type SpotOptionsParameters struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy"`

	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`

	InstancePoolsToUseCount *int64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count"`

	MaintenanceStrategies []MaintenanceStrategiesParameters `json:"maintenanceStrategies,omitempty" tf:"maintenance_strategies"`
}

type TargetCapacitySpecificationObservation struct {
}

type TargetCapacitySpecificationParameters struct {
	DefaultTargetCapacityType string `json:"defaultTargetCapacityType" tf:"default_target_capacity_type"`

	OnDemandTargetCapacity *int64 `json:"onDemandTargetCapacity,omitempty" tf:"on_demand_target_capacity"`

	SpotTargetCapacity *int64 `json:"spotTargetCapacity,omitempty" tf:"spot_target_capacity"`

	TotalTargetCapacity int64 `json:"totalTargetCapacity" tf:"total_target_capacity"`
}

// Ec2FleetSpec defines the desired state of Ec2Fleet
type Ec2FleetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2FleetParameters `json:"forProvider"`
}

// Ec2FleetStatus defines the observed state of Ec2Fleet.
type Ec2FleetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2FleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2Fleet is the Schema for the Ec2Fleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2FleetSpec   `json:"spec"`
	Status            Ec2FleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2FleetList contains a list of Ec2Fleets
type Ec2FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2Fleet `json:"items"`
}

// Repository type metadata.
var (
	Ec2FleetKind             = "Ec2Fleet"
	Ec2FleetGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2FleetKind}.String()
	Ec2FleetKindAPIVersion   = Ec2FleetKind + "." + v1alpha1.GroupVersion.String()
	Ec2FleetGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2FleetKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2Fleet{}, &Ec2FleetList{})
}
