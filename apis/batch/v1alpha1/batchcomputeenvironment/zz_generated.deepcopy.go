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
func (in *BatchComputeEnvironment) DeepCopyInto(out *BatchComputeEnvironment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironment.
func (in *BatchComputeEnvironment) DeepCopy() *BatchComputeEnvironment {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchComputeEnvironment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchComputeEnvironmentList) DeepCopyInto(out *BatchComputeEnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchComputeEnvironment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironmentList.
func (in *BatchComputeEnvironmentList) DeepCopy() *BatchComputeEnvironmentList {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchComputeEnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchComputeEnvironmentObservation) DeepCopyInto(out *BatchComputeEnvironmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironmentObservation.
func (in *BatchComputeEnvironmentObservation) DeepCopy() *BatchComputeEnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchComputeEnvironmentParameters) DeepCopyInto(out *BatchComputeEnvironmentParameters) {
	*out = *in
	if in.ComputeEnvironmentName != nil {
		in, out := &in.ComputeEnvironmentName, &out.ComputeEnvironmentName
		*out = new(string)
		**out = **in
	}
	if in.ComputeEnvironmentNamePrefix != nil {
		in, out := &in.ComputeEnvironmentNamePrefix, &out.ComputeEnvironmentNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ComputeResources != nil {
		in, out := &in.ComputeResources, &out.ComputeResources
		*out = make([]ComputeResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceRole != nil {
		in, out := &in.ServiceRole, &out.ServiceRole
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironmentParameters.
func (in *BatchComputeEnvironmentParameters) DeepCopy() *BatchComputeEnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchComputeEnvironmentSpec) DeepCopyInto(out *BatchComputeEnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironmentSpec.
func (in *BatchComputeEnvironmentSpec) DeepCopy() *BatchComputeEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchComputeEnvironmentStatus) DeepCopyInto(out *BatchComputeEnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchComputeEnvironmentStatus.
func (in *BatchComputeEnvironmentStatus) DeepCopy() *BatchComputeEnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(BatchComputeEnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeResourcesObservation) DeepCopyInto(out *ComputeResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeResourcesObservation.
func (in *ComputeResourcesObservation) DeepCopy() *ComputeResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ComputeResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeResourcesParameters) DeepCopyInto(out *ComputeResourcesParameters) {
	*out = *in
	if in.AllocationStrategy != nil {
		in, out := &in.AllocationStrategy, &out.AllocationStrategy
		*out = new(string)
		**out = **in
	}
	if in.BidPercentage != nil {
		in, out := &in.BidPercentage, &out.BidPercentage
		*out = new(int64)
		**out = **in
	}
	if in.DesiredVcpus != nil {
		in, out := &in.DesiredVcpus, &out.DesiredVcpus
		*out = new(int64)
		**out = **in
	}
	if in.Ec2KeyPair != nil {
		in, out := &in.Ec2KeyPair, &out.Ec2KeyPair
		*out = new(string)
		**out = **in
	}
	if in.ImageId != nil {
		in, out := &in.ImageId, &out.ImageId
		*out = new(string)
		**out = **in
	}
	if in.InstanceRole != nil {
		in, out := &in.InstanceRole, &out.InstanceRole
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LaunchTemplate != nil {
		in, out := &in.LaunchTemplate, &out.LaunchTemplate
		*out = make([]LaunchTemplateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MinVcpus != nil {
		in, out := &in.MinVcpus, &out.MinVcpus
		*out = new(int64)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SpotIamFleetRole != nil {
		in, out := &in.SpotIamFleetRole, &out.SpotIamFleetRole
		*out = new(string)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeResourcesParameters.
func (in *ComputeResourcesParameters) DeepCopy() *ComputeResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeResourcesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateObservation) DeepCopyInto(out *LaunchTemplateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateObservation.
func (in *LaunchTemplateObservation) DeepCopy() *LaunchTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateParameters) DeepCopyInto(out *LaunchTemplateParameters) {
	*out = *in
	if in.LaunchTemplateId != nil {
		in, out := &in.LaunchTemplateId, &out.LaunchTemplateId
		*out = new(string)
		**out = **in
	}
	if in.LaunchTemplateName != nil {
		in, out := &in.LaunchTemplateName, &out.LaunchTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateParameters.
func (in *LaunchTemplateParameters) DeepCopy() *LaunchTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateParameters)
	in.DeepCopyInto(out)
	return out
}
