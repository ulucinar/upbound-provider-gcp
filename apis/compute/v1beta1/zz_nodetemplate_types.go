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

type NodeTemplateInitParameters struct {

	// CPU overcommit.
	// Default value is NONE.
	// Possible values are: ENABLED, NONE.
	CPUOvercommitType *string `json:"cpuOvercommitType,omitempty" tf:"cpu_overcommit_type,omitempty"`

	// An optional textual description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	// +mapType=granular
	NodeAffinityLabels map[string]*string `json:"nodeAffinityLabels,omitempty" tf:"node_affinity_labels,omitempty"`

	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility []NodeTypeFlexibilityInitParameters `json:"nodeTypeFlexibility,omitempty" tf:"node_type_flexibility,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding []ServerBindingInitParameters `json:"serverBinding,omitempty" tf:"server_binding,omitempty"`
}

type NodeTemplateObservation struct {

	// CPU overcommit.
	// Default value is NONE.
	// Possible values are: ENABLED, NONE.
	CPUOvercommitType *string `json:"cpuOvercommitType,omitempty" tf:"cpu_overcommit_type,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional textual description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	// +mapType=granular
	NodeAffinityLabels map[string]*string `json:"nodeAffinityLabels,omitempty" tf:"node_affinity_labels,omitempty"`

	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility []NodeTypeFlexibilityObservation `json:"nodeTypeFlexibility,omitempty" tf:"node_type_flexibility,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding []ServerBindingObservation `json:"serverBinding,omitempty" tf:"server_binding,omitempty"`
}

type NodeTemplateParameters struct {

	// CPU overcommit.
	// Default value is NONE.
	// Possible values are: ENABLED, NONE.
	// +kubebuilder:validation:Optional
	CPUOvercommitType *string `json:"cpuOvercommitType,omitempty" tf:"cpu_overcommit_type,omitempty"`

	// An optional textual description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	NodeAffinityLabels map[string]*string `json:"nodeAffinityLabels,omitempty" tf:"node_affinity_labels,omitempty"`

	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NodeTypeFlexibility []NodeTypeFlexibilityParameters `json:"nodeTypeFlexibility,omitempty" tf:"node_type_flexibility,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServerBinding []ServerBindingParameters `json:"serverBinding,omitempty" tf:"server_binding,omitempty"`
}

type NodeTypeFlexibilityInitParameters struct {

	// Number of virtual CPUs to use.
	Cpus *string `json:"cpus,omitempty" tf:"cpus,omitempty"`

	// Physical memory available to the node, defined in MB.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type NodeTypeFlexibilityObservation struct {

	// Number of virtual CPUs to use.
	Cpus *string `json:"cpus,omitempty" tf:"cpus,omitempty"`

	// (Output)
	// Use local SSD
	LocalSsd *string `json:"localSsd,omitempty" tf:"local_ssd,omitempty"`

	// Physical memory available to the node, defined in MB.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type NodeTypeFlexibilityParameters struct {

	// Number of virtual CPUs to use.
	// +kubebuilder:validation:Optional
	Cpus *string `json:"cpus,omitempty" tf:"cpus,omitempty"`

	// Physical memory available to the node, defined in MB.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type ServerBindingInitParameters struct {

	// Type of server binding policy. If RESTART_NODE_ON_ANY_SERVER,
	// nodes using this template will restart on any physical server
	// following a maintenance event.
	// If RESTART_NODE_ON_MINIMAL_SERVER, nodes using this template
	// will restart on the same physical server following a maintenance
	// event, instead of being live migrated to or restarted on a new
	// physical server. This option may be useful if you are using
	// software licenses tied to the underlying server characteristics
	// such as physical sockets or cores, to avoid the need for
	// additional licenses when maintenance occurs. However, VMs on such
	// nodes will experience outages while maintenance is applied.
	// Possible values are: RESTART_NODE_ON_ANY_SERVER, RESTART_NODE_ON_MINIMAL_SERVERS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServerBindingObservation struct {

	// Type of server binding policy. If RESTART_NODE_ON_ANY_SERVER,
	// nodes using this template will restart on any physical server
	// following a maintenance event.
	// If RESTART_NODE_ON_MINIMAL_SERVER, nodes using this template
	// will restart on the same physical server following a maintenance
	// event, instead of being live migrated to or restarted on a new
	// physical server. This option may be useful if you are using
	// software licenses tied to the underlying server characteristics
	// such as physical sockets or cores, to avoid the need for
	// additional licenses when maintenance occurs. However, VMs on such
	// nodes will experience outages while maintenance is applied.
	// Possible values are: RESTART_NODE_ON_ANY_SERVER, RESTART_NODE_ON_MINIMAL_SERVERS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServerBindingParameters struct {

	// Type of server binding policy. If RESTART_NODE_ON_ANY_SERVER,
	// nodes using this template will restart on any physical server
	// following a maintenance event.
	// If RESTART_NODE_ON_MINIMAL_SERVER, nodes using this template
	// will restart on the same physical server following a maintenance
	// event, instead of being live migrated to or restarted on a new
	// physical server. This option may be useful if you are using
	// software licenses tied to the underlying server characteristics
	// such as physical sockets or cores, to avoid the need for
	// additional licenses when maintenance occurs. However, VMs on such
	// nodes will experience outages while maintenance is applied.
	// Possible values are: RESTART_NODE_ON_ANY_SERVER, RESTART_NODE_ON_MINIMAL_SERVERS.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// NodeTemplateSpec defines the desired state of NodeTemplate
type NodeTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeTemplateParameters `json:"forProvider"`
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
	InitProvider NodeTemplateInitParameters `json:"initProvider,omitempty"`
}

// NodeTemplateStatus defines the observed state of NodeTemplate.
type NodeTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodeTemplate is the Schema for the NodeTemplates API. Represents a NodeTemplate resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeTemplateSpec   `json:"spec"`
	Status            NodeTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeTemplateList contains a list of NodeTemplates
type NodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeTemplate `json:"items"`
}

// Repository type metadata.
var (
	NodeTemplate_Kind             = "NodeTemplate"
	NodeTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeTemplate_Kind}.String()
	NodeTemplate_KindAPIVersion   = NodeTemplate_Kind + "." + CRDGroupVersion.String()
	NodeTemplate_GroupVersionKind = CRDGroupVersion.WithKind(NodeTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeTemplate{}, &NodeTemplateList{})
}
