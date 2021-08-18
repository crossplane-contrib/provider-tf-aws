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
func (in *ConversationLogsObservation) DeepCopyInto(out *ConversationLogsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversationLogsObservation.
func (in *ConversationLogsObservation) DeepCopy() *ConversationLogsObservation {
	if in == nil {
		return nil
	}
	out := new(ConversationLogsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversationLogsParameters) DeepCopyInto(out *ConversationLogsParameters) {
	*out = *in
	if in.LogSettings != nil {
		in, out := &in.LogSettings, &out.LogSettings
		*out = make([]LogSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversationLogsParameters.
func (in *ConversationLogsParameters) DeepCopy() *ConversationLogsParameters {
	if in == nil {
		return nil
	}
	out := new(ConversationLogsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAlias) DeepCopyInto(out *LexBotAlias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAlias.
func (in *LexBotAlias) DeepCopy() *LexBotAlias {
	if in == nil {
		return nil
	}
	out := new(LexBotAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LexBotAlias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAliasList) DeepCopyInto(out *LexBotAliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LexBotAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAliasList.
func (in *LexBotAliasList) DeepCopy() *LexBotAliasList {
	if in == nil {
		return nil
	}
	out := new(LexBotAliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LexBotAliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAliasObservation) DeepCopyInto(out *LexBotAliasObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAliasObservation.
func (in *LexBotAliasObservation) DeepCopy() *LexBotAliasObservation {
	if in == nil {
		return nil
	}
	out := new(LexBotAliasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAliasParameters) DeepCopyInto(out *LexBotAliasParameters) {
	*out = *in
	if in.ConversationLogs != nil {
		in, out := &in.ConversationLogs, &out.ConversationLogs
		*out = make([]ConversationLogsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAliasParameters.
func (in *LexBotAliasParameters) DeepCopy() *LexBotAliasParameters {
	if in == nil {
		return nil
	}
	out := new(LexBotAliasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAliasSpec) DeepCopyInto(out *LexBotAliasSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAliasSpec.
func (in *LexBotAliasSpec) DeepCopy() *LexBotAliasSpec {
	if in == nil {
		return nil
	}
	out := new(LexBotAliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LexBotAliasStatus) DeepCopyInto(out *LexBotAliasStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LexBotAliasStatus.
func (in *LexBotAliasStatus) DeepCopy() *LexBotAliasStatus {
	if in == nil {
		return nil
	}
	out := new(LexBotAliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettingsObservation) DeepCopyInto(out *LogSettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettingsObservation.
func (in *LogSettingsObservation) DeepCopy() *LogSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(LogSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSettingsParameters) DeepCopyInto(out *LogSettingsParameters) {
	*out = *in
	if in.KmsKeyArn != nil {
		in, out := &in.KmsKeyArn, &out.KmsKeyArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSettingsParameters.
func (in *LogSettingsParameters) DeepCopy() *LogSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(LogSettingsParameters)
	in.DeepCopyInto(out)
	return out
}
