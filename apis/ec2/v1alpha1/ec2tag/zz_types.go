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

type Ec2TagObservation struct {
}

type Ec2TagParameters struct {
	Key string `json:"key" tf:"key"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	Value string `json:"value" tf:"value"`
}

// Ec2TagSpec defines the desired state of Ec2Tag
type Ec2TagSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TagParameters `json:"forProvider"`
}

// Ec2TagStatus defines the observed state of Ec2Tag.
type Ec2TagStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2Tag is the Schema for the Ec2Tags API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Ec2Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TagSpec   `json:"spec"`
	Status            Ec2TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TagList contains a list of Ec2Tags
type Ec2TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2Tag `json:"items"`
}

// Repository type metadata.
var (
	Ec2TagKind             = "Ec2Tag"
	Ec2TagGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: Ec2TagKind}.String()
	Ec2TagKindAPIVersion   = Ec2TagKind + "." + v1alpha1.GroupVersion.String()
	Ec2TagGroupVersionKind = v1alpha1.GroupVersion.WithKind(Ec2TagKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&Ec2Tag{}, &Ec2TagList{})
}
