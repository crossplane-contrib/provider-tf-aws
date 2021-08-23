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
// +groupName=pinpoint.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/pinpoint/v1alpha1"
)

type PinpointApnsSandboxChannelObservation struct {
}

type PinpointApnsSandboxChannelParameters struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	BundleId *string `json:"bundleId,omitempty" tf:"bundle_id"`

	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	DefaultAuthenticationMethod *string `json:"defaultAuthenticationMethod,omitempty" tf:"default_authentication_method"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	PrivateKey *string `json:"privateKey,omitempty" tf:"private_key"`

	TeamId *string `json:"teamId,omitempty" tf:"team_id"`

	TokenKey *string `json:"tokenKey,omitempty" tf:"token_key"`

	TokenKeyId *string `json:"tokenKeyId,omitempty" tf:"token_key_id"`
}

// PinpointApnsSandboxChannelSpec defines the desired state of PinpointApnsSandboxChannel
type PinpointApnsSandboxChannelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointApnsSandboxChannelParameters `json:"forProvider"`
}

// PinpointApnsSandboxChannelStatus defines the observed state of PinpointApnsSandboxChannel.
type PinpointApnsSandboxChannelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointApnsSandboxChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsSandboxChannel is the Schema for the PinpointApnsSandboxChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointApnsSandboxChannelSpec   `json:"spec"`
	Status            PinpointApnsSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsSandboxChannelList contains a list of PinpointApnsSandboxChannels
type PinpointApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsSandboxChannel `json:"items"`
}

// Repository type metadata.
var (
	PinpointApnsSandboxChannelKind             = "PinpointApnsSandboxChannel"
	PinpointApnsSandboxChannelGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointApnsSandboxChannelKind}.String()
	PinpointApnsSandboxChannelKindAPIVersion   = PinpointApnsSandboxChannelKind + "." + v1alpha1.GroupVersion.String()
	PinpointApnsSandboxChannelGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointApnsSandboxChannelKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&PinpointApnsSandboxChannel{}, &PinpointApnsSandboxChannelList{})
}
