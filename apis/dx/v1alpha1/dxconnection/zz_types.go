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
// +groupName=dx.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/dx/v1alpha1"
)

type DxConnectionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	AwsDevice string `json:"awsDevice" tf:"aws_device"`

	HasLogicalRedundancy string `json:"hasLogicalRedundancy" tf:"has_logical_redundancy"`

	JumboFrameCapable bool `json:"jumboFrameCapable" tf:"jumbo_frame_capable"`
}

type DxConnectionParameters struct {
	Bandwidth string `json:"bandwidth" tf:"bandwidth"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DxConnectionSpec defines the desired state of DxConnection
type DxConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxConnectionParameters `json:"forProvider"`
}

// DxConnectionStatus defines the observed state of DxConnection.
type DxConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxConnection is the Schema for the DxConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DxConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxConnectionSpec   `json:"spec"`
	Status            DxConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxConnectionList contains a list of DxConnections
type DxConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxConnection `json:"items"`
}

// Repository type metadata.
var (
	DxConnectionKind             = "DxConnection"
	DxConnectionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DxConnectionKind}.String()
	DxConnectionKindAPIVersion   = DxConnectionKind + "." + v1alpha1.GroupVersion.String()
	DxConnectionGroupVersionKind = v1alpha1.GroupVersion.WithKind(DxConnectionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DxConnection{}, &DxConnectionList{})
}
