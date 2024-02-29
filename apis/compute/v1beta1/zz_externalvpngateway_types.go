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

type ExternalVPNGatewayInitParameters struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interface []InterfaceInitParameters `json:"interface,omitempty" tf:"interface,omitempty"`

	// Labels for the external VPN gateway resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: FOUR_IPS_REDUNDANCY, SINGLE_IP_INTERNALLY_REDUNDANT, TWO_IPS_REDUNDANCY.
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type,omitempty"`
}

type ExternalVPNGatewayObservation struct {

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/externalVpnGateways/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interface []InterfaceObservation `json:"interface,omitempty" tf:"interface,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels for the external VPN gateway resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: FOUR_IPS_REDUNDANCY, SINGLE_IP_INTERNALLY_REDUNDANT, TWO_IPS_REDUNDANCY.
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ExternalVPNGatewayParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Interface []InterfaceParameters `json:"interface,omitempty" tf:"interface,omitempty"`

	// Labels for the external VPN gateway resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates the redundancy type of this external VPN gateway
	// Possible values are: FOUR_IPS_REDUNDANCY, SINGLE_IP_INTERNALLY_REDUNDANT, TWO_IPS_REDUNDANCY.
	// +kubebuilder:validation:Optional
	RedundancyType *string `json:"redundancyType,omitempty" tf:"redundancy_type,omitempty"`
}

type InterfaceInitParameters struct {

	// The numeric ID for this interface. Allowed values are based on the redundancy type
	// of this external VPN gateway
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the interface in the external VPN gateway.
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type InterfaceObservation struct {

	// The numeric ID for this interface. Allowed values are based on the redundancy type
	// of this external VPN gateway
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the interface in the external VPN gateway.
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type InterfaceParameters struct {

	// The numeric ID for this interface. Allowed values are based on the redundancy type
	// of this external VPN gateway
	// +kubebuilder:validation:Optional
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// IP address of the interface in the external VPN gateway.
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

// ExternalVPNGatewaySpec defines the desired state of ExternalVPNGateway
type ExternalVPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalVPNGatewayParameters `json:"forProvider"`
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
	InitProvider ExternalVPNGatewayInitParameters `json:"initProvider,omitempty"`
}

// ExternalVPNGatewayStatus defines the observed state of ExternalVPNGateway.
type ExternalVPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalVPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ExternalVPNGateway is the Schema for the ExternalVPNGateways API. Represents a VPN gateway managed outside of GCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ExternalVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalVPNGatewaySpec   `json:"spec"`
	Status            ExternalVPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalVPNGatewayList contains a list of ExternalVPNGateways
type ExternalVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalVPNGateway `json:"items"`
}

// Repository type metadata.
var (
	ExternalVPNGateway_Kind             = "ExternalVPNGateway"
	ExternalVPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalVPNGateway_Kind}.String()
	ExternalVPNGateway_KindAPIVersion   = ExternalVPNGateway_Kind + "." + CRDGroupVersion.String()
	ExternalVPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(ExternalVPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalVPNGateway{}, &ExternalVPNGatewayList{})
}
