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
func (in *SsoadminManagedPolicyAttachment) DeepCopyInto(out *SsoadminManagedPolicyAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachment.
func (in *SsoadminManagedPolicyAttachment) DeepCopy() *SsoadminManagedPolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsoadminManagedPolicyAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminManagedPolicyAttachmentList) DeepCopyInto(out *SsoadminManagedPolicyAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SsoadminManagedPolicyAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachmentList.
func (in *SsoadminManagedPolicyAttachmentList) DeepCopy() *SsoadminManagedPolicyAttachmentList {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsoadminManagedPolicyAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminManagedPolicyAttachmentObservation) DeepCopyInto(out *SsoadminManagedPolicyAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachmentObservation.
func (in *SsoadminManagedPolicyAttachmentObservation) DeepCopy() *SsoadminManagedPolicyAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminManagedPolicyAttachmentParameters) DeepCopyInto(out *SsoadminManagedPolicyAttachmentParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachmentParameters.
func (in *SsoadminManagedPolicyAttachmentParameters) DeepCopy() *SsoadminManagedPolicyAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminManagedPolicyAttachmentSpec) DeepCopyInto(out *SsoadminManagedPolicyAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachmentSpec.
func (in *SsoadminManagedPolicyAttachmentSpec) DeepCopy() *SsoadminManagedPolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoadminManagedPolicyAttachmentStatus) DeepCopyInto(out *SsoadminManagedPolicyAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoadminManagedPolicyAttachmentStatus.
func (in *SsoadminManagedPolicyAttachmentStatus) DeepCopy() *SsoadminManagedPolicyAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(SsoadminManagedPolicyAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
