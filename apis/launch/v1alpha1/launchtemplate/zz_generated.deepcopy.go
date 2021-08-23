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
func (in *BlockDeviceMappingsObservation) DeepCopyInto(out *BlockDeviceMappingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMappingsObservation.
func (in *BlockDeviceMappingsObservation) DeepCopy() *BlockDeviceMappingsObservation {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMappingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMappingsParameters) DeepCopyInto(out *BlockDeviceMappingsParameters) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Ebs != nil {
		in, out := &in.Ebs, &out.Ebs
		*out = make([]EbsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMappingsParameters.
func (in *BlockDeviceMappingsParameters) DeepCopy() *BlockDeviceMappingsParameters {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMappingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationSpecificationObservation) DeepCopyInto(out *CapacityReservationSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationSpecificationObservation.
func (in *CapacityReservationSpecificationObservation) DeepCopy() *CapacityReservationSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationSpecificationParameters) DeepCopyInto(out *CapacityReservationSpecificationParameters) {
	*out = *in
	if in.CapacityReservationPreference != nil {
		in, out := &in.CapacityReservationPreference, &out.CapacityReservationPreference
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationTarget != nil {
		in, out := &in.CapacityReservationTarget, &out.CapacityReservationTarget
		*out = make([]CapacityReservationTargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationSpecificationParameters.
func (in *CapacityReservationSpecificationParameters) DeepCopy() *CapacityReservationSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationTargetObservation) DeepCopyInto(out *CapacityReservationTargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationTargetObservation.
func (in *CapacityReservationTargetObservation) DeepCopy() *CapacityReservationTargetObservation {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationTargetParameters) DeepCopyInto(out *CapacityReservationTargetParameters) {
	*out = *in
	if in.CapacityReservationId != nil {
		in, out := &in.CapacityReservationId, &out.CapacityReservationId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationTargetParameters.
func (in *CapacityReservationTargetParameters) DeepCopy() *CapacityReservationTargetParameters {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuOptionsObservation) DeepCopyInto(out *CpuOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuOptionsObservation.
func (in *CpuOptionsObservation) DeepCopy() *CpuOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(CpuOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CpuOptionsParameters) DeepCopyInto(out *CpuOptionsParameters) {
	*out = *in
	if in.CoreCount != nil {
		in, out := &in.CoreCount, &out.CoreCount
		*out = new(int64)
		**out = **in
	}
	if in.ThreadsPerCore != nil {
		in, out := &in.ThreadsPerCore, &out.ThreadsPerCore
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CpuOptionsParameters.
func (in *CpuOptionsParameters) DeepCopy() *CpuOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(CpuOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreditSpecificationObservation) DeepCopyInto(out *CreditSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreditSpecificationObservation.
func (in *CreditSpecificationObservation) DeepCopy() *CreditSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CreditSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreditSpecificationParameters) DeepCopyInto(out *CreditSpecificationParameters) {
	*out = *in
	if in.CpuCredits != nil {
		in, out := &in.CpuCredits, &out.CpuCredits
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreditSpecificationParameters.
func (in *CreditSpecificationParameters) DeepCopy() *CreditSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CreditSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsObservation) DeepCopyInto(out *EbsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsObservation.
func (in *EbsObservation) DeepCopy() *EbsObservation {
	if in == nil {
		return nil
	}
	out := new(EbsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EbsParameters) DeepCopyInto(out *EbsParameters) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.SnapshotId != nil {
		in, out := &in.SnapshotId, &out.SnapshotId
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EbsParameters.
func (in *EbsParameters) DeepCopy() *EbsParameters {
	if in == nil {
		return nil
	}
	out := new(EbsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticGpuSpecificationsObservation) DeepCopyInto(out *ElasticGpuSpecificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticGpuSpecificationsObservation.
func (in *ElasticGpuSpecificationsObservation) DeepCopy() *ElasticGpuSpecificationsObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticGpuSpecificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticGpuSpecificationsParameters) DeepCopyInto(out *ElasticGpuSpecificationsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticGpuSpecificationsParameters.
func (in *ElasticGpuSpecificationsParameters) DeepCopy() *ElasticGpuSpecificationsParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticGpuSpecificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticInferenceAcceleratorObservation) DeepCopyInto(out *ElasticInferenceAcceleratorObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticInferenceAcceleratorObservation.
func (in *ElasticInferenceAcceleratorObservation) DeepCopy() *ElasticInferenceAcceleratorObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticInferenceAcceleratorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticInferenceAcceleratorParameters) DeepCopyInto(out *ElasticInferenceAcceleratorParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticInferenceAcceleratorParameters.
func (in *ElasticInferenceAcceleratorParameters) DeepCopy() *ElasticInferenceAcceleratorParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticInferenceAcceleratorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnclaveOptionsObservation) DeepCopyInto(out *EnclaveOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnclaveOptionsObservation.
func (in *EnclaveOptionsObservation) DeepCopy() *EnclaveOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(EnclaveOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnclaveOptionsParameters) DeepCopyInto(out *EnclaveOptionsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnclaveOptionsParameters.
func (in *EnclaveOptionsParameters) DeepCopy() *EnclaveOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(EnclaveOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HibernationOptionsObservation) DeepCopyInto(out *HibernationOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HibernationOptionsObservation.
func (in *HibernationOptionsObservation) DeepCopy() *HibernationOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(HibernationOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HibernationOptionsParameters) DeepCopyInto(out *HibernationOptionsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HibernationOptionsParameters.
func (in *HibernationOptionsParameters) DeepCopy() *HibernationOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(HibernationOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamInstanceProfileObservation) DeepCopyInto(out *IamInstanceProfileObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamInstanceProfileObservation.
func (in *IamInstanceProfileObservation) DeepCopy() *IamInstanceProfileObservation {
	if in == nil {
		return nil
	}
	out := new(IamInstanceProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamInstanceProfileParameters) DeepCopyInto(out *IamInstanceProfileParameters) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamInstanceProfileParameters.
func (in *IamInstanceProfileParameters) DeepCopy() *IamInstanceProfileParameters {
	if in == nil {
		return nil
	}
	out := new(IamInstanceProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMarketOptionsObservation) DeepCopyInto(out *InstanceMarketOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMarketOptionsObservation.
func (in *InstanceMarketOptionsObservation) DeepCopy() *InstanceMarketOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceMarketOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMarketOptionsParameters) DeepCopyInto(out *InstanceMarketOptionsParameters) {
	*out = *in
	if in.MarketType != nil {
		in, out := &in.MarketType, &out.MarketType
		*out = new(string)
		**out = **in
	}
	if in.SpotOptions != nil {
		in, out := &in.SpotOptions, &out.SpotOptions
		*out = make([]SpotOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMarketOptionsParameters.
func (in *InstanceMarketOptionsParameters) DeepCopy() *InstanceMarketOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceMarketOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplate) DeepCopyInto(out *LaunchTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplate.
func (in *LaunchTemplate) DeepCopy() *LaunchTemplate {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateList) DeepCopyInto(out *LaunchTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LaunchTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateList.
func (in *LaunchTemplateList) DeepCopy() *LaunchTemplateList {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LaunchTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
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
	if in.BlockDeviceMappings != nil {
		in, out := &in.BlockDeviceMappings, &out.BlockDeviceMappings
		*out = make([]BlockDeviceMappingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CapacityReservationSpecification != nil {
		in, out := &in.CapacityReservationSpecification, &out.CapacityReservationSpecification
		*out = make([]CapacityReservationSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CpuOptions != nil {
		in, out := &in.CpuOptions, &out.CpuOptions
		*out = make([]CpuOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreditSpecification != nil {
		in, out := &in.CreditSpecification, &out.CreditSpecification
		*out = make([]CreditSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableApiTermination != nil {
		in, out := &in.DisableApiTermination, &out.DisableApiTermination
		*out = new(bool)
		**out = **in
	}
	if in.EbsOptimized != nil {
		in, out := &in.EbsOptimized, &out.EbsOptimized
		*out = new(string)
		**out = **in
	}
	if in.ElasticGpuSpecifications != nil {
		in, out := &in.ElasticGpuSpecifications, &out.ElasticGpuSpecifications
		*out = make([]ElasticGpuSpecificationsParameters, len(*in))
		copy(*out, *in)
	}
	if in.ElasticInferenceAccelerator != nil {
		in, out := &in.ElasticInferenceAccelerator, &out.ElasticInferenceAccelerator
		*out = make([]ElasticInferenceAcceleratorParameters, len(*in))
		copy(*out, *in)
	}
	if in.EnclaveOptions != nil {
		in, out := &in.EnclaveOptions, &out.EnclaveOptions
		*out = make([]EnclaveOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HibernationOptions != nil {
		in, out := &in.HibernationOptions, &out.HibernationOptions
		*out = make([]HibernationOptionsParameters, len(*in))
		copy(*out, *in)
	}
	if in.IamInstanceProfile != nil {
		in, out := &in.IamInstanceProfile, &out.IamInstanceProfile
		*out = make([]IamInstanceProfileParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageId != nil {
		in, out := &in.ImageId, &out.ImageId
		*out = new(string)
		**out = **in
	}
	if in.InstanceInitiatedShutdownBehavior != nil {
		in, out := &in.InstanceInitiatedShutdownBehavior, &out.InstanceInitiatedShutdownBehavior
		*out = new(string)
		**out = **in
	}
	if in.InstanceMarketOptions != nil {
		in, out := &in.InstanceMarketOptions, &out.InstanceMarketOptions
		*out = make([]InstanceMarketOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KernelId != nil {
		in, out := &in.KernelId, &out.KernelId
		*out = new(string)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.LicenseSpecification != nil {
		in, out := &in.LicenseSpecification, &out.LicenseSpecification
		*out = make([]LicenseSpecificationParameters, len(*in))
		copy(*out, *in)
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = make([]MetadataOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = make([]MonitoringParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]NetworkInterfacesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = make([]PlacementParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RamDiskId != nil {
		in, out := &in.RamDiskId, &out.RamDiskId
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TagSpecifications != nil {
		in, out := &in.TagSpecifications, &out.TagSpecifications
		*out = make([]TagSpecificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.UpdateDefaultVersion != nil {
		in, out := &in.UpdateDefaultVersion, &out.UpdateDefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VpcSecurityGroupIds != nil {
		in, out := &in.VpcSecurityGroupIds, &out.VpcSecurityGroupIds
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateSpec) DeepCopyInto(out *LaunchTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateSpec.
func (in *LaunchTemplateSpec) DeepCopy() *LaunchTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateStatus) DeepCopyInto(out *LaunchTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateStatus.
func (in *LaunchTemplateStatus) DeepCopy() *LaunchTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseSpecificationObservation) DeepCopyInto(out *LicenseSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseSpecificationObservation.
func (in *LicenseSpecificationObservation) DeepCopy() *LicenseSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(LicenseSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseSpecificationParameters) DeepCopyInto(out *LicenseSpecificationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseSpecificationParameters.
func (in *LicenseSpecificationParameters) DeepCopy() *LicenseSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(LicenseSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsObservation) DeepCopyInto(out *MetadataOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsObservation.
func (in *MetadataOptionsObservation) DeepCopy() *MetadataOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetadataOptionsParameters) DeepCopyInto(out *MetadataOptionsParameters) {
	*out = *in
	if in.HttpEndpoint != nil {
		in, out := &in.HttpEndpoint, &out.HttpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HttpPutResponseHopLimit != nil {
		in, out := &in.HttpPutResponseHopLimit, &out.HttpPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HttpTokens != nil {
		in, out := &in.HttpTokens, &out.HttpTokens
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetadataOptionsParameters.
func (in *MetadataOptionsParameters) DeepCopy() *MetadataOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(MetadataOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringObservation) DeepCopyInto(out *MonitoringObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringObservation.
func (in *MonitoringObservation) DeepCopy() *MonitoringObservation {
	if in == nil {
		return nil
	}
	out := new(MonitoringObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringParameters) DeepCopyInto(out *MonitoringParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringParameters.
func (in *MonitoringParameters) DeepCopy() *MonitoringParameters {
	if in == nil {
		return nil
	}
	out := new(MonitoringParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfacesObservation) DeepCopyInto(out *NetworkInterfacesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfacesObservation.
func (in *NetworkInterfacesObservation) DeepCopy() *NetworkInterfacesObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfacesParameters) DeepCopyInto(out *NetworkInterfacesParameters) {
	*out = *in
	if in.AssociateCarrierIpAddress != nil {
		in, out := &in.AssociateCarrierIpAddress, &out.AssociateCarrierIpAddress
		*out = new(string)
		**out = **in
	}
	if in.AssociatePublicIpAddress != nil {
		in, out := &in.AssociatePublicIpAddress, &out.AssociatePublicIpAddress
		*out = new(string)
		**out = **in
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceIndex != nil {
		in, out := &in.DeviceIndex, &out.DeviceIndex
		*out = new(int64)
		**out = **in
	}
	if in.InterfaceType != nil {
		in, out := &in.InterfaceType, &out.InterfaceType
		*out = new(string)
		**out = **in
	}
	if in.Ipv4AddressCount != nil {
		in, out := &in.Ipv4AddressCount, &out.Ipv4AddressCount
		*out = new(int64)
		**out = **in
	}
	if in.Ipv4Addresses != nil {
		in, out := &in.Ipv4Addresses, &out.Ipv4Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ipv6AddressCount != nil {
		in, out := &in.Ipv6AddressCount, &out.Ipv6AddressCount
		*out = new(int64)
		**out = **in
	}
	if in.Ipv6Addresses != nil {
		in, out := &in.Ipv6Addresses, &out.Ipv6Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkInterfaceId != nil {
		in, out := &in.NetworkInterfaceId, &out.NetworkInterfaceId
		*out = new(string)
		**out = **in
	}
	if in.PrivateIpAddress != nil {
		in, out := &in.PrivateIpAddress, &out.PrivateIpAddress
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfacesParameters.
func (in *NetworkInterfacesParameters) DeepCopy() *NetworkInterfacesParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfacesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementObservation) DeepCopyInto(out *PlacementObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementObservation.
func (in *PlacementObservation) DeepCopy() *PlacementObservation {
	if in == nil {
		return nil
	}
	out := new(PlacementObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementParameters) DeepCopyInto(out *PlacementParameters) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.HostId != nil {
		in, out := &in.HostId, &out.HostId
		*out = new(string)
		**out = **in
	}
	if in.HostResourceGroupArn != nil {
		in, out := &in.HostResourceGroupArn, &out.HostResourceGroupArn
		*out = new(string)
		**out = **in
	}
	if in.PartitionNumber != nil {
		in, out := &in.PartitionNumber, &out.PartitionNumber
		*out = new(int64)
		**out = **in
	}
	if in.SpreadDomain != nil {
		in, out := &in.SpreadDomain, &out.SpreadDomain
		*out = new(string)
		**out = **in
	}
	if in.Tenancy != nil {
		in, out := &in.Tenancy, &out.Tenancy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementParameters.
func (in *PlacementParameters) DeepCopy() *PlacementParameters {
	if in == nil {
		return nil
	}
	out := new(PlacementParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotOptionsObservation) DeepCopyInto(out *SpotOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotOptionsObservation.
func (in *SpotOptionsObservation) DeepCopy() *SpotOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(SpotOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotOptionsParameters) DeepCopyInto(out *SpotOptionsParameters) {
	*out = *in
	if in.BlockDurationMinutes != nil {
		in, out := &in.BlockDurationMinutes, &out.BlockDurationMinutes
		*out = new(int64)
		**out = **in
	}
	if in.InstanceInterruptionBehavior != nil {
		in, out := &in.InstanceInterruptionBehavior, &out.InstanceInterruptionBehavior
		*out = new(string)
		**out = **in
	}
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(string)
		**out = **in
	}
	if in.SpotInstanceType != nil {
		in, out := &in.SpotInstanceType, &out.SpotInstanceType
		*out = new(string)
		**out = **in
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotOptionsParameters.
func (in *SpotOptionsParameters) DeepCopy() *SpotOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(SpotOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpecificationsObservation) DeepCopyInto(out *TagSpecificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpecificationsObservation.
func (in *TagSpecificationsObservation) DeepCopy() *TagSpecificationsObservation {
	if in == nil {
		return nil
	}
	out := new(TagSpecificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpecificationsParameters) DeepCopyInto(out *TagSpecificationsParameters) {
	*out = *in
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpecificationsParameters.
func (in *TagSpecificationsParameters) DeepCopy() *TagSpecificationsParameters {
	if in == nil {
		return nil
	}
	out := new(TagSpecificationsParameters)
	in.DeepCopyInto(out)
	return out
}
