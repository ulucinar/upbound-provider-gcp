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

type LabelsInitParameters struct {

	// A human-readable description for the label.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// Default value is STRING.
	// Possible values are: STRING, BOOL, INT64.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LabelsObservation struct {

	// A human-readable description for the label.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// Default value is STRING.
	// Possible values are: STRING, BOOL, INT64.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LabelsParameters struct {

	// A human-readable description for the label.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The type of data that can be assigned to the label.
	// Default value is STRING.
	// Possible values are: STRING, BOOL, INT64.
	// +kubebuilder:validation:Optional
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetadataInitParameters struct {

	// The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	IngestDelay *string `json:"ingestDelay,omitempty" tf:"ingest_delay,omitempty"`

	// The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	SamplePeriod *string `json:"samplePeriod,omitempty" tf:"sample_period,omitempty"`
}

type MetadataObservation struct {

	// The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	IngestDelay *string `json:"ingestDelay,omitempty" tf:"ingest_delay,omitempty"`

	// The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	SamplePeriod *string `json:"samplePeriod,omitempty" tf:"sample_period,omitempty"`
}

type MetadataParameters struct {

	// The delay of data points caused by ingestion. Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	// +kubebuilder:validation:Optional
	IngestDelay *string `json:"ingestDelay,omitempty" tf:"ingest_delay,omitempty"`

	// The sampling period of metric data points. For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In [duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration).
	// +kubebuilder:validation:Optional
	SamplePeriod *string `json:"samplePeriod,omitempty" tf:"sample_period,omitempty"`
}

type MetricDescriptorInitParameters struct {

	// A detailed description of the metric, which can be used in documentation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels []LabelsInitParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// The launch stage of the metric definition.
	// Possible values are: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED.
	LaunchStage *string `json:"launchStage,omitempty" tf:"launch_stage,omitempty"`

	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: METRIC_KIND_UNSPECIFIED, GAUGE, DELTA, CUMULATIVE.
	MetricKind *string `json:"metricKind,omitempty" tf:"metric_kind,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricDescriptorObservation struct {

	// A detailed description of the metric, which can be used in documentation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	Labels []LabelsObservation `json:"labels,omitempty" tf:"labels,omitempty"`

	// The launch stage of the metric definition.
	// Possible values are: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED.
	LaunchStage *string `json:"launchStage,omitempty" tf:"launch_stage,omitempty"`

	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: METRIC_KIND_UNSPECIFIED, GAUGE, DELTA, CUMULATIVE.
	MetricKind *string `json:"metricKind,omitempty" tf:"metric_kind,omitempty"`

	// If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here. This field allows time series to be associated with the intersection of this metric type and the monitored resource types in this list.
	MonitoredResourceTypes []*string `json:"monitoredResourceTypes,omitempty" tf:"monitored_resource_types,omitempty"`

	// The resource name of the metric descriptor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION.
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type MetricDescriptorParameters struct {

	// A detailed description of the metric, which can be used in documentation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count".
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The set of labels that can be used to describe a specific instance of this metric type. In order to delete a label, the entire resource must be deleted, then created with the desired labels.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// The launch stage of the metric definition.
	// Possible values are: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED.
	// +kubebuilder:validation:Optional
	LaunchStage *string `json:"launchStage,omitempty" tf:"launch_stage,omitempty"`

	// Metadata which can be used to guide usage of the metric.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: METRIC_KIND_UNSPECIFIED, GAUGE, DELTA, CUMULATIVE.
	// +kubebuilder:validation:Optional
	MetricKind *string `json:"metricKind,omitempty" tf:"metric_kind,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The metric type, including its DNS name prefix. The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The units in which the metric value is reported. It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`

	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported.
	// Possible values are: BOOL, INT64, DOUBLE, STRING, DISTRIBUTION.
	// +kubebuilder:validation:Optional
	ValueType *string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

// MetricDescriptorSpec defines the desired state of MetricDescriptor
type MetricDescriptorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricDescriptorParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MetricDescriptorInitParameters `json:"initProvider,omitempty"`
}

// MetricDescriptorStatus defines the observed state of MetricDescriptor.
type MetricDescriptorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricDescriptorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MetricDescriptor is the Schema for the MetricDescriptors API. Defines a metric type and its schema.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type MetricDescriptor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricKind) || (has(self.initProvider) && has(self.initProvider.metricKind))",message="spec.forProvider.metricKind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.valueType) || (has(self.initProvider) && has(self.initProvider.valueType))",message="spec.forProvider.valueType is a required parameter"
	Spec   MetricDescriptorSpec   `json:"spec"`
	Status MetricDescriptorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricDescriptorList contains a list of MetricDescriptors
type MetricDescriptorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricDescriptor `json:"items"`
}

// Repository type metadata.
var (
	MetricDescriptor_Kind             = "MetricDescriptor"
	MetricDescriptor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricDescriptor_Kind}.String()
	MetricDescriptor_KindAPIVersion   = MetricDescriptor_Kind + "." + CRDGroupVersion.String()
	MetricDescriptor_GroupVersionKind = CRDGroupVersion.WithKind(MetricDescriptor_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricDescriptor{}, &MetricDescriptorList{})
}
