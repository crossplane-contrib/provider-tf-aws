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
// +groupName=imagebuilder.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/imagebuilder/v1alpha1"
)

type ImageTestsConfigurationObservation struct {
}

type ImageTestsConfigurationParameters struct {
	ImageTestsEnabled *bool `json:"imageTestsEnabled,omitempty" tf:"image_tests_enabled"`

	TimeoutMinutes *int64 `json:"timeoutMinutes,omitempty" tf:"timeout_minutes"`
}

type ImagebuilderImagePipelineObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DateCreated string `json:"dateCreated" tf:"date_created"`

	DateLastRun string `json:"dateLastRun" tf:"date_last_run"`

	DateNextRun string `json:"dateNextRun" tf:"date_next_run"`

	DateUpdated string `json:"dateUpdated" tf:"date_updated"`

	Platform string `json:"platform" tf:"platform"`
}

type ImagebuilderImagePipelineParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DistributionConfigurationArn *string `json:"distributionConfigurationArn,omitempty" tf:"distribution_configuration_arn"`

	EnhancedImageMetadataEnabled *bool `json:"enhancedImageMetadataEnabled,omitempty" tf:"enhanced_image_metadata_enabled"`

	ImageRecipeArn string `json:"imageRecipeArn" tf:"image_recipe_arn"`

	ImageTestsConfiguration []ImageTestsConfigurationParameters `json:"imageTestsConfiguration,omitempty" tf:"image_tests_configuration"`

	InfrastructureConfigurationArn string `json:"infrastructureConfigurationArn" tf:"infrastructure_configuration_arn"`

	Name string `json:"name" tf:"name"`

	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule"`

	Status *string `json:"status,omitempty" tf:"status"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {
	PipelineExecutionStartCondition *string `json:"pipelineExecutionStartCondition,omitempty" tf:"pipeline_execution_start_condition"`

	ScheduleExpression string `json:"scheduleExpression" tf:"schedule_expression"`
}

// ImagebuilderImagePipelineSpec defines the desired state of ImagebuilderImagePipeline
type ImagebuilderImagePipelineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ImagebuilderImagePipelineParameters `json:"forProvider"`
}

// ImagebuilderImagePipelineStatus defines the observed state of ImagebuilderImagePipeline.
type ImagebuilderImagePipelineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ImagebuilderImagePipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImagebuilderImagePipeline is the Schema for the ImagebuilderImagePipelines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ImagebuilderImagePipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImagebuilderImagePipelineSpec   `json:"spec"`
	Status            ImagebuilderImagePipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImagebuilderImagePipelineList contains a list of ImagebuilderImagePipelines
type ImagebuilderImagePipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagebuilderImagePipeline `json:"items"`
}

// Repository type metadata.
var (
	ImagebuilderImagePipelineKind             = "ImagebuilderImagePipeline"
	ImagebuilderImagePipelineGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ImagebuilderImagePipelineKind}.String()
	ImagebuilderImagePipelineKindAPIVersion   = ImagebuilderImagePipelineKind + "." + v1alpha1.GroupVersion.String()
	ImagebuilderImagePipelineGroupVersionKind = v1alpha1.GroupVersion.WithKind(ImagebuilderImagePipelineKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ImagebuilderImagePipeline{}, &ImagebuilderImagePipelineList{})
}
