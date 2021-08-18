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
func (in *AccessLogSettingsObservation) DeepCopyInto(out *AccessLogSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogSettingsObservation.
func (in *AccessLogSettingsObservation) DeepCopy() *AccessLogSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(AccessLogSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLogSettingsParameters) DeepCopyInto(out *AccessLogSettingsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLogSettingsParameters.
func (in *AccessLogSettingsParameters) DeepCopy() *AccessLogSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(AccessLogSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStage) DeepCopyInto(out *ApiGatewayStage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStage.
func (in *ApiGatewayStage) DeepCopy() *ApiGatewayStage {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayStage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStageList) DeepCopyInto(out *ApiGatewayStageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiGatewayStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStageList.
func (in *ApiGatewayStageList) DeepCopy() *ApiGatewayStageList {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiGatewayStageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStageObservation) DeepCopyInto(out *ApiGatewayStageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStageObservation.
func (in *ApiGatewayStageObservation) DeepCopy() *ApiGatewayStageObservation {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStageParameters) DeepCopyInto(out *ApiGatewayStageParameters) {
	*out = *in
	if in.AccessLogSettings != nil {
		in, out := &in.AccessLogSettings, &out.AccessLogSettings
		*out = make([]AccessLogSettingsParameters, len(*in))
		copy(*out, *in)
	}
	if in.CacheClusterEnabled != nil {
		in, out := &in.CacheClusterEnabled, &out.CacheClusterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CacheClusterSize != nil {
		in, out := &in.CacheClusterSize, &out.CacheClusterSize
		*out = new(string)
		**out = **in
	}
	if in.ClientCertificateId != nil {
		in, out := &in.ClientCertificateId, &out.ClientCertificateId
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DocumentationVersion != nil {
		in, out := &in.DocumentationVersion, &out.DocumentationVersion
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.XrayTracingEnabled != nil {
		in, out := &in.XrayTracingEnabled, &out.XrayTracingEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStageParameters.
func (in *ApiGatewayStageParameters) DeepCopy() *ApiGatewayStageParameters {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStageSpec) DeepCopyInto(out *ApiGatewayStageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStageSpec.
func (in *ApiGatewayStageSpec) DeepCopy() *ApiGatewayStageSpec {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiGatewayStageStatus) DeepCopyInto(out *ApiGatewayStageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiGatewayStageStatus.
func (in *ApiGatewayStageStatus) DeepCopy() *ApiGatewayStageStatus {
	if in == nil {
		return nil
	}
	out := new(ApiGatewayStageStatus)
	in.DeepCopyInto(out)
	return out
}
