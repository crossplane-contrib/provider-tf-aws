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
// +groupName=cognito.aws.tf.crossplane.io
// +versionName=v1alpha1

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/cognito/v1alpha1"
)

type AccountRecoverySettingObservation struct {
}

type AccountRecoverySettingParameters struct {
	RecoveryMechanism []RecoveryMechanismParameters `json:"recoveryMechanism" tf:"recovery_mechanism"`
}

type AdminCreateUserConfigObservation struct {
}

type AdminCreateUserConfigParameters struct {
	AllowAdminCreateUserOnly *bool `json:"allowAdminCreateUserOnly,omitempty" tf:"allow_admin_create_user_only"`

	InviteMessageTemplate []InviteMessageTemplateParameters `json:"inviteMessageTemplate,omitempty" tf:"invite_message_template"`
}

type CognitoUserPoolObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreationDate string `json:"creationDate" tf:"creation_date"`

	CustomDomain string `json:"customDomain" tf:"custom_domain"`

	Domain string `json:"domain" tf:"domain"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	EstimatedNumberOfUsers int64 `json:"estimatedNumberOfUsers" tf:"estimated_number_of_users"`

	LastModifiedDate string `json:"lastModifiedDate" tf:"last_modified_date"`
}

type CognitoUserPoolParameters struct {
	AccountRecoverySetting []AccountRecoverySettingParameters `json:"accountRecoverySetting,omitempty" tf:"account_recovery_setting"`

	AdminCreateUserConfig []AdminCreateUserConfigParameters `json:"adminCreateUserConfig,omitempty" tf:"admin_create_user_config"`

	AliasAttributes []string `json:"aliasAttributes,omitempty" tf:"alias_attributes"`

	AutoVerifiedAttributes []string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes"`

	DeviceConfiguration []DeviceConfigurationParameters `json:"deviceConfiguration,omitempty" tf:"device_configuration"`

	EmailConfiguration []EmailConfigurationParameters `json:"emailConfiguration,omitempty" tf:"email_configuration"`

	EmailVerificationMessage *string `json:"emailVerificationMessage,omitempty" tf:"email_verification_message"`

	EmailVerificationSubject *string `json:"emailVerificationSubject,omitempty" tf:"email_verification_subject"`

	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config"`

	MfaConfiguration *string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration"`

	Name string `json:"name" tf:"name"`

	PasswordPolicy []PasswordPolicyParameters `json:"passwordPolicy,omitempty" tf:"password_policy"`

	Schema []SchemaParameters `json:"schema,omitempty" tf:"schema"`

	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message"`

	SmsConfiguration []SmsConfigurationParameters `json:"smsConfiguration,omitempty" tf:"sms_configuration"`

	SmsVerificationMessage *string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message"`

	SoftwareTokenMfaConfiguration []SoftwareTokenMfaConfigurationParameters `json:"softwareTokenMfaConfiguration,omitempty" tf:"software_token_mfa_configuration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UserPoolAddOns []UserPoolAddOnsParameters `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons"`

	UsernameAttributes []string `json:"usernameAttributes,omitempty" tf:"username_attributes"`

	UsernameConfiguration []UsernameConfigurationParameters `json:"usernameConfiguration,omitempty" tf:"username_configuration"`

	VerificationMessageTemplate []VerificationMessageTemplateParameters `json:"verificationMessageTemplate,omitempty" tf:"verification_message_template"`
}

type CustomEmailSenderObservation struct {
}

type CustomEmailSenderParameters struct {
	LambdaArn string `json:"lambdaArn" tf:"lambda_arn"`

	LambdaVersion string `json:"lambdaVersion" tf:"lambda_version"`
}

type CustomSmsSenderObservation struct {
}

type CustomSmsSenderParameters struct {
	LambdaArn string `json:"lambdaArn" tf:"lambda_arn"`

	LambdaVersion string `json:"lambdaVersion" tf:"lambda_version"`
}

type DeviceConfigurationObservation struct {
}

type DeviceConfigurationParameters struct {
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device"`

	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt"`
}

type EmailConfigurationObservation struct {
}

type EmailConfigurationParameters struct {
	ConfigurationSet *string `json:"configurationSet,omitempty" tf:"configuration_set"`

	EmailSendingAccount *string `json:"emailSendingAccount,omitempty" tf:"email_sending_account"`

	FromEmailAddress *string `json:"fromEmailAddress,omitempty" tf:"from_email_address"`

	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address"`

	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn"`
}

type InviteMessageTemplateObservation struct {
}

type InviteMessageTemplateParameters struct {
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message"`

	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject"`

	SmsMessage *string `json:"smsMessage,omitempty" tf:"sms_message"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {
	CreateAuthChallenge *string `json:"createAuthChallenge,omitempty" tf:"create_auth_challenge"`

	CustomEmailSender []CustomEmailSenderParameters `json:"customEmailSender,omitempty" tf:"custom_email_sender"`

	CustomMessage *string `json:"customMessage,omitempty" tf:"custom_message"`

	CustomSmsSender []CustomSmsSenderParameters `json:"customSmsSender,omitempty" tf:"custom_sms_sender"`

	DefineAuthChallenge *string `json:"defineAuthChallenge,omitempty" tf:"define_auth_challenge"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	PostAuthentication *string `json:"postAuthentication,omitempty" tf:"post_authentication"`

	PostConfirmation *string `json:"postConfirmation,omitempty" tf:"post_confirmation"`

	PreAuthentication *string `json:"preAuthentication,omitempty" tf:"pre_authentication"`

	PreSignUp *string `json:"preSignUp,omitempty" tf:"pre_sign_up"`

	PreTokenGeneration *string `json:"preTokenGeneration,omitempty" tf:"pre_token_generation"`

	UserMigration *string `json:"userMigration,omitempty" tf:"user_migration"`

	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse,omitempty" tf:"verify_auth_challenge_response"`
}

type NumberAttributeConstraintsObservation struct {
}

type NumberAttributeConstraintsParameters struct {
	MaxValue *string `json:"maxValue,omitempty" tf:"max_value"`

	MinValue *string `json:"minValue,omitempty" tf:"min_value"`
}

type PasswordPolicyObservation struct {
}

type PasswordPolicyParameters struct {
	MinimumLength *int64 `json:"minimumLength,omitempty" tf:"minimum_length"`

	RequireLowercase *bool `json:"requireLowercase,omitempty" tf:"require_lowercase"`

	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers"`

	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols"`

	RequireUppercase *bool `json:"requireUppercase,omitempty" tf:"require_uppercase"`

	TemporaryPasswordValidityDays *int64 `json:"temporaryPasswordValidityDays,omitempty" tf:"temporary_password_validity_days"`
}

type RecoveryMechanismObservation struct {
}

type RecoveryMechanismParameters struct {
	Name string `json:"name" tf:"name"`

	Priority int64 `json:"priority" tf:"priority"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {
	AttributeDataType string `json:"attributeDataType" tf:"attribute_data_type"`

	DeveloperOnlyAttribute *bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute"`

	Mutable *bool `json:"mutable,omitempty" tf:"mutable"`

	Name string `json:"name" tf:"name"`

	NumberAttributeConstraints []NumberAttributeConstraintsParameters `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints"`

	Required *bool `json:"required,omitempty" tf:"required"`

	StringAttributeConstraints []StringAttributeConstraintsParameters `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints"`
}

type SmsConfigurationObservation struct {
}

type SmsConfigurationParameters struct {
	ExternalId string `json:"externalId" tf:"external_id"`

	SnsCallerArn string `json:"snsCallerArn" tf:"sns_caller_arn"`
}

type SoftwareTokenMfaConfigurationObservation struct {
}

type SoftwareTokenMfaConfigurationParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type StringAttributeConstraintsObservation struct {
}

type StringAttributeConstraintsParameters struct {
	MaxLength *string `json:"maxLength,omitempty" tf:"max_length"`

	MinLength *string `json:"minLength,omitempty" tf:"min_length"`
}

type UserPoolAddOnsObservation struct {
}

type UserPoolAddOnsParameters struct {
	AdvancedSecurityMode string `json:"advancedSecurityMode" tf:"advanced_security_mode"`
}

type UsernameConfigurationObservation struct {
}

type UsernameConfigurationParameters struct {
	CaseSensitive bool `json:"caseSensitive" tf:"case_sensitive"`
}

type VerificationMessageTemplateObservation struct {
}

type VerificationMessageTemplateParameters struct {
	DefaultEmailOption *string `json:"defaultEmailOption,omitempty" tf:"default_email_option"`

	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message"`

	EmailMessageByLink *string `json:"emailMessageByLink,omitempty" tf:"email_message_by_link"`

	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject"`

	EmailSubjectByLink *string `json:"emailSubjectByLink,omitempty" tf:"email_subject_by_link"`

	SmsMessage *string `json:"smsMessage,omitempty" tf:"sms_message"`
}

// CognitoUserPoolSpec defines the desired state of CognitoUserPool
type CognitoUserPoolSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoUserPoolParameters `json:"forProvider"`
}

// CognitoUserPoolStatus defines the observed state of CognitoUserPool.
type CognitoUserPoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoUserPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPool is the Schema for the CognitoUserPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolSpec   `json:"spec"`
	Status            CognitoUserPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoUserPoolList contains a list of CognitoUserPools
type CognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoUserPool `json:"items"`
}

// Repository type metadata.
var (
	CognitoUserPoolKind             = "CognitoUserPool"
	CognitoUserPoolGroupKind        = schema.GroupKind{Group: v1alpha1.Group, Kind: CognitoUserPoolKind}.String()
	CognitoUserPoolKindAPIVersion   = CognitoUserPoolKind + "." + v1alpha1.GroupVersion.String()
	CognitoUserPoolGroupVersionKind = v1alpha1.GroupVersion.WithKind(CognitoUserPoolKind)
)

func init() {
	v1alpha1.SchemeBuilder.Register(&CognitoUserPool{}, &CognitoUserPoolList{})
}
