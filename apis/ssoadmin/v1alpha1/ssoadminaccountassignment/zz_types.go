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
// +groupName=ssoadmin.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ssoadmin/v1alpha1"
)

type SsoadminAccountAssignmentObservation struct {
}

type SsoadminAccountAssignmentParameters struct {
	InstanceArn string `json:"instanceArn" tf:"instance_arn"`

	PermissionSetArn string `json:"permissionSetArn" tf:"permission_set_arn"`

	PrincipalId string `json:"principalId" tf:"principal_id"`

	PrincipalType string `json:"principalType" tf:"principal_type"`

	TargetId string `json:"targetId" tf:"target_id"`

	TargetType *string `json:"targetType,omitempty" tf:"target_type"`
}

// SsoadminAccountAssignmentSpec defines the desired state of SsoadminAccountAssignment
type SsoadminAccountAssignmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsoadminAccountAssignmentParameters `json:"forProvider"`
}

// SsoadminAccountAssignmentStatus defines the observed state of SsoadminAccountAssignment.
type SsoadminAccountAssignmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsoadminAccountAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminAccountAssignment is the Schema for the SsoadminAccountAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SsoadminAccountAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsoadminAccountAssignmentSpec   `json:"spec"`
	Status            SsoadminAccountAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminAccountAssignmentList contains a list of SsoadminAccountAssignments
type SsoadminAccountAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsoadminAccountAssignment `json:"items"`
}

// Repository type metadata.
var (
	SsoadminAccountAssignmentKind             = "SsoadminAccountAssignment"
	SsoadminAccountAssignmentGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SsoadminAccountAssignmentKind}.String()
	SsoadminAccountAssignmentKindAPIVersion   = SsoadminAccountAssignmentKind + "." + v1alpha1.GroupVersion.String()
	SsoadminAccountAssignmentGroupVersionKind = v1alpha1.GroupVersion.WithKind(SsoadminAccountAssignmentKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SsoadminAccountAssignment{}, &SsoadminAccountAssignmentList{})
}
