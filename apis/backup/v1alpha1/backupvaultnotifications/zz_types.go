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
// +groupName=backup.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/backup/v1alpha1"
)

type BackupVaultNotificationsObservation struct {
	BackupVaultArn string `json:"backupVaultArn" tf:"backup_vault_arn"`
}

type BackupVaultNotificationsParameters struct {
	BackupVaultEvents []string `json:"backupVaultEvents" tf:"backup_vault_events"`

	BackupVaultName string `json:"backupVaultName" tf:"backup_vault_name"`

	SnsTopicArn string `json:"snsTopicArn" tf:"sns_topic_arn"`
}

// BackupVaultNotificationsSpec defines the desired state of BackupVaultNotifications
type BackupVaultNotificationsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupVaultNotificationsParameters `json:"forProvider"`
}

// BackupVaultNotificationsStatus defines the observed state of BackupVaultNotifications.
type BackupVaultNotificationsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupVaultNotificationsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultNotifications is the Schema for the BackupVaultNotificationss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BackupVaultNotifications struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupVaultNotificationsSpec   `json:"spec"`
	Status            BackupVaultNotificationsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultNotificationsList contains a list of BackupVaultNotificationss
type BackupVaultNotificationsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVaultNotifications `json:"items"`
}

// Repository type metadata.
var (
	BackupVaultNotificationsKind             = "BackupVaultNotifications"
	BackupVaultNotificationsGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: BackupVaultNotificationsKind}.String()
	BackupVaultNotificationsKindAPIVersion   = BackupVaultNotificationsKind + "." + v1alpha1.GroupVersion.String()
	BackupVaultNotificationsGroupVersionKind = v1alpha1.GroupVersion.WithKind(BackupVaultNotificationsKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&BackupVaultNotifications{}, &BackupVaultNotificationsList{})
}
