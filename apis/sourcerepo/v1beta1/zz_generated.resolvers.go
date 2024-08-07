// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Repository) ResolveReferences( // ResolveReferences of this Repository.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PubsubConfigs); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmail")
		}
		mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PubsubConfigs[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.PubsubConfigs); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PubsubConfigs[i3].Topic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.PubsubConfigs[i3].TopicRef,
				Selector:     mg.Spec.ForProvider.PubsubConfigs[i3].TopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PubsubConfigs[i3].Topic")
		}
		mg.Spec.ForProvider.PubsubConfigs[i3].Topic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PubsubConfigs[i3].TopicRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.PubsubConfigs); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmail),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmailRef,
				Selector:     mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmailSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmail")
		}
		mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmail = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PubsubConfigs[i3].ServiceAccountEmailRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.PubsubConfigs); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("pubsub.gcp.upbound.io", "v1beta2", "Topic", "TopicList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PubsubConfigs[i3].Topic),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.PubsubConfigs[i3].TopicRef,
				Selector:     mg.Spec.InitProvider.PubsubConfigs[i3].TopicSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PubsubConfigs[i3].Topic")
		}
		mg.Spec.InitProvider.PubsubConfigs[i3].Topic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PubsubConfigs[i3].TopicRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RepositoryIAMMember.
func (mg *RepositoryIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sourcerepo.gcp.upbound.io", "v1beta1", "Repository", "RepositoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RepositoryRef,
			Selector:     mg.Spec.ForProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Repository")
	}
	mg.Spec.ForProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sourcerepo.gcp.upbound.io", "v1beta1", "Repository", "RepositoryList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Repository),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RepositoryRef,
			Selector:     mg.Spec.InitProvider.RepositorySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Repository")
	}
	mg.Spec.InitProvider.Repository = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryRef = rsp.ResolvedReference

	return nil
}
