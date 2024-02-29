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

type GlobalNetworkEndpointGroupInitParameters struct {

	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Type of network endpoints in this network endpoint group.
	// Possible values are: INTERNET_IP_PORT, INTERNET_FQDN_PORT.
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type GlobalNetworkEndpointGroupObservation struct {

	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/networkEndpointGroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of network endpoints in this network endpoint group.
	// Possible values are: INTERNET_IP_PORT, INTERNET_FQDN_PORT.
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type GlobalNetworkEndpointGroupParameters struct {

	// The default port used if the port number is not specified in the
	// network endpoint.
	// +kubebuilder:validation:Optional
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Type of network endpoints in this network endpoint group.
	// Possible values are: INTERNET_IP_PORT, INTERNET_FQDN_PORT.
	// +kubebuilder:validation:Optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// GlobalNetworkEndpointGroupSpec defines the desired state of GlobalNetworkEndpointGroup
type GlobalNetworkEndpointGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalNetworkEndpointGroupParameters `json:"forProvider"`
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
	InitProvider GlobalNetworkEndpointGroupInitParameters `json:"initProvider,omitempty"`
}

// GlobalNetworkEndpointGroupStatus defines the observed state of GlobalNetworkEndpointGroup.
type GlobalNetworkEndpointGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalNetworkEndpointGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GlobalNetworkEndpointGroup is the Schema for the GlobalNetworkEndpointGroups API. A global network endpoint group contains endpoints that reside outside of Google Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GlobalNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkEndpointType) || (has(self.initProvider) && has(self.initProvider.networkEndpointType))",message="spec.forProvider.networkEndpointType is a required parameter"
	Spec   GlobalNetworkEndpointGroupSpec   `json:"spec"`
	Status GlobalNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalNetworkEndpointGroupList contains a list of GlobalNetworkEndpointGroups
type GlobalNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalNetworkEndpointGroup `json:"items"`
}

// Repository type metadata.
var (
	GlobalNetworkEndpointGroup_Kind             = "GlobalNetworkEndpointGroup"
	GlobalNetworkEndpointGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalNetworkEndpointGroup_Kind}.String()
	GlobalNetworkEndpointGroup_KindAPIVersion   = GlobalNetworkEndpointGroup_Kind + "." + CRDGroupVersion.String()
	GlobalNetworkEndpointGroup_GroupVersionKind = CRDGroupVersion.WithKind(GlobalNetworkEndpointGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalNetworkEndpointGroup{}, &GlobalNetworkEndpointGroupList{})
}
