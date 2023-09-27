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
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.CodeSHA256, b.ko.Spec.CodeSHA256) {
		delta.Add("Spec.CodeSHA256", a.ko.Spec.CodeSHA256, b.ko.Spec.CodeSHA256)
	} else if a.ko.Spec.CodeSHA256 != nil && b.ko.Spec.CodeSHA256 != nil {
		if *a.ko.Spec.CodeSHA256 != *b.ko.Spec.CodeSHA256 {
			delta.Add("Spec.CodeSHA256", a.ko.Spec.CodeSHA256, b.ko.Spec.CodeSHA256)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FunctionName, b.ko.Spec.FunctionName) {
		delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
	} else if a.ko.Spec.FunctionName != nil && b.ko.Spec.FunctionName != nil {
		if *a.ko.Spec.FunctionName != *b.ko.Spec.FunctionName {
			delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RevisionID, b.ko.Spec.RevisionID) {
		delta.Add("Spec.RevisionID", a.ko.Spec.RevisionID, b.ko.Spec.RevisionID)
	} else if a.ko.Spec.RevisionID != nil && b.ko.Spec.RevisionID != nil {
		if *a.ko.Spec.RevisionID != *b.ko.Spec.RevisionID {
			delta.Add("Spec.RevisionID", a.ko.Spec.RevisionID, b.ko.Spec.RevisionID)
		}
	}

	return delta
}