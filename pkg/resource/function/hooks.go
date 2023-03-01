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

package function

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"

	svcapitypes "github.com/aws-controllers-k8s/lambda-controller/apis/v1alpha1"
)

var (
	ErrFunctionPending      = errors.New("Function in 'Pending' state, cannot be modified or deleted")
	ErrCannotSetFunctionCSC = errors.New("cannot set function code signing config when package type is Image")
)

var (
	requeueWaitWhilePending = ackrequeue.NeededAfter(
		ErrFunctionPending,
		5*time.Second,
	)
)

// isFunctionPending returns true if the supplied Lambda Function is in a pending
// state
func isFunctionPending(r *resource) bool {
	if r.ko.Status.State == nil {
		return false
	}
	state := *r.ko.Status.State
	return state == string(svcapitypes.State_Pending)
}

// customUpdateFunction patches each of the resource properties in the backend AWS
// service API and returns a new resource with updated fields.
func (rm *resourceManager) customUpdateFunction(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.customUpdateFunction")
	defer exit(err)

	if isFunctionPending(desired) {
		return nil, requeueWaitWhilePending
	}

	if delta.DifferentAt("Spec.Tags") {
		err = rm.updateFunctionTags(ctx, latest, desired)
		if err != nil {
			return nil, err
		}
	}
	if delta.DifferentAt("Spec.ReservedConcurrentExecutions") {
		err = rm.updateFunctionConcurrency(ctx, desired)
		if err != nil {
			return nil, err
		}
	}
	if delta.DifferentAt("Spec.CodeSigningConfigARN") {
		if desired.ko.Spec.PackageType != nil && *desired.ko.Spec.PackageType == "Image" &&
			desired.ko.Spec.CodeSigningConfigARN != nil && *desired.ko.Spec.CodeSigningConfigARN != "" {
			return nil, ackerr.NewTerminalError(ErrCannotSetFunctionCSC)
		} else {
			err = rm.updateFunctionCodeSigningConfig(ctx, desired)
			if err != nil {
				return nil, err
			}
		}
	}

	// Only try to update Spec.Code or Spec.Configuration at once. It is
	// not correct to sequentially call UpdateFunctionConfiguration and
	// UpdateFunctionCode because both of them can put the function in a
	// Pending state.
	switch {
	case delta.DifferentAt("Spec.Code"):
		err = rm.updateFunctionCode(ctx, desired, delta)
		if err != nil {
			return nil, err
		}
	case delta.DifferentExcept(
		"Spec.Code",
		"Spec.Tags",
		"Spec.ReservedConcurrentExecutions",
		"Spec.CodeSigningConfigARN"):
		err = rm.updateFunctionConfiguration(ctx, desired, delta)
		if err != nil {
			return nil, err
		}
	}

	bytes, _ := json.Marshal(desired.ko.ObjectMeta.Annotations)
	fmt.Println("My annotation before is:", string(bytes))

	readOneLatest, err := rm.ReadOne(ctx, desired)
	if err != nil {
		return nil, err
	}

	latest.ko.ObjectMeta.Annotations = nil
	bytes1, _ := json.Marshal(readOneLatest.MetaObject().GetAnnotations())
	fmt.Println("My annotation after is:", string(bytes1))

	return rm.concreteResource(readOneLatest), nil
}

// updateFunctionConfiguration calls the UpdateFunctionConfiguration to edit a
// specific lambda function configuration.
func (rm *resourceManager) updateFunctionConfiguration(
	ctx context.Context,
	desired *resource,
	delta *ackcompare.Delta,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.updateFunctionConfiguration")
	defer exit(err)

	dspec := desired.ko.Spec
	input := &svcsdk.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(*dspec.Name),
	}

	if delta.DifferentAt("Spec.Description") {
		if dspec.Description != nil {
			input.Description = aws.String(*dspec.Description)
		} else {
			input.Description = aws.String("")
		}
	}

	if delta.DifferentAt("Spec.Handler") {
		if dspec.Handler != nil {
			input.Handler = aws.String(*dspec.Handler)
		} else {
			input.Handler = aws.String("")
		}
	}

	if delta.DifferentAt("Spec.KMSKeyARN") {
		if dspec.KMSKeyARN != nil {
			input.KMSKeyArn = aws.String(*dspec.KMSKeyARN)
		} else {
			input.KMSKeyArn = aws.String("")
		}
	}

	if delta.DifferentAt("Spec.Role") {
		if dspec.Role != nil {
			input.Role = aws.String(*dspec.Role)
		} else {
			input.Role = aws.String("")
		}
	}

	if delta.DifferentAt("Spec.MemorySize") {
		if dspec.MemorySize != nil {
			input.MemorySize = aws.Int64(*dspec.MemorySize)
		} else {
			input.MemorySize = aws.Int64(0)
		}
	}

	if delta.DifferentAt("Spec.Timeout") {
		if dspec.Timeout != nil {
			input.Timeout = aws.Int64(*dspec.Timeout)
		} else {
			input.Timeout = aws.Int64(0)
		}
	}

	if delta.DifferentAt("Spec.Environment") {
		environment := &svcsdk.Environment{}
		if dspec.Environment != nil {
			environment.Variables = dspec.Environment.DeepCopy().Variables
		}
		input.Environment = environment
	}

	if delta.DifferentAt("Spec.FileSystemConfigs") {
		fileSystemConfigs := []*svcsdk.FileSystemConfig{}
		if len(dspec.FileSystemConfigs) > 0 {
			for _, elem := range dspec.FileSystemConfigs {
				elemCopy := elem.DeepCopy()
				fscElem := &svcsdk.FileSystemConfig{
					Arn:            elemCopy.ARN,
					LocalMountPath: elemCopy.LocalMountPath,
				}
				fileSystemConfigs = append(fileSystemConfigs, fscElem)
			}
			input.FileSystemConfigs = fileSystemConfigs
		}
	}

	if delta.DifferentAt("Spec.ImageConfig") {
		if dspec.ImageConfig != nil && dspec.Code.ImageURI != nil && *dspec.Code.ImageURI != "" {
			imageConfig := &svcsdk.ImageConfig{}
			if dspec.ImageConfig != nil {
				imageConfigCopy := dspec.ImageConfig.DeepCopy()
				imageConfig.Command = imageConfigCopy.Command
				imageConfig.EntryPoint = imageConfigCopy.EntryPoint
				imageConfig.WorkingDirectory = imageConfigCopy.WorkingDirectory
			}
			input.ImageConfig = imageConfig
		}
	}

	if delta.DifferentAt("Spec.TracingConfig") {
		tracingConfig := &svcsdk.TracingConfig{}
		if dspec.TracingConfig != nil {
			tracingConfig.Mode = aws.String(*dspec.TracingConfig.Mode)
		}
		input.TracingConfig = tracingConfig
	}

	if delta.DifferentAt("Spec.VPCConfig") {
		VPCConfig := &svcsdk.VpcConfig{}
		if dspec.VPCConfig != nil {
			vpcConfigCopy := dspec.VPCConfig.DeepCopy()
			VPCConfig.SubnetIds = vpcConfigCopy.SubnetIDs
			VPCConfig.SecurityGroupIds = vpcConfigCopy.SecurityGroupIDs
		}
		input.VpcConfig = VPCConfig
	}

	if delta.DifferentAt("Spec.EphemeralStorage") {
		ephemeralStorage := &svcsdk.EphemeralStorage{}
		if dspec.EphemeralStorage != nil {
			ephemeralStorageCopy := dspec.EphemeralStorage.DeepCopy()
			ephemeralStorage.Size = ephemeralStorageCopy.Size
		}
		input.EphemeralStorage = ephemeralStorage
	}

	if delta.DifferentAt(("Spec.SnapStart")) {
		snapStart := &svcsdk.SnapStart{}
		if dspec.SnapStart != nil {
			snapStartCopy := dspec.SnapStart.DeepCopy()
			snapStart.ApplyOn = snapStartCopy.ApplyOn
		}
		input.SnapStart = snapStart
	}

	_, err = rm.sdkapi.UpdateFunctionConfigurationWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateFunctionConfiguration", err)
	if err != nil {
		return err
	}

	return nil
}

// updateFunctionsTags uses TagResource and UntagResource to add, remove and update
// a lambda function tags.
func (rm *resourceManager) updateFunctionTags(
	ctx context.Context,
	latest *resource,
	desired *resource,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.updateFunctionTags")
	defer exit(err)

	added, removed, updated := compareMaps(latest.ko.Spec.Tags, desired.ko.Spec.Tags)

	// There is no api call to update tags, so we need to remove them and add them later
	// with their new values.
	if len(removed)+len(updated) > 0 {
		removeTags := []*string{}
		for k := range updated {
			removeTags = append(removeTags, &k)
		}
		for _, k := range removed {
			removeTags = append(removeTags, &k)
		}
		input := &svcsdk.UntagResourceInput{
			Resource: (*string)(desired.ko.Status.ACKResourceMetadata.ARN),
			TagKeys:  removeTags,
		}
		_, err = rm.sdkapi.UntagResourceWithContext(ctx, input)
		rm.metrics.RecordAPICall("UPDATE", "UntagResource", err)
		if err != nil {
			return err
		}
	}

	if len(updated)+len(added) > 0 {
		addedTags := map[string]*string{}
		for k, v := range added {
			addedTags[k] = v
		}
		for k, v := range updated {
			addedTags[k] = v
		}

		input := &svcsdk.TagResourceInput{
			Resource: (*string)(desired.ko.Status.ACKResourceMetadata.ARN),
			Tags:     addedTags,
		}
		_, err = rm.sdkapi.TagResourceWithContext(ctx, input)
		rm.metrics.RecordAPICall("UPDATE", "TagResource", err)
		if err != nil {
			return err
		}
	}

	return nil
}

// updateFunctionsCode calls UpdateFunctionCode to update a specific lambda
// function code.
func (rm *resourceManager) updateFunctionCode(
	ctx context.Context,
	desired *resource,
	delta *ackcompare.Delta,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.updateFunctionCode")
	defer exit(err)

	dspec := desired.ko.Spec
	input := &svcsdk.UpdateFunctionCodeInput{
		FunctionName: aws.String(*dspec.Name),
	}

	if delta.DifferentAt("Spec.ImageURI") {
		input.ImageUri = dspec.Code.ImageURI
	} else if delta.DifferentAt("Spec.Code.S3Key") || delta.DifferentAt("Spec.Code.S3Bucket") {
		input.S3Key = aws.String(*dspec.Code.S3Key)
		input.S3Bucket = aws.String(*dspec.Code.S3Bucket)
	}

	_, err = rm.sdkapi.UpdateFunctionCodeWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateFunctionCode", err)
	if err != nil {
		return err
	}

	if dspec.Code != nil {
		if dspec.Code.S3Key != nil && *dspec.Code.S3Key != "" {
			if desired.ko.ObjectMeta.Annotations["lambda.services.k8s.aws/s3-key"] != *dspec.Code.S3Key {
				desired.ko.ObjectMeta.Annotations["lambda.services.k8s.aws/s3-key"] = *dspec.Code.S3Key
			}
		}
	}
	return nil
}

// compareMaps compares two string to string maps and returns three outputs: a
// map of the new key/values observed, a list of the keys of the removed values
// and a map containing the updated keys and their new values.
func compareMaps(
	a map[string]*string,
	b map[string]*string,
) (added map[string]*string, removed []string, updated map[string]*string) {
	added = map[string]*string{}
	updated = map[string]*string{}
	visited := make(map[string]bool, len(a))
	for keyA, valueA := range a {
		valueB, found := b[keyA]
		if !found {
			removed = append(removed, keyA)
			continue
		}
		if *valueA != *valueB {
			updated[keyA] = valueB
		}
		visited[keyA] = true
	}
	for keyB, valueB := range b {
		_, found := a[keyB]
		if !found {
			added[keyB] = valueB
		}
	}
	return
}

func customPreCompare(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) {
	s3KeyAnnot, exist := a.ko.GetAnnotations()["lambda.services.k8s.aws/s3-key"]
	fmt.Println("Checking annotations"+s3KeyAnnot, exist)
	bytes, _ := json.Marshal(a.ko.GetAnnotations())
	fmt.Println("My object is:", string(bytes))
	if !exist {
		if a.ko.Spec.Code != nil {
			if a.ko.Spec.Code.S3Key != nil && *a.ko.Spec.Code.S3Key != "" {
				delta.Add("", nil, nil)
			}
		}
	} else if exist {
		if s3KeyAnnot == "" {
			if a.ko.Spec.Code != nil && a.ko.Spec.Code.S3Key != nil && *a.ko.Spec.Code.S3Key != "" {
				delta.Add("Spec.Code.S3Key", a.ko.ObjectMeta.Annotations["lambda.services.k8s.aws/s3-key"], b.ko.Spec.Code.S3Key)
			}
		} else {
			fmt.Println("checking my delta is working"+s3KeyAnnot, a.ko.Spec.Code.S3Key)
			if a.ko.Spec.Code == nil {
				delta.Add("Spec.Code", a.ko.ObjectMeta.Annotations["lambda.services.k8s.aws/s3-key"], b.ko.Spec.Code)
			}
			if a.ko.Spec.Code.S3Key == nil || *a.ko.Spec.Code.S3Key != s3KeyAnnot {
				delta.Add("Spec.Code.S3Key", a.ko.ObjectMeta.Annotations["lambda.services.k8s.aws/s3-key"], b.ko.Spec.Code.S3Key)
			}
		}
	}

	if ackcompare.HasNilDifference(a.ko.Spec.Code, b.ko.Spec.Code) {
		delta.Add("Spec.Code", a.ko.Spec.Code, b.ko.Spec.Code)
	} else if a.ko.Spec.Code != nil && b.ko.Spec.Code != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Code.ImageURI, b.ko.Spec.Code.ImageURI) {
			delta.Add("Spec.Code.ImageURI", a.ko.Spec.Code.ImageURI, b.ko.Spec.Code.ImageURI)
		} else if a.ko.Spec.Code.ImageURI != nil && b.ko.Spec.Code.ImageURI != nil {
			if *a.ko.Spec.Code.ImageURI != *b.ko.Spec.Code.ImageURI {
				delta.Add("Spec.Code.ImageURI", a.ko.Spec.Code.ImageURI, b.ko.Spec.Code.ImageURI)
			}
		}
		//TODO(hialylmh) handle Spec.Code.S3bucket changes
		//  if ackcompare.HasNilDifference(a.ko.Spec.Code.S3Bucket, b.ko.Spec.Code.S3Bucket) {
		//  	delta.Add("Spec.Code.S3Bucket", a.ko.Spec.Code.S3Bucket, b.ko.Spec.Code.S3Bucket)
		//  } else if a.ko.Spec.Code.S3Bucket != nil && b.ko.Spec.Code.S3Bucket != nil {
		//  	if *a.ko.Spec.Code.S3Bucket != *b.ko.Spec.Code.S3Bucket {
		//  		delta.Add("Spec.Code.S3Bucket", a.ko.Spec.Code.S3Bucket, b.ko.Spec.Code.S3Bucket)
		//  	}
		//  }
		if ackcompare.HasNilDifference(a.ko.Spec.Code.S3Key, b.ko.Spec.Code.S3Key) {
			delta.Add("Spec.Code.S3Key", a.ko.Spec.Code.S3Key, b.ko.Spec.Code.S3Key)
		} else if a.ko.Spec.Code.S3Key != nil && b.ko.Spec.Code.S3Key != nil {
			if *a.ko.Spec.Code.S3Key != *b.ko.Spec.Code.S3Key {
				delta.Add("Spec.Code.S3Key", a.ko.Spec.Code.S3Key, b.ko.Spec.Code.S3Key)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Code.S3ObjectVersion, b.ko.Spec.Code.S3ObjectVersion) {
			delta.Add("Spec.Code.S3ObjectVersion", a.ko.Spec.Code.S3ObjectVersion, b.ko.Spec.Code.S3ObjectVersion)
		} else if a.ko.Spec.Code.S3ObjectVersion != nil && b.ko.Spec.Code.S3ObjectVersion != nil {
			if *a.ko.Spec.Code.S3ObjectVersion != *b.ko.Spec.Code.S3ObjectVersion {
				delta.Add("Spec.Code.S3ObjectVersion", a.ko.Spec.Code.S3ObjectVersion, b.ko.Spec.Code.S3ObjectVersion)
			}
		}
	}
}

// updateFunctionConcurrency calls UpdateFunctionConcurrency to update a specific
// lambda function reserved concurrent executions.
func (rm *resourceManager) updateFunctionConcurrency(
	ctx context.Context,
	desired *resource,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.updateFunctionConcurrency")
	defer exit(err)

	dspec := desired.ko.Spec
	input := &svcsdk.PutFunctionConcurrencyInput{
		FunctionName: aws.String(*dspec.Name),
	}

	if desired.ko.Spec.ReservedConcurrentExecutions != nil {
		input.ReservedConcurrentExecutions = aws.Int64(*desired.ko.Spec.ReservedConcurrentExecutions)
	} else {
		input.ReservedConcurrentExecutions = aws.Int64(0)
	}

	_, err = rm.sdkapi.PutFunctionConcurrencyWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "PutFunctionConcurrency", err)
	if err != nil {
		return err
	}
	return nil
}

// updateFunctionCodeSigningConfig calls PutFunctionCodeSigningConfig to update
// it code signing configuration
func (rm *resourceManager) updateFunctionCodeSigningConfig(
	ctx context.Context,
	desired *resource,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.updateFunctionCodeSigningConfig")
	defer exit(err)

	if desired.ko.Spec.CodeSigningConfigARN == nil || *desired.ko.Spec.CodeSigningConfigARN == "" {
		return rm.deleteFunctionCodeSigningConfig(ctx, desired)
	}

	dspec := desired.ko.Spec
	input := &svcsdk.PutFunctionCodeSigningConfigInput{
		FunctionName:         aws.String(*dspec.Name),
		CodeSigningConfigArn: aws.String(*dspec.CodeSigningConfigARN),
	}

	_, err = rm.sdkapi.PutFunctionCodeSigningConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "PutFunctionCodeSigningConfig", err)
	if err != nil {
		return err
	}
	return nil
}

// deleteFunctionCodeSigningConfig calls deleteFunctionCodeSigningConfig to update
// it code signing configuration
func (rm *resourceManager) deleteFunctionCodeSigningConfig(
	ctx context.Context,
	desired *resource,
) error {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.deleteFunctionCodeSigningConfig")
	defer exit(err)

	dspec := desired.ko.Spec
	input := &svcsdk.DeleteFunctionCodeSigningConfigInput{
		FunctionName: aws.String(*dspec.Name),
	}

	_, err = rm.sdkapi.DeleteFunctionCodeSigningConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "DeleteFunctionCodeSigningConfig", err)
	if err != nil {
		return err
	}
	return nil
}

// setResourceAdditionalFields will describe the fields that are not return by
// GetFunctionConcurrency calls
func (rm *resourceManager) setResourceAdditionalFields(
	ctx context.Context,
	ko *svcapitypes.Function,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.setResourceAdditionalFields")
	defer exit(err)

	var getFunctionConcurrencyOutput *svcsdk.GetFunctionConcurrencyOutput
	getFunctionConcurrencyOutput, err = rm.sdkapi.GetFunctionConcurrencyWithContext(
		ctx,
		&svcsdk.GetFunctionConcurrencyInput{
			FunctionName: ko.Spec.Name,
		},
	)
	rm.metrics.RecordAPICall("GET", "GetFunctionConcurrency", err)
	if err != nil {
		return err
	}
	ko.Spec.ReservedConcurrentExecutions = getFunctionConcurrencyOutput.ReservedConcurrentExecutions

	if ko.Spec.PackageType != nil && *ko.Spec.PackageType == "Zip" {
		var getFunctionCodeSigningConfigOutput *svcsdk.GetFunctionCodeSigningConfigOutput
		getFunctionCodeSigningConfigOutput, err = rm.sdkapi.GetFunctionCodeSigningConfigWithContext(
			ctx,
			&svcsdk.GetFunctionCodeSigningConfigInput{
				FunctionName: ko.Spec.Name,
			},
		)
		rm.metrics.RecordAPICall("GET", "GetFunctionCodeSigningConfig", err)
		if err != nil {
			return err
		}
		ko.Spec.CodeSigningConfigARN = getFunctionCodeSigningConfigOutput.CodeSigningConfigArn
	}
	if ko.Spec.PackageType != nil && *ko.Spec.PackageType == "Image" &&
		ko.Spec.CodeSigningConfigARN != nil && *ko.Spec.CodeSigningConfigARN != "" {
		return ackerr.NewTerminalError(ErrCannotSetFunctionCSC)
	}
	return nil
}
