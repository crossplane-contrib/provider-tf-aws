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
func (in *PinpointApnsChannel) DeepCopyInto(out *PinpointApnsChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannel.
func (in *PinpointApnsChannel) DeepCopy() *PinpointApnsChannel {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointApnsChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointApnsChannelList) DeepCopyInto(out *PinpointApnsChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PinpointApnsChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannelList.
func (in *PinpointApnsChannelList) DeepCopy() *PinpointApnsChannelList {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointApnsChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointApnsChannelObservation) DeepCopyInto(out *PinpointApnsChannelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannelObservation.
func (in *PinpointApnsChannelObservation) DeepCopy() *PinpointApnsChannelObservation {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointApnsChannelParameters) DeepCopyInto(out *PinpointApnsChannelParameters) {
	*out = *in
	if in.BundleId != nil {
		in, out := &in.BundleId, &out.BundleId
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DefaultAuthenticationMethod != nil {
		in, out := &in.DefaultAuthenticationMethod, &out.DefaultAuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.TeamId != nil {
		in, out := &in.TeamId, &out.TeamId
		*out = new(string)
		**out = **in
	}
	if in.TokenKey != nil {
		in, out := &in.TokenKey, &out.TokenKey
		*out = new(string)
		**out = **in
	}
	if in.TokenKeyId != nil {
		in, out := &in.TokenKeyId, &out.TokenKeyId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannelParameters.
func (in *PinpointApnsChannelParameters) DeepCopy() *PinpointApnsChannelParameters {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointApnsChannelSpec) DeepCopyInto(out *PinpointApnsChannelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannelSpec.
func (in *PinpointApnsChannelSpec) DeepCopy() *PinpointApnsChannelSpec {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointApnsChannelStatus) DeepCopyInto(out *PinpointApnsChannelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointApnsChannelStatus.
func (in *PinpointApnsChannelStatus) DeepCopy() *PinpointApnsChannelStatus {
	if in == nil {
		return nil
	}
	out := new(PinpointApnsChannelStatus)
	in.DeepCopyInto(out)
	return out
}
