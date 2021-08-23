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
func (in *CognitoConfigObservation) DeepCopyInto(out *CognitoConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoConfigObservation.
func (in *CognitoConfigObservation) DeepCopy() *CognitoConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CognitoConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CognitoConfigParameters) DeepCopyInto(out *CognitoConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CognitoConfigParameters.
func (in *CognitoConfigParameters) DeepCopy() *CognitoConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CognitoConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcConfigObservation) DeepCopyInto(out *OidcConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcConfigObservation.
func (in *OidcConfigObservation) DeepCopy() *OidcConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OidcConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcConfigParameters) DeepCopyInto(out *OidcConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcConfigParameters.
func (in *OidcConfigParameters) DeepCopy() *OidcConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OidcConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforce) DeepCopyInto(out *SagemakerWorkforce) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforce.
func (in *SagemakerWorkforce) DeepCopy() *SagemakerWorkforce {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforce)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerWorkforce) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforceList) DeepCopyInto(out *SagemakerWorkforceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SagemakerWorkforce, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforceList.
func (in *SagemakerWorkforceList) DeepCopy() *SagemakerWorkforceList {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SagemakerWorkforceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforceObservation) DeepCopyInto(out *SagemakerWorkforceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforceObservation.
func (in *SagemakerWorkforceObservation) DeepCopy() *SagemakerWorkforceObservation {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforceParameters) DeepCopyInto(out *SagemakerWorkforceParameters) {
	*out = *in
	if in.CognitoConfig != nil {
		in, out := &in.CognitoConfig, &out.CognitoConfig
		*out = make([]CognitoConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.OidcConfig != nil {
		in, out := &in.OidcConfig, &out.OidcConfig
		*out = make([]OidcConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.SourceIpConfig != nil {
		in, out := &in.SourceIpConfig, &out.SourceIpConfig
		*out = make([]SourceIpConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforceParameters.
func (in *SagemakerWorkforceParameters) DeepCopy() *SagemakerWorkforceParameters {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforceSpec) DeepCopyInto(out *SagemakerWorkforceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforceSpec.
func (in *SagemakerWorkforceSpec) DeepCopy() *SagemakerWorkforceSpec {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SagemakerWorkforceStatus) DeepCopyInto(out *SagemakerWorkforceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SagemakerWorkforceStatus.
func (in *SagemakerWorkforceStatus) DeepCopy() *SagemakerWorkforceStatus {
	if in == nil {
		return nil
	}
	out := new(SagemakerWorkforceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceIpConfigObservation) DeepCopyInto(out *SourceIpConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceIpConfigObservation.
func (in *SourceIpConfigObservation) DeepCopy() *SourceIpConfigObservation {
	if in == nil {
		return nil
	}
	out := new(SourceIpConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceIpConfigParameters) DeepCopyInto(out *SourceIpConfigParameters) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceIpConfigParameters.
func (in *SourceIpConfigParameters) DeepCopy() *SourceIpConfigParameters {
	if in == nil {
		return nil
	}
	out := new(SourceIpConfigParameters)
	in.DeepCopyInto(out)
	return out
}
