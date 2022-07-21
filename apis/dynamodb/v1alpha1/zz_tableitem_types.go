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

type TableItemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TableItemParameters struct {

	// +kubebuilder:validation:Required
	HashKey *string `json:"hashKey" tf:"hash_key,omitempty"`

	// +kubebuilder:validation:Required
	Item *string `json:"item" tf:"item,omitempty"`

	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

// TableItemSpec defines the desired state of TableItem
type TableItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableItemParameters `json:"forProvider"`
}

// TableItemStatus defines the observed state of TableItem.
type TableItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableItem is the Schema for the TableItems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TableItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableItemSpec   `json:"spec"`
	Status            TableItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableItemList contains a list of TableItems
type TableItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableItem `json:"items"`
}

// Repository type metadata.
var (
	TableItem_Kind             = "TableItem"
	TableItem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableItem_Kind}.String()
	TableItem_KindAPIVersion   = TableItem_Kind + "." + CRDGroupVersion.String()
	TableItem_GroupVersionKind = CRDGroupVersion.WithKind(TableItem_Kind)
)

func init() {
	SchemeBuilder.Register(&TableItem{}, &TableItemList{})
}
