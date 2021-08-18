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
// +groupName=load.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/load/v1alpha1"
)

type LoadBalancerBackendServerPolicyObservation struct {
}

type LoadBalancerBackendServerPolicyParameters struct {
	InstancePort int64 `json:"instancePort" tf:"instance_port"`

	LoadBalancerName string `json:"loadBalancerName" tf:"load_balancer_name"`

	PolicyNames []string `json:"policyNames,omitempty" tf:"policy_names"`
}

// LoadBalancerBackendServerPolicySpec defines the desired state of LoadBalancerBackendServerPolicy
type LoadBalancerBackendServerPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LoadBalancerBackendServerPolicyParameters `json:"forProvider"`
}

// LoadBalancerBackendServerPolicyStatus defines the observed state of LoadBalancerBackendServerPolicy.
type LoadBalancerBackendServerPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LoadBalancerBackendServerPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerBackendServerPolicy is the Schema for the LoadBalancerBackendServerPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LoadBalancerBackendServerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerBackendServerPolicySpec   `json:"spec"`
	Status            LoadBalancerBackendServerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerBackendServerPolicyList contains a list of LoadBalancerBackendServerPolicys
type LoadBalancerBackendServerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerBackendServerPolicy `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerBackendServerPolicyKind             = "LoadBalancerBackendServerPolicy"
	LoadBalancerBackendServerPolicyGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LoadBalancerBackendServerPolicyKind}.String()
	LoadBalancerBackendServerPolicyKindAPIVersion   = LoadBalancerBackendServerPolicyKind + "." + v1alpha1.GroupVersion.String()
	LoadBalancerBackendServerPolicyGroupVersionKind = v1alpha1.GroupVersion.WithKind(LoadBalancerBackendServerPolicyKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LoadBalancerBackendServerPolicy{}, &LoadBalancerBackendServerPolicyList{})
}
