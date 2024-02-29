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

type CollectorIlbInitParameters struct {

	// The URL of the forwarding rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a ForwardingRule in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type CollectorIlbObservation struct {

	// The URL of the forwarding rule.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type CollectorIlbParameters struct {

	// The URL of the forwarding rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a ForwardingRule in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type FilterInitParameters struct {

	// IP CIDR ranges that apply as a filter on the source (ingress) or
	// destination (egress) IP in the IP header. Only IPv4 is supported.
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Direction of traffic to mirror.
	// Default value is BOTH.
	// Possible values are: INGRESS, EGRESS, BOTH.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Possible IP protocols including tcp, udp, icmp and esp
	IPProtocols []*string `json:"ipProtocols,omitempty" tf:"ip_protocols,omitempty"`
}

type FilterObservation struct {

	// IP CIDR ranges that apply as a filter on the source (ingress) or
	// destination (egress) IP in the IP header. Only IPv4 is supported.
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Direction of traffic to mirror.
	// Default value is BOTH.
	// Possible values are: INGRESS, EGRESS, BOTH.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Possible IP protocols including tcp, udp, icmp and esp
	IPProtocols []*string `json:"ipProtocols,omitempty" tf:"ip_protocols,omitempty"`
}

type FilterParameters struct {

	// IP CIDR ranges that apply as a filter on the source (ingress) or
	// destination (egress) IP in the IP header. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Direction of traffic to mirror.
	// Default value is BOTH.
	// Possible values are: INGRESS, EGRESS, BOTH.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Possible IP protocols including tcp, udp, icmp and esp
	// +kubebuilder:validation:Optional
	IPProtocols []*string `json:"ipProtocols,omitempty" tf:"ip_protocols,omitempty"`
}

type InstancesInitParameters struct {

	// The URL of the subnetwork where this rule should be active.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a Instance in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type InstancesObservation struct {

	// The URL of the subnetwork where this rule should be active.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type InstancesParameters struct {

	// The URL of the subnetwork where this rule should be active.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a Instance in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type MirroredResourcesInitParameters struct {

	// All the listed instances will be mirrored.  Specify at most 50.
	// Structure is documented below.
	Instances []InstancesInitParameters `json:"instances,omitempty" tf:"instances,omitempty"`

	// All instances in one of these subnetworks will be mirrored.
	// Structure is documented below.
	Subnetworks []SubnetworksInitParameters `json:"subnetworks,omitempty" tf:"subnetworks,omitempty"`

	// All instances with these tags will be mirrored.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MirroredResourcesObservation struct {

	// All the listed instances will be mirrored.  Specify at most 50.
	// Structure is documented below.
	Instances []InstancesObservation `json:"instances,omitempty" tf:"instances,omitempty"`

	// All instances in one of these subnetworks will be mirrored.
	// Structure is documented below.
	Subnetworks []SubnetworksObservation `json:"subnetworks,omitempty" tf:"subnetworks,omitempty"`

	// All instances with these tags will be mirrored.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MirroredResourcesParameters struct {

	// All the listed instances will be mirrored.  Specify at most 50.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Instances []InstancesParameters `json:"instances,omitempty" tf:"instances,omitempty"`

	// All instances in one of these subnetworks will be mirrored.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Subnetworks []SubnetworksParameters `json:"subnetworks,omitempty" tf:"subnetworks,omitempty"`

	// All instances with these tags will be mirrored.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PacketMirroringInitParameters struct {

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have is_mirroring_collector
	// set to true.
	// Structure is documented below.
	CollectorIlb []CollectorIlbInitParameters `json:"collectorIlb,omitempty" tf:"collector_ilb,omitempty"`

	// A human-readable description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter []FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources []MirroredResourcesInitParameters `json:"mirroredResources,omitempty" tf:"mirrored_resources,omitempty"`

	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network []PacketMirroringNetworkInitParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type PacketMirroringNetworkInitParameters struct {

	// The full self_link URL of the network where this rule is active.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a Network in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type PacketMirroringNetworkObservation struct {

	// The full self_link URL of the network where this rule is active.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PacketMirroringNetworkParameters struct {

	// The full self_link URL of the network where this rule is active.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Reference to a Network in compute to populate url.
	// +kubebuilder:validation:Optional
	URLRef *v1.Reference `json:"urlRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate url.
	// +kubebuilder:validation:Optional
	URLSelector *v1.Selector `json:"urlSelector,omitempty" tf:"-"`
}

type PacketMirroringObservation struct {

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have is_mirroring_collector
	// set to true.
	// Structure is documented below.
	CollectorIlb []CollectorIlbObservation `json:"collectorIlb,omitempty" tf:"collector_ilb,omitempty"`

	// A human-readable description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources []MirroredResourcesObservation `json:"mirroredResources,omitempty" tf:"mirrored_resources,omitempty"`

	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network []PacketMirroringNetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PacketMirroringParameters struct {

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have is_mirroring_collector
	// set to true.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CollectorIlb []CollectorIlbParameters `json:"collectorIlb,omitempty" tf:"collector_ilb,omitempty"`

	// A human-readable description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// A means of specifying which resources to mirror.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MirroredResources []MirroredResourcesParameters `json:"mirroredResources,omitempty" tf:"mirrored_resources,omitempty"`

	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Network []PacketMirroringNetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type SubnetworksInitParameters struct {

	// The URL of the subnetwork where this rule should be active.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SubnetworksObservation struct {

	// The URL of the subnetwork where this rule should be active.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SubnetworksParameters struct {

	// The URL of the subnetwork where this rule should be active.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// PacketMirroringSpec defines the desired state of PacketMirroring
type PacketMirroringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PacketMirroringParameters `json:"forProvider"`
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
	InitProvider PacketMirroringInitParameters `json:"initProvider,omitempty"`
}

// PacketMirroringStatus defines the observed state of PacketMirroring.
type PacketMirroringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PacketMirroringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PacketMirroring is the Schema for the PacketMirrorings API. Packet Mirroring mirrors traffic to and from particular VM instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type PacketMirroring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.collectorIlb) || (has(self.initProvider) && has(self.initProvider.collectorIlb))",message="spec.forProvider.collectorIlb is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mirroredResources) || (has(self.initProvider) && has(self.initProvider.mirroredResources))",message="spec.forProvider.mirroredResources is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.network) || (has(self.initProvider) && has(self.initProvider.network))",message="spec.forProvider.network is a required parameter"
	Spec   PacketMirroringSpec   `json:"spec"`
	Status PacketMirroringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PacketMirroringList contains a list of PacketMirrorings
type PacketMirroringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PacketMirroring `json:"items"`
}

// Repository type metadata.
var (
	PacketMirroring_Kind             = "PacketMirroring"
	PacketMirroring_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PacketMirroring_Kind}.String()
	PacketMirroring_KindAPIVersion   = PacketMirroring_Kind + "." + CRDGroupVersion.String()
	PacketMirroring_GroupVersionKind = CRDGroupVersion.WithKind(PacketMirroring_Kind)
)

func init() {
	SchemeBuilder.Register(&PacketMirroring{}, &PacketMirroringList{})
}
