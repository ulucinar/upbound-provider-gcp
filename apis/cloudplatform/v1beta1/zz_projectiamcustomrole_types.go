//go:build (cloudplatform || all) && !ignore_autogenerated

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

type ProjectIAMCustomRoleInitParameters struct {

	// A human-readable description for the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ProjectIAMCustomRoleObservation struct {

	// The current deleted state of the role.
	Deleted *bool `json:"deleted,omitempty" tf:"deleted,omitempty"`

	// A human-readable description for the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with the format projects/{{project}}/roles/{{role_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the role in the format projects/{{project}}/roles/{{role_id}}. Like id, this field can be used as a reference in other resources such as IAM role bindings.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ProjectIAMCustomRoleParameters struct {

	// A human-readable description for the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +kubebuilder:validation:Optional
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The project that the service account will be created in.
	// Defaults to the provider project configuration.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// ProjectIAMCustomRoleSpec defines the desired state of ProjectIAMCustomRole
type ProjectIAMCustomRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMCustomRoleParameters `json:"forProvider"`
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
	InitProvider ProjectIAMCustomRoleInitParameters `json:"initProvider,omitempty"`
}

// ProjectIAMCustomRoleStatus defines the observed state of ProjectIAMCustomRole.
type ProjectIAMCustomRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMCustomRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectIAMCustomRole is the Schema for the ProjectIAMCustomRoles API. Allows management of a customized Cloud IAM project role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectIAMCustomRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	Spec   ProjectIAMCustomRoleSpec   `json:"spec"`
	Status ProjectIAMCustomRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMCustomRoleList contains a list of ProjectIAMCustomRoles
type ProjectIAMCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMCustomRole `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMCustomRole_Kind             = "ProjectIAMCustomRole"
	ProjectIAMCustomRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectIAMCustomRole_Kind}.String()
	ProjectIAMCustomRole_KindAPIVersion   = ProjectIAMCustomRole_Kind + "." + CRDGroupVersion.String()
	ProjectIAMCustomRole_GroupVersionKind = CRDGroupVersion.WithKind(ProjectIAMCustomRole_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMCustomRole{}, &ProjectIAMCustomRoleList{})
}
