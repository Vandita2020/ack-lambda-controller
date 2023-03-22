// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FunctionSpec defines the desired state of Function.
type FunctionSpec struct {

	// The instruction set architecture that the function supports. Enter a string
	// array with one of the valid values (arm64 or x86_64). The default value is
	// x86_64.
	Architectures []*string `json:"architectures,omitempty"`
	// The code for the function.
	// +kubebuilder:validation:Required
	Code *FunctionCode `json:"code"`
	// To enable code signing for this function, specify the ARN of a code-signing
	// configuration. A code-signing configuration includes a set of signing profiles,
	// which define the trusted publishers for this function.
	CodeSigningConfigARN *string `json:"codeSigningConfigARN,omitempty"`
	// A dead-letter queue configuration that specifies the queue or topic where
	// Lambda sends asynchronous events when they fail processing. For more information,
	// see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).
	DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
	// A description of the function.
	Description *string `json:"description,omitempty"`
	// Environment variables that are accessible from function code during execution.
	Environment *Environment `json:"environment,omitempty"`
	// The size of the function's /tmp directory in MB. The default value is 512,
	// but can be any whole number between 512 and 10,240 MB.
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`
	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*FileSystemConfig `json:"fileSystemConfigs,omitempty"`
	// Configures options for asynchronous invocation on a function.
	//
	// - DestinationConfig
	// A destination for events after they have been sent to a function for processing.
	//
	// Types of Destinations:
	// Function - The Amazon Resource Name (ARN) of a Lambda function.
	// Queue - The ARN of a standard SQS queue.
	// Topic - The ARN of a standard SNS topic.
	// Event Bus - The ARN of an Amazon EventBridge event bus.
	//
	// - MaximumEventAgeInSeconds
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// - MaximumRetryAttempts
	// The maximum number of times to retry when the function returns an error.
	FunctionEventInvokeConfig *PutFunctionEventInvokeConfigInput `json:"functionEventInvokeConfig,omitempty"`
	// The name of the method within your code that Lambda calls to run your function.
	// Handler is required if the deployment package is a .zip file archive. The
	// format includes the file name. It can also include namespaces and other qualifiers,
	// depending on the runtime. For more information, see Lambda programming model
	// (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).
	Handler *string `json:"handler,omitempty"`
	// Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings)
	// that override the values in the container image Dockerfile.
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
	// The ARN of the Key Management Service (KMS) key that's used to encrypt your
	// function's environment variables. If it's not provided, Lambda uses a default
	// service key.
	KMSKeyARN *string                                  `json:"kmsKeyARN,omitempty"`
	KMSKeyRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"kmsKeyRef,omitempty"`
	// A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	// to add to the function's execution environment. Specify each layer by its
	// ARN, including the version.
	Layers []*string `json:"layers,omitempty"`
	// The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console)
	// at runtime. Increasing the function memory also increases its CPU allocation.
	// The default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize *int64 `json:"memorySize,omitempty"`
	// The name of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – my-function.
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//   - Partial ARN – 123456789012:function:my-function.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The type of deployment package. Set to Image for container image and set
	// to Zip for .zip file archive.
	PackageType *string `json:"packageType,omitempty"`
	// Set to true to publish the first version of the function during creation.
	Publish *bool `json:"publish,omitempty"`
	// The Amazon Resource Name (ARN) of the function's execution role.
	Role    *string                                  `json:"role,omitempty"`
	RoleRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"roleRef,omitempty"`
	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	// Runtime is required if the deployment package is a .zip file archive.
	Runtime *string `json:"runtime,omitempty"`
	// The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
	// setting.
	SnapStart *SnapStart `json:"snapStart,omitempty"`
	// A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html)
	// to apply to the function.
	Tags map[string]*string `json:"tags,omitempty"`
	// The amount of time (in seconds) that Lambda allows a function to run before
	// stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
	// For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).
	Timeout *int64 `json:"timeout,omitempty"`
	// Set Mode to Active to sample and trace a subset of incoming requests with
	// X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty"`
	// For network connectivity to Amazon Web Services resources in a VPC, specify
	// a list of security groups and subnets in the VPC. When you connect a function
	// to a VPC, it can access resources and the internet only through that VPC.
	// For more information, see Configuring a Lambda function to access resources
	// in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
	VPCConfig *VPCConfig `json:"vpcConfig,omitempty"`
}

// FunctionStatus defines the observed state of Function
type FunctionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The SHA256 hash of the function's deployment package.
	// +kubebuilder:validation:Optional
	CodeSHA256 *string `json:"codeSHA256,omitempty"`
	// The size of the function's deployment package, in bytes.
	// +kubebuilder:validation:Optional
	CodeSize *int64 `json:"codeSize,omitempty"`
	// The function's image configuration values.
	// +kubebuilder:validation:Optional
	ImageConfigResponse *ImageConfigResponse `json:"imageConfigResponse,omitempty"`
	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	// +kubebuilder:validation:Optional
	LastModified *string `json:"lastModified,omitempty"`
	// The status of the last update that was performed on the function. This is
	// first set to Successful after function creation completes.
	// +kubebuilder:validation:Optional
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty"`
	// The reason for the last update that was performed on the function.
	// +kubebuilder:validation:Optional
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty"`
	// The reason code for the last update that was performed on the function.
	// +kubebuilder:validation:Optional
	LastUpdateStatusReasonCode *string `json:"lastUpdateStatusReasonCode,omitempty"`
	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	// +kubebuilder:validation:Optional
	LayerStatuses []*Layer `json:"layerStatuses,omitempty"`
	// For Lambda@Edge functions, the ARN of the main function.
	// +kubebuilder:validation:Optional
	MasterARN *string `json:"masterARN,omitempty"`
	// The latest updated revision of the function or alias.
	// +kubebuilder:validation:Optional
	RevisionID *string `json:"revisionID,omitempty"`
	// The ARN of the signing job.
	// +kubebuilder:validation:Optional
	SigningJobARN *string `json:"signingJobARN,omitempty"`
	// The ARN of the signing profile version.
	// +kubebuilder:validation:Optional
	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
	// The current state of the function. When the state is Inactive, you can reactivate
	// the function by invoking it.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
	// The reason for the function's current state.
	// +kubebuilder:validation:Optional
	StateReason *string `json:"stateReason,omitempty"`
	// The reason code for the function's current state. When the code is Creating,
	// you can't invoke or modify the function.
	// +kubebuilder:validation:Optional
	StateReasonCode *string `json:"stateReasonCode,omitempty"`
	// The version of the Lambda function.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty"`
}

// Function is the Schema for the Functions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec,omitempty"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// FunctionList contains a list of Function
// +kubebuilder:object:root=true
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
