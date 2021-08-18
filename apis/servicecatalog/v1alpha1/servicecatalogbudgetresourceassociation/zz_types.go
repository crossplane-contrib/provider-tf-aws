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

type ServicecatalogBudgetResourceAssociationObservation struct {
}

type ServicecatalogBudgetResourceAssociationParameters struct {
	BudgetName string `json:"budgetName" tf:"budget_name"`

	ResourceId string `json:"resourceId" tf:"resource_id"`
}

// ServicecatalogBudgetResourceAssociationSpec defines the desired state of ServicecatalogBudgetResourceAssociation
type ServicecatalogBudgetResourceAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogBudgetResourceAssociationParameters `json:"forProvider"`
}

// ServicecatalogBudgetResourceAssociationStatus defines the observed state of ServicecatalogBudgetResourceAssociation.
type ServicecatalogBudgetResourceAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogBudgetResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogBudgetResourceAssociation is the Schema for the ServicecatalogBudgetResourceAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServicecatalogBudgetResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogBudgetResourceAssociationSpec   `json:"spec"`
	Status            ServicecatalogBudgetResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogBudgetResourceAssociationList contains a list of ServicecatalogBudgetResourceAssociations
type ServicecatalogBudgetResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogBudgetResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogBudgetResourceAssociationKind             = "ServicecatalogBudgetResourceAssociation"
	ServicecatalogBudgetResourceAssociationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: ServicecatalogBudgetResourceAssociationKind}.String()
	ServicecatalogBudgetResourceAssociationKindAPIVersion   = ServicecatalogBudgetResourceAssociationKind + "." + v1alpha1.GroupVersion.String()
	ServicecatalogBudgetResourceAssociationGroupVersionKind = v1alpha1.GroupVersion.WithKind(ServicecatalogBudgetResourceAssociationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&ServicecatalogBudgetResourceAssociation{}, &ServicecatalogBudgetResourceAssociationList{})
}
