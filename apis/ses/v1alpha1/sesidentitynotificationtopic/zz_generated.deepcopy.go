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
func (in *SesIdentityNotificationTopic) DeepCopyInto(out *SesIdentityNotificationTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopic.
func (in *SesIdentityNotificationTopic) DeepCopy() *SesIdentityNotificationTopic {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesIdentityNotificationTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityNotificationTopicList) DeepCopyInto(out *SesIdentityNotificationTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SesIdentityNotificationTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopicList.
func (in *SesIdentityNotificationTopicList) DeepCopy() *SesIdentityNotificationTopicList {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SesIdentityNotificationTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityNotificationTopicObservation) DeepCopyInto(out *SesIdentityNotificationTopicObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopicObservation.
func (in *SesIdentityNotificationTopicObservation) DeepCopy() *SesIdentityNotificationTopicObservation {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityNotificationTopicParameters) DeepCopyInto(out *SesIdentityNotificationTopicParameters) {
	*out = *in
	if in.IncludeOriginalHeaders != nil {
		in, out := &in.IncludeOriginalHeaders, &out.IncludeOriginalHeaders
		*out = new(bool)
		**out = **in
	}
	if in.TopicArn != nil {
		in, out := &in.TopicArn, &out.TopicArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopicParameters.
func (in *SesIdentityNotificationTopicParameters) DeepCopy() *SesIdentityNotificationTopicParameters {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityNotificationTopicSpec) DeepCopyInto(out *SesIdentityNotificationTopicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopicSpec.
func (in *SesIdentityNotificationTopicSpec) DeepCopy() *SesIdentityNotificationTopicSpec {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SesIdentityNotificationTopicStatus) DeepCopyInto(out *SesIdentityNotificationTopicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SesIdentityNotificationTopicStatus.
func (in *SesIdentityNotificationTopicStatus) DeepCopy() *SesIdentityNotificationTopicStatus {
	if in == nil {
		return nil
	}
	out := new(SesIdentityNotificationTopicStatus)
	in.DeepCopyInto(out)
	return out
}
