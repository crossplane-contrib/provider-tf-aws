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
// +groupName=quicksight.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/quicksight/v1alpha1"
)

type QuicksightGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type QuicksightGroupParameters struct {
	AwsAccountId *string `json:"awsAccountId,omitempty" tf:"aws_account_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	GroupName string `json:"groupName" tf:"group_name"`

	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
}

// QuicksightGroupSpec defines the desired state of QuicksightGroup
type QuicksightGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       QuicksightGroupParameters `json:"forProvider"`
}

// QuicksightGroupStatus defines the observed state of QuicksightGroup.
type QuicksightGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          QuicksightGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QuicksightGroup is the Schema for the QuicksightGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type QuicksightGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuicksightGroupSpec   `json:"spec"`
	Status            QuicksightGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuicksightGroupList contains a list of QuicksightGroups
type QuicksightGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuicksightGroup `json:"items"`
}

// Repository type metadata.
var (
	QuicksightGroupKind             = "QuicksightGroup"
	QuicksightGroupGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: QuicksightGroupKind}.String()
	QuicksightGroupKindAPIVersion   = QuicksightGroupKind + "." + v1alpha1.GroupVersion.String()
	QuicksightGroupGroupVersionKind = v1alpha1.GroupVersion.WithKind(QuicksightGroupKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&QuicksightGroup{}, &QuicksightGroupList{})
}
