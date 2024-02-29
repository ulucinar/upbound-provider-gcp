//go:build (compute || all) && !ignore_autogenerated

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

type SharedVPCServiceProjectInitParameters struct {

	// The deletion policy for the shared VPC service. Setting ABANDON allows the resource to be abandoned rather than deleted. Possible values are: "ABANDON".
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The ID of a host project to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractProjectID()
	HostProject *string `json:"hostProject,omitempty" tf:"host_project,omitempty"`

	// Reference to a Project in cloudplatform to populate hostProject.
	// +kubebuilder:validation:Optional
	HostProjectRef *v1.Reference `json:"hostProjectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate hostProject.
	// +kubebuilder:validation:Optional
	HostProjectSelector *v1.Selector `json:"hostProjectSelector,omitempty" tf:"-"`

	// The ID of the project that will serve as a Shared VPC service project.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractProjectID()
	ServiceProject *string `json:"serviceProject,omitempty" tf:"service_project,omitempty"`

	// Reference to a Project in cloudplatform to populate serviceProject.
	// +kubebuilder:validation:Optional
	ServiceProjectRef *v1.Reference `json:"serviceProjectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate serviceProject.
	// +kubebuilder:validation:Optional
	ServiceProjectSelector *v1.Selector `json:"serviceProjectSelector,omitempty" tf:"-"`
}

type SharedVPCServiceProjectObservation struct {

	// The deletion policy for the shared VPC service. Setting ABANDON allows the resource to be abandoned rather than deleted. Possible values are: "ABANDON".
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The ID of a host project to associate.
	HostProject *string `json:"hostProject,omitempty" tf:"host_project,omitempty"`

	// an identifier for the resource with format {{host_project}}/{{service_project}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project that will serve as a Shared VPC service project.
	ServiceProject *string `json:"serviceProject,omitempty" tf:"service_project,omitempty"`
}

type SharedVPCServiceProjectParameters struct {

	// The deletion policy for the shared VPC service. Setting ABANDON allows the resource to be abandoned rather than deleted. Possible values are: "ABANDON".
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// The ID of a host project to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractProjectID()
	// +kubebuilder:validation:Optional
	HostProject *string `json:"hostProject,omitempty" tf:"host_project,omitempty"`

	// Reference to a Project in cloudplatform to populate hostProject.
	// +kubebuilder:validation:Optional
	HostProjectRef *v1.Reference `json:"hostProjectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate hostProject.
	// +kubebuilder:validation:Optional
	HostProjectSelector *v1.Selector `json:"hostProjectSelector,omitempty" tf:"-"`

	// The ID of the project that will serve as a Shared VPC service project.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractProjectID()
	// +kubebuilder:validation:Optional
	ServiceProject *string `json:"serviceProject,omitempty" tf:"service_project,omitempty"`

	// Reference to a Project in cloudplatform to populate serviceProject.
	// +kubebuilder:validation:Optional
	ServiceProjectRef *v1.Reference `json:"serviceProjectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate serviceProject.
	// +kubebuilder:validation:Optional
	ServiceProjectSelector *v1.Selector `json:"serviceProjectSelector,omitempty" tf:"-"`
}

// SharedVPCServiceProjectSpec defines the desired state of SharedVPCServiceProject
type SharedVPCServiceProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedVPCServiceProjectParameters `json:"forProvider"`
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
	InitProvider SharedVPCServiceProjectInitParameters `json:"initProvider,omitempty"`
}

// SharedVPCServiceProjectStatus defines the observed state of SharedVPCServiceProject.
type SharedVPCServiceProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedVPCServiceProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SharedVPCServiceProject is the Schema for the SharedVPCServiceProjects API. Enables the Google Compute Engine Shared VPC feature for a project, assigning it as a service project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SharedVPCServiceProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SharedVPCServiceProjectSpec   `json:"spec"`
	Status            SharedVPCServiceProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedVPCServiceProjectList contains a list of SharedVPCServiceProjects
type SharedVPCServiceProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedVPCServiceProject `json:"items"`
}

// Repository type metadata.
var (
	SharedVPCServiceProject_Kind             = "SharedVPCServiceProject"
	SharedVPCServiceProject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedVPCServiceProject_Kind}.String()
	SharedVPCServiceProject_KindAPIVersion   = SharedVPCServiceProject_Kind + "." + CRDGroupVersion.String()
	SharedVPCServiceProject_GroupVersionKind = CRDGroupVersion.WithKind(SharedVPCServiceProject_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedVPCServiceProject{}, &SharedVPCServiceProjectList{})
}
