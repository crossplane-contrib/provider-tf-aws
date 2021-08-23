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
func (in *Apigatewayv2IntegrationResponse) DeepCopyInto(out *Apigatewayv2IntegrationResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponse.
func (in *Apigatewayv2IntegrationResponse) DeepCopy() *Apigatewayv2IntegrationResponse {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2IntegrationResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2IntegrationResponseList) DeepCopyInto(out *Apigatewayv2IntegrationResponseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Apigatewayv2IntegrationResponse, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponseList.
func (in *Apigatewayv2IntegrationResponseList) DeepCopy() *Apigatewayv2IntegrationResponseList {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Apigatewayv2IntegrationResponseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2IntegrationResponseObservation) DeepCopyInto(out *Apigatewayv2IntegrationResponseObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponseObservation.
func (in *Apigatewayv2IntegrationResponseObservation) DeepCopy() *Apigatewayv2IntegrationResponseObservation {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2IntegrationResponseParameters) DeepCopyInto(out *Apigatewayv2IntegrationResponseParameters) {
	*out = *in
	if in.ContentHandlingStrategy != nil {
		in, out := &in.ContentHandlingStrategy, &out.ContentHandlingStrategy
		*out = new(string)
		**out = **in
	}
	if in.ResponseTemplates != nil {
		in, out := &in.ResponseTemplates, &out.ResponseTemplates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TemplateSelectionExpression != nil {
		in, out := &in.TemplateSelectionExpression, &out.TemplateSelectionExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponseParameters.
func (in *Apigatewayv2IntegrationResponseParameters) DeepCopy() *Apigatewayv2IntegrationResponseParameters {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2IntegrationResponseSpec) DeepCopyInto(out *Apigatewayv2IntegrationResponseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponseSpec.
func (in *Apigatewayv2IntegrationResponseSpec) DeepCopy() *Apigatewayv2IntegrationResponseSpec {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Apigatewayv2IntegrationResponseStatus) DeepCopyInto(out *Apigatewayv2IntegrationResponseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Apigatewayv2IntegrationResponseStatus.
func (in *Apigatewayv2IntegrationResponseStatus) DeepCopy() *Apigatewayv2IntegrationResponseStatus {
	if in == nil {
		return nil
	}
	out := new(Apigatewayv2IntegrationResponseStatus)
	in.DeepCopyInto(out)
	return out
}
