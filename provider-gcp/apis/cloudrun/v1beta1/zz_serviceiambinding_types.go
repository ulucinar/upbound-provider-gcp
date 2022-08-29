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

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type ServiceIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a Service to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a Service to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

// ServiceIAMBindingSpec defines the desired state of ServiceIAMBinding
type ServiceIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceIAMBindingParameters `json:"forProvider"`
}

// ServiceIAMBindingStatus defines the observed state of ServiceIAMBinding.
type ServiceIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMBinding is the Schema for the ServiceIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceIAMBindingSpec   `json:"spec"`
	Status            ServiceIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIAMBindingList contains a list of ServiceIAMBindings
type ServiceIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	ServiceIAMBinding_Kind             = "ServiceIAMBinding"
	ServiceIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceIAMBinding_Kind}.String()
	ServiceIAMBinding_KindAPIVersion   = ServiceIAMBinding_Kind + "." + CRDGroupVersion.String()
	ServiceIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(ServiceIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceIAMBinding{}, &ServiceIAMBindingList{})
}
