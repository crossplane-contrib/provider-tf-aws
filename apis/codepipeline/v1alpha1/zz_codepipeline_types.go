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

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// +kubebuilder:validation:Optional
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// +kubebuilder:validation:Required
	Provider *string `json:"provider" tf:"provider,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ArtifactStoreObservation struct {
}

type ArtifactStoreParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionKey []EncryptionKeyParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CodepipelineObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CodepipelineParameters struct {

	// +kubebuilder:validation:Required
	ArtifactStore []ArtifactStoreParameters `json:"artifactStore" tf:"artifact_store,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Stage []StageParameters `json:"stage" tf:"stage,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EncryptionKeyObservation struct {
}

type EncryptionKeyParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type StageObservation struct {
}

type StageParameters struct {

	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// CodepipelineSpec defines the desired state of Codepipeline
type CodepipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodepipelineParameters `json:"forProvider"`
}

// CodepipelineStatus defines the observed state of Codepipeline.
type CodepipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodepipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Codepipeline is the Schema for the Codepipelines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Codepipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodepipelineSpec   `json:"spec"`
	Status            CodepipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineList contains a list of Codepipelines
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Codepipeline `json:"items"`
}

// Repository type metadata.
var (
	Codepipeline_Kind             = "Codepipeline"
	Codepipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Codepipeline_Kind}.String()
	Codepipeline_KindAPIVersion   = Codepipeline_Kind + "." + CRDGroupVersion.String()
	Codepipeline_GroupVersionKind = CRDGroupVersion.WithKind(Codepipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&Codepipeline{}, &CodepipelineList{})
}
