//go:build (eventarc || all) && !ignore_autogenerated

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

type ChannelInitParameters struct {

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: projects/{project}/locations/{location}/providers/{provider_id}.
	ThirdPartyProvider *string `json:"thirdPartyProvider,omitempty" tf:"third_party_provider,omitempty"`
}

type ChannelObservation struct {

	// Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
	ActivationToken *string `json:"activationToken,omitempty" tf:"activation_token,omitempty"`

	// Output only. The creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/channels/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: projects/{project}/topics/{topic_id}.
	PubsubTopic *string `json:"pubsubTopic,omitempty" tf:"pubsub_topic,omitempty"`

	// Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: projects/{project}/locations/{location}/providers/{provider_id}.
	ThirdPartyProvider *string `json:"thirdPartyProvider,omitempty" tf:"third_party_provider,omitempty"`

	// Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ChannelParameters struct {

	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern projects/*/locations/*/keyRings/*/cryptoKeys/*.
	// +kubebuilder:validation:Optional
	CryptoKeyName *string `json:"cryptoKeyName,omitempty" tf:"crypto_key_name,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: projects/{project}/locations/{location}/providers/{provider_id}.
	// +kubebuilder:validation:Optional
	ThirdPartyProvider *string `json:"thirdPartyProvider,omitempty" tf:"third_party_provider,omitempty"`
}

// ChannelSpec defines the desired state of Channel
type ChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelParameters `json:"forProvider"`
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
	InitProvider ChannelInitParameters `json:"initProvider,omitempty"`
}

// ChannelStatus defines the observed state of Channel.
type ChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Channel is the Schema for the Channels API. The Eventarc Channel resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Channel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelSpec   `json:"spec"`
	Status            ChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelList contains a list of Channels
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Channel `json:"items"`
}

// Repository type metadata.
var (
	Channel_Kind             = "Channel"
	Channel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Channel_Kind}.String()
	Channel_KindAPIVersion   = Channel_Kind + "." + CRDGroupVersion.String()
	Channel_GroupVersionKind = CRDGroupVersion.WithKind(Channel_Kind)
)

func init() {
	SchemeBuilder.Register(&Channel{}, &ChannelList{})
}
