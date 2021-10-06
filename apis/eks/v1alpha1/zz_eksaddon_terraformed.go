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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this EksAddon
func (mg *EksAddon) GetTerraformResourceType() string {
	return "aws_eks_addon"
}

// GetTerraformResourceIDField returns Terraform identifier field for this EksAddon
func (tr *EksAddon) GetTerraformResourceIDField() string {
	return "id"
}

// GetObservation of this EksAddon
func (tr *EksAddon) GetObservation(ctx context.Context, c resource.SecretClient) (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	if err := json.TFParser.Unmarshal(o, &base); err != nil {
		return nil, err
	}
	return base, nil
}

// SetObservation for this EksAddon
func (tr *EksAddon) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this EksAddon
func (tr *EksAddon) GetParameters(ctx context.Context, c resource.SecretClient) (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	if err := json.TFParser.Unmarshal(p, &base); err != nil {
		return nil, err
	}
	return base, nil
}

// SetParameters for this EksAddon
func (tr *EksAddon) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this EksAddon using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *EksAddon) LateInitialize(attrs []byte) (bool, error) {
	params := &EksAddonParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	li := resource.NewGenericLateInitializer(resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard),
		resource.WithZeroElemPtrFilter(resource.CNameWildcard))
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetConnectionDetails of this EksAddon
func (tr *EksAddon) GetConnectionDetails(obs map[string]interface{}) (map[string][]byte, error) {
	return nil, nil
}
