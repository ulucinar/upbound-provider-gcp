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

type RouterInterfaceInitParameters struct {

	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`

	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the project in which this interface's router belongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface *string `json:"redundantInterface,omitempty" tf:"redundant_interface,omitempty"`

	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	// +crossplane:generate:reference:type=Router
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	// +crossplane:generate:reference:type=VPNTunnel
	VPNTunnel *string `json:"vpnTunnel,omitempty" tf:"vpn_tunnel,omitempty"`

	// Reference to a VPNTunnel to populate vpnTunnel.
	// +kubebuilder:validation:Optional
	VPNTunnelRef *v1.Reference `json:"vpnTunnelRef,omitempty" tf:"-"`

	// Selector for a VPNTunnel to populate vpnTunnel.
	// +kubebuilder:validation:Optional
	VPNTunnelSelector *v1.Selector `json:"vpnTunnelSelector,omitempty" tf:"-"`
}

type RouterInterfaceObservation struct {

	// an identifier for the resource with format {{region}}/{{router}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`

	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the project in which this interface's router belongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface *string `json:"redundantInterface,omitempty" tf:"redundant_interface,omitempty"`

	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	VPNTunnel *string `json:"vpnTunnel,omitempty" tf:"vpn_tunnel,omitempty"`
}

type RouterInterfaceParameters struct {

	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	// +kubebuilder:validation:Optional
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty" tf:"interconnect_attachment,omitempty"`

	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The ID of the project in which this interface's router belongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	RedundantInterface *string `json:"redundantInterface,omitempty" tf:"redundant_interface,omitempty"`

	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	// +crossplane:generate:reference:type=Router
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Reference to a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterRef *v1.Reference `json:"routerRef,omitempty" tf:"-"`

	// Selector for a Router to populate router.
	// +kubebuilder:validation:Optional
	RouterSelector *v1.Selector `json:"routerSelector,omitempty" tf:"-"`

	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	// +crossplane:generate:reference:type=VPNTunnel
	// +kubebuilder:validation:Optional
	VPNTunnel *string `json:"vpnTunnel,omitempty" tf:"vpn_tunnel,omitempty"`

	// Reference to a VPNTunnel to populate vpnTunnel.
	// +kubebuilder:validation:Optional
	VPNTunnelRef *v1.Reference `json:"vpnTunnelRef,omitempty" tf:"-"`

	// Selector for a VPNTunnel to populate vpnTunnel.
	// +kubebuilder:validation:Optional
	VPNTunnelSelector *v1.Selector `json:"vpnTunnelSelector,omitempty" tf:"-"`
}

// RouterInterfaceSpec defines the desired state of RouterInterface
type RouterInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterInterfaceParameters `json:"forProvider"`
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
	InitProvider RouterInterfaceInitParameters `json:"initProvider,omitempty"`
}

// RouterInterfaceStatus defines the observed state of RouterInterface.
type RouterInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouterInterface is the Schema for the RouterInterfaces API. Manages a Cloud Router interface.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RouterInterfaceSpec   `json:"spec"`
	Status RouterInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterfaceList contains a list of RouterInterfaces
type RouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterInterface `json:"items"`
}

// Repository type metadata.
var (
	RouterInterface_Kind             = "RouterInterface"
	RouterInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterInterface_Kind}.String()
	RouterInterface_KindAPIVersion   = RouterInterface_Kind + "." + CRDGroupVersion.String()
	RouterInterface_GroupVersionKind = CRDGroupVersion.WithKind(RouterInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterInterface{}, &RouterInterfaceList{})
}
