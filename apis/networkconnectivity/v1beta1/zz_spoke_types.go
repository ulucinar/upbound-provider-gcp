//go:build (networkconnectivity || all) && !ignore_autogenerated

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

type InstancesInitParameters struct {

	// The IP address on the VM to use for peering.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The URI of the virtual machine resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	VirtualMachine *string `json:"virtualMachine,omitempty" tf:"virtual_machine,omitempty"`

	// Reference to a Instance in compute to populate virtualMachine.
	// +kubebuilder:validation:Optional
	VirtualMachineRef *v1.Reference `json:"virtualMachineRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate virtualMachine.
	// +kubebuilder:validation:Optional
	VirtualMachineSelector *v1.Selector `json:"virtualMachineSelector,omitempty" tf:"-"`
}

type InstancesObservation struct {

	// The IP address on the VM to use for peering.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The URI of the virtual machine resource
	VirtualMachine *string `json:"virtualMachine,omitempty" tf:"virtual_machine,omitempty"`
}

type InstancesParameters struct {

	// The IP address on the VM to use for peering.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The URI of the virtual machine resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("self_link",true)
	// +kubebuilder:validation:Optional
	VirtualMachine *string `json:"virtualMachine,omitempty" tf:"virtual_machine,omitempty"`

	// Reference to a Instance in compute to populate virtualMachine.
	// +kubebuilder:validation:Optional
	VirtualMachineRef *v1.Reference `json:"virtualMachineRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate virtualMachine.
	// +kubebuilder:validation:Optional
	VirtualMachineSelector *v1.Selector `json:"virtualMachineSelector,omitempty" tf:"-"`
}

type LinkedInterconnectAttachmentsInitParameters struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked interconnect attachment resources
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type LinkedInterconnectAttachmentsObservation struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked interconnect attachment resources
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type LinkedInterconnectAttachmentsParameters struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	// +kubebuilder:validation:Optional
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked interconnect attachment resources
	// +kubebuilder:validation:Optional
	Uris []*string `json:"uris" tf:"uris,omitempty"`
}

type LinkedRouterApplianceInstancesInitParameters struct {

	// The list of router appliance instances
	Instances []InstancesInitParameters `json:"instances,omitempty" tf:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`
}

type LinkedRouterApplianceInstancesObservation struct {

	// The list of router appliance instances
	Instances []InstancesObservation `json:"instances,omitempty" tf:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`
}

type LinkedRouterApplianceInstancesParameters struct {

	// The list of router appliance instances
	// +kubebuilder:validation:Optional
	Instances []InstancesParameters `json:"instances" tf:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	// +kubebuilder:validation:Optional
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer" tf:"site_to_site_data_transfer,omitempty"`
}

type LinkedVPNTunnelsInitParameters struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked VPN tunnel resources.
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type LinkedVPNTunnelsObservation struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked VPN tunnel resources.
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type LinkedVPNTunnelsParameters struct {

	// A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.
	// +kubebuilder:validation:Optional
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer" tf:"site_to_site_data_transfer,omitempty"`

	// The URIs of linked VPN tunnel resources.
	// +kubebuilder:validation:Optional
	Uris []*string `json:"uris" tf:"uris,omitempty"`
}

type SpokeInitParameters struct {

	// An optional description of the spoke.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Immutable. The URI of the hub that this spoke is attached to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/networkconnectivity/v1beta1.Hub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Hub *string `json:"hub,omitempty" tf:"hub,omitempty"`

	// Reference to a Hub in networkconnectivity to populate hub.
	// +kubebuilder:validation:Optional
	HubRef *v1.Reference `json:"hubRef,omitempty" tf:"-"`

	// Selector for a Hub in networkconnectivity to populate hub.
	// +kubebuilder:validation:Optional
	HubSelector *v1.Selector `json:"hubSelector,omitempty" tf:"-"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments []LinkedInterconnectAttachmentsInitParameters `json:"linkedInterconnectAttachments,omitempty" tf:"linked_interconnect_attachments,omitempty"`

	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances []LinkedRouterApplianceInstancesInitParameters `json:"linkedRouterApplianceInstances,omitempty" tf:"linked_router_appliance_instances,omitempty"`

	// The URIs of linked VPN tunnel resources
	LinkedVPNTunnels []LinkedVPNTunnelsInitParameters `json:"linkedVpnTunnels,omitempty" tf:"linked_vpn_tunnels,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Immutable. The name of the spoke. Spoke names must be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type SpokeObservation struct {

	// Output only. The time the spoke was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// An optional description of the spoke.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Immutable. The URI of the hub that this spoke is attached to.
	Hub *string `json:"hub,omitempty" tf:"hub,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/spokes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments []LinkedInterconnectAttachmentsObservation `json:"linkedInterconnectAttachments,omitempty" tf:"linked_interconnect_attachments,omitempty"`

	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances []LinkedRouterApplianceInstancesObservation `json:"linkedRouterApplianceInstances,omitempty" tf:"linked_router_appliance_instances,omitempty"`

	// The URIs of linked VPN tunnel resources
	LinkedVPNTunnels []LinkedVPNTunnelsObservation `json:"linkedVpnTunnels,omitempty" tf:"linked_vpn_tunnels,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Immutable. The name of the spoke. Spoke names must be unique.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`

	// Output only. The time the spoke was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type SpokeParameters struct {

	// An optional description of the spoke.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Immutable. The URI of the hub that this spoke is attached to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/networkconnectivity/v1beta1.Hub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Hub *string `json:"hub,omitempty" tf:"hub,omitempty"`

	// Reference to a Hub in networkconnectivity to populate hub.
	// +kubebuilder:validation:Optional
	HubRef *v1.Reference `json:"hubRef,omitempty" tf:"-"`

	// Selector for a Hub in networkconnectivity to populate hub.
	// +kubebuilder:validation:Optional
	HubSelector *v1.Selector `json:"hubSelector,omitempty" tf:"-"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	// +kubebuilder:validation:Optional
	LinkedInterconnectAttachments []LinkedInterconnectAttachmentsParameters `json:"linkedInterconnectAttachments,omitempty" tf:"linked_interconnect_attachments,omitempty"`

	// The URIs of linked Router appliance resources
	// +kubebuilder:validation:Optional
	LinkedRouterApplianceInstances []LinkedRouterApplianceInstancesParameters `json:"linkedRouterApplianceInstances,omitempty" tf:"linked_router_appliance_instances,omitempty"`

	// The URIs of linked VPN tunnel resources
	// +kubebuilder:validation:Optional
	LinkedVPNTunnels []LinkedVPNTunnelsParameters `json:"linkedVpnTunnels,omitempty" tf:"linked_vpn_tunnels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Immutable. The name of the spoke. Spoke names must be unique.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// SpokeSpec defines the desired state of Spoke
type SpokeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpokeParameters `json:"forProvider"`
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
	InitProvider SpokeInitParameters `json:"initProvider,omitempty"`
}

// SpokeStatus defines the observed state of Spoke.
type SpokeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpokeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Spoke is the Schema for the Spokes API. The NetworkConnectivity Spoke resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Spoke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SpokeSpec   `json:"spec"`
	Status SpokeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpokeList contains a list of Spokes
type SpokeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spoke `json:"items"`
}

// Repository type metadata.
var (
	Spoke_Kind             = "Spoke"
	Spoke_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Spoke_Kind}.String()
	Spoke_KindAPIVersion   = Spoke_Kind + "." + CRDGroupVersion.String()
	Spoke_GroupVersionKind = CRDGroupVersion.WithKind(Spoke_Kind)
)

func init() {
	SchemeBuilder.Register(&Spoke{}, &SpokeList{})
}
