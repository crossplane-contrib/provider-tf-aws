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
// +groupName=ebs.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/ebs/v1alpha1"
)

type ClientDataObservation struct {
}

type ClientDataParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`

	UploadEnd *string `json:"uploadEnd,omitempty" tf:"upload_end"`

	UploadSize *float64 `json:"uploadSize,omitempty" tf:"upload_size"`

	UploadStart *string `json:"uploadStart,omitempty" tf:"upload_start"`
}

type DiskContainerObservation struct {
}

type DiskContainerParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Format string `json:"format" tf:"format"`

	Url *string `json:"url,omitempty" tf:"url"`

	UserBucket []UserBucketParameters `json:"userBucket,omitempty" tf:"user_bucket"`
}

type EbsSnapshotImportObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DataEncryptionKeyId string `json:"dataEncryptionKeyId" tf:"data_encryption_key_id"`

	OwnerAlias string `json:"ownerAlias" tf:"owner_alias"`

	OwnerId string `json:"ownerId" tf:"owner_id"`

	VolumeSize int64 `json:"volumeSize" tf:"volume_size"`
}

type EbsSnapshotImportParameters struct {
	ClientData []ClientDataParameters `json:"clientData,omitempty" tf:"client_data"`

	Description *string `json:"description,omitempty" tf:"description"`

	DiskContainer []DiskContainerParameters `json:"diskContainer" tf:"disk_container"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	RoleName *string `json:"roleName,omitempty" tf:"role_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type UserBucketObservation struct {
}

type UserBucketParameters struct {
	S3Bucket string `json:"s3Bucket" tf:"s3_bucket"`

	S3Key string `json:"s3Key" tf:"s3_key"`
}

// EbsSnapshotImportSpec defines the desired state of EbsSnapshotImport
type EbsSnapshotImportSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EbsSnapshotImportParameters `json:"forProvider"`
}

// EbsSnapshotImportStatus defines the observed state of EbsSnapshotImport.
type EbsSnapshotImportStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EbsSnapshotImportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EbsSnapshotImport is the Schema for the EbsSnapshotImports API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EbsSnapshotImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EbsSnapshotImportSpec   `json:"spec"`
	Status            EbsSnapshotImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsSnapshotImportList contains a list of EbsSnapshotImports
type EbsSnapshotImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsSnapshotImport `json:"items"`
}

// Repository type metadata.
var (
	EbsSnapshotImportKind             = "EbsSnapshotImport"
	EbsSnapshotImportGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: EbsSnapshotImportKind}.String()
	EbsSnapshotImportKindAPIVersion   = EbsSnapshotImportKind + "." + v1alpha1.GroupVersion.String()
	EbsSnapshotImportGroupVersionKind = v1alpha1.GroupVersion.WithKind(EbsSnapshotImportKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&EbsSnapshotImport{}, &EbsSnapshotImportList{})
}
