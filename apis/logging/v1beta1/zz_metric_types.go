//go:build (logging || all) && !ignore_autogenerated

// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketOptionsInitParameters struct {

	// Specifies a set of buckets with arbitrary widths.
	// Structure is documented below.
	ExplicitBuckets []ExplicitBucketsInitParameters `json:"explicitBuckets,omitempty" tf:"explicit_buckets,omitempty"`

	// Specifies an exponential sequence of buckets that have a width that is proportional to the value of
	// the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
	// Structure is documented below.
	ExponentialBuckets []ExponentialBucketsInitParameters `json:"exponentialBuckets,omitempty" tf:"exponential_buckets,omitempty"`

	// Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
	// Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
	// Structure is documented below.
	LinearBuckets []LinearBucketsInitParameters `json:"linearBuckets,omitempty" tf:"linear_buckets,omitempty"`
}

type BucketOptionsObservation struct {

	// Specifies a set of buckets with arbitrary widths.
	// Structure is documented below.
	ExplicitBuckets []ExplicitBucketsObservation `json:"explicitBuckets,omitempty" tf:"explicit_buckets,omitempty"`

	// Specifies an exponential sequence of buckets that have a width that is proportional to the value of
	// the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
	// Structure is documented below.
	ExponentialBuckets []ExponentialBucketsObservation `json:"exponentialBuckets,omitempty" tf:"exponential_buckets,omitempty"`

	// Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
	// Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
	// Structure is documented below.
	LinearBuckets []LinearBucketsObservation `json:"linearBuckets,omitempty" tf:"linear_buckets,omitempty"`
}

type BucketOptionsParameters struct {

	// Specifies a set of buckets with arbitrary widths.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ExplicitBuckets []ExplicitBucketsParameters `json:"explicitBuckets,omitempty" tf:"explicit_buckets,omitempty"`

	// Specifies an exponential sequence of buckets that have a width that is proportional to the value of
	// the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ExponentialBuckets []ExponentialBucketsParameters `json:"exponentialBuckets,omitempty" tf:"exponential_buckets,omitempty"`

	// Specifies a linear sequence of buckets that all have the same width (except overflow and underflow).
	// Each bucket represents a constant absolute uncertainty on the specific value in the bucket.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LinearBuckets []LinearBucketsParameters `json:"linearBuckets,omitempty" tf:"linear_buckets,omitempty"`
}

type ExplicitBucketsInitParameters struct {

	// The values must be monotonically increasing.
	Bounds []*float64 `json:"bounds,omitempty" tf:"bounds,omitempty"`
}

type ExplicitBucketsObservation struct {

	// The values must be monotonically increasing.
	Bounds []*float64 `json:"bounds,omitempty" tf:"bounds,omitempty"`
}

type ExplicitBucketsParameters struct {

	// The values must be monotonically increasing.
	// +kubebuilder:validation:Optional
	Bounds []*float64 `json:"bounds" tf:"bounds,omitempty"`
}

type ExponentialBucketsInitParameters struct {

	// Must be greater than 1.
	GrowthFactor *float64 `json:"growthFactor,omitempty" tf:"growth_factor,omitempty"`

	// Must be greater than 0.
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Must be greater than 0.
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`
}

type ExponentialBucketsObservation struct {

	// Must be greater than 1.
	GrowthFactor *float64 `json:"growthFactor,omitempty" tf:"growth_factor,omitempty"`

	// Must be greater than 0.
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Must be greater than 0.
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`
}

type ExponentialBucketsParameters struct {

	// Must be greater than 1.
	// +kubebuilder:validation:Optional
	GrowthFactor *float64 `json:"growthFactor,omitempty" tf:"growth_factor,omitempty"`

	// Must be greater than 0.
	// +kubebuilder:validation:Optional
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Must be greater than 0.
	// +kubebuilder:validation:Optional
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`
}

type LabelsInitParameters struct {

	// A human-readable description for the label.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The label key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LabelsObservation struct {

	// A human-readable description for the label.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The label key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LabelsParameters struct {

	// A human-readable description for the label.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The label key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	// +kubebuilder:validation:Optional
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LinearBucketsInitParameters struct {

	// Must be greater than 0.
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Lower bound of the first bucket.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Must be greater than 0.
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

type LinearBucketsObservation struct {

	// Must be greater than 0.
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Lower bound of the first bucket.
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Must be greater than 0.
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

type LinearBucketsParameters struct {

	// Must be greater than 0.
	// +kubebuilder:validation:Optional
	NumFiniteBuckets *float64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`

	// Lower bound of the first bucket.
	// +kubebuilder:validation:Optional
	Offset *float64 `json:"offset,omitempty" tf:"offset,omitempty"`

	// Must be greater than 0.
	// +kubebuilder:validation:Optional
	Width *float64 `json:"width,omitempty" tf:"width,omitempty"`
}

type MetricDescriptorInitParameters struct {

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case
	// without an ending period, for example "Request count". This field is optional but it is
	// recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. For
	// example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
	// for the HTTP response code, response_code, so you can look at latencies for successful responses
	// or just for responses that failed.
	// Structure is documented below.
	Labels []LabelsInitParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to DELTA.
	// Possible values are: DELTA, GAUGE, CUMULATIVE.
	MetricKind *string `json:"metricKind,omitempty" tf:"metric_kind,omitempty"`

	// The unit in which the metric value is reported. It is only applicable if the valueType is
	// INT64, DOUBLE, or DISTRIBUTION. The supported units are a subset of
	// The Unified Code for Units of Measure standard
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricDescriptorObservation struct {

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case
	// without an ending period, for example "Request count". This field is optional but it is
	// recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. For
	// example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
	// for the HTTP response code, response_code, so you can look at latencies for successful responses
	// or just for responses that failed.
	// Structure is documented below.
	Labels []LabelsObservation `json:"labels,omitempty" tf:"labels,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to DELTA.
	// Possible values are: DELTA, GAUGE, CUMULATIVE.
	MetricKind *string `json:"metricKind,omitempty" tf:"metric_kind,omitempty"`

	// The unit in which the metric value is reported. It is only applicable if the valueType is
	// INT64, DOUBLE, or DISTRIBUTION. The supported units are a subset of
	// The Unified Code for Units of Measure standard
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricDescriptorParameters struct {

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case
	// without an ending period, for example "Request count". This field is optional but it is
	// recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. For
	// example, the appengine.googleapis.com/http/server/response_latencies metric type has a label
	// for the HTTP response code, response_code, so you can look at latencies for successful responses
	// or just for responses that failed.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to DELTA.
	// Possible values are: DELTA, GAUGE, CUMULATIVE.
	// +kubebuilder:validation:Optional
	MetricKind *string `json:"metricKind" tf:"metric_kind,omitempty"`

	// The unit in which the metric value is reported. It is only applicable if the valueType is
	// INT64, DOUBLE, or DISTRIBUTION. The supported units are a subset of
	// The Unified Code for Units of Measure standard
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of metricKind and valueType might not be supported.
	// For counter metrics, set this to INT64.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION, MONEY.
	// +kubebuilder:validation:Optional
	ValueType *string `json:"valueType" tf:"value_type,omitempty"`
}

type MetricInitParameters struct {

	// The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
	// are supported. The bucket has to be in the same project as the metric.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/logging/v1beta1.ProjectBucketConfig
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a ProjectBucketConfig in logging to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a ProjectBucketConfig in logging to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions []BucketOptionsInitParameters `json:"bucketOptions,omitempty" tf:"bucket_options,omitempty"`

	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this metric is disabled and it does not generate any points.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	// +mapType=granular
	LabelExtractors map[string]*string `json:"labelExtractors,omitempty" tf:"label_extractors,omitempty"`

	// The optional metric descriptor associated with the logs-based metric.
	// If unspecified, it uses a default metric descriptor with a DELTA metric kind,
	// INT64 value type, with no labels and a unit of "1". Such a metric counts the
	// number of log entries matching the filter expression.
	// Structure is documented below.
	MetricDescriptor []MetricDescriptorInitParameters `json:"metricDescriptor,omitempty" tf:"metric_descriptor,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor *string `json:"valueExtractor,omitempty" tf:"value_extractor,omitempty"`
}

type MetricObservation struct {

	// The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
	// are supported. The bucket has to be in the same project as the metric.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions []BucketOptionsObservation `json:"bucketOptions,omitempty" tf:"bucket_options,omitempty"`

	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this metric is disabled and it does not generate any points.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	// +mapType=granular
	LabelExtractors map[string]*string `json:"labelExtractors,omitempty" tf:"label_extractors,omitempty"`

	// The optional metric descriptor associated with the logs-based metric.
	// If unspecified, it uses a default metric descriptor with a DELTA metric kind,
	// INT64 value type, with no labels and a unit of "1". Such a metric counts the
	// number of log entries matching the filter expression.
	// Structure is documented below.
	MetricDescriptor []MetricDescriptorObservation `json:"metricDescriptor,omitempty" tf:"metric_descriptor,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor *string `json:"valueExtractor,omitempty" tf:"value_extractor,omitempty"`
}

type MetricParameters struct {

	// The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
	// are supported. The bucket has to be in the same project as the metric.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/logging/v1beta1.ProjectBucketConfig
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a ProjectBucketConfig in logging to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a ProjectBucketConfig in logging to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BucketOptions []BucketOptionsParameters `json:"bucketOptions,omitempty" tf:"bucket_options,omitempty"`

	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this metric is disabled and it does not generate any points.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	LabelExtractors map[string]*string `json:"labelExtractors,omitempty" tf:"label_extractors,omitempty"`

	// The optional metric descriptor associated with the logs-based metric.
	// If unspecified, it uses a default metric descriptor with a DELTA metric kind,
	// INT64 value type, with no labels and a unit of "1". Such a metric counts the
	// number of log entries matching the filter expression.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MetricDescriptor []MetricDescriptorParameters `json:"metricDescriptor,omitempty" tf:"metric_descriptor,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	// +kubebuilder:validation:Optional
	ValueExtractor *string `json:"valueExtractor,omitempty" tf:"value_extractor,omitempty"`
}

// MetricSpec defines the desired state of Metric
type MetricSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MetricInitParameters `json:"initProvider,omitempty"`
}

// MetricStatus defines the observed state of Metric.
type MetricStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Metric is the Schema for the Metrics API. Logs-based metric can also be used to extract values from logs and create a a distribution of the values.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Metric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || (has(self.initProvider) && has(self.initProvider.filter))",message="spec.forProvider.filter is a required parameter"
	Spec   MetricSpec   `json:"spec"`
	Status MetricStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricList contains a list of Metrics
type MetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Metric `json:"items"`
}

// Repository type metadata.
var (
	Metric_Kind             = "Metric"
	Metric_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Metric_Kind}.String()
	Metric_KindAPIVersion   = Metric_Kind + "." + CRDGroupVersion.String()
	Metric_GroupVersionKind = CRDGroupVersion.WithKind(Metric_Kind)
)

func init() {
	SchemeBuilder.Register(&Metric{}, &MetricList{})
}
