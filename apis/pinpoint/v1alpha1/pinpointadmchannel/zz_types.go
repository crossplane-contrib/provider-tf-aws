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

type PinpointAdmChannelObservation struct {
}

type PinpointAdmChannelParameters struct {
	ApplicationId string `json:"applicationId" tf:"application_id"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

// PinpointAdmChannelSpec defines the desired state of PinpointAdmChannel
type PinpointAdmChannelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PinpointAdmChannelParameters `json:"forProvider"`
}

// PinpointAdmChannelStatus defines the observed state of PinpointAdmChannel.
type PinpointAdmChannelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PinpointAdmChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointAdmChannel is the Schema for the PinpointAdmChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PinpointAdmChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointAdmChannelSpec   `json:"spec"`
	Status            PinpointAdmChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointAdmChannelList contains a list of PinpointAdmChannels
type PinpointAdmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointAdmChannel `json:"items"`
}

// Repository type metadata.
var (
	PinpointAdmChannelKind             = "PinpointAdmChannel"
	PinpointAdmChannelGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: PinpointAdmChannelKind}.String()
	PinpointAdmChannelKindAPIVersion   = PinpointAdmChannelKind + "." + v1alpha1.GroupVersion.String()
	PinpointAdmChannelGroupVersionKind = v1alpha1.GroupVersion.WithKind(PinpointAdmChannelKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&PinpointAdmChannel{}, &PinpointAdmChannelList{})
}
