//go:build !ignore_autogenerated
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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasList) DeepCopyInto(out *AliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasList.
func (in *AliasList) DeepCopy() *AliasList {
	if in == nil {
		return nil
	}
	out := new(AliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasObservation) DeepCopyInto(out *AliasObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TargetKeyArn != nil {
		in, out := &in.TargetKeyArn, &out.TargetKeyArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasObservation.
func (in *AliasObservation) DeepCopy() *AliasObservation {
	if in == nil {
		return nil
	}
	out := new(AliasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasParameters) DeepCopyInto(out *AliasParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TargetKeyID != nil {
		in, out := &in.TargetKeyID, &out.TargetKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasParameters.
func (in *AliasParameters) DeepCopy() *AliasParameters {
	if in == nil {
		return nil
	}
	out := new(AliasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpec) DeepCopyInto(out *AliasSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpec.
func (in *AliasSpec) DeepCopy() *AliasSpec {
	if in == nil {
		return nil
	}
	out := new(AliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStatus) DeepCopyInto(out *AliasStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStatus.
func (in *AliasStatus) DeepCopy() *AliasStatus {
	if in == nil {
		return nil
	}
	out := new(AliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ciphertext) DeepCopyInto(out *Ciphertext) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ciphertext.
func (in *Ciphertext) DeepCopy() *Ciphertext {
	if in == nil {
		return nil
	}
	out := new(Ciphertext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ciphertext) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiphertextList) DeepCopyInto(out *CiphertextList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ciphertext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiphertextList.
func (in *CiphertextList) DeepCopy() *CiphertextList {
	if in == nil {
		return nil
	}
	out := new(CiphertextList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiphertextList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiphertextObservation) DeepCopyInto(out *CiphertextObservation) {
	*out = *in
	if in.CiphertextBlob != nil {
		in, out := &in.CiphertextBlob, &out.CiphertextBlob
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiphertextObservation.
func (in *CiphertextObservation) DeepCopy() *CiphertextObservation {
	if in == nil {
		return nil
	}
	out := new(CiphertextObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiphertextParameters) DeepCopyInto(out *CiphertextParameters) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	out.PlaintextSecretRef = in.PlaintextSecretRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiphertextParameters.
func (in *CiphertextParameters) DeepCopy() *CiphertextParameters {
	if in == nil {
		return nil
	}
	out := new(CiphertextParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiphertextSpec) DeepCopyInto(out *CiphertextSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiphertextSpec.
func (in *CiphertextSpec) DeepCopy() *CiphertextSpec {
	if in == nil {
		return nil
	}
	out := new(CiphertextSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiphertextStatus) DeepCopyInto(out *CiphertextStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiphertextStatus.
func (in *CiphertextStatus) DeepCopy() *CiphertextStatus {
	if in == nil {
		return nil
	}
	out := new(CiphertextStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsObservation) DeepCopyInto(out *ConstraintsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsObservation.
func (in *ConstraintsObservation) DeepCopy() *ConstraintsObservation {
	if in == nil {
		return nil
	}
	out := new(ConstraintsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintsParameters) DeepCopyInto(out *ConstraintsParameters) {
	*out = *in
	if in.EncryptionContextEquals != nil {
		in, out := &in.EncryptionContextEquals, &out.EncryptionContextEquals
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.EncryptionContextSubset != nil {
		in, out := &in.EncryptionContextSubset, &out.EncryptionContextSubset
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintsParameters.
func (in *ConstraintsParameters) DeepCopy() *ConstraintsParameters {
	if in == nil {
		return nil
	}
	out := new(ConstraintsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKey) DeepCopyInto(out *ExternalKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKey.
func (in *ExternalKey) DeepCopy() *ExternalKey {
	if in == nil {
		return nil
	}
	out := new(ExternalKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKeyList) DeepCopyInto(out *ExternalKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKeyList.
func (in *ExternalKeyList) DeepCopy() *ExternalKeyList {
	if in == nil {
		return nil
	}
	out := new(ExternalKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKeyObservation) DeepCopyInto(out *ExternalKeyObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ExpirationModel != nil {
		in, out := &in.ExpirationModel, &out.ExpirationModel
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeyState != nil {
		in, out := &in.KeyState, &out.KeyState
		*out = new(string)
		**out = **in
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKeyObservation.
func (in *ExternalKeyObservation) DeepCopy() *ExternalKeyObservation {
	if in == nil {
		return nil
	}
	out := new(ExternalKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKeyParameters) DeepCopyInto(out *ExternalKeyParameters) {
	*out = *in
	if in.BypassPolicyLockoutSafetyCheck != nil {
		in, out := &in.BypassPolicyLockoutSafetyCheck, &out.BypassPolicyLockoutSafetyCheck
		*out = new(bool)
		**out = **in
	}
	if in.DeletionWindowInDays != nil {
		in, out := &in.DeletionWindowInDays, &out.DeletionWindowInDays
		*out = new(float64)
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
	if in.KeyMaterialBase64SecretRef != nil {
		in, out := &in.KeyMaterialBase64SecretRef, &out.KeyMaterialBase64SecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ValidTo != nil {
		in, out := &in.ValidTo, &out.ValidTo
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKeyParameters.
func (in *ExternalKeyParameters) DeepCopy() *ExternalKeyParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKeySpec) DeepCopyInto(out *ExternalKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKeySpec.
func (in *ExternalKeySpec) DeepCopy() *ExternalKeySpec {
	if in == nil {
		return nil
	}
	out := new(ExternalKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalKeyStatus) DeepCopyInto(out *ExternalKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalKeyStatus.
func (in *ExternalKeyStatus) DeepCopy() *ExternalKeyStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grant) DeepCopyInto(out *Grant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grant.
func (in *Grant) DeepCopy() *Grant {
	if in == nil {
		return nil
	}
	out := new(Grant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantList) DeepCopyInto(out *GrantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantList.
func (in *GrantList) DeepCopy() *GrantList {
	if in == nil {
		return nil
	}
	out := new(GrantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantObservation) DeepCopyInto(out *GrantObservation) {
	*out = *in
	if in.GrantID != nil {
		in, out := &in.GrantID, &out.GrantID
		*out = new(string)
		**out = **in
	}
	if in.GrantToken != nil {
		in, out := &in.GrantToken, &out.GrantToken
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantObservation.
func (in *GrantObservation) DeepCopy() *GrantObservation {
	if in == nil {
		return nil
	}
	out := new(GrantObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantParameters) DeepCopyInto(out *GrantParameters) {
	*out = *in
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]ConstraintsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GrantCreationTokens != nil {
		in, out := &in.GrantCreationTokens, &out.GrantCreationTokens
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.GranteePrincipal != nil {
		in, out := &in.GranteePrincipal, &out.GranteePrincipal
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RetireOnDelete != nil {
		in, out := &in.RetireOnDelete, &out.RetireOnDelete
		*out = new(bool)
		**out = **in
	}
	if in.RetiringPrincipal != nil {
		in, out := &in.RetiringPrincipal, &out.RetiringPrincipal
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantParameters.
func (in *GrantParameters) DeepCopy() *GrantParameters {
	if in == nil {
		return nil
	}
	out := new(GrantParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantSpec) DeepCopyInto(out *GrantSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantSpec.
func (in *GrantSpec) DeepCopy() *GrantSpec {
	if in == nil {
		return nil
	}
	out := new(GrantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantStatus) DeepCopyInto(out *GrantStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantStatus.
func (in *GrantStatus) DeepCopy() *GrantStatus {
	if in == nil {
		return nil
	}
	out := new(GrantStatus)
	in.DeepCopyInto(out)
	return out
}
