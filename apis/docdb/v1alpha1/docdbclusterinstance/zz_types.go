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
// +groupName=docdb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/docdb/v1alpha1"
)

type DocdbClusterInstanceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DbSubnetGroupName string `json:"dbSubnetGroupName" tf:"db_subnet_group_name"`

	DbiResourceId string `json:"dbiResourceId" tf:"dbi_resource_id"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	EngineVersion string `json:"engineVersion" tf:"engine_version"`

	KmsKeyId string `json:"kmsKeyId" tf:"kms_key_id"`

	Port int64 `json:"port" tf:"port"`

	PreferredBackupWindow string `json:"preferredBackupWindow" tf:"preferred_backup_window"`

	PubliclyAccessible bool `json:"publiclyAccessible" tf:"publicly_accessible"`

	StorageEncrypted bool `json:"storageEncrypted" tf:"storage_encrypted"`

	Writer bool `json:"writer" tf:"writer"`
}

type DocdbClusterInstanceParameters struct {
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier"`

	ClusterIdentifier string `json:"clusterIdentifier" tf:"cluster_identifier"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	Identifier *string `json:"identifier,omitempty" tf:"identifier"`

	IdentifierPrefix *string `json:"identifierPrefix,omitempty" tf:"identifier_prefix"`

	InstanceClass string `json:"instanceClass" tf:"instance_class"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`

	PromotionTier *int64 `json:"promotionTier,omitempty" tf:"promotion_tier"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DocdbClusterInstanceSpec defines the desired state of DocdbClusterInstance
type DocdbClusterInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DocdbClusterInstanceParameters `json:"forProvider"`
}

// DocdbClusterInstanceStatus defines the observed state of DocdbClusterInstance.
type DocdbClusterInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DocdbClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterInstance is the Schema for the DocdbClusterInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocdbClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterInstanceSpec   `json:"spec"`
	Status            DocdbClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterInstanceList contains a list of DocdbClusterInstances
type DocdbClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	DocdbClusterInstanceKind             = "DocdbClusterInstance"
	DocdbClusterInstanceGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: DocdbClusterInstanceKind}.String()
	DocdbClusterInstanceKindAPIVersion   = DocdbClusterInstanceKind + "." + v1alpha1.GroupVersion.String()
	DocdbClusterInstanceGroupVersionKind = v1alpha1.GroupVersion.WithKind(DocdbClusterInstanceKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&DocdbClusterInstance{}, &DocdbClusterInstanceList{})
}
