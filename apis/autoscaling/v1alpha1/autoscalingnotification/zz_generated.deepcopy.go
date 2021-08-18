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
func (in *AutoscalingNotification) DeepCopyInto(out *AutoscalingNotification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotification.
func (in *AutoscalingNotification) DeepCopy() *AutoscalingNotification {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingNotification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingNotificationList) DeepCopyInto(out *AutoscalingNotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoscalingNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotificationList.
func (in *AutoscalingNotificationList) DeepCopy() *AutoscalingNotificationList {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoscalingNotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingNotificationObservation) DeepCopyInto(out *AutoscalingNotificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotificationObservation.
func (in *AutoscalingNotificationObservation) DeepCopy() *AutoscalingNotificationObservation {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingNotificationParameters) DeepCopyInto(out *AutoscalingNotificationParameters) {
	*out = *in
	if in.GroupNames != nil {
		in, out := &in.GroupNames, &out.GroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotificationParameters.
func (in *AutoscalingNotificationParameters) DeepCopy() *AutoscalingNotificationParameters {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingNotificationSpec) DeepCopyInto(out *AutoscalingNotificationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotificationSpec.
func (in *AutoscalingNotificationSpec) DeepCopy() *AutoscalingNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalingNotificationStatus) DeepCopyInto(out *AutoscalingNotificationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalingNotificationStatus.
func (in *AutoscalingNotificationStatus) DeepCopy() *AutoscalingNotificationStatus {
	if in == nil {
		return nil
	}
	out := new(AutoscalingNotificationStatus)
	in.DeepCopyInto(out)
	return out
}
