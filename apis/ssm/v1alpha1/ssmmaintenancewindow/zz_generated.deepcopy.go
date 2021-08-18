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
func (in *SsmMaintenanceWindow) DeepCopyInto(out *SsmMaintenanceWindow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindow.
func (in *SsmMaintenanceWindow) DeepCopy() *SsmMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsmMaintenanceWindow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmMaintenanceWindowList) DeepCopyInto(out *SsmMaintenanceWindowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SsmMaintenanceWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindowList.
func (in *SsmMaintenanceWindowList) DeepCopy() *SsmMaintenanceWindowList {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SsmMaintenanceWindowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmMaintenanceWindowObservation) DeepCopyInto(out *SsmMaintenanceWindowObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindowObservation.
func (in *SsmMaintenanceWindowObservation) DeepCopy() *SsmMaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmMaintenanceWindowParameters) DeepCopyInto(out *SsmMaintenanceWindowParameters) {
	*out = *in
	if in.AllowUnassociatedTargets != nil {
		in, out := &in.AllowUnassociatedTargets, &out.AllowUnassociatedTargets
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.ScheduleOffset != nil {
		in, out := &in.ScheduleOffset, &out.ScheduleOffset
		*out = new(int64)
		**out = **in
	}
	if in.ScheduleTimezone != nil {
		in, out := &in.ScheduleTimezone, &out.ScheduleTimezone
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindowParameters.
func (in *SsmMaintenanceWindowParameters) DeepCopy() *SsmMaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmMaintenanceWindowSpec) DeepCopyInto(out *SsmMaintenanceWindowSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindowSpec.
func (in *SsmMaintenanceWindowSpec) DeepCopy() *SsmMaintenanceWindowSpec {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmMaintenanceWindowStatus) DeepCopyInto(out *SsmMaintenanceWindowStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmMaintenanceWindowStatus.
func (in *SsmMaintenanceWindowStatus) DeepCopy() *SsmMaintenanceWindowStatus {
	if in == nil {
		return nil
	}
	out := new(SsmMaintenanceWindowStatus)
	in.DeepCopyInto(out)
	return out
}
