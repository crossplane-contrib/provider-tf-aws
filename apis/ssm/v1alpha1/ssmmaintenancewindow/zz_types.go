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
// +groupName=ssm.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ssm/v1alpha1"
)

type SsmMaintenanceWindowObservation struct {
}

type SsmMaintenanceWindowParameters struct {
	AllowUnassociatedTargets *bool `json:"allowUnassociatedTargets,omitempty" tf:"allow_unassociated_targets"`

	Cutoff int64 `json:"cutoff" tf:"cutoff"`

	Description *string `json:"description,omitempty" tf:"description"`

	Duration int64 `json:"duration" tf:"duration"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EndDate *string `json:"endDate,omitempty" tf:"end_date"`

	Name string `json:"name" tf:"name"`

	Schedule string `json:"schedule" tf:"schedule"`

	ScheduleOffset *int64 `json:"scheduleOffset,omitempty" tf:"schedule_offset"`

	ScheduleTimezone *string `json:"scheduleTimezone,omitempty" tf:"schedule_timezone"`

	StartDate *string `json:"startDate,omitempty" tf:"start_date"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SsmMaintenanceWindowSpec defines the desired state of SsmMaintenanceWindow
type SsmMaintenanceWindowSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsmMaintenanceWindowParameters `json:"forProvider"`
}

// SsmMaintenanceWindowStatus defines the observed state of SsmMaintenanceWindow.
type SsmMaintenanceWindowStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsmMaintenanceWindowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindow is the Schema for the SsmMaintenanceWindows API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SsmMaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowSpec   `json:"spec"`
	Status            SsmMaintenanceWindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowList contains a list of SsmMaintenanceWindows
type SsmMaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindow `json:"items"`
}

// Repository type metadata.
var (
	SsmMaintenanceWindowKind             = "SsmMaintenanceWindow"
	SsmMaintenanceWindowGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SsmMaintenanceWindowKind}.String()
	SsmMaintenanceWindowKindAPIVersion   = SsmMaintenanceWindowKind + "." + v1alpha1.GroupVersion.String()
	SsmMaintenanceWindowGroupVersionKind = v1alpha1.GroupVersion.WithKind(SsmMaintenanceWindowKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SsmMaintenanceWindow{}, &SsmMaintenanceWindowList{})
}
