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
func (in *Vpc) DeepCopyInto(out *Vpc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpc.
func (in *Vpc) DeepCopy() *Vpc {
	if in == nil {
		return nil
	}
	out := new(Vpc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vpc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcList) DeepCopyInto(out *VpcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vpc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcList.
func (in *VpcList) DeepCopy() *VpcList {
	if in == nil {
		return nil
	}
	out := new(VpcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcObservation) DeepCopyInto(out *VpcObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcObservation.
func (in *VpcObservation) DeepCopy() *VpcObservation {
	if in == nil {
		return nil
	}
	out := new(VpcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcParameters) DeepCopyInto(out *VpcParameters) {
	*out = *in
	if in.AssignGeneratedIpv6CidrBlock != nil {
		in, out := &in.AssignGeneratedIpv6CidrBlock, &out.AssignGeneratedIpv6CidrBlock
		*out = new(bool)
		**out = **in
	}
	if in.EnableClassiclink != nil {
		in, out := &in.EnableClassiclink, &out.EnableClassiclink
		*out = new(bool)
		**out = **in
	}
	if in.EnableClassiclinkDnsSupport != nil {
		in, out := &in.EnableClassiclinkDnsSupport, &out.EnableClassiclinkDnsSupport
		*out = new(bool)
		**out = **in
	}
	if in.EnableDnsHostnames != nil {
		in, out := &in.EnableDnsHostnames, &out.EnableDnsHostnames
		*out = new(bool)
		**out = **in
	}
	if in.EnableDnsSupport != nil {
		in, out := &in.EnableDnsSupport, &out.EnableDnsSupport
		*out = new(bool)
		**out = **in
	}
	if in.InstanceTenancy != nil {
		in, out := &in.InstanceTenancy, &out.InstanceTenancy
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcParameters.
func (in *VpcParameters) DeepCopy() *VpcParameters {
	if in == nil {
		return nil
	}
	out := new(VpcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcSpec) DeepCopyInto(out *VpcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcSpec.
func (in *VpcSpec) DeepCopy() *VpcSpec {
	if in == nil {
		return nil
	}
	out := new(VpcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcStatus) DeepCopyInto(out *VpcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcStatus.
func (in *VpcStatus) DeepCopy() *VpcStatus {
	if in == nil {
		return nil
	}
	out := new(VpcStatus)
	in.DeepCopyInto(out)
	return out
}
