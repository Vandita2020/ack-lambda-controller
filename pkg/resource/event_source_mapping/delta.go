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

package event_source_mapping

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
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
	customPreCompare(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AmazonManagedKafkaEventSourceConfig, b.ko.Spec.AmazonManagedKafkaEventSourceConfig) {
		delta.Add("Spec.AmazonManagedKafkaEventSourceConfig", a.ko.Spec.AmazonManagedKafkaEventSourceConfig, b.ko.Spec.AmazonManagedKafkaEventSourceConfig)
	} else if a.ko.Spec.AmazonManagedKafkaEventSourceConfig != nil && b.ko.Spec.AmazonManagedKafkaEventSourceConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID) {
			delta.Add("Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID", a.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID)
		} else if a.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID != nil && b.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID != nil {
			if *a.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID != *b.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID {
				delta.Add("Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID", a.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.AmazonManagedKafkaEventSourceConfig.ConsumerGroupID)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BatchSize, b.ko.Spec.BatchSize) {
		delta.Add("Spec.BatchSize", a.ko.Spec.BatchSize, b.ko.Spec.BatchSize)
	} else if a.ko.Spec.BatchSize != nil && b.ko.Spec.BatchSize != nil {
		if *a.ko.Spec.BatchSize != *b.ko.Spec.BatchSize {
			delta.Add("Spec.BatchSize", a.ko.Spec.BatchSize, b.ko.Spec.BatchSize)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BisectBatchOnFunctionError, b.ko.Spec.BisectBatchOnFunctionError) {
		delta.Add("Spec.BisectBatchOnFunctionError", a.ko.Spec.BisectBatchOnFunctionError, b.ko.Spec.BisectBatchOnFunctionError)
	} else if a.ko.Spec.BisectBatchOnFunctionError != nil && b.ko.Spec.BisectBatchOnFunctionError != nil {
		if *a.ko.Spec.BisectBatchOnFunctionError != *b.ko.Spec.BisectBatchOnFunctionError {
			delta.Add("Spec.BisectBatchOnFunctionError", a.ko.Spec.BisectBatchOnFunctionError, b.ko.Spec.BisectBatchOnFunctionError)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DestinationConfig, b.ko.Spec.DestinationConfig) {
		delta.Add("Spec.DestinationConfig", a.ko.Spec.DestinationConfig, b.ko.Spec.DestinationConfig)
	} else if a.ko.Spec.DestinationConfig != nil && b.ko.Spec.DestinationConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.DestinationConfig.OnFailure, b.ko.Spec.DestinationConfig.OnFailure) {
			delta.Add("Spec.DestinationConfig.OnFailure", a.ko.Spec.DestinationConfig.OnFailure, b.ko.Spec.DestinationConfig.OnFailure)
		} else if a.ko.Spec.DestinationConfig.OnFailure != nil && b.ko.Spec.DestinationConfig.OnFailure != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.DestinationConfig.OnFailure.Destination, b.ko.Spec.DestinationConfig.OnFailure.Destination) {
				delta.Add("Spec.DestinationConfig.OnFailure.Destination", a.ko.Spec.DestinationConfig.OnFailure.Destination, b.ko.Spec.DestinationConfig.OnFailure.Destination)
			} else if a.ko.Spec.DestinationConfig.OnFailure.Destination != nil && b.ko.Spec.DestinationConfig.OnFailure.Destination != nil {
				if *a.ko.Spec.DestinationConfig.OnFailure.Destination != *b.ko.Spec.DestinationConfig.OnFailure.Destination {
					delta.Add("Spec.DestinationConfig.OnFailure.Destination", a.ko.Spec.DestinationConfig.OnFailure.Destination, b.ko.Spec.DestinationConfig.OnFailure.Destination)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DestinationConfig.OnSuccess, b.ko.Spec.DestinationConfig.OnSuccess) {
			delta.Add("Spec.DestinationConfig.OnSuccess", a.ko.Spec.DestinationConfig.OnSuccess, b.ko.Spec.DestinationConfig.OnSuccess)
		} else if a.ko.Spec.DestinationConfig.OnSuccess != nil && b.ko.Spec.DestinationConfig.OnSuccess != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.DestinationConfig.OnSuccess.Destination, b.ko.Spec.DestinationConfig.OnSuccess.Destination) {
				delta.Add("Spec.DestinationConfig.OnSuccess.Destination", a.ko.Spec.DestinationConfig.OnSuccess.Destination, b.ko.Spec.DestinationConfig.OnSuccess.Destination)
			} else if a.ko.Spec.DestinationConfig.OnSuccess.Destination != nil && b.ko.Spec.DestinationConfig.OnSuccess.Destination != nil {
				if *a.ko.Spec.DestinationConfig.OnSuccess.Destination != *b.ko.Spec.DestinationConfig.OnSuccess.Destination {
					delta.Add("Spec.DestinationConfig.OnSuccess.Destination", a.ko.Spec.DestinationConfig.OnSuccess.Destination, b.ko.Spec.DestinationConfig.OnSuccess.Destination)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Enabled, b.ko.Spec.Enabled) {
		delta.Add("Spec.Enabled", a.ko.Spec.Enabled, b.ko.Spec.Enabled)
	} else if a.ko.Spec.Enabled != nil && b.ko.Spec.Enabled != nil {
		if *a.ko.Spec.Enabled != *b.ko.Spec.Enabled {
			delta.Add("Spec.Enabled", a.ko.Spec.Enabled, b.ko.Spec.Enabled)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EventSourceARN, b.ko.Spec.EventSourceARN) {
		delta.Add("Spec.EventSourceARN", a.ko.Spec.EventSourceARN, b.ko.Spec.EventSourceARN)
	} else if a.ko.Spec.EventSourceARN != nil && b.ko.Spec.EventSourceARN != nil {
		if *a.ko.Spec.EventSourceARN != *b.ko.Spec.EventSourceARN {
			delta.Add("Spec.EventSourceARN", a.ko.Spec.EventSourceARN, b.ko.Spec.EventSourceARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FunctionName, b.ko.Spec.FunctionName) {
		delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
	} else if a.ko.Spec.FunctionName != nil && b.ko.Spec.FunctionName != nil {
		if *a.ko.Spec.FunctionName != *b.ko.Spec.FunctionName {
			delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.FunctionRef, b.ko.Spec.FunctionRef) {
		delta.Add("Spec.FunctionRef", a.ko.Spec.FunctionRef, b.ko.Spec.FunctionRef)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.FunctionResponseTypes, b.ko.Spec.FunctionResponseTypes) {
		delta.Add("Spec.FunctionResponseTypes", a.ko.Spec.FunctionResponseTypes, b.ko.Spec.FunctionResponseTypes)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaximumBatchingWindowInSeconds, b.ko.Spec.MaximumBatchingWindowInSeconds) {
		delta.Add("Spec.MaximumBatchingWindowInSeconds", a.ko.Spec.MaximumBatchingWindowInSeconds, b.ko.Spec.MaximumBatchingWindowInSeconds)
	} else if a.ko.Spec.MaximumBatchingWindowInSeconds != nil && b.ko.Spec.MaximumBatchingWindowInSeconds != nil {
		if *a.ko.Spec.MaximumBatchingWindowInSeconds != *b.ko.Spec.MaximumBatchingWindowInSeconds {
			delta.Add("Spec.MaximumBatchingWindowInSeconds", a.ko.Spec.MaximumBatchingWindowInSeconds, b.ko.Spec.MaximumBatchingWindowInSeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaximumRecordAgeInSeconds, b.ko.Spec.MaximumRecordAgeInSeconds) {
		delta.Add("Spec.MaximumRecordAgeInSeconds", a.ko.Spec.MaximumRecordAgeInSeconds, b.ko.Spec.MaximumRecordAgeInSeconds)
	} else if a.ko.Spec.MaximumRecordAgeInSeconds != nil && b.ko.Spec.MaximumRecordAgeInSeconds != nil {
		if *a.ko.Spec.MaximumRecordAgeInSeconds != *b.ko.Spec.MaximumRecordAgeInSeconds {
			delta.Add("Spec.MaximumRecordAgeInSeconds", a.ko.Spec.MaximumRecordAgeInSeconds, b.ko.Spec.MaximumRecordAgeInSeconds)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MaximumRetryAttempts, b.ko.Spec.MaximumRetryAttempts) {
		delta.Add("Spec.MaximumRetryAttempts", a.ko.Spec.MaximumRetryAttempts, b.ko.Spec.MaximumRetryAttempts)
	} else if a.ko.Spec.MaximumRetryAttempts != nil && b.ko.Spec.MaximumRetryAttempts != nil {
		if *a.ko.Spec.MaximumRetryAttempts != *b.ko.Spec.MaximumRetryAttempts {
			delta.Add("Spec.MaximumRetryAttempts", a.ko.Spec.MaximumRetryAttempts, b.ko.Spec.MaximumRetryAttempts)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ParallelizationFactor, b.ko.Spec.ParallelizationFactor) {
		delta.Add("Spec.ParallelizationFactor", a.ko.Spec.ParallelizationFactor, b.ko.Spec.ParallelizationFactor)
	} else if a.ko.Spec.ParallelizationFactor != nil && b.ko.Spec.ParallelizationFactor != nil {
		if *a.ko.Spec.ParallelizationFactor != *b.ko.Spec.ParallelizationFactor {
			delta.Add("Spec.ParallelizationFactor", a.ko.Spec.ParallelizationFactor, b.ko.Spec.ParallelizationFactor)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.QueueRefs, b.ko.Spec.QueueRefs) {
		delta.Add("Spec.QueueRefs", a.ko.Spec.QueueRefs, b.ko.Spec.QueueRefs)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.Queues, b.ko.Spec.Queues) {
		delta.Add("Spec.Queues", a.ko.Spec.Queues, b.ko.Spec.Queues)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfig, b.ko.Spec.ScalingConfig) {
		delta.Add("Spec.ScalingConfig", a.ko.Spec.ScalingConfig, b.ko.Spec.ScalingConfig)
	} else if a.ko.Spec.ScalingConfig != nil && b.ko.Spec.ScalingConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ScalingConfig.MaximumConcurrency, b.ko.Spec.ScalingConfig.MaximumConcurrency) {
			delta.Add("Spec.ScalingConfig.MaximumConcurrency", a.ko.Spec.ScalingConfig.MaximumConcurrency, b.ko.Spec.ScalingConfig.MaximumConcurrency)
		} else if a.ko.Spec.ScalingConfig.MaximumConcurrency != nil && b.ko.Spec.ScalingConfig.MaximumConcurrency != nil {
			if *a.ko.Spec.ScalingConfig.MaximumConcurrency != *b.ko.Spec.ScalingConfig.MaximumConcurrency {
				delta.Add("Spec.ScalingConfig.MaximumConcurrency", a.ko.Spec.ScalingConfig.MaximumConcurrency, b.ko.Spec.ScalingConfig.MaximumConcurrency)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SelfManagedEventSource, b.ko.Spec.SelfManagedEventSource) {
		delta.Add("Spec.SelfManagedEventSource", a.ko.Spec.SelfManagedEventSource, b.ko.Spec.SelfManagedEventSource)
	} else if a.ko.Spec.SelfManagedEventSource != nil && b.ko.Spec.SelfManagedEventSource != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.SelfManagedEventSource.Endpoints, b.ko.Spec.SelfManagedEventSource.Endpoints) {
			delta.Add("Spec.SelfManagedEventSource.Endpoints", a.ko.Spec.SelfManagedEventSource.Endpoints, b.ko.Spec.SelfManagedEventSource.Endpoints)
		} else if a.ko.Spec.SelfManagedEventSource.Endpoints != nil && b.ko.Spec.SelfManagedEventSource.Endpoints != nil {
			if !reflect.DeepEqual(a.ko.Spec.SelfManagedEventSource.Endpoints, b.ko.Spec.SelfManagedEventSource.Endpoints) {
				delta.Add("Spec.SelfManagedEventSource.Endpoints", a.ko.Spec.SelfManagedEventSource.Endpoints, b.ko.Spec.SelfManagedEventSource.Endpoints)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SelfManagedKafkaEventSourceConfig, b.ko.Spec.SelfManagedKafkaEventSourceConfig) {
		delta.Add("Spec.SelfManagedKafkaEventSourceConfig", a.ko.Spec.SelfManagedKafkaEventSourceConfig, b.ko.Spec.SelfManagedKafkaEventSourceConfig)
	} else if a.ko.Spec.SelfManagedKafkaEventSourceConfig != nil && b.ko.Spec.SelfManagedKafkaEventSourceConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID) {
			delta.Add("Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID", a.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID)
		} else if a.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID != nil && b.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID != nil {
			if *a.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID != *b.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID {
				delta.Add("Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID", a.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID, b.ko.Spec.SelfManagedKafkaEventSourceConfig.ConsumerGroupID)
			}
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.SourceAccessConfigurations, b.ko.Spec.SourceAccessConfigurations) {
		delta.Add("Spec.SourceAccessConfigurations", a.ko.Spec.SourceAccessConfigurations, b.ko.Spec.SourceAccessConfigurations)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StartingPosition, b.ko.Spec.StartingPosition) {
		delta.Add("Spec.StartingPosition", a.ko.Spec.StartingPosition, b.ko.Spec.StartingPosition)
	} else if a.ko.Spec.StartingPosition != nil && b.ko.Spec.StartingPosition != nil {
		if *a.ko.Spec.StartingPosition != *b.ko.Spec.StartingPosition {
			delta.Add("Spec.StartingPosition", a.ko.Spec.StartingPosition, b.ko.Spec.StartingPosition)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StartingPositionTimestamp, b.ko.Spec.StartingPositionTimestamp) {
		delta.Add("Spec.StartingPositionTimestamp", a.ko.Spec.StartingPositionTimestamp, b.ko.Spec.StartingPositionTimestamp)
	} else if a.ko.Spec.StartingPositionTimestamp != nil && b.ko.Spec.StartingPositionTimestamp != nil {
		if !a.ko.Spec.StartingPositionTimestamp.Equal(b.ko.Spec.StartingPositionTimestamp) {
			delta.Add("Spec.StartingPositionTimestamp", a.ko.Spec.StartingPositionTimestamp, b.ko.Spec.StartingPositionTimestamp)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.Topics, b.ko.Spec.Topics) {
		delta.Add("Spec.Topics", a.ko.Spec.Topics, b.ko.Spec.Topics)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TumblingWindowInSeconds, b.ko.Spec.TumblingWindowInSeconds) {
		delta.Add("Spec.TumblingWindowInSeconds", a.ko.Spec.TumblingWindowInSeconds, b.ko.Spec.TumblingWindowInSeconds)
	} else if a.ko.Spec.TumblingWindowInSeconds != nil && b.ko.Spec.TumblingWindowInSeconds != nil {
		if *a.ko.Spec.TumblingWindowInSeconds != *b.ko.Spec.TumblingWindowInSeconds {
			delta.Add("Spec.TumblingWindowInSeconds", a.ko.Spec.TumblingWindowInSeconds, b.ko.Spec.TumblingWindowInSeconds)
		}
	}

	return delta
}
