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

type DiskIAMPolicyObservation struct {

	// The etag of the IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskIAMPolicyParameters struct {

	// Used to find the parent resource to bind the IAM policy to
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Disk
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Disk in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The policy data generated by
	// a google_iam_policy data source.
	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the zone where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskIAMPolicySpec defines the desired state of DiskIAMPolicy
type DiskIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskIAMPolicyParameters `json:"forProvider"`
}

// DiskIAMPolicyStatus defines the observed state of DiskIAMPolicy.
type DiskIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIAMPolicy is the Schema for the DiskIAMPolicys API. Collection of resources to manage IAM policy for Compute Engine Disk
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DiskIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskIAMPolicySpec   `json:"spec"`
	Status            DiskIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIAMPolicyList contains a list of DiskIAMPolicys
type DiskIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	DiskIAMPolicy_Kind             = "DiskIAMPolicy"
	DiskIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskIAMPolicy_Kind}.String()
	DiskIAMPolicy_KindAPIVersion   = DiskIAMPolicy_Kind + "." + CRDGroupVersion.String()
	DiskIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DiskIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskIAMPolicy{}, &DiskIAMPolicyList{})
}
