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
// +groupName=iot.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iot/v1alpha1"
)

type CloudwatchAlarmObservation struct {
}

type CloudwatchAlarmParameters struct {
	AlarmName string `json:"alarmName" tf:"alarm_name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	StateReason string `json:"stateReason" tf:"state_reason"`

	StateValue string `json:"stateValue" tf:"state_value"`
}

type CloudwatchMetricObservation struct {
}

type CloudwatchMetricParameters struct {
	MetricName string `json:"metricName" tf:"metric_name"`

	MetricNamespace string `json:"metricNamespace" tf:"metric_namespace"`

	MetricTimestamp *string `json:"metricTimestamp,omitempty" tf:"metric_timestamp"`

	MetricUnit string `json:"metricUnit" tf:"metric_unit"`

	MetricValue string `json:"metricValue" tf:"metric_value"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type DynamodbObservation struct {
}

type DynamodbParameters struct {
	HashKeyField string `json:"hashKeyField" tf:"hash_key_field"`

	HashKeyType *string `json:"hashKeyType,omitempty" tf:"hash_key_type"`

	HashKeyValue string `json:"hashKeyValue" tf:"hash_key_value"`

	Operation *string `json:"operation,omitempty" tf:"operation"`

	PayloadField *string `json:"payloadField,omitempty" tf:"payload_field"`

	RangeKeyField *string `json:"rangeKeyField,omitempty" tf:"range_key_field"`

	RangeKeyType *string `json:"rangeKeyType,omitempty" tf:"range_key_type"`

	RangeKeyValue *string `json:"rangeKeyValue,omitempty" tf:"range_key_value"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	TableName string `json:"tableName" tf:"table_name"`
}

type Dynamodbv2Observation struct {
}

type Dynamodbv2Parameters struct {
	PutItem []PutItemParameters `json:"putItem,omitempty" tf:"put_item"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type ElasticsearchObservation struct {
}

type ElasticsearchParameters struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`

	Id string `json:"id" tf:"id"`

	Index string `json:"index" tf:"index"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Type string `json:"type" tf:"type"`
}

type ErrorActionObservation struct {
}

type ErrorActionParameters struct {
	CloudwatchAlarm []CloudwatchAlarmParameters `json:"cloudwatchAlarm,omitempty" tf:"cloudwatch_alarm"`

	CloudwatchMetric []CloudwatchMetricParameters `json:"cloudwatchMetric,omitempty" tf:"cloudwatch_metric"`

	Dynamodb []DynamodbParameters `json:"dynamodb,omitempty" tf:"dynamodb"`

	Dynamodbv2 []Dynamodbv2Parameters `json:"dynamodbv2,omitempty" tf:"dynamodbv2"`

	Elasticsearch []ElasticsearchParameters `json:"elasticsearch,omitempty" tf:"elasticsearch"`

	Firehose []FirehoseParameters `json:"firehose,omitempty" tf:"firehose"`

	IotAnalytics []IotAnalyticsParameters `json:"iotAnalytics,omitempty" tf:"iot_analytics"`

	IotEvents []IotEventsParameters `json:"iotEvents,omitempty" tf:"iot_events"`

	Kinesis []KinesisParameters `json:"kinesis,omitempty" tf:"kinesis"`

	Lambda []LambdaParameters `json:"lambda,omitempty" tf:"lambda"`

	Republish []RepublishParameters `json:"republish,omitempty" tf:"republish"`

	S3 []S3Parameters `json:"s3,omitempty" tf:"s3"`

	Sns []SnsParameters `json:"sns,omitempty" tf:"sns"`

	Sqs []SqsParameters `json:"sqs,omitempty" tf:"sqs"`

	StepFunctions []StepFunctionsParameters `json:"stepFunctions,omitempty" tf:"step_functions"`
}

type FirehoseObservation struct {
}

type FirehoseParameters struct {
	DeliveryStreamName string `json:"deliveryStreamName" tf:"delivery_stream_name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Separator *string `json:"separator,omitempty" tf:"separator"`
}

type IotAnalyticsObservation struct {
}

type IotAnalyticsParameters struct {
	ChannelName string `json:"channelName" tf:"channel_name"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type IotEventsObservation struct {
}

type IotEventsParameters struct {
	InputName string `json:"inputName" tf:"input_name"`

	MessageId *string `json:"messageId,omitempty" tf:"message_id"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type IotTopicRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type IotTopicRuleParameters struct {
	CloudwatchAlarm []CloudwatchAlarmParameters `json:"cloudwatchAlarm,omitempty" tf:"cloudwatch_alarm"`

	CloudwatchMetric []CloudwatchMetricParameters `json:"cloudwatchMetric,omitempty" tf:"cloudwatch_metric"`

	Description *string `json:"description,omitempty" tf:"description"`

	Dynamodb []DynamodbParameters `json:"dynamodb,omitempty" tf:"dynamodb"`

	Dynamodbv2 []Dynamodbv2Parameters `json:"dynamodbv2,omitempty" tf:"dynamodbv2"`

	Elasticsearch []ElasticsearchParameters `json:"elasticsearch,omitempty" tf:"elasticsearch"`

	Enabled bool `json:"enabled" tf:"enabled"`

	ErrorAction []ErrorActionParameters `json:"errorAction,omitempty" tf:"error_action"`

	Firehose []FirehoseParameters `json:"firehose,omitempty" tf:"firehose"`

	IotAnalytics []IotAnalyticsParameters `json:"iotAnalytics,omitempty" tf:"iot_analytics"`

	IotEvents []IotEventsParameters `json:"iotEvents,omitempty" tf:"iot_events"`

	Kinesis []KinesisParameters `json:"kinesis,omitempty" tf:"kinesis"`

	Lambda []LambdaParameters `json:"lambda,omitempty" tf:"lambda"`

	Name string `json:"name" tf:"name"`

	Republish []RepublishParameters `json:"republish,omitempty" tf:"republish"`

	S3 []S3Parameters `json:"s3,omitempty" tf:"s3"`

	Sns []SnsParameters `json:"sns,omitempty" tf:"sns"`

	Sql string `json:"sql" tf:"sql"`

	SqlVersion string `json:"sqlVersion" tf:"sql_version"`

	Sqs []SqsParameters `json:"sqs,omitempty" tf:"sqs"`

	StepFunctions []StepFunctionsParameters `json:"stepFunctions,omitempty" tf:"step_functions"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type KinesisObservation struct {
}

type KinesisParameters struct {
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	StreamName string `json:"streamName" tf:"stream_name"`
}

type LambdaObservation struct {
}

type LambdaParameters struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

type PutItemObservation struct {
}

type PutItemParameters struct {
	TableName string `json:"tableName" tf:"table_name"`
}

type RepublishObservation struct {
}

type RepublishParameters struct {
	Qos *int64 `json:"qos,omitempty" tf:"qos"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	Topic string `json:"topic" tf:"topic"`
}

type S3Observation struct {
}

type S3Parameters struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`

	Key string `json:"key" tf:"key"`

	RoleArn string `json:"roleArn" tf:"role_arn"`
}

type SnsObservation struct {
}

type SnsParameters struct {
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	TargetArn string `json:"targetArn" tf:"target_arn"`
}

type SqsObservation struct {
}

type SqsParameters struct {
	QueueUrl string `json:"queueUrl" tf:"queue_url"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	UseBase64 bool `json:"useBase64" tf:"use_base64"`
}

type StepFunctionsObservation struct {
}

type StepFunctionsParameters struct {
	ExecutionNamePrefix *string `json:"executionNamePrefix,omitempty" tf:"execution_name_prefix"`

	RoleArn string `json:"roleArn" tf:"role_arn"`

	StateMachineName string `json:"stateMachineName" tf:"state_machine_name"`
}

// IotTopicRuleSpec defines the desired state of IotTopicRule
type IotTopicRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IotTopicRuleParameters `json:"forProvider"`
}

// IotTopicRuleStatus defines the observed state of IotTopicRule.
type IotTopicRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IotTopicRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IotTopicRule is the Schema for the IotTopicRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IotTopicRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotTopicRuleSpec   `json:"spec"`
	Status            IotTopicRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotTopicRuleList contains a list of IotTopicRules
type IotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotTopicRule `json:"items"`
}

// Repository type metadata.
var (
	IotTopicRuleKind             = "IotTopicRule"
	IotTopicRuleGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: IotTopicRuleKind}.String()
	IotTopicRuleKindAPIVersion   = IotTopicRuleKind + "." + v1alpha1.GroupVersion.String()
	IotTopicRuleGroupVersionKind = v1alpha1.GroupVersion.WithKind(IotTopicRuleKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&IotTopicRule{}, &IotTopicRuleList{})
}
