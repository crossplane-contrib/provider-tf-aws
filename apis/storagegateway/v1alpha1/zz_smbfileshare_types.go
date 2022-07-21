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

type SMBFileShareCacheAttributesObservation struct {
}

type SMBFileShareCacheAttributesParameters struct {

	// +kubebuilder:validation:Optional
	CacheStaleTimeoutInSeconds *float64 `json:"cacheStaleTimeoutInSeconds,omitempty" tf:"cache_stale_timeout_in_seconds,omitempty"`
}

type SMBFileShareObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	FileshareID *string `json:"fileshareId,omitempty" tf:"fileshare_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SMBFileShareParameters struct {

	// +kubebuilder:validation:Optional
	AccessBasedEnumeration *bool `json:"accessBasedEnumeration,omitempty" tf:"access_based_enumeration,omitempty"`

	// +kubebuilder:validation:Optional
	AdminUserList []*string `json:"adminUserList,omitempty" tf:"admin_user_list,omitempty"`

	// +kubebuilder:validation:Optional
	AuditDestinationArn *string `json:"auditDestinationArn,omitempty" tf:"audit_destination_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Authentication *string `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRegion *string `json:"bucketRegion,omitempty" tf:"bucket_region,omitempty"`

	// +kubebuilder:validation:Optional
	CacheAttributes []SMBFileShareCacheAttributesParameters `json:"cacheAttributes,omitempty" tf:"cache_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	CaseSensitivity *string `json:"caseSensitivity,omitempty" tf:"case_sensitivity,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultStorageClass *string `json:"defaultStorageClass,omitempty" tf:"default_storage_class,omitempty"`

	// +kubebuilder:validation:Optional
	FileShareName *string `json:"fileShareName,omitempty" tf:"file_share_name,omitempty"`

	// +kubebuilder:validation:Required
	GatewayArn *string `json:"gatewayArn" tf:"gateway_arn,omitempty"`

	// +kubebuilder:validation:Optional
	GuessMimeTypeEnabled *bool `json:"guessMimeTypeEnabled,omitempty" tf:"guess_mime_type_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	InvalidUserList []*string `json:"invalidUserList,omitempty" tf:"invalid_user_list,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	LocationArn *string `json:"locationArn" tf:"location_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationPolicy *string `json:"notificationPolicy,omitempty" tf:"notification_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectACL *string `json:"objectAcl,omitempty" tf:"object_acl,omitempty"`

	// +kubebuilder:validation:Optional
	OplocksEnabled *bool `json:"oplocksEnabled,omitempty" tf:"oplocks_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SMBACLEnabled *bool `json:"smbAclEnabled,omitempty" tf:"smb_acl_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCEndpointDNSName *string `json:"vpcEndpointDnsName,omitempty" tf:"vpc_endpoint_dns_name,omitempty"`

	// +kubebuilder:validation:Optional
	ValidUserList []*string `json:"validUserList,omitempty" tf:"valid_user_list,omitempty"`
}

// SMBFileShareSpec defines the desired state of SMBFileShare
type SMBFileShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SMBFileShareParameters `json:"forProvider"`
}

// SMBFileShareStatus defines the observed state of SMBFileShare.
type SMBFileShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SMBFileShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SMBFileShare is the Schema for the SMBFileShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SMBFileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SMBFileShareSpec   `json:"spec"`
	Status            SMBFileShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SMBFileShareList contains a list of SMBFileShares
type SMBFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMBFileShare `json:"items"`
}

// Repository type metadata.
var (
	SMBFileShare_Kind             = "SMBFileShare"
	SMBFileShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SMBFileShare_Kind}.String()
	SMBFileShare_KindAPIVersion   = SMBFileShare_Kind + "." + CRDGroupVersion.String()
	SMBFileShare_GroupVersionKind = CRDGroupVersion.WithKind(SMBFileShare_Kind)
)

func init() {
	SchemeBuilder.Register(&SMBFileShare{}, &SMBFileShareList{})
}
