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
func (in *TimestreamwriteDatabase) DeepCopyInto(out *TimestreamwriteDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabase.
func (in *TimestreamwriteDatabase) DeepCopy() *TimestreamwriteDatabase {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimestreamwriteDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestreamwriteDatabaseList) DeepCopyInto(out *TimestreamwriteDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TimestreamwriteDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabaseList.
func (in *TimestreamwriteDatabaseList) DeepCopy() *TimestreamwriteDatabaseList {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimestreamwriteDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestreamwriteDatabaseObservation) DeepCopyInto(out *TimestreamwriteDatabaseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabaseObservation.
func (in *TimestreamwriteDatabaseObservation) DeepCopy() *TimestreamwriteDatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestreamwriteDatabaseParameters) DeepCopyInto(out *TimestreamwriteDatabaseParameters) {
	*out = *in
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabaseParameters.
func (in *TimestreamwriteDatabaseParameters) DeepCopy() *TimestreamwriteDatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestreamwriteDatabaseSpec) DeepCopyInto(out *TimestreamwriteDatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabaseSpec.
func (in *TimestreamwriteDatabaseSpec) DeepCopy() *TimestreamwriteDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimestreamwriteDatabaseStatus) DeepCopyInto(out *TimestreamwriteDatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimestreamwriteDatabaseStatus.
func (in *TimestreamwriteDatabaseStatus) DeepCopy() *TimestreamwriteDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(TimestreamwriteDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}
