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

type ServiceNetworkingPeeredDNSDomainInitParameters struct {

	// The DNS domain suffix of the peered DNS domain. Make sure to suffix with a . (dot).
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// The producer project number. If not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ServiceNetworkingPeeredDNSDomainObservation struct {

	// The DNS domain suffix of the peered DNS domain. Make sure to suffix with a . (dot).
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// an identifier for the resource with format services/{{service}}/projects/{{project}}/global/networks/{{network}}/peeredDnsDomains/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The network in the consumer project.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// an identifier for the resource with format services/{{service}}/projects/{{project}}/global/networks/{{network}}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// The producer project number. If not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Private service connection between service and consumer network, defaults to servicenetworking.googleapis.com
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceNetworkingPeeredDNSDomainParameters struct {

	// The DNS domain suffix of the peered DNS domain. Make sure to suffix with a . (dot).
	// +kubebuilder:validation:Optional
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// The network in the consumer project.
	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`

	// The producer project number. If not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Private service connection between service and consumer network, defaults to servicenetworking.googleapis.com
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainSpec defines the desired state of ServiceNetworkingPeeredDNSDomain
type ServiceNetworkingPeeredDNSDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceNetworkingPeeredDNSDomainParameters `json:"forProvider"`
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
	InitProvider ServiceNetworkingPeeredDNSDomainInitParameters `json:"initProvider,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainStatus defines the observed state of ServiceNetworkingPeeredDNSDomain.
type ServiceNetworkingPeeredDNSDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceNetworkingPeeredDNSDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceNetworkingPeeredDNSDomain is the Schema for the ServiceNetworkingPeeredDNSDomains API. Allows management of a single peered DNS domain on a project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceNetworkingPeeredDNSDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsSuffix) || (has(self.initProvider) && has(self.initProvider.dnsSuffix))",message="spec.forProvider.dnsSuffix is a required parameter"
	Spec   ServiceNetworkingPeeredDNSDomainSpec   `json:"spec"`
	Status ServiceNetworkingPeeredDNSDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkingPeeredDNSDomainList contains a list of ServiceNetworkingPeeredDNSDomains
type ServiceNetworkingPeeredDNSDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkingPeeredDNSDomain `json:"items"`
}

// Repository type metadata.
var (
	ServiceNetworkingPeeredDNSDomain_Kind             = "ServiceNetworkingPeeredDNSDomain"
	ServiceNetworkingPeeredDNSDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceNetworkingPeeredDNSDomain_Kind}.String()
	ServiceNetworkingPeeredDNSDomain_KindAPIVersion   = ServiceNetworkingPeeredDNSDomain_Kind + "." + CRDGroupVersion.String()
	ServiceNetworkingPeeredDNSDomain_GroupVersionKind = CRDGroupVersion.WithKind(ServiceNetworkingPeeredDNSDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceNetworkingPeeredDNSDomain{}, &ServiceNetworkingPeeredDNSDomainList{})
}
