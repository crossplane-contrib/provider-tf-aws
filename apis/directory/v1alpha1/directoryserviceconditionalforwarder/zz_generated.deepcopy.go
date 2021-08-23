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
func (in *DirectoryServiceConditionalForwarder) DeepCopyInto(out *DirectoryServiceConditionalForwarder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarder.
func (in *DirectoryServiceConditionalForwarder) DeepCopy() *DirectoryServiceConditionalForwarder {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryServiceConditionalForwarder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceConditionalForwarderList) DeepCopyInto(out *DirectoryServiceConditionalForwarderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DirectoryServiceConditionalForwarder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarderList.
func (in *DirectoryServiceConditionalForwarderList) DeepCopy() *DirectoryServiceConditionalForwarderList {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryServiceConditionalForwarderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceConditionalForwarderObservation) DeepCopyInto(out *DirectoryServiceConditionalForwarderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarderObservation.
func (in *DirectoryServiceConditionalForwarderObservation) DeepCopy() *DirectoryServiceConditionalForwarderObservation {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceConditionalForwarderParameters) DeepCopyInto(out *DirectoryServiceConditionalForwarderParameters) {
	*out = *in
	if in.DnsIps != nil {
		in, out := &in.DnsIps, &out.DnsIps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarderParameters.
func (in *DirectoryServiceConditionalForwarderParameters) DeepCopy() *DirectoryServiceConditionalForwarderParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceConditionalForwarderSpec) DeepCopyInto(out *DirectoryServiceConditionalForwarderSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarderSpec.
func (in *DirectoryServiceConditionalForwarderSpec) DeepCopy() *DirectoryServiceConditionalForwarderSpec {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryServiceConditionalForwarderStatus) DeepCopyInto(out *DirectoryServiceConditionalForwarderStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryServiceConditionalForwarderStatus.
func (in *DirectoryServiceConditionalForwarderStatus) DeepCopy() *DirectoryServiceConditionalForwarderStatus {
	if in == nil {
		return nil
	}
	out := new(DirectoryServiceConditionalForwarderStatus)
	in.DeepCopyInto(out)
	return out
}
