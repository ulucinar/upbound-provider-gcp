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

type ProjectSinkBigqueryOptionsInitParameters struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	UsePartitionedTables *bool `json:"usePartitionedTables,omitempty" tf:"use_partitioned_tables,omitempty"`
}

type ProjectSinkBigqueryOptionsObservation struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	UsePartitionedTables *bool `json:"usePartitionedTables,omitempty" tf:"use_partitioned_tables,omitempty"`
}

type ProjectSinkBigqueryOptionsParameters struct {

	// Whether to use BigQuery's partition tables.
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned
	// tables the date suffix is no longer present and special query syntax
	// has to be used instead. In both cases, tables are sharded based on UTC timezone.
	// +kubebuilder:validation:Optional
	UsePartitionedTables *bool `json:"usePartitionedTables" tf:"use_partitioned_tables,omitempty"`
}

type ProjectSinkExclusionsInitParameters struct {

	// A description of this exclusion.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProjectSinkExclusionsObservation struct {

	// A description of this exclusion.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProjectSinkExclusionsParameters struct {

	// A description of this exclusion.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// A client-assigned identifier, such as load-balancer-exclusion. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ProjectSinkInitParameters struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions []ProjectSinkBigqueryOptionsInitParameters `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusionsInitParameters `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Whether or not to create a unique identity associated with this sink. If false
	// (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// bigquery_options, you must set unique_writer_identity to true.
	UniqueWriterIdentity *bool `json:"uniqueWriterIdentity,omitempty" tf:"unique_writer_identity,omitempty"`
}

type ProjectSinkObservation struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions []ProjectSinkBigqueryOptionsObservation `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusionsObservation `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// an identifier for the resource with format projects/{{project}}/sinks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Whether or not to create a unique identity associated with this sink. If false
	// (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// bigquery_options, you must set unique_writer_identity to true.
	UniqueWriterIdentity *bool `json:"uniqueWriterIdentity,omitempty" tf:"unique_writer_identity,omitempty"`

	// The identity associated with this sink. This identity must be granted write access to the
	// configured destination.
	WriterIdentity *string `json:"writerIdentity,omitempty" tf:"writer_identity,omitempty"`
}

type ProjectSinkParameters struct {

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	// +kubebuilder:validation:Optional
	BigqueryOptions []ProjectSinkBigqueryOptionsParameters `json:"bigqueryOptions,omitempty" tf:"bigquery_options,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusions.filter, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	// +kubebuilder:validation:Optional
	Exclusions []ProjectSinkExclusionsParameters `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Whether or not to create a unique identity associated with this sink. If false
	// (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// bigquery_options, you must set unique_writer_identity to true.
	// +kubebuilder:validation:Optional
	UniqueWriterIdentity *bool `json:"uniqueWriterIdentity,omitempty" tf:"unique_writer_identity,omitempty"`
}

// ProjectSinkSpec defines the desired state of ProjectSink
type ProjectSinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectSinkParameters `json:"forProvider"`
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
	InitProvider ProjectSinkInitParameters `json:"initProvider,omitempty"`
}

// ProjectSinkStatus defines the observed state of ProjectSink.
type ProjectSinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectSinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectSink is the Schema for the ProjectSinks API. Manages a project-level logging sink.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	Spec   ProjectSinkSpec   `json:"spec"`
	Status ProjectSinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectSinkList contains a list of ProjectSinks
type ProjectSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectSink `json:"items"`
}

// Repository type metadata.
var (
	ProjectSink_Kind             = "ProjectSink"
	ProjectSink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectSink_Kind}.String()
	ProjectSink_KindAPIVersion   = ProjectSink_Kind + "." + CRDGroupVersion.String()
	ProjectSink_GroupVersionKind = CRDGroupVersion.WithKind(ProjectSink_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectSink{}, &ProjectSinkList{})
}
