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
// +groupName=cloudfront.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cloudfront/v1alpha1"
)

type CloudfrontOriginAccessIdentityObservation struct {
	CallerReference string `json:"callerReference" tf:"caller_reference"`

	CloudfrontAccessIdentityPath string `json:"cloudfrontAccessIdentityPath" tf:"cloudfront_access_identity_path"`

	Etag string `json:"etag" tf:"etag"`

	IamArn string `json:"iamArn" tf:"iam_arn"`

	S3CanonicalUserId string `json:"s3CanonicalUserId" tf:"s3_canonical_user_id"`
}

type CloudfrontOriginAccessIdentityParameters struct {
	Comment *string `json:"comment,omitempty" tf:"comment"`
}

// CloudfrontOriginAccessIdentitySpec defines the desired state of CloudfrontOriginAccessIdentity
type CloudfrontOriginAccessIdentitySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudfrontOriginAccessIdentityParameters `json:"forProvider"`
}

// CloudfrontOriginAccessIdentityStatus defines the observed state of CloudfrontOriginAccessIdentity.
type CloudfrontOriginAccessIdentityStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudfrontOriginAccessIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontOriginAccessIdentity is the Schema for the CloudfrontOriginAccessIdentitys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CloudfrontOriginAccessIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontOriginAccessIdentitySpec   `json:"spec"`
	Status            CloudfrontOriginAccessIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontOriginAccessIdentityList contains a list of CloudfrontOriginAccessIdentitys
type CloudfrontOriginAccessIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontOriginAccessIdentity `json:"items"`
}

// Repository type metadata.
var (
	CloudfrontOriginAccessIdentityKind             = "CloudfrontOriginAccessIdentity"
	CloudfrontOriginAccessIdentityGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CloudfrontOriginAccessIdentityKind}.String()
	CloudfrontOriginAccessIdentityKindAPIVersion   = CloudfrontOriginAccessIdentityKind + "." + v1alpha1.GroupVersion.String()
	CloudfrontOriginAccessIdentityGroupVersionKind = v1alpha1.GroupVersion.WithKind(CloudfrontOriginAccessIdentityKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CloudfrontOriginAccessIdentity{}, &CloudfrontOriginAccessIdentityList{})
}
