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

package version

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/lambda-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.Lambda{}
	_ = &svcapitypes.Version{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	if r.ko.Status.Version == nil {
		return nil, ackerr.NotFound
	}
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.FunctionConfiguration
	resp, err = rm.sdkapi.GetFunctionConfigurationWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetFunctionConfiguration", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Architectures != nil {
		f0 := []*string{}
		for _, f0iter := range resp.Architectures {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		ko.Status.Architectures = f0
	} else {
		ko.Status.Architectures = nil
	}
	if resp.CodeSha256 != nil {
		ko.Spec.CodeSHA256 = resp.CodeSha256
	} else {
		ko.Spec.CodeSHA256 = nil
	}
	if resp.CodeSize != nil {
		ko.Status.CodeSize = resp.CodeSize
	} else {
		ko.Status.CodeSize = nil
	}
	if resp.DeadLetterConfig != nil {
		f3 := &svcapitypes.DeadLetterConfig{}
		if resp.DeadLetterConfig.TargetArn != nil {
			f3.TargetARN = resp.DeadLetterConfig.TargetArn
		}
		ko.Status.DeadLetterConfig = f3
	} else {
		ko.Status.DeadLetterConfig = nil
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.Environment != nil {
		f5 := &svcapitypes.EnvironmentResponse{}
		if resp.Environment.Error != nil {
			f5f0 := &svcapitypes.EnvironmentError{}
			if resp.Environment.Error.ErrorCode != nil {
				f5f0.ErrorCode = resp.Environment.Error.ErrorCode
			}
			if resp.Environment.Error.Message != nil {
				f5f0.Message = resp.Environment.Error.Message
			}
			f5.Error = f5f0
		}
		if resp.Environment.Variables != nil {
			f5f1 := map[string]*string{}
			for f5f1key, f5f1valiter := range resp.Environment.Variables {
				var f5f1val string
				f5f1val = *f5f1valiter
				f5f1[f5f1key] = &f5f1val
			}
			f5.Variables = f5f1
		}
		ko.Status.Environment = f5
	} else {
		ko.Status.Environment = nil
	}
	if resp.EphemeralStorage != nil {
		f6 := &svcapitypes.EphemeralStorage{}
		if resp.EphemeralStorage.Size != nil {
			f6.Size = resp.EphemeralStorage.Size
		}
		ko.Status.EphemeralStorage = f6
	} else {
		ko.Status.EphemeralStorage = nil
	}
	if resp.FileSystemConfigs != nil {
		f7 := []*svcapitypes.FileSystemConfig{}
		for _, f7iter := range resp.FileSystemConfigs {
			f7elem := &svcapitypes.FileSystemConfig{}
			if f7iter.Arn != nil {
				f7elem.ARN = f7iter.Arn
			}
			if f7iter.LocalMountPath != nil {
				f7elem.LocalMountPath = f7iter.LocalMountPath
			}
			f7 = append(f7, f7elem)
		}
		ko.Status.FileSystemConfigs = f7
	} else {
		ko.Status.FileSystemConfigs = nil
	}
	if resp.FunctionArn != nil {
		ko.Status.FunctionARN = resp.FunctionArn
	} else {
		ko.Status.FunctionARN = nil
	}
	if resp.FunctionName != nil {
		ko.Spec.FunctionName = resp.FunctionName
	} else {
		ko.Spec.FunctionName = nil
	}
	if resp.Handler != nil {
		ko.Status.Handler = resp.Handler
	} else {
		ko.Status.Handler = nil
	}
	if resp.ImageConfigResponse != nil {
		f11 := &svcapitypes.ImageConfigResponse{}
		if resp.ImageConfigResponse.Error != nil {
			f11f0 := &svcapitypes.ImageConfigError{}
			if resp.ImageConfigResponse.Error.ErrorCode != nil {
				f11f0.ErrorCode = resp.ImageConfigResponse.Error.ErrorCode
			}
			if resp.ImageConfigResponse.Error.Message != nil {
				f11f0.Message = resp.ImageConfigResponse.Error.Message
			}
			f11.Error = f11f0
		}
		if resp.ImageConfigResponse.ImageConfig != nil {
			f11f1 := &svcapitypes.ImageConfig{}
			if resp.ImageConfigResponse.ImageConfig.Command != nil {
				f11f1f0 := []*string{}
				for _, f11f1f0iter := range resp.ImageConfigResponse.ImageConfig.Command {
					var f11f1f0elem string
					f11f1f0elem = *f11f1f0iter
					f11f1f0 = append(f11f1f0, &f11f1f0elem)
				}
				f11f1.Command = f11f1f0
			}
			if resp.ImageConfigResponse.ImageConfig.EntryPoint != nil {
				f11f1f1 := []*string{}
				for _, f11f1f1iter := range resp.ImageConfigResponse.ImageConfig.EntryPoint {
					var f11f1f1elem string
					f11f1f1elem = *f11f1f1iter
					f11f1f1 = append(f11f1f1, &f11f1f1elem)
				}
				f11f1.EntryPoint = f11f1f1
			}
			if resp.ImageConfigResponse.ImageConfig.WorkingDirectory != nil {
				f11f1.WorkingDirectory = resp.ImageConfigResponse.ImageConfig.WorkingDirectory
			}
			f11.ImageConfig = f11f1
		}
		ko.Status.ImageConfigResponse = f11
	} else {
		ko.Status.ImageConfigResponse = nil
	}
	if resp.KMSKeyArn != nil {
		ko.Status.KMSKeyARN = resp.KMSKeyArn
	} else {
		ko.Status.KMSKeyARN = nil
	}
	if resp.LastModified != nil {
		ko.Status.LastModified = resp.LastModified
	} else {
		ko.Status.LastModified = nil
	}
	if resp.LastUpdateStatus != nil {
		ko.Status.LastUpdateStatus = resp.LastUpdateStatus
	} else {
		ko.Status.LastUpdateStatus = nil
	}
	if resp.LastUpdateStatusReason != nil {
		ko.Status.LastUpdateStatusReason = resp.LastUpdateStatusReason
	} else {
		ko.Status.LastUpdateStatusReason = nil
	}
	if resp.LastUpdateStatusReasonCode != nil {
		ko.Status.LastUpdateStatusReasonCode = resp.LastUpdateStatusReasonCode
	} else {
		ko.Status.LastUpdateStatusReasonCode = nil
	}
	if resp.Layers != nil {
		f17 := []*svcapitypes.Layer{}
		for _, f17iter := range resp.Layers {
			f17elem := &svcapitypes.Layer{}
			if f17iter.Arn != nil {
				f17elem.ARN = f17iter.Arn
			}
			if f17iter.CodeSize != nil {
				f17elem.CodeSize = f17iter.CodeSize
			}
			if f17iter.SigningJobArn != nil {
				f17elem.SigningJobARN = f17iter.SigningJobArn
			}
			if f17iter.SigningProfileVersionArn != nil {
				f17elem.SigningProfileVersionARN = f17iter.SigningProfileVersionArn
			}
			f17 = append(f17, f17elem)
		}
		ko.Status.Layers = f17
	} else {
		ko.Status.Layers = nil
	}
	if resp.MasterArn != nil {
		ko.Status.MasterARN = resp.MasterArn
	} else {
		ko.Status.MasterARN = nil
	}
	if resp.MemorySize != nil {
		ko.Status.MemorySize = resp.MemorySize
	} else {
		ko.Status.MemorySize = nil
	}
	if resp.PackageType != nil {
		ko.Status.PackageType = resp.PackageType
	} else {
		ko.Status.PackageType = nil
	}
	if resp.RevisionId != nil {
		ko.Spec.RevisionID = resp.RevisionId
	} else {
		ko.Spec.RevisionID = nil
	}
	if resp.Role != nil {
		ko.Status.Role = resp.Role
	} else {
		ko.Status.Role = nil
	}
	if resp.Runtime != nil {
		ko.Status.Runtime = resp.Runtime
	} else {
		ko.Status.Runtime = nil
	}
	if resp.SigningJobArn != nil {
		ko.Status.SigningJobARN = resp.SigningJobArn
	} else {
		ko.Status.SigningJobARN = nil
	}
	if resp.SigningProfileVersionArn != nil {
		ko.Status.SigningProfileVersionARN = resp.SigningProfileVersionArn
	} else {
		ko.Status.SigningProfileVersionARN = nil
	}
	if resp.SnapStart != nil {
		f26 := &svcapitypes.SnapStartResponse{}
		if resp.SnapStart.ApplyOn != nil {
			f26.ApplyOn = resp.SnapStart.ApplyOn
		}
		if resp.SnapStart.OptimizationStatus != nil {
			f26.OptimizationStatus = resp.SnapStart.OptimizationStatus
		}
		ko.Status.SnapStart = f26
	} else {
		ko.Status.SnapStart = nil
	}
	if resp.State != nil {
		ko.Status.State = resp.State
	} else {
		ko.Status.State = nil
	}
	if resp.StateReason != nil {
		ko.Status.StateReason = resp.StateReason
	} else {
		ko.Status.StateReason = nil
	}
	if resp.StateReasonCode != nil {
		ko.Status.StateReasonCode = resp.StateReasonCode
	} else {
		ko.Status.StateReasonCode = nil
	}
	if resp.Timeout != nil {
		ko.Status.Timeout = resp.Timeout
	} else {
		ko.Status.Timeout = nil
	}
	if resp.TracingConfig != nil {
		f31 := &svcapitypes.TracingConfigResponse{}
		if resp.TracingConfig.Mode != nil {
			f31.Mode = resp.TracingConfig.Mode
		}
		ko.Status.TracingConfig = f31
	} else {
		ko.Status.TracingConfig = nil
	}
	if resp.Version != nil {
		ko.Status.Version = resp.Version
	} else {
		ko.Status.Version = nil
	}
	if resp.VpcConfig != nil {
		f33 := &svcapitypes.VPCConfigResponse{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f33f0 := []*string{}
			for _, f33f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f33f0elem string
				f33f0elem = *f33f0iter
				f33f0 = append(f33f0, &f33f0elem)
			}
			f33.SecurityGroupIDs = f33f0
		}
		if resp.VpcConfig.SubnetIds != nil {
			f33f1 := []*string{}
			for _, f33f1iter := range resp.VpcConfig.SubnetIds {
				var f33f1elem string
				f33f1elem = *f33f1iter
				f33f1 = append(f33f1, &f33f1elem)
			}
			f33.SubnetIDs = f33f1
		}
		if resp.VpcConfig.VpcId != nil {
			f33.VPCID = resp.VpcConfig.VpcId
		}
		ko.Status.VPCConfig = f33
	} else {
		ko.Status.VPCConfig = nil
	}

	rm.setStatusDefaults(ko)
	if err := rm.setResourceAdditionalFields(ctx, ko); err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.FunctionName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetFunctionConfigurationInput, error) {
	res := &svcsdk.GetFunctionConfigurationInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Status.Version != nil {
		res.SetQualifier(*r.ko.Status.Version)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	var marker *string = nil
	var versionList []*svcsdk.FunctionConfiguration
	for {
		listVersionsInput := &svcsdk.ListVersionsByFunctionInput{
			FunctionName: desired.ko.Spec.FunctionName,
			Marker:       marker,
		}
		listVersionResponse, err := rm.sdkapi.ListVersionsByFunctionWithContext(ctx, listVersionsInput)
		if err != nil {
			return nil, err
		}
		versionList = append(versionList, listVersionResponse.Versions...)

		if listVersionResponse.NextMarker == nil {
			break
		}
		marker = listVersionResponse.NextMarker
	}
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.FunctionConfiguration
	_ = resp
	resp, err = rm.sdkapi.PublishVersionWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "PublishVersion", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()
	for _, version := range versionList {
		if *version.Version == *resp.Version {
			ErrCannotCreateResource := errors.New("No changes were made to $LATEST since publishing last version, so no version was published.")
			return nil, ackerr.NewTerminalError(ErrCannotCreateResource)
		}
	}

	if resp.Architectures != nil {
		f0 := []*string{}
		for _, f0iter := range resp.Architectures {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		ko.Status.Architectures = f0
	} else {
		ko.Status.Architectures = nil
	}
	if resp.CodeSha256 != nil {
		ko.Spec.CodeSHA256 = resp.CodeSha256
	} else {
		ko.Spec.CodeSHA256 = nil
	}
	if resp.CodeSize != nil {
		ko.Status.CodeSize = resp.CodeSize
	} else {
		ko.Status.CodeSize = nil
	}
	if resp.DeadLetterConfig != nil {
		f3 := &svcapitypes.DeadLetterConfig{}
		if resp.DeadLetterConfig.TargetArn != nil {
			f3.TargetARN = resp.DeadLetterConfig.TargetArn
		}
		ko.Status.DeadLetterConfig = f3
	} else {
		ko.Status.DeadLetterConfig = nil
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.Environment != nil {
		f5 := &svcapitypes.EnvironmentResponse{}
		if resp.Environment.Error != nil {
			f5f0 := &svcapitypes.EnvironmentError{}
			if resp.Environment.Error.ErrorCode != nil {
				f5f0.ErrorCode = resp.Environment.Error.ErrorCode
			}
			if resp.Environment.Error.Message != nil {
				f5f0.Message = resp.Environment.Error.Message
			}
			f5.Error = f5f0
		}
		if resp.Environment.Variables != nil {
			f5f1 := map[string]*string{}
			for f5f1key, f5f1valiter := range resp.Environment.Variables {
				var f5f1val string
				f5f1val = *f5f1valiter
				f5f1[f5f1key] = &f5f1val
			}
			f5.Variables = f5f1
		}
		ko.Status.Environment = f5
	} else {
		ko.Status.Environment = nil
	}
	if resp.EphemeralStorage != nil {
		f6 := &svcapitypes.EphemeralStorage{}
		if resp.EphemeralStorage.Size != nil {
			f6.Size = resp.EphemeralStorage.Size
		}
		ko.Status.EphemeralStorage = f6
	} else {
		ko.Status.EphemeralStorage = nil
	}
	if resp.FileSystemConfigs != nil {
		f7 := []*svcapitypes.FileSystemConfig{}
		for _, f7iter := range resp.FileSystemConfigs {
			f7elem := &svcapitypes.FileSystemConfig{}
			if f7iter.Arn != nil {
				f7elem.ARN = f7iter.Arn
			}
			if f7iter.LocalMountPath != nil {
				f7elem.LocalMountPath = f7iter.LocalMountPath
			}
			f7 = append(f7, f7elem)
		}
		ko.Status.FileSystemConfigs = f7
	} else {
		ko.Status.FileSystemConfigs = nil
	}
	if resp.FunctionArn != nil {
		ko.Status.FunctionARN = resp.FunctionArn
	} else {
		ko.Status.FunctionARN = nil
	}
	if resp.FunctionName != nil {
		ko.Spec.FunctionName = resp.FunctionName
	} else {
		ko.Spec.FunctionName = nil
	}
	if resp.Handler != nil {
		ko.Status.Handler = resp.Handler
	} else {
		ko.Status.Handler = nil
	}
	if resp.ImageConfigResponse != nil {
		f11 := &svcapitypes.ImageConfigResponse{}
		if resp.ImageConfigResponse.Error != nil {
			f11f0 := &svcapitypes.ImageConfigError{}
			if resp.ImageConfigResponse.Error.ErrorCode != nil {
				f11f0.ErrorCode = resp.ImageConfigResponse.Error.ErrorCode
			}
			if resp.ImageConfigResponse.Error.Message != nil {
				f11f0.Message = resp.ImageConfigResponse.Error.Message
			}
			f11.Error = f11f0
		}
		if resp.ImageConfigResponse.ImageConfig != nil {
			f11f1 := &svcapitypes.ImageConfig{}
			if resp.ImageConfigResponse.ImageConfig.Command != nil {
				f11f1f0 := []*string{}
				for _, f11f1f0iter := range resp.ImageConfigResponse.ImageConfig.Command {
					var f11f1f0elem string
					f11f1f0elem = *f11f1f0iter
					f11f1f0 = append(f11f1f0, &f11f1f0elem)
				}
				f11f1.Command = f11f1f0
			}
			if resp.ImageConfigResponse.ImageConfig.EntryPoint != nil {
				f11f1f1 := []*string{}
				for _, f11f1f1iter := range resp.ImageConfigResponse.ImageConfig.EntryPoint {
					var f11f1f1elem string
					f11f1f1elem = *f11f1f1iter
					f11f1f1 = append(f11f1f1, &f11f1f1elem)
				}
				f11f1.EntryPoint = f11f1f1
			}
			if resp.ImageConfigResponse.ImageConfig.WorkingDirectory != nil {
				f11f1.WorkingDirectory = resp.ImageConfigResponse.ImageConfig.WorkingDirectory
			}
			f11.ImageConfig = f11f1
		}
		ko.Status.ImageConfigResponse = f11
	} else {
		ko.Status.ImageConfigResponse = nil
	}
	if resp.KMSKeyArn != nil {
		ko.Status.KMSKeyARN = resp.KMSKeyArn
	} else {
		ko.Status.KMSKeyARN = nil
	}
	if resp.LastModified != nil {
		ko.Status.LastModified = resp.LastModified
	} else {
		ko.Status.LastModified = nil
	}
	if resp.LastUpdateStatus != nil {
		ko.Status.LastUpdateStatus = resp.LastUpdateStatus
	} else {
		ko.Status.LastUpdateStatus = nil
	}
	if resp.LastUpdateStatusReason != nil {
		ko.Status.LastUpdateStatusReason = resp.LastUpdateStatusReason
	} else {
		ko.Status.LastUpdateStatusReason = nil
	}
	if resp.LastUpdateStatusReasonCode != nil {
		ko.Status.LastUpdateStatusReasonCode = resp.LastUpdateStatusReasonCode
	} else {
		ko.Status.LastUpdateStatusReasonCode = nil
	}
	if resp.Layers != nil {
		f17 := []*svcapitypes.Layer{}
		for _, f17iter := range resp.Layers {
			f17elem := &svcapitypes.Layer{}
			if f17iter.Arn != nil {
				f17elem.ARN = f17iter.Arn
			}
			if f17iter.CodeSize != nil {
				f17elem.CodeSize = f17iter.CodeSize
			}
			if f17iter.SigningJobArn != nil {
				f17elem.SigningJobARN = f17iter.SigningJobArn
			}
			if f17iter.SigningProfileVersionArn != nil {
				f17elem.SigningProfileVersionARN = f17iter.SigningProfileVersionArn
			}
			f17 = append(f17, f17elem)
		}
		ko.Status.Layers = f17
	} else {
		ko.Status.Layers = nil
	}
	if resp.MasterArn != nil {
		ko.Status.MasterARN = resp.MasterArn
	} else {
		ko.Status.MasterARN = nil
	}
	if resp.MemorySize != nil {
		ko.Status.MemorySize = resp.MemorySize
	} else {
		ko.Status.MemorySize = nil
	}
	if resp.PackageType != nil {
		ko.Status.PackageType = resp.PackageType
	} else {
		ko.Status.PackageType = nil
	}
	if resp.RevisionId != nil {
		ko.Spec.RevisionID = resp.RevisionId
	} else {
		ko.Spec.RevisionID = nil
	}
	if resp.Role != nil {
		ko.Status.Role = resp.Role
	} else {
		ko.Status.Role = nil
	}
	if resp.Runtime != nil {
		ko.Status.Runtime = resp.Runtime
	} else {
		ko.Status.Runtime = nil
	}
	if resp.SigningJobArn != nil {
		ko.Status.SigningJobARN = resp.SigningJobArn
	} else {
		ko.Status.SigningJobARN = nil
	}
	if resp.SigningProfileVersionArn != nil {
		ko.Status.SigningProfileVersionARN = resp.SigningProfileVersionArn
	} else {
		ko.Status.SigningProfileVersionARN = nil
	}
	if resp.SnapStart != nil {
		f26 := &svcapitypes.SnapStartResponse{}
		if resp.SnapStart.ApplyOn != nil {
			f26.ApplyOn = resp.SnapStart.ApplyOn
		}
		if resp.SnapStart.OptimizationStatus != nil {
			f26.OptimizationStatus = resp.SnapStart.OptimizationStatus
		}
		ko.Status.SnapStart = f26
	} else {
		ko.Status.SnapStart = nil
	}
	if resp.State != nil {
		ko.Status.State = resp.State
	} else {
		ko.Status.State = nil
	}
	if resp.StateReason != nil {
		ko.Status.StateReason = resp.StateReason
	} else {
		ko.Status.StateReason = nil
	}
	if resp.StateReasonCode != nil {
		ko.Status.StateReasonCode = resp.StateReasonCode
	} else {
		ko.Status.StateReasonCode = nil
	}
	if resp.Timeout != nil {
		ko.Status.Timeout = resp.Timeout
	} else {
		ko.Status.Timeout = nil
	}
	if resp.TracingConfig != nil {
		f31 := &svcapitypes.TracingConfigResponse{}
		if resp.TracingConfig.Mode != nil {
			f31.Mode = resp.TracingConfig.Mode
		}
		ko.Status.TracingConfig = f31
	} else {
		ko.Status.TracingConfig = nil
	}
	if resp.Version != nil {
		ko.Status.Version = resp.Version
	} else {
		ko.Status.Version = nil
	}
	if resp.VpcConfig != nil {
		f33 := &svcapitypes.VPCConfigResponse{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f33f0 := []*string{}
			for _, f33f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f33f0elem string
				f33f0elem = *f33f0iter
				f33f0 = append(f33f0, &f33f0elem)
			}
			f33.SecurityGroupIDs = f33f0
		}
		if resp.VpcConfig.SubnetIds != nil {
			f33f1 := []*string{}
			for _, f33f1iter := range resp.VpcConfig.SubnetIds {
				var f33f1elem string
				f33f1elem = *f33f1iter
				f33f1 = append(f33f1, &f33f1elem)
			}
			f33.SubnetIDs = f33f1
		}
		if resp.VpcConfig.VpcId != nil {
			f33.VPCID = resp.VpcConfig.VpcId
		}
		ko.Status.VPCConfig = f33
	} else {
		ko.Status.VPCConfig = nil
	}

	rm.setStatusDefaults(ko)
	if ko.Spec.FunctionEventInvokeConfig != nil {
		err = rm.syncEventInvokeConfig(ctx, desired)
		if err != nil {
			return nil, err
		}
	}
	if ko.Spec.ProvisionedConcurrencyConfig != nil {
		err = rm.updateProvisionedConcurrency(ctx, desired)
		if err != nil {
			return nil, err
		}
	}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PublishVersionInput, error) {
	res := &svcsdk.PublishVersionInput{}

	if r.ko.Spec.CodeSHA256 != nil {
		res.SetCodeSha256(*r.ko.Spec.CodeSHA256)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.RevisionID != nil {
		res.SetRevisionId(*r.ko.Spec.RevisionID)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateVersion(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteFunctionOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteFunctionWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFunction", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFunctionInput, error) {
	res := &svcsdk.DeleteFunctionInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Status.Version != nil {
		res.SetQualifier(*r.ko.Status.Version)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Version,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
