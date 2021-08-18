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
func (in *PinpointAdmChannel) DeepCopyInto(out *PinpointAdmChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannel.
func (in *PinpointAdmChannel) DeepCopy() *PinpointAdmChannel {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointAdmChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointAdmChannelList) DeepCopyInto(out *PinpointAdmChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PinpointAdmChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannelList.
func (in *PinpointAdmChannelList) DeepCopy() *PinpointAdmChannelList {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PinpointAdmChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointAdmChannelObservation) DeepCopyInto(out *PinpointAdmChannelObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannelObservation.
func (in *PinpointAdmChannelObservation) DeepCopy() *PinpointAdmChannelObservation {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointAdmChannelParameters) DeepCopyInto(out *PinpointAdmChannelParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannelParameters.
func (in *PinpointAdmChannelParameters) DeepCopy() *PinpointAdmChannelParameters {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointAdmChannelSpec) DeepCopyInto(out *PinpointAdmChannelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannelSpec.
func (in *PinpointAdmChannelSpec) DeepCopy() *PinpointAdmChannelSpec {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PinpointAdmChannelStatus) DeepCopyInto(out *PinpointAdmChannelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PinpointAdmChannelStatus.
func (in *PinpointAdmChannelStatus) DeepCopy() *PinpointAdmChannelStatus {
	if in == nil {
		return nil
	}
	out := new(PinpointAdmChannelStatus)
	in.DeepCopyInto(out)
	return out
}
