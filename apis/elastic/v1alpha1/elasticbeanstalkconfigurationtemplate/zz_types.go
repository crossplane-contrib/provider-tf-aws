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
// +groupName=elastic.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/elastic/v1alpha1"
)

type ElasticBeanstalkConfigurationTemplateObservation struct {
}

type ElasticBeanstalkConfigurationTemplateParameters struct {
	Application string `json:"application" tf:"application"`

	Description *string `json:"description,omitempty" tf:"description"`

	EnvironmentId *string `json:"environmentId,omitempty" tf:"environment_id"`

	Name string `json:"name" tf:"name"`

	Setting []SettingParameters `json:"setting,omitempty" tf:"setting"`

	SolutionStackName *string `json:"solutionStackName,omitempty" tf:"solution_stack_name"`
}

type SettingObservation struct {
}

type SettingParameters struct {
	Name string `json:"name" tf:"name"`

	Namespace string `json:"namespace" tf:"namespace"`

	Resource *string `json:"resource,omitempty" tf:"resource"`

	Value string `json:"value" tf:"value"`
}

// ElasticBeanstalkConfigurationTemplateSpec defines the desired state of ElasticBeanstalkConfigurationTemplate
type ElasticBeanstalkConfigurationTemplateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticBeanstalkConfigurationTemplateParameters `json:"forProvider"`
}

// ElasticBeanstalkConfigurationTemplateStatus defines the observed state of ElasticBeanstalkConfigurationTemplate.
type ElasticBeanstalkConfigurationTemplateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticBeanstalkConfigurationTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkConfigurationTemplate is the Schema for the ElasticBeanstalkConfigurationTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ElasticBeanstalkConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkConfigurationTemplateSpec   `json:"spec"`
	Status            ElasticBeanstalkConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkConfigurationTemplateList contains a list of ElasticBeanstalkConfigurationTemplates
type ElasticBeanstalkConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkConfigurationTemplate `json:"items"`
}

// Repository type metadata.
var (
	ElasticBeanstalkConfigurationTemplateKind             = "ElasticBeanstalkConfigurationTemplate"
	ElasticBeanstalkConfigurationTemplateGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ElasticBeanstalkConfigurationTemplateKind}.String()
	ElasticBeanstalkConfigurationTemplateKindAPIVersion   = ElasticBeanstalkConfigurationTemplateKind + "." + v1alpha1.GroupVersion.String()
	ElasticBeanstalkConfigurationTemplateGroupVersionKind = v1alpha1.GroupVersion.WithKind(ElasticBeanstalkConfigurationTemplateKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ElasticBeanstalkConfigurationTemplate{}, &ElasticBeanstalkConfigurationTemplateList{})
}
