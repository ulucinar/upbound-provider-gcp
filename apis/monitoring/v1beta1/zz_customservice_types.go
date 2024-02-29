//go:build (monitoring || all) && !ignore_autogenerated

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

type CustomServiceInitParameters struct {

	// Name used for UI elements listing this Service.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry []TelemetryInitParameters `json:"telemetry,omitempty" tf:"telemetry,omitempty"`

	// Labels which have been used to annotate the service. Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	// +mapType=granular
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type CustomServiceObservation struct {

	// Name used for UI elements listing this Service.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID]/services/[SERVICE_ID].
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	Telemetry []TelemetryObservation `json:"telemetry,omitempty" tf:"telemetry,omitempty"`

	// Labels which have been used to annotate the service. Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	// +mapType=granular
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type CustomServiceParameters struct {

	// Name used for UI elements listing this Service.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Configuration for how to query telemetry on a Service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Telemetry []TelemetryParameters `json:"telemetry,omitempty" tf:"telemetry,omitempty"`

	// Labels which have been used to annotate the service. Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

type TelemetryInitParameters struct {

	// The full name of the resource that defines this service.
	// Formatted as described in
	// https://cloud.google.com/apis/design/resource_names.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

type TelemetryObservation struct {

	// The full name of the resource that defines this service.
	// Formatted as described in
	// https://cloud.google.com/apis/design/resource_names.
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

type TelemetryParameters struct {

	// The full name of the resource that defines this service.
	// Formatted as described in
	// https://cloud.google.com/apis/design/resource_names.
	// +kubebuilder:validation:Optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`
}

// CustomServiceSpec defines the desired state of CustomService
type CustomServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomServiceParameters `json:"forProvider"`
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
	InitProvider CustomServiceInitParameters `json:"initProvider,omitempty"`
}

// CustomServiceStatus defines the observed state of CustomService.
type CustomServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomService is the Schema for the CustomServices API. A Service is a discrete, autonomous, and network-accessible unit, designed to solve an individual concern (Wikipedia).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CustomService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomServiceSpec   `json:"spec"`
	Status            CustomServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomServiceList contains a list of CustomServices
type CustomServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomService `json:"items"`
}

// Repository type metadata.
var (
	CustomService_Kind             = "CustomService"
	CustomService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomService_Kind}.String()
	CustomService_KindAPIVersion   = CustomService_Kind + "." + CRDGroupVersion.String()
	CustomService_GroupVersionKind = CRDGroupVersion.WithKind(CustomService_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomService{}, &CustomServiceList{})
}
