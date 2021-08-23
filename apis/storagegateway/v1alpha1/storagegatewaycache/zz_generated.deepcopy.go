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
func (in *StoragegatewayCache) DeepCopyInto(out *StoragegatewayCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCache.
func (in *StoragegatewayCache) DeepCopy() *StoragegatewayCache {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragegatewayCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewayCacheList) DeepCopyInto(out *StoragegatewayCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragegatewayCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCacheList.
func (in *StoragegatewayCacheList) DeepCopy() *StoragegatewayCacheList {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragegatewayCacheList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewayCacheObservation) DeepCopyInto(out *StoragegatewayCacheObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCacheObservation.
func (in *StoragegatewayCacheObservation) DeepCopy() *StoragegatewayCacheObservation {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCacheObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewayCacheParameters) DeepCopyInto(out *StoragegatewayCacheParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCacheParameters.
func (in *StoragegatewayCacheParameters) DeepCopy() *StoragegatewayCacheParameters {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCacheParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewayCacheSpec) DeepCopyInto(out *StoragegatewayCacheSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCacheSpec.
func (in *StoragegatewayCacheSpec) DeepCopy() *StoragegatewayCacheSpec {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragegatewayCacheStatus) DeepCopyInto(out *StoragegatewayCacheStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragegatewayCacheStatus.
func (in *StoragegatewayCacheStatus) DeepCopy() *StoragegatewayCacheStatus {
	if in == nil {
		return nil
	}
	out := new(StoragegatewayCacheStatus)
	in.DeepCopyInto(out)
	return out
}
