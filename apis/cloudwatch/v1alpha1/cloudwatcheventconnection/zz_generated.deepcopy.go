// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeyObservation) DeepCopyInto(out *ApiKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeyObservation.
func (in *ApiKeyObservation) DeepCopy() *ApiKeyObservation {
	if in == nil {
		return nil
	}
	out := new(ApiKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiKeyParameters) DeepCopyInto(out *ApiKeyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiKeyParameters.
func (in *ApiKeyParameters) DeepCopy() *ApiKeyParameters {
	if in == nil {
		return nil
	}
	out := new(ApiKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthParametersObservation) DeepCopyInto(out *AuthParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthParametersObservation.
func (in *AuthParametersObservation) DeepCopy() *AuthParametersObservation {
	if in == nil {
		return nil
	}
	out := new(AuthParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthParametersParameters) DeepCopyInto(out *AuthParametersParameters) {
	*out = *in
	if in.ApiKey != nil {
		in, out := &in.ApiKey, &out.ApiKey
		*out = make([]ApiKeyParameters, len(*in))
		copy(*out, *in)
	}
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		*out = make([]BasicParameters, len(*in))
		copy(*out, *in)
	}
	if in.InvocationHttpParameters != nil {
		in, out := &in.InvocationHttpParameters, &out.InvocationHttpParameters
		*out = make([]InvocationHttpParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Oauth != nil {
		in, out := &in.Oauth, &out.Oauth
		*out = make([]OauthParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthParametersParameters.
func (in *AuthParametersParameters) DeepCopy() *AuthParametersParameters {
	if in == nil {
		return nil
	}
	out := new(AuthParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicObservation) DeepCopyInto(out *BasicObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicObservation.
func (in *BasicObservation) DeepCopy() *BasicObservation {
	if in == nil {
		return nil
	}
	out := new(BasicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicParameters) DeepCopyInto(out *BasicParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicParameters.
func (in *BasicParameters) DeepCopy() *BasicParameters {
	if in == nil {
		return nil
	}
	out := new(BasicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BodyObservation) DeepCopyInto(out *BodyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BodyObservation.
func (in *BodyObservation) DeepCopy() *BodyObservation {
	if in == nil {
		return nil
	}
	out := new(BodyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BodyParameters) DeepCopyInto(out *BodyParameters) {
	*out = *in
	if in.IsValueSecret != nil {
		in, out := &in.IsValueSecret, &out.IsValueSecret
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BodyParameters.
func (in *BodyParameters) DeepCopy() *BodyParameters {
	if in == nil {
		return nil
	}
	out := new(BodyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientParametersObservation) DeepCopyInto(out *ClientParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientParametersObservation.
func (in *ClientParametersObservation) DeepCopy() *ClientParametersObservation {
	if in == nil {
		return nil
	}
	out := new(ClientParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientParametersParameters) DeepCopyInto(out *ClientParametersParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientParametersParameters.
func (in *ClientParametersParameters) DeepCopy() *ClientParametersParameters {
	if in == nil {
		return nil
	}
	out := new(ClientParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnection) DeepCopyInto(out *CloudwatchEventConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnection.
func (in *CloudwatchEventConnection) DeepCopy() *CloudwatchEventConnection {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchEventConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnectionList) DeepCopyInto(out *CloudwatchEventConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudwatchEventConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnectionList.
func (in *CloudwatchEventConnectionList) DeepCopy() *CloudwatchEventConnectionList {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudwatchEventConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnectionObservation) DeepCopyInto(out *CloudwatchEventConnectionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnectionObservation.
func (in *CloudwatchEventConnectionObservation) DeepCopy() *CloudwatchEventConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnectionParameters) DeepCopyInto(out *CloudwatchEventConnectionParameters) {
	*out = *in
	if in.AuthParameters != nil {
		in, out := &in.AuthParameters, &out.AuthParameters
		*out = make([]AuthParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnectionParameters.
func (in *CloudwatchEventConnectionParameters) DeepCopy() *CloudwatchEventConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnectionSpec) DeepCopyInto(out *CloudwatchEventConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnectionSpec.
func (in *CloudwatchEventConnectionSpec) DeepCopy() *CloudwatchEventConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudwatchEventConnectionStatus) DeepCopyInto(out *CloudwatchEventConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudwatchEventConnectionStatus.
func (in *CloudwatchEventConnectionStatus) DeepCopy() *CloudwatchEventConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(CloudwatchEventConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderObservation) DeepCopyInto(out *HeaderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderObservation.
func (in *HeaderObservation) DeepCopy() *HeaderObservation {
	if in == nil {
		return nil
	}
	out := new(HeaderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderParameters) DeepCopyInto(out *HeaderParameters) {
	*out = *in
	if in.IsValueSecret != nil {
		in, out := &in.IsValueSecret, &out.IsValueSecret
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderParameters.
func (in *HeaderParameters) DeepCopy() *HeaderParameters {
	if in == nil {
		return nil
	}
	out := new(HeaderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvocationHttpParametersObservation) DeepCopyInto(out *InvocationHttpParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvocationHttpParametersObservation.
func (in *InvocationHttpParametersObservation) DeepCopy() *InvocationHttpParametersObservation {
	if in == nil {
		return nil
	}
	out := new(InvocationHttpParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvocationHttpParametersParameters) DeepCopyInto(out *InvocationHttpParametersParameters) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = make([]BodyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = make([]HeaderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueryString != nil {
		in, out := &in.QueryString, &out.QueryString
		*out = make([]QueryStringParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvocationHttpParametersParameters.
func (in *InvocationHttpParametersParameters) DeepCopy() *InvocationHttpParametersParameters {
	if in == nil {
		return nil
	}
	out := new(InvocationHttpParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthHttpParametersObservation) DeepCopyInto(out *OauthHttpParametersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthHttpParametersObservation.
func (in *OauthHttpParametersObservation) DeepCopy() *OauthHttpParametersObservation {
	if in == nil {
		return nil
	}
	out := new(OauthHttpParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthHttpParametersParameters) DeepCopyInto(out *OauthHttpParametersParameters) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = make([]BodyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = make([]HeaderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueryString != nil {
		in, out := &in.QueryString, &out.QueryString
		*out = make([]QueryStringParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthHttpParametersParameters.
func (in *OauthHttpParametersParameters) DeepCopy() *OauthHttpParametersParameters {
	if in == nil {
		return nil
	}
	out := new(OauthHttpParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthObservation) DeepCopyInto(out *OauthObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthObservation.
func (in *OauthObservation) DeepCopy() *OauthObservation {
	if in == nil {
		return nil
	}
	out := new(OauthObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthParameters) DeepCopyInto(out *OauthParameters) {
	*out = *in
	if in.ClientParameters != nil {
		in, out := &in.ClientParameters, &out.ClientParameters
		*out = make([]ClientParametersParameters, len(*in))
		copy(*out, *in)
	}
	if in.OauthHttpParameters != nil {
		in, out := &in.OauthHttpParameters, &out.OauthHttpParameters
		*out = make([]OauthHttpParametersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthParameters.
func (in *OauthParameters) DeepCopy() *OauthParameters {
	if in == nil {
		return nil
	}
	out := new(OauthParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryStringObservation) DeepCopyInto(out *QueryStringObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryStringObservation.
func (in *QueryStringObservation) DeepCopy() *QueryStringObservation {
	if in == nil {
		return nil
	}
	out := new(QueryStringObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryStringParameters) DeepCopyInto(out *QueryStringParameters) {
	*out = *in
	if in.IsValueSecret != nil {
		in, out := &in.IsValueSecret, &out.IsValueSecret
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryStringParameters.
func (in *QueryStringParameters) DeepCopy() *QueryStringParameters {
	if in == nil {
		return nil
	}
	out := new(QueryStringParameters)
	in.DeepCopyInto(out)
	return out
}
