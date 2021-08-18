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

// +kubebuilder:object:generate=true
// +groupName=kinesis.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kinesis/v1alpha1"
)

type CloudwatchLoggingOptionsObservation struct {
	Id string `json:"id" tf:"id"`
}

type CloudwatchLoggingOptionsParameters struct {
	LogStreamArn string `json:"logStreamArn" tf:"log_stream_arn"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type CsvObservation struct {
}

type CsvParameters struct {
	RecordColumnDelimiter string `json:"recordColumnDelimiter" tf:"record_column_delimiter"`

	RecordRowDelimiter string `json:"recordRowDelimiter" tf:"record_row_delimiter"`
}

type InputsObservation struct {
	Id string `json:"id" tf:"id"`

	StreamNames []string `json:"streamNames" tf:"stream_names"`
}

type InputsParameters struct {
	KinesisFirehose []KinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose"`

	KinesisStream []KinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream"`

	NamePrefix string `json:"namePrefix" tf:"name_prefix"`

	Parallelism []ParallelismParameters `json:"parallelism,omitempty" tf:"parallelism"`

	ProcessingConfiguration []ProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration"`

	Schema []SchemaParameters `json:"schema" tf:"schema"`

	StartingPositionConfiguration []StartingPositionConfigurationParameters `json:"startingPositionConfiguration,omitempty" tf:"starting_position_configuration"`
}

type JsonObservation struct {
}

type JsonParameters struct {
	RecordRowPath string `json:"recordRowPath" tf:"record_row_path"`
}

type KinesisAnalyticsApplicationObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateTimestamp string `json:"createTimestamp" tf:"create_timestamp"`

	LastUpdateTimestamp string `json:"lastUpdateTimestamp" tf:"last_update_timestamp"`

	Status string `json:"status" tf:"status"`

	Version int64 `json:"version" tf:"version"`
}

type KinesisAnalyticsApplicationParameters struct {
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options"`

	Code *string `json:"code,omitempty" tf:"code"`

	Description *string `json:"description,omitempty" tf:"description"`

	Inputs []InputsParameters `json:"inputs,omitempty" tf:"inputs"`

	Name string `json:"name" tf:"name"`

	Outputs []OutputsParameters `json:"outputs,omitempty" tf:"outputs"`

	ReferenceDataSources []ReferenceDataSourcesParameters `json:"referenceDataSources,omitempty" tf:"reference_data_sources"`

	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type KinesisFirehoseObservation struct {
}

type KinesisFirehoseParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type KinesisStreamObservation struct {
}

type KinesisStreamParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type LambdaObservation struct {
}

type LambdaParameters struct {
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type MappingParametersObservation struct {
}

type MappingParametersParameters struct {
	Csv []CsvParameters `json:"csv,omitempty" tf:"csv"`

	Json []JsonParameters `json:"json,omitempty" tf:"json"`
}

type OutputsObservation struct {
	Id string `json:"id" tf:"id"`
}

type OutputsParameters struct {
	KinesisFirehose []KinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose"`

	KinesisStream []KinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream"`

	Lambda []LambdaParameters `json:"lambda,omitempty" tf:"lambda"`

	Name string `json:"name" tf:"name"`

	Schema []SchemaParameters `json:"schema" tf:"schema"`
}

type ParallelismObservation struct {
}

type ParallelismParameters struct {
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type ProcessingConfigurationObservation struct {
}

type ProcessingConfigurationParameters struct {
	Lambda []LambdaParameters `json:"lambda" tf:"lambda"`
}

type RecordColumnsObservation struct {
}

type RecordColumnsParameters struct {
	Mapping *string `json:"mapping,omitempty" tf:"mapping"`

	Name string `json:"name" tf:"name"`

	SqlType string `json:"sqlType" tf:"sql_type"`
}

type RecordFormatObservation struct {
	RecordFormatType string `json:"recordFormatType" tf:"record_format_type"`
}

type RecordFormatParameters struct {
	MappingParameters []MappingParametersParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters"`
}

type ReferenceDataSourcesObservation struct {
	Id string `json:"id" tf:"id"`
}

type ReferenceDataSourcesParameters struct {
	S3 []S3Parameters `json:"s3" tf:"s3"`

	Schema []SchemaParameters `json:"schema" tf:"schema"`

	TableName string `json:"tableName" tf:"table_name"`
}

type S3Observation struct {
}

type S3Parameters struct {
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`

	FileKey string `json:"fileKey" tf:"file_key"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {
	RecordColumns []RecordColumnsParameters `json:"recordColumns" tf:"record_columns"`

	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding"`

	RecordFormat []RecordFormatParameters `json:"recordFormat" tf:"record_format"`
}

type StartingPositionConfigurationObservation struct {
}

type StartingPositionConfigurationParameters struct {
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position"`
}

// KinesisAnalyticsApplicationSpec defines the desired state of KinesisAnalyticsApplication
type KinesisAnalyticsApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KinesisAnalyticsApplicationParameters `json:"forProvider"`
}

// KinesisAnalyticsApplicationStatus defines the observed state of KinesisAnalyticsApplication.
type KinesisAnalyticsApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KinesisAnalyticsApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisAnalyticsApplication is the Schema for the KinesisAnalyticsApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type KinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KinesisAnalyticsApplicationSpec   `json:"spec"`
	Status            KinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisAnalyticsApplicationList contains a list of KinesisAnalyticsApplications
type KinesisAnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisAnalyticsApplication `json:"items"`
}

// Repository type metadata.
var (
	KinesisAnalyticsApplicationKind             = "KinesisAnalyticsApplication"
	KinesisAnalyticsApplicationGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: KinesisAnalyticsApplicationKind}.String()
	KinesisAnalyticsApplicationKindAPIVersion   = KinesisAnalyticsApplicationKind + "." + v1alpha1.GroupVersion.String()
	KinesisAnalyticsApplicationGroupVersionKind = v1alpha1.GroupVersion.WithKind(KinesisAnalyticsApplicationKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&KinesisAnalyticsApplication{}, &KinesisAnalyticsApplicationList{})
}
