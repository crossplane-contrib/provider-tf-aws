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
func (in *ServicecatalogConstraint) DeepCopyInto(out *ServicecatalogConstraint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraint.
func (in *ServicecatalogConstraint) DeepCopy() *ServicecatalogConstraint {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogConstraint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogConstraintList) DeepCopyInto(out *ServicecatalogConstraintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicecatalogConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraintList.
func (in *ServicecatalogConstraintList) DeepCopy() *ServicecatalogConstraintList {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServicecatalogConstraintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogConstraintObservation) DeepCopyInto(out *ServicecatalogConstraintObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraintObservation.
func (in *ServicecatalogConstraintObservation) DeepCopy() *ServicecatalogConstraintObservation {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraintObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogConstraintParameters) DeepCopyInto(out *ServicecatalogConstraintParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraintParameters.
func (in *ServicecatalogConstraintParameters) DeepCopy() *ServicecatalogConstraintParameters {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraintParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogConstraintSpec) DeepCopyInto(out *ServicecatalogConstraintSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraintSpec.
func (in *ServicecatalogConstraintSpec) DeepCopy() *ServicecatalogConstraintSpec {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicecatalogConstraintStatus) DeepCopyInto(out *ServicecatalogConstraintStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicecatalogConstraintStatus.
func (in *ServicecatalogConstraintStatus) DeepCopy() *ServicecatalogConstraintStatus {
	if in == nil {
		return nil
	}
	out := new(ServicecatalogConstraintStatus)
	in.DeepCopyInto(out)
	return out
}
