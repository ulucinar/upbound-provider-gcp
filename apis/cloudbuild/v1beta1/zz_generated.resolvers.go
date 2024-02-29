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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"

	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1"
	v1beta13 "github.com/upbound/provider-gcp/apis/compute/v1beta1"
	v1beta1 "github.com/upbound/provider-gcp/apis/pubsub/v1beta1"
	v1beta12 "github.com/upbound/provider-gcp/apis/secretmanager/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PubsubConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PubsubConfig[i3].Topic),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PubsubConfig[i3].TopicRef,
			Selector:     mg.Spec.ForProvider.PubsubConfig[i3].TopicSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PubsubConfig[i3].Topic")
		}
		mg.Spec.ForProvider.PubsubConfig[i3].Topic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PubsubConfig[i3].TopicRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccount),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &v1beta11.ServiceAccountList{},
			Managed: &v1beta11.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccount")
	}
	mg.Spec.ForProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.WebhookConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WebhookConfig[i3].Secret),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WebhookConfig[i3].SecretRef,
			Selector:     mg.Spec.ForProvider.WebhookConfig[i3].SecretSelector,
			To: reference.To{
				List:    &v1beta12.SecretVersionList{},
				Managed: &v1beta12.SecretVersion{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.WebhookConfig[i3].Secret")
		}
		mg.Spec.ForProvider.WebhookConfig[i3].Secret = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.WebhookConfig[i3].SecretRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.PubsubConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PubsubConfig[i3].Topic),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PubsubConfig[i3].TopicRef,
			Selector:     mg.Spec.InitProvider.PubsubConfig[i3].TopicSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PubsubConfig[i3].Topic")
		}
		mg.Spec.InitProvider.PubsubConfig[i3].Topic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PubsubConfig[i3].TopicRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccount),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &v1beta11.ServiceAccountList{},
			Managed: &v1beta11.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccount")
	}
	mg.Spec.InitProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.WebhookConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WebhookConfig[i3].Secret),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WebhookConfig[i3].SecretRef,
			Selector:     mg.Spec.InitProvider.WebhookConfig[i3].SecretSelector,
			To: reference.To{
				List:    &v1beta12.SecretVersionList{},
				Managed: &v1beta12.SecretVersion{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.WebhookConfig[i3].Secret")
		}
		mg.Spec.InitProvider.WebhookConfig[i3].Secret = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.WebhookConfig[i3].SecretRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this WorkerPool.
func (mg *WorkerPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetwork),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetworkRef,
			Selector:     mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetworkSelector,
			To: reference.To{
				List:    &v1beta13.NetworkList{},
				Managed: &v1beta13.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetwork")
		}
		mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetwork = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkConfig[i3].PeeredNetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetwork),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetworkRef,
			Selector:     mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetworkSelector,
			To: reference.To{
				List:    &v1beta13.NetworkList{},
				Managed: &v1beta13.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetwork")
		}
		mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetwork = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NetworkConfig[i3].PeeredNetworkRef = rsp.ResolvedReference

	}

	return nil
}
