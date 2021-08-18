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
// +groupName=cloudwatch.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudwatch/v1alpha1"
)

type CloudwatchEventBusPolicyObservation struct {
}

type CloudwatchEventBusPolicyParameters struct {
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name"`

	Policy string `json:"policy" tf:"policy"`
}

// CloudwatchEventBusPolicySpec defines the desired state of CloudwatchEventBusPolicy
type CloudwatchEventBusPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventBusPolicyParameters `json:"forProvider"`
}

// CloudwatchEventBusPolicyStatus defines the observed state of CloudwatchEventBusPolicy.
type CloudwatchEventBusPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventBusPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventBusPolicy is the Schema for the CloudwatchEventBusPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudwatchEventBusPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventBusPolicySpec   `json:"spec"`
	Status            CloudwatchEventBusPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventBusPolicyList contains a list of CloudwatchEventBusPolicys
type CloudwatchEventBusPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventBusPolicy `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventBusPolicyKind             = "CloudwatchEventBusPolicy"
	CloudwatchEventBusPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudwatchEventBusPolicyKind}.String()
	CloudwatchEventBusPolicyKindAPIVersion   = CloudwatchEventBusPolicyKind + "." + v1alpha1.GroupVersion.String()
	CloudwatchEventBusPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudwatchEventBusPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudwatchEventBusPolicy{}, &CloudwatchEventBusPolicyList{})
}
