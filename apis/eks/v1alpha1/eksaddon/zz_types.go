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
// +groupName=eks.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/eks/v1alpha1"
)

type EksAddonObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedAt string `json:"createdAt" tf:"created_at"`

	ModifiedAt string `json:"modifiedAt" tf:"modified_at"`
}

type EksAddonParameters struct {
	AddonName string `json:"addonName" tf:"addon_name"`

	AddonVersion *string `json:"addonVersion,omitempty" tf:"addon_version"`

	ClusterName string `json:"clusterName" tf:"cluster_name"`

	ResolveConflicts *string `json:"resolveConflicts,omitempty" tf:"resolve_conflicts"`

	ServiceAccountRoleArn *string `json:"serviceAccountRoleArn,omitempty" tf:"service_account_role_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// EksAddonSpec defines the desired state of EksAddon
type EksAddonSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EksAddonParameters `json:"forProvider"`
}

// EksAddonStatus defines the observed state of EksAddon.
type EksAddonStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EksAddonObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksAddon is the Schema for the EksAddons API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EksAddon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksAddonSpec   `json:"spec"`
	Status            EksAddonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksAddonList contains a list of EksAddons
type EksAddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksAddon `json:"items"`
}

// Repository type metadata.
var (
	EksAddonKind             = "EksAddon"
	EksAddonGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EksAddonKind}.String()
	EksAddonKindAPIVersion   = EksAddonKind + "." + v1alpha1.GroupVersion.String()
	EksAddonGroupVersionKind = v1alpha1.GroupVersion.WithKind(EksAddonKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EksAddon{}, &EksAddonList{})
}
