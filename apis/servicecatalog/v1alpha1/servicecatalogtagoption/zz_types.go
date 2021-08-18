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
// +groupName=servicecatalog.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/servicecatalog/v1alpha1"
)

type ServicecatalogTagOptionObservation struct {
	Owner string `json:"owner" tf:"owner"`
}

type ServicecatalogTagOptionParameters struct {
	Active *bool `json:"active,omitempty" tf:"active"`

	Key string `json:"key" tf:"key"`

	Value string `json:"value" tf:"value"`
}

// ServicecatalogTagOptionSpec defines the desired state of ServicecatalogTagOption
type ServicecatalogTagOptionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogTagOptionParameters `json:"forProvider"`
}

// ServicecatalogTagOptionStatus defines the observed state of ServicecatalogTagOption.
type ServicecatalogTagOptionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogTagOptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogTagOption is the Schema for the ServicecatalogTagOptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogTagOption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogTagOptionSpec   `json:"spec"`
	Status            ServicecatalogTagOptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogTagOptionList contains a list of ServicecatalogTagOptions
type ServicecatalogTagOptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogTagOption `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogTagOptionKind             = "ServicecatalogTagOption"
	ServicecatalogTagOptionGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServicecatalogTagOptionKind}.String()
	ServicecatalogTagOptionKindAPIVersion   = ServicecatalogTagOptionKind + "." + v1alpha1.GroupVersion.String()
	ServicecatalogTagOptionGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServicecatalogTagOptionKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ServicecatalogTagOption{}, &ServicecatalogTagOptionList{})
}
