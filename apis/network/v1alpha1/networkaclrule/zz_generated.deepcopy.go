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
func (in *NetworkAclRule) DeepCopyInto(out *NetworkAclRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRule.
func (in *NetworkAclRule) DeepCopy() *NetworkAclRule {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAclRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclRuleList) DeepCopyInto(out *NetworkAclRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkAclRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRuleList.
func (in *NetworkAclRuleList) DeepCopy() *NetworkAclRuleList {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAclRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclRuleObservation) DeepCopyInto(out *NetworkAclRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRuleObservation.
func (in *NetworkAclRuleObservation) DeepCopy() *NetworkAclRuleObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclRuleParameters) DeepCopyInto(out *NetworkAclRuleParameters) {
	*out = *in
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = new(bool)
		**out = **in
	}
	if in.FromPort != nil {
		in, out := &in.FromPort, &out.FromPort
		*out = new(int64)
		**out = **in
	}
	if in.IcmpCode != nil {
		in, out := &in.IcmpCode, &out.IcmpCode
		*out = new(string)
		**out = **in
	}
	if in.IcmpType != nil {
		in, out := &in.IcmpType, &out.IcmpType
		*out = new(string)
		**out = **in
	}
	if in.Ipv6CidrBlock != nil {
		in, out := &in.Ipv6CidrBlock, &out.Ipv6CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.ToPort != nil {
		in, out := &in.ToPort, &out.ToPort
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRuleParameters.
func (in *NetworkAclRuleParameters) DeepCopy() *NetworkAclRuleParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclRuleSpec) DeepCopyInto(out *NetworkAclRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRuleSpec.
func (in *NetworkAclRuleSpec) DeepCopy() *NetworkAclRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAclRuleStatus) DeepCopyInto(out *NetworkAclRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAclRuleStatus.
func (in *NetworkAclRuleStatus) DeepCopy() *NetworkAclRuleStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkAclRuleStatus)
	in.DeepCopyInto(out)
	return out
}
