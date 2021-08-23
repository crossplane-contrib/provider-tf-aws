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
func (in *ByteMatchTuplesObservation) DeepCopyInto(out *ByteMatchTuplesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByteMatchTuplesObservation.
func (in *ByteMatchTuplesObservation) DeepCopy() *ByteMatchTuplesObservation {
	if in == nil {
		return nil
	}
	out := new(ByteMatchTuplesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ByteMatchTuplesParameters) DeepCopyInto(out *ByteMatchTuplesParameters) {
	*out = *in
	if in.FieldToMatch != nil {
		in, out := &in.FieldToMatch, &out.FieldToMatch
		*out = make([]FieldToMatchParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetString != nil {
		in, out := &in.TargetString, &out.TargetString
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByteMatchTuplesParameters.
func (in *ByteMatchTuplesParameters) DeepCopy() *ByteMatchTuplesParameters {
	if in == nil {
		return nil
	}
	out := new(ByteMatchTuplesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldToMatchObservation) DeepCopyInto(out *FieldToMatchObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldToMatchObservation.
func (in *FieldToMatchObservation) DeepCopy() *FieldToMatchObservation {
	if in == nil {
		return nil
	}
	out := new(FieldToMatchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldToMatchParameters) DeepCopyInto(out *FieldToMatchParameters) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldToMatchParameters.
func (in *FieldToMatchParameters) DeepCopy() *FieldToMatchParameters {
	if in == nil {
		return nil
	}
	out := new(FieldToMatchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSet) DeepCopyInto(out *WafByteMatchSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSet.
func (in *WafByteMatchSet) DeepCopy() *WafByteMatchSet {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafByteMatchSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSetList) DeepCopyInto(out *WafByteMatchSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WafByteMatchSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSetList.
func (in *WafByteMatchSetList) DeepCopy() *WafByteMatchSetList {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WafByteMatchSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSetObservation) DeepCopyInto(out *WafByteMatchSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSetObservation.
func (in *WafByteMatchSetObservation) DeepCopy() *WafByteMatchSetObservation {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSetParameters) DeepCopyInto(out *WafByteMatchSetParameters) {
	*out = *in
	if in.ByteMatchTuples != nil {
		in, out := &in.ByteMatchTuples, &out.ByteMatchTuples
		*out = make([]ByteMatchTuplesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSetParameters.
func (in *WafByteMatchSetParameters) DeepCopy() *WafByteMatchSetParameters {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSetSpec) DeepCopyInto(out *WafByteMatchSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSetSpec.
func (in *WafByteMatchSetSpec) DeepCopy() *WafByteMatchSetSpec {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WafByteMatchSetStatus) DeepCopyInto(out *WafByteMatchSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WafByteMatchSetStatus.
func (in *WafByteMatchSetStatus) DeepCopy() *WafByteMatchSetStatus {
	if in == nil {
		return nil
	}
	out := new(WafByteMatchSetStatus)
	in.DeepCopyInto(out)
	return out
}
