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

type JobQueueObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type JobQueueParameters struct {

	// +kubebuilder:validation:Required
	ComputeEnvironments []*string `json:"computeEnvironments" tf:"compute_environments,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	State *string `json:"state" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// JobQueueSpec defines the desired state of JobQueue
type JobQueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobQueueParameters `json:"forProvider"`
}

// JobQueueStatus defines the observed state of JobQueue.
type JobQueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobQueue is the Schema for the JobQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type JobQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobQueueSpec   `json:"spec"`
	Status            JobQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobQueueList contains a list of JobQueues
type JobQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobQueue `json:"items"`
}

// Repository type metadata.
var (
	JobQueue_Kind             = "JobQueue"
	JobQueue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobQueue_Kind}.String()
	JobQueue_KindAPIVersion   = JobQueue_Kind + "." + CRDGroupVersion.String()
	JobQueue_GroupVersionKind = CRDGroupVersion.WithKind(JobQueue_Kind)
)

func init() {
	SchemeBuilder.Register(&JobQueue{}, &JobQueueList{})
}
