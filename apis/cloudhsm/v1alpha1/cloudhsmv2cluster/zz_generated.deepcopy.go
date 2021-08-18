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
func (in *CloudhsmV2Cluster) DeepCopyInto(out *CloudhsmV2Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2Cluster.
func (in *CloudhsmV2Cluster) DeepCopy() *CloudhsmV2Cluster {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudhsmV2Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudhsmV2ClusterList) DeepCopyInto(out *CloudhsmV2ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudhsmV2Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2ClusterList.
func (in *CloudhsmV2ClusterList) DeepCopy() *CloudhsmV2ClusterList {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudhsmV2ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudhsmV2ClusterObservation) DeepCopyInto(out *CloudhsmV2ClusterObservation) {
	*out = *in
	if in.ClusterCertificates != nil {
		in, out := &in.ClusterCertificates, &out.ClusterCertificates
		*out = make([]ClusterCertificatesObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2ClusterObservation.
func (in *CloudhsmV2ClusterObservation) DeepCopy() *CloudhsmV2ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudhsmV2ClusterParameters) DeepCopyInto(out *CloudhsmV2ClusterParameters) {
	*out = *in
	if in.SourceBackupIdentifier != nil {
		in, out := &in.SourceBackupIdentifier, &out.SourceBackupIdentifier
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2ClusterParameters.
func (in *CloudhsmV2ClusterParameters) DeepCopy() *CloudhsmV2ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudhsmV2ClusterSpec) DeepCopyInto(out *CloudhsmV2ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2ClusterSpec.
func (in *CloudhsmV2ClusterSpec) DeepCopy() *CloudhsmV2ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudhsmV2ClusterStatus) DeepCopyInto(out *CloudhsmV2ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudhsmV2ClusterStatus.
func (in *CloudhsmV2ClusterStatus) DeepCopy() *CloudhsmV2ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CloudhsmV2ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCertificatesObservation) DeepCopyInto(out *ClusterCertificatesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCertificatesObservation.
func (in *ClusterCertificatesObservation) DeepCopy() *ClusterCertificatesObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterCertificatesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCertificatesParameters) DeepCopyInto(out *ClusterCertificatesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCertificatesParameters.
func (in *ClusterCertificatesParameters) DeepCopy() *ClusterCertificatesParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterCertificatesParameters)
	in.DeepCopyInto(out)
	return out
}
