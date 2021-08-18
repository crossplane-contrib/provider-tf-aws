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
func (in *FsxLustreFileSystem) DeepCopyInto(out *FsxLustreFileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystem.
func (in *FsxLustreFileSystem) DeepCopy() *FsxLustreFileSystem {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FsxLustreFileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxLustreFileSystemList) DeepCopyInto(out *FsxLustreFileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FsxLustreFileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystemList.
func (in *FsxLustreFileSystemList) DeepCopy() *FsxLustreFileSystemList {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FsxLustreFileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxLustreFileSystemObservation) DeepCopyInto(out *FsxLustreFileSystemObservation) {
	*out = *in
	if in.NetworkInterfaceIds != nil {
		in, out := &in.NetworkInterfaceIds, &out.NetworkInterfaceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystemObservation.
func (in *FsxLustreFileSystemObservation) DeepCopy() *FsxLustreFileSystemObservation {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxLustreFileSystemParameters) DeepCopyInto(out *FsxLustreFileSystemParameters) {
	*out = *in
	if in.AutoImportPolicy != nil {
		in, out := &in.AutoImportPolicy, &out.AutoImportPolicy
		*out = new(string)
		**out = **in
	}
	if in.AutomaticBackupRetentionDays != nil {
		in, out := &in.AutomaticBackupRetentionDays, &out.AutomaticBackupRetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.CopyTagsToBackups != nil {
		in, out := &in.CopyTagsToBackups, &out.CopyTagsToBackups
		*out = new(bool)
		**out = **in
	}
	if in.DailyAutomaticBackupStartTime != nil {
		in, out := &in.DailyAutomaticBackupStartTime, &out.DailyAutomaticBackupStartTime
		*out = new(string)
		**out = **in
	}
	if in.DataCompressionType != nil {
		in, out := &in.DataCompressionType, &out.DataCompressionType
		*out = new(string)
		**out = **in
	}
	if in.DeploymentType != nil {
		in, out := &in.DeploymentType, &out.DeploymentType
		*out = new(string)
		**out = **in
	}
	if in.DriveCacheType != nil {
		in, out := &in.DriveCacheType, &out.DriveCacheType
		*out = new(string)
		**out = **in
	}
	if in.ExportPath != nil {
		in, out := &in.ExportPath, &out.ExportPath
		*out = new(string)
		**out = **in
	}
	if in.ImportPath != nil {
		in, out := &in.ImportPath, &out.ImportPath
		*out = new(string)
		**out = **in
	}
	if in.ImportedFileChunkSize != nil {
		in, out := &in.ImportedFileChunkSize, &out.ImportedFileChunkSize
		*out = new(int64)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.PerUnitStorageThroughput != nil {
		in, out := &in.PerUnitStorageThroughput, &out.PerUnitStorageThroughput
		*out = new(int64)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.WeeklyMaintenanceStartTime != nil {
		in, out := &in.WeeklyMaintenanceStartTime, &out.WeeklyMaintenanceStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystemParameters.
func (in *FsxLustreFileSystemParameters) DeepCopy() *FsxLustreFileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxLustreFileSystemSpec) DeepCopyInto(out *FsxLustreFileSystemSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystemSpec.
func (in *FsxLustreFileSystemSpec) DeepCopy() *FsxLustreFileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsxLustreFileSystemStatus) DeepCopyInto(out *FsxLustreFileSystemStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsxLustreFileSystemStatus.
func (in *FsxLustreFileSystemStatus) DeepCopy() *FsxLustreFileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(FsxLustreFileSystemStatus)
	in.DeepCopyInto(out)
	return out
}
