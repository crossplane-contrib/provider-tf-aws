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

import "github.com/crossplane-contrib/terrajet/pkg/conversion"

func (mg *Apigatewayv2IntegrationResponse) GetTerraformResourceType() string {
	return "aws_apigatewayv2_integration_response"
}

func (tr *Apigatewayv2IntegrationResponse) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this Apigatewayv2IntegrationResponse
func (tr *Apigatewayv2IntegrationResponse) GetObservation() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Status.AtProvider)
}

// SetObservation for this Apigatewayv2IntegrationResponse
func (tr *Apigatewayv2IntegrationResponse) SetObservation(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Status.AtProvider)
}

// GetParameters of this Apigatewayv2IntegrationResponse
func (tr *Apigatewayv2IntegrationResponse) GetParameters() ([]byte, error) {
	return conversion.TFParser.Marshal(tr.Spec.ForProvider)
}

// SetParameters for this Apigatewayv2IntegrationResponse
func (tr *Apigatewayv2IntegrationResponse) SetParameters(data []byte) error {
	return conversion.TFParser.Unmarshal(data, &tr.Spec.ForProvider)
}
