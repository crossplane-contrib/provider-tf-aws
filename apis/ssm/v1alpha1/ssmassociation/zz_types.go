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

type OutputLocationObservation struct {
}

type OutputLocationParameters struct {
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`

	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix"`
}

type SsmAssociationObservation struct {
	AssociationId string `json:"associationId" tf:"association_id"`
}

type SsmAssociationParameters struct {
	ApplyOnlyAtCronInterval *bool `json:"applyOnlyAtCronInterval,omitempty" tf:"apply_only_at_cron_interval"`

	AssociationName *string `json:"associationName,omitempty" tf:"association_name"`

	AutomationTargetParameterName *string `json:"automationTargetParameterName,omitempty" tf:"automation_target_parameter_name"`

	ComplianceSeverity *string `json:"complianceSeverity,omitempty" tf:"compliance_severity"`

	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version"`

	InstanceId *string `json:"instanceId,omitempty" tf:"instance_id"`

	MaxConcurrency *string `json:"maxConcurrency,omitempty" tf:"max_concurrency"`

	MaxErrors *string `json:"maxErrors,omitempty" tf:"max_errors"`

	Name string `json:"name" tf:"name"`

	OutputLocation []OutputLocationParameters `json:"outputLocation,omitempty" tf:"output_location"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	ScheduleExpression *string `json:"scheduleExpression,omitempty" tf:"schedule_expression"`

	Targets []TargetsParameters `json:"targets,omitempty" tf:"targets"`
}

type TargetsObservation struct {
}

type TargetsParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

// SsmAssociationSpec defines the desired state of SsmAssociation
type SsmAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsmAssociationParameters `json:"forProvider"`
}

// SsmAssociationStatus defines the observed state of SsmAssociation.
type SsmAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsmAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsmAssociation is the Schema for the SsmAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SsmAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmAssociationSpec   `json:"spec"`
	Status            SsmAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmAssociationList contains a list of SsmAssociations
type SsmAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmAssociation `json:"items"`
}

// Repository type metadata.
var (
	SsmAssociationKind             = "SsmAssociation"
	SsmAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: SsmAssociationKind}.String()
	SsmAssociationKindAPIVersion   = SsmAssociationKind + "." + v1alpha1.GroupVersion.String()
	SsmAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(SsmAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&SsmAssociation{}, &SsmAssociationList{})
}
