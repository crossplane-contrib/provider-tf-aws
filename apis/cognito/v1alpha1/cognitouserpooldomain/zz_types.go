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
// +groupName=cognito.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cognito/v1alpha1"
)

type CognitoUserPoolDomainObservation struct {
	AwsAccountId string `json:"awsAccountId" tf:"aws_account_id"`

	CloudfrontDistributionArn string `json:"cloudfrontDistributionArn" tf:"cloudfront_distribution_arn"`

	S3Bucket string `json:"s3Bucket" tf:"s3_bucket"`

	Version string `json:"version" tf:"version"`
}

type CognitoUserPoolDomainParameters struct {
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn"`

	Domain string `json:"domain" tf:"domain"`

	UserPoolId string `json:"userPoolId" tf:"user_pool_id"`
}

// CognitoUserPoolDomainSpec defines the desired state of CognitoUserPoolDomain
type CognitoUserPoolDomainSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserPoolDomainParameters `json:"forProvider"`
}

// CognitoUserPoolDomainStatus defines the observed state of CognitoUserPoolDomain.
type CognitoUserPoolDomainStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserPoolDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolDomain is the Schema for the CognitoUserPoolDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolDomainSpec   `json:"spec"`
	Status            CognitoUserPoolDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolDomainList contains a list of CognitoUserPoolDomains
type CognitoUserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPoolDomain `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserPoolDomainKind             = "CognitoUserPoolDomain"
	CognitoUserPoolDomainGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CognitoUserPoolDomainKind}.String()
	CognitoUserPoolDomainKindAPIVersion   = CognitoUserPoolDomainKind + "." + v1alpha1.GroupVersion.String()
	CognitoUserPoolDomainGroupVersionKind = v1alpha1.GroupVersion.WithKind(CognitoUserPoolDomainKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CognitoUserPoolDomain{}, &CognitoUserPoolDomainList{})
}
