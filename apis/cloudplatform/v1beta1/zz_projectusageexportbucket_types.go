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

type ProjectUsageExportBucketInitParameters struct {

	// :  The bucket to store reports in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// :  A prefix for the reports, for instance, the project name.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// :  The project to set the export bucket on. If it is not provided, the provider project is used.
	// +crossplane:generate:reference:type=Project
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

type ProjectUsageExportBucketObservation struct {

	// :  The bucket to store reports in.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// :  A prefix for the reports, for instance, the project name.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// :  The project to set the export bucket on. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectUsageExportBucketParameters struct {

	// :  The bucket to store reports in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// :  A prefix for the reports, for instance, the project name.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// :  The project to set the export bucket on. If it is not provided, the provider project is used.
	// +crossplane:generate:reference:type=Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`
}

// ProjectUsageExportBucketSpec defines the desired state of ProjectUsageExportBucket
type ProjectUsageExportBucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectUsageExportBucketParameters `json:"forProvider"`
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
	InitProvider ProjectUsageExportBucketInitParameters `json:"initProvider,omitempty"`
}

// ProjectUsageExportBucketStatus defines the observed state of ProjectUsageExportBucket.
type ProjectUsageExportBucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectUsageExportBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectUsageExportBucket is the Schema for the ProjectUsageExportBuckets API. Manages a project's usage export bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectUsageExportBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectUsageExportBucketSpec   `json:"spec"`
	Status            ProjectUsageExportBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectUsageExportBucketList contains a list of ProjectUsageExportBuckets
type ProjectUsageExportBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectUsageExportBucket `json:"items"`
}

// Repository type metadata.
var (
	ProjectUsageExportBucket_Kind             = "ProjectUsageExportBucket"
	ProjectUsageExportBucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectUsageExportBucket_Kind}.String()
	ProjectUsageExportBucket_KindAPIVersion   = ProjectUsageExportBucket_Kind + "." + CRDGroupVersion.String()
	ProjectUsageExportBucket_GroupVersionKind = CRDGroupVersion.WithKind(ProjectUsageExportBucket_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectUsageExportBucket{}, &ProjectUsageExportBucketList{})
}
