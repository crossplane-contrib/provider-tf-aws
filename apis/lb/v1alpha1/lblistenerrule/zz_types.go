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
// +groupName=lb.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/lb/v1alpha1"
)

type ActionObservation struct {
}

type ActionParameters struct {
	AuthenticateCognito []AuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito"`

	AuthenticateOidc []AuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc"`

	FixedResponse []FixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response"`

	Forward []ForwardParameters `json:"forward,omitempty" tf:"forward"`

	Order *int64 `json:"order,omitempty" tf:"order"`

	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect"`

	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn"`

	Type string `json:"type" tf:"type"`
}

type AuthenticateCognitoObservation struct {
}

type AuthenticateCognitoParameters struct {
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`

	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`

	Scope *string `json:"scope,omitempty" tf:"scope"`

	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`

	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout"`

	UserPoolArn string `json:"userPoolArn" tf:"user_pool_arn"`

	UserPoolClientId string `json:"userPoolClientId" tf:"user_pool_client_id"`

	UserPoolDomain string `json:"userPoolDomain" tf:"user_pool_domain"`
}

type AuthenticateOidcObservation struct {
}

type AuthenticateOidcParameters struct {
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params"`

	AuthorizationEndpoint string `json:"authorizationEndpoint" tf:"authorization_endpoint"`

	ClientId string `json:"clientId" tf:"client_id"`

	ClientSecret string `json:"clientSecret" tf:"client_secret"`

	Issuer string `json:"issuer" tf:"issuer"`

	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request"`

	Scope *string `json:"scope,omitempty" tf:"scope"`

	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name"`

	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout"`

	TokenEndpoint string `json:"tokenEndpoint" tf:"token_endpoint"`

	UserInfoEndpoint string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {
	HostHeader []HostHeaderParameters `json:"hostHeader,omitempty" tf:"host_header"`

	HttpHeader []HttpHeaderParameters `json:"httpHeader,omitempty" tf:"http_header"`

	HttpRequestMethod []HttpRequestMethodParameters `json:"httpRequestMethod,omitempty" tf:"http_request_method"`

	PathPattern []PathPatternParameters `json:"pathPattern,omitempty" tf:"path_pattern"`

	QueryString []QueryStringParameters `json:"queryString,omitempty" tf:"query_string"`

	SourceIp []SourceIpParameters `json:"sourceIp,omitempty" tf:"source_ip"`
}

type FixedResponseObservation struct {
}

type FixedResponseParameters struct {
	ContentType string `json:"contentType" tf:"content_type"`

	MessageBody *string `json:"messageBody,omitempty" tf:"message_body"`

	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
}

type ForwardObservation struct {
}

type ForwardParameters struct {
	Stickiness []StickinessParameters `json:"stickiness,omitempty" tf:"stickiness"`

	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group"`
}

type HostHeaderObservation struct {
}

type HostHeaderParameters struct {
	Values []string `json:"values" tf:"values"`
}

type HttpHeaderObservation struct {
}

type HttpHeaderParameters struct {
	HttpHeaderName string `json:"httpHeaderName" tf:"http_header_name"`

	Values []string `json:"values" tf:"values"`
}

type HttpRequestMethodObservation struct {
}

type HttpRequestMethodParameters struct {
	Values []string `json:"values" tf:"values"`
}

type LbListenerRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type LbListenerRuleParameters struct {
	Action []ActionParameters `json:"action" tf:"action"`

	Condition []ConditionParameters `json:"condition" tf:"condition"`

	ListenerArn string `json:"listenerArn" tf:"listener_arn"`

	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type PathPatternObservation struct {
}

type PathPatternParameters struct {
	Values []string `json:"values" tf:"values"`
}

type QueryStringObservation struct {
}

type QueryStringParameters struct {
	Key *string `json:"key,omitempty" tf:"key"`

	Value string `json:"value" tf:"value"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {
	Host *string `json:"host,omitempty" tf:"host"`

	Path *string `json:"path,omitempty" tf:"path"`

	Port *string `json:"port,omitempty" tf:"port"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	Query *string `json:"query,omitempty" tf:"query"`

	StatusCode string `json:"statusCode" tf:"status_code"`
}

type SourceIpObservation struct {
}

type SourceIpParameters struct {
	Values []string `json:"values" tf:"values"`
}

type StickinessObservation struct {
}

type StickinessParameters struct {
	Duration int64 `json:"duration" tf:"duration"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
}

type TargetGroupObservation struct {
}

type TargetGroupParameters struct {
	Arn string `json:"arn" tf:"arn"`

	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

// LbListenerRuleSpec defines the desired state of LbListenerRule
type LbListenerRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LbListenerRuleParameters `json:"forProvider"`
}

// LbListenerRuleStatus defines the observed state of LbListenerRule.
type LbListenerRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LbListenerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerRule is the Schema for the LbListenerRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LbListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerRuleSpec   `json:"spec"`
	Status            LbListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerRuleList contains a list of LbListenerRules
type LbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbListenerRule `json:"items"`
}

// Repository type metadata.
var (
	LbListenerRuleKind             = "LbListenerRule"
	LbListenerRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: LbListenerRuleKind}.String()
	LbListenerRuleKindAPIVersion   = LbListenerRuleKind + "." + v1alpha1.GroupVersion.String()
	LbListenerRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(LbListenerRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&LbListenerRule{}, &LbListenerRuleList{})
}
