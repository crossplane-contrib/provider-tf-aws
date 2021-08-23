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
func (in *OpsworksPermission) DeepCopyInto(out *OpsworksPermission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermission.
func (in *OpsworksPermission) DeepCopy() *OpsworksPermission {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksPermission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksPermissionList) DeepCopyInto(out *OpsworksPermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpsworksPermission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermissionList.
func (in *OpsworksPermissionList) DeepCopy() *OpsworksPermissionList {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpsworksPermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksPermissionObservation) DeepCopyInto(out *OpsworksPermissionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermissionObservation.
func (in *OpsworksPermissionObservation) DeepCopy() *OpsworksPermissionObservation {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermissionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksPermissionParameters) DeepCopyInto(out *OpsworksPermissionParameters) {
	*out = *in
	if in.AllowSsh != nil {
		in, out := &in.AllowSsh, &out.AllowSsh
		*out = new(bool)
		**out = **in
	}
	if in.AllowSudo != nil {
		in, out := &in.AllowSudo, &out.AllowSudo
		*out = new(bool)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.StackId != nil {
		in, out := &in.StackId, &out.StackId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermissionParameters.
func (in *OpsworksPermissionParameters) DeepCopy() *OpsworksPermissionParameters {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermissionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksPermissionSpec) DeepCopyInto(out *OpsworksPermissionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermissionSpec.
func (in *OpsworksPermissionSpec) DeepCopy() *OpsworksPermissionSpec {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsworksPermissionStatus) DeepCopyInto(out *OpsworksPermissionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsworksPermissionStatus.
func (in *OpsworksPermissionStatus) DeepCopy() *OpsworksPermissionStatus {
	if in == nil {
		return nil
	}
	out := new(OpsworksPermissionStatus)
	in.DeepCopyInto(out)
	return out
}
