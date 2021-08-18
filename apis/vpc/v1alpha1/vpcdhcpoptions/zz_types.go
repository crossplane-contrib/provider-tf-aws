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
// +groupName=vpc.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1"
)

type VpcDhcpOptionsObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type VpcDhcpOptionsParameters struct {
	DomainName *string `json:"domainName,omitempty" tf:"domain_name"`

	DomainNameServers []string `json:"domainNameServers,omitempty" tf:"domain_name_servers"`

	NetbiosNameServers []string `json:"netbiosNameServers,omitempty" tf:"netbios_name_servers"`

	NetbiosNodeType *string `json:"netbiosNodeType,omitempty" tf:"netbios_node_type"`

	NtpServers []string `json:"ntpServers,omitempty" tf:"ntp_servers"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// VpcDhcpOptionsSpec defines the desired state of VpcDhcpOptions
type VpcDhcpOptionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcDhcpOptionsParameters `json:"forProvider"`
}

// VpcDhcpOptionsStatus defines the observed state of VpcDhcpOptions.
type VpcDhcpOptionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcDhcpOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcDhcpOptions is the Schema for the VpcDhcpOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VpcDhcpOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcDhcpOptionsSpec   `json:"spec"`
	Status            VpcDhcpOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcDhcpOptionsList contains a list of VpcDhcpOptionss
type VpcDhcpOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcDhcpOptions `json:"items"`
}

// Repository type metadata.
var (
	VpcDhcpOptionsKind             = "VpcDhcpOptions"
	VpcDhcpOptionsGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: VpcDhcpOptionsKind}.String()
	VpcDhcpOptionsKindAPIVersion   = VpcDhcpOptionsKind + "." + v1alpha1.GroupVersion.String()
	VpcDhcpOptionsGroupVersionKind = v1alpha1.GroupVersion.WithKind(VpcDhcpOptionsKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&VpcDhcpOptions{}, &VpcDhcpOptionsList{})
}
