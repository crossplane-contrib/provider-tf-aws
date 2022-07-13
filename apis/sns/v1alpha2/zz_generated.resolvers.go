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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import (
	"context"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2"
	common "github.com/crossplane-contrib/provider-jet-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArn")
	}
	mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationFailureFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArn")
	}
	mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationSuccessFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArn")
	}
	mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirehoseFailureFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArn")
	}
	mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FirehoseSuccessFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPFailureFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.HTTPFailureFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.HTTPFailureFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HTTPFailureFeedbackRoleArn")
	}
	mg.Spec.ForProvider.HTTPFailureFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HTTPFailureFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArn")
	}
	mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HTTPSuccessFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LambdaFailureFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.LambdaFailureFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.LambdaFailureFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LambdaFailureFeedbackRoleArn")
	}
	mg.Spec.ForProvider.LambdaFailureFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LambdaFailureFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArn")
	}
	mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LambdaSuccessFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SqsFailureFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.SqsFailureFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.SqsFailureFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SqsFailureFeedbackRoleArn")
	}
	mg.Spec.ForProvider.SqsFailureFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SqsFailureFeedbackRoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SqsSuccessFeedbackRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.SqsSuccessFeedbackRoleArnRef,
		Selector:     mg.Spec.ForProvider.SqsSuccessFeedbackRoleArnSelector,
		To: reference.To{
			List:    &v1alpha2.RoleList{},
			Managed: &v1alpha2.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SqsSuccessFeedbackRoleArn")
	}
	mg.Spec.ForProvider.SqsSuccessFeedbackRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SqsSuccessFeedbackRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicPolicy.
func (mg *TopicPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Arn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ArnRef,
		Selector:     mg.Spec.ForProvider.ArnSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Arn")
	}
	mg.Spec.ForProvider.Arn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicSubscription.
func (mg *TopicSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.TopicArnRef,
		Selector:     mg.Spec.ForProvider.TopicArnSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicArn")
	}
	mg.Spec.ForProvider.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicArnRef = rsp.ResolvedReference

	return nil
}
