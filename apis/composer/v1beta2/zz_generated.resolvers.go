// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Environment) ResolveReferences( // ResolveReferences of this Environment.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.Config != nil {
		if mg.Spec.ForProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Config.NodeConfig.Network),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.Config.NodeConfig.NetworkRef,
					Selector:     mg.Spec.ForProvider.Config.NodeConfig.NetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Config.NodeConfig.Network")
			}
			mg.Spec.ForProvider.Config.NodeConfig.Network = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Config.NodeConfig.NetworkRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.Config != nil {
		if mg.Spec.ForProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Config.NodeConfig.ServiceAccount),
					Extract:      resource.ExtractParamPath("name", true),
					Reference:    mg.Spec.ForProvider.Config.NodeConfig.ServiceAccountRef,
					Selector:     mg.Spec.ForProvider.Config.NodeConfig.ServiceAccountSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Config.NodeConfig.ServiceAccount")
			}
			mg.Spec.ForProvider.Config.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Config.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.Config != nil {
		if mg.Spec.ForProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Config.NodeConfig.Subnetwork),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.Config.NodeConfig.SubnetworkRef,
					Selector:     mg.Spec.ForProvider.Config.NodeConfig.SubnetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Config.NodeConfig.Subnetwork")
			}
			mg.Spec.ForProvider.Config.NodeConfig.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Config.NodeConfig.SubnetworkRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ProjectRef,
			Selector:     mg.Spec.ForProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.Config != nil {
		if mg.Spec.InitProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Config.NodeConfig.Network),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.Config.NodeConfig.NetworkRef,
					Selector:     mg.Spec.InitProvider.Config.NodeConfig.NetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Config.NodeConfig.Network")
			}
			mg.Spec.InitProvider.Config.NodeConfig.Network = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Config.NodeConfig.NetworkRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.Config != nil {
		if mg.Spec.InitProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Config.NodeConfig.ServiceAccount),
					Extract:      resource.ExtractParamPath("name", true),
					Reference:    mg.Spec.InitProvider.Config.NodeConfig.ServiceAccountRef,
					Selector:     mg.Spec.InitProvider.Config.NodeConfig.ServiceAccountSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Config.NodeConfig.ServiceAccount")
			}
			mg.Spec.InitProvider.Config.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Config.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.Config != nil {
		if mg.Spec.InitProvider.Config.NodeConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Config.NodeConfig.Subnetwork),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.Config.NodeConfig.SubnetworkRef,
					Selector:     mg.Spec.InitProvider.Config.NodeConfig.SubnetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Config.NodeConfig.Subnetwork")
			}
			mg.Spec.InitProvider.Config.NodeConfig.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Config.NodeConfig.SubnetworkRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ProjectRef,
			Selector:     mg.Spec.InitProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}