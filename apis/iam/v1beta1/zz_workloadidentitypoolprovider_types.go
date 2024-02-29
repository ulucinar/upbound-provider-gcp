//go:build (iam || all) && !ignore_autogenerated

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

type AwsInitParameters struct {

	// The AWS account ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`
}

type AwsObservation struct {

	// The AWS account ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`
}

type AwsParameters struct {

	// The AWS account ID.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`
}

type OidcInitParameters struct {

	// Acceptable values for the aud field (audience) in the OIDC token. Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	AllowedAudiences []*string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`

	// The OIDC issuer URL.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// OIDC JWKs in JSON String format. For details on definition of a
	// JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we
	// use the jwks_uri from the discovery document fetched from the
	// .well-known path for the issuer_uri. Currently, RSA and EC asymmetric
	// keys are supported. The JWK must use following format and include only
	// the following fields:
	JwksJSON *string `json:"jwksJson,omitempty" tf:"jwks_json,omitempty"`
}

type OidcObservation struct {

	// Acceptable values for the aud field (audience) in the OIDC token. Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	AllowedAudiences []*string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`

	// The OIDC issuer URL.
	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	// OIDC JWKs in JSON String format. For details on definition of a
	// JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we
	// use the jwks_uri from the discovery document fetched from the
	// .well-known path for the issuer_uri. Currently, RSA and EC asymmetric
	// keys are supported. The JWK must use following format and include only
	// the following fields:
	JwksJSON *string `json:"jwksJson,omitempty" tf:"jwks_json,omitempty"`
}

type OidcParameters struct {

	// Acceptable values for the aud field (audience) in the OIDC token. Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	// +kubebuilder:validation:Optional
	AllowedAudiences []*string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`

	// The OIDC issuer URL.
	// +kubebuilder:validation:Optional
	IssuerURI *string `json:"issuerUri" tf:"issuer_uri,omitempty"`

	// OIDC JWKs in JSON String format. For details on definition of a
	// JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we
	// use the jwks_uri from the discovery document fetched from the
	// .well-known path for the issuer_uri. Currently, RSA and EC asymmetric
	// keys are supported. The JWK must use following format and include only
	// the following fields:
	// +kubebuilder:validation:Optional
	JwksJSON *string `json:"jwksJson,omitempty" tf:"jwks_json,omitempty"`
}

type WorkloadIdentityPoolProviderInitParameters struct {

	// A Common Expression Language expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	AttributeCondition *string `json:"attributeCondition,omitempty" tf:"attribute_condition,omitempty"`

	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as subject and segment.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// +mapType=granular
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws []AwsInitParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// A description for the provider. Cannot exceed 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc []OidcInitParameters `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type WorkloadIdentityPoolProviderObservation struct {

	// A Common Expression Language expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	AttributeCondition *string `json:"attributeCondition,omitempty" tf:"attribute_condition,omitempty"`

	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as subject and segment.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// +mapType=granular
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws []AwsObservation `json:"aws,omitempty" tf:"aws,omitempty"`

	// A description for the provider. Cannot exceed 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the provider as
	// projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc []OidcObservation `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The state of the provider.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// gcp- is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolID *string `json:"workloadIdentityPoolId,omitempty" tf:"workload_identity_pool_id,omitempty"`
}

type WorkloadIdentityPoolProviderParameters struct {

	// A Common Expression Language expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// +kubebuilder:validation:Optional
	AttributeCondition *string `json:"attributeCondition,omitempty" tf:"attribute_condition,omitempty"`

	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as subject and segment.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Aws []AwsParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// A description for the provider. Cannot exceed 256 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A display name for the provider. Cannot exceed 32 characters.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Oidc []OidcParameters `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// gcp- is reserved for use by Google, and may not be specified.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/iam/v1beta1.WorkloadIdentityPool
	// +kubebuilder:validation:Optional
	WorkloadIdentityPoolID *string `json:"workloadIdentityPoolId,omitempty" tf:"workload_identity_pool_id,omitempty"`

	// Reference to a WorkloadIdentityPool in iam to populate workloadIdentityPoolId.
	// +kubebuilder:validation:Optional
	WorkloadIdentityPoolIDRef *v1.Reference `json:"workloadIdentityPoolIdRef,omitempty" tf:"-"`

	// Selector for a WorkloadIdentityPool in iam to populate workloadIdentityPoolId.
	// +kubebuilder:validation:Optional
	WorkloadIdentityPoolIDSelector *v1.Selector `json:"workloadIdentityPoolIdSelector,omitempty" tf:"-"`
}

// WorkloadIdentityPoolProviderSpec defines the desired state of WorkloadIdentityPoolProvider
type WorkloadIdentityPoolProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkloadIdentityPoolProviderParameters `json:"forProvider"`
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
	InitProvider WorkloadIdentityPoolProviderInitParameters `json:"initProvider,omitempty"`
}

// WorkloadIdentityPoolProviderStatus defines the observed state of WorkloadIdentityPoolProvider.
type WorkloadIdentityPoolProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkloadIdentityPoolProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WorkloadIdentityPoolProvider is the Schema for the WorkloadIdentityPoolProviders API. A configuration for an external identity provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type WorkloadIdentityPoolProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkloadIdentityPoolProviderSpec   `json:"spec"`
	Status            WorkloadIdentityPoolProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkloadIdentityPoolProviderList contains a list of WorkloadIdentityPoolProviders
type WorkloadIdentityPoolProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadIdentityPoolProvider `json:"items"`
}

// Repository type metadata.
var (
	WorkloadIdentityPoolProvider_Kind             = "WorkloadIdentityPoolProvider"
	WorkloadIdentityPoolProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkloadIdentityPoolProvider_Kind}.String()
	WorkloadIdentityPoolProvider_KindAPIVersion   = WorkloadIdentityPoolProvider_Kind + "." + CRDGroupVersion.String()
	WorkloadIdentityPoolProvider_GroupVersionKind = CRDGroupVersion.WithKind(WorkloadIdentityPoolProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkloadIdentityPoolProvider{}, &WorkloadIdentityPoolProviderList{})
}
