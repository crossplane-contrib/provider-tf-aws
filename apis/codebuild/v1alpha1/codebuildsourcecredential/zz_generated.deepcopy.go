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
func (in *CodebuildSourceCredential) DeepCopyInto(out *CodebuildSourceCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredential.
func (in *CodebuildSourceCredential) DeepCopy() *CodebuildSourceCredential {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebuildSourceCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebuildSourceCredentialList) DeepCopyInto(out *CodebuildSourceCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CodebuildSourceCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredentialList.
func (in *CodebuildSourceCredentialList) DeepCopy() *CodebuildSourceCredentialList {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CodebuildSourceCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebuildSourceCredentialObservation) DeepCopyInto(out *CodebuildSourceCredentialObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredentialObservation.
func (in *CodebuildSourceCredentialObservation) DeepCopy() *CodebuildSourceCredentialObservation {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredentialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebuildSourceCredentialParameters) DeepCopyInto(out *CodebuildSourceCredentialParameters) {
	*out = *in
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredentialParameters.
func (in *CodebuildSourceCredentialParameters) DeepCopy() *CodebuildSourceCredentialParameters {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredentialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebuildSourceCredentialSpec) DeepCopyInto(out *CodebuildSourceCredentialSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredentialSpec.
func (in *CodebuildSourceCredentialSpec) DeepCopy() *CodebuildSourceCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CodebuildSourceCredentialStatus) DeepCopyInto(out *CodebuildSourceCredentialStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CodebuildSourceCredentialStatus.
func (in *CodebuildSourceCredentialStatus) DeepCopy() *CodebuildSourceCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(CodebuildSourceCredentialStatus)
	in.DeepCopyInto(out)
	return out
}
