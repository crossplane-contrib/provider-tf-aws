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
func (in *SimpledbDomain) DeepCopyInto(out *SimpledbDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomain.
func (in *SimpledbDomain) DeepCopy() *SimpledbDomain {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimpledbDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpledbDomainList) DeepCopyInto(out *SimpledbDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SimpledbDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomainList.
func (in *SimpledbDomainList) DeepCopy() *SimpledbDomainList {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimpledbDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpledbDomainObservation) DeepCopyInto(out *SimpledbDomainObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomainObservation.
func (in *SimpledbDomainObservation) DeepCopy() *SimpledbDomainObservation {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpledbDomainParameters) DeepCopyInto(out *SimpledbDomainParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomainParameters.
func (in *SimpledbDomainParameters) DeepCopy() *SimpledbDomainParameters {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpledbDomainSpec) DeepCopyInto(out *SimpledbDomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomainSpec.
func (in *SimpledbDomainSpec) DeepCopy() *SimpledbDomainSpec {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpledbDomainStatus) DeepCopyInto(out *SimpledbDomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpledbDomainStatus.
func (in *SimpledbDomainStatus) DeepCopy() *SimpledbDomainStatus {
	if in == nil {
		return nil
	}
	out := new(SimpledbDomainStatus)
	in.DeepCopyInto(out)
	return out
}
