//go:build (firebaserules || all) && !ignore_autogenerated

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

type ReleaseInitParameters struct {

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Ruleset referred to by this Release. The Ruleset must exist for the Release to be created.
	// +crossplane:generate:reference:type=Ruleset
	RulesetName *string `json:"rulesetName,omitempty" tf:"ruleset_name,omitempty"`

	// Reference to a Ruleset to populate rulesetName.
	// +kubebuilder:validation:Optional
	RulesetNameRef *v1.Reference `json:"rulesetNameRef,omitempty" tf:"-"`

	// Selector for a Ruleset to populate rulesetName.
	// +kubebuilder:validation:Optional
	RulesetNameSelector *v1.Selector `json:"rulesetNameSelector,omitempty" tf:"-"`
}

type ReleaseObservation struct {

	// Output only. Time the release was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// an identifier for the resource with format projects/{{project}}/releases/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Ruleset referred to by this Release. The Ruleset must exist for the Release to be created.
	RulesetName *string `json:"rulesetName,omitempty" tf:"ruleset_name,omitempty"`

	// Output only. Time the release was updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ReleaseParameters struct {

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the Ruleset referred to by this Release. The Ruleset must exist for the Release to be created.
	// +crossplane:generate:reference:type=Ruleset
	// +kubebuilder:validation:Optional
	RulesetName *string `json:"rulesetName,omitempty" tf:"ruleset_name,omitempty"`

	// Reference to a Ruleset to populate rulesetName.
	// +kubebuilder:validation:Optional
	RulesetNameRef *v1.Reference `json:"rulesetNameRef,omitempty" tf:"-"`

	// Selector for a Ruleset to populate rulesetName.
	// +kubebuilder:validation:Optional
	RulesetNameSelector *v1.Selector `json:"rulesetNameSelector,omitempty" tf:"-"`
}

// ReleaseSpec defines the desired state of Release
type ReleaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReleaseParameters `json:"forProvider"`
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
	InitProvider ReleaseInitParameters `json:"initProvider,omitempty"`
}

// ReleaseStatus defines the observed state of Release.
type ReleaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReleaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Release is the Schema for the Releases API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReleaseSpec   `json:"spec"`
	Status            ReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReleaseList contains a list of Releases
type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Release `json:"items"`
}

// Repository type metadata.
var (
	Release_Kind             = "Release"
	Release_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Release_Kind}.String()
	Release_KindAPIVersion   = Release_Kind + "." + CRDGroupVersion.String()
	Release_GroupVersionKind = CRDGroupVersion.WithKind(Release_Kind)
)

func init() {
	SchemeBuilder.Register(&Release{}, &ReleaseList{})
}
