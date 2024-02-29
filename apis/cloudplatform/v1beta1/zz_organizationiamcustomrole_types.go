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

type OrganizationIAMCustomRoleInitParameters struct {

	// A human-readable description for the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The numeric ID of the organization in which you want to create a custom role.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The role id to use for this role.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type OrganizationIAMCustomRoleObservation struct {

	// The current deleted state of the role.
	Deleted *bool `json:"deleted,omitempty" tf:"deleted,omitempty"`

	// A human-readable description for the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with the format organizations/{{org_id}}/roles/{{role_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the role in the format organizations/{{org_id}}/roles/{{role_id}}. Like id, this field can be used as a reference in other resources such as IAM role bindings.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The numeric ID of the organization in which you want to create a custom role.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The role id to use for this role.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type OrganizationIAMCustomRoleParameters struct {

	// A human-readable description for the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The numeric ID of the organization in which you want to create a custom role.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	// +kubebuilder:validation:Optional
	// +listType=set
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The role id to use for this role.
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The current launch stage of the role.
	// Defaults to GA.
	// List of possible stages is here.
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// A human-readable title for the role.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// OrganizationIAMCustomRoleSpec defines the desired state of OrganizationIAMCustomRole
type OrganizationIAMCustomRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationIAMCustomRoleParameters `json:"forProvider"`
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
	InitProvider OrganizationIAMCustomRoleInitParameters `json:"initProvider,omitempty"`
}

// OrganizationIAMCustomRoleStatus defines the observed state of OrganizationIAMCustomRole.
type OrganizationIAMCustomRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationIAMCustomRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrganizationIAMCustomRole is the Schema for the OrganizationIAMCustomRoles API. Allows management of a customized Cloud IAM organization role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type OrganizationIAMCustomRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.orgId) || (has(self.initProvider) && has(self.initProvider.orgId))",message="spec.forProvider.orgId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleId) || (has(self.initProvider) && has(self.initProvider.roleId))",message="spec.forProvider.roleId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	Spec   OrganizationIAMCustomRoleSpec   `json:"spec"`
	Status OrganizationIAMCustomRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIAMCustomRoleList contains a list of OrganizationIAMCustomRoles
type OrganizationIAMCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationIAMCustomRole `json:"items"`
}

// Repository type metadata.
var (
	OrganizationIAMCustomRole_Kind             = "OrganizationIAMCustomRole"
	OrganizationIAMCustomRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationIAMCustomRole_Kind}.String()
	OrganizationIAMCustomRole_KindAPIVersion   = OrganizationIAMCustomRole_Kind + "." + CRDGroupVersion.String()
	OrganizationIAMCustomRole_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationIAMCustomRole_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationIAMCustomRole{}, &OrganizationIAMCustomRoleList{})
}
