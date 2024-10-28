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
	common "github.com/upbound/provider-gcp/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Cluster.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
			Extract:      common.SelfLinkExtractor(),
			Reference:    mg.Spec.ForProvider.NetworkRef,
			Selector:     mg.Spec.ForProvider.NetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.NodeConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig.ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.NodeConfig.ServiceAccountRef,
				Selector:     mg.Spec.ForProvider.NodeConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig.ServiceAccount")
		}
		mg.Spec.ForProvider.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.PrivateClusterConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetwork),
				Extract:      common.SelfLinkExtractor(),
				Reference:    mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetworkRef,
				Selector:     mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetwork")
		}
		mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetwork = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateClusterConfig.PrivateEndpointSubnetworkRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subnetwork),
			Extract:      common.SelfLinkExtractor(),
			Reference:    mg.Spec.ForProvider.SubnetworkRef,
			Selector:     mg.Spec.ForProvider.SubnetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnetwork")
	}
	mg.Spec.ForProvider.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetworkRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Network),
			Extract:      common.SelfLinkExtractor(),
			Reference:    mg.Spec.InitProvider.NetworkRef,
			Selector:     mg.Spec.InitProvider.NetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Network")
	}
	mg.Spec.InitProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.NodeConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NodeConfig.ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.NodeConfig.ServiceAccountRef,
				Selector:     mg.Spec.InitProvider.NodeConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NodeConfig.ServiceAccount")
		}
		mg.Spec.InitProvider.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.PrivateClusterConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetwork),
				Extract:      common.SelfLinkExtractor(),
				Reference:    mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetworkRef,
				Selector:     mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetwork")
		}
		mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetwork = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PrivateClusterConfig.PrivateEndpointSubnetworkRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta2", "Subnetwork", "SubnetworkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Subnetwork),
			Extract:      common.SelfLinkExtractor(),
			Reference:    mg.Spec.InitProvider.SubnetworkRef,
			Selector:     mg.Spec.InitProvider.SubnetworkSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subnetwork")
	}
	mg.Spec.InitProvider.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodePool.
func (mg *NodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("container.gcp.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ClusterRef,
			Selector:     mg.Spec.ForProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.NodeConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig.ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.NodeConfig.ServiceAccountRef,
				Selector:     mg.Spec.ForProvider.NodeConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig.ServiceAccount")
		}
		mg.Spec.ForProvider.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.NodeConfig != nil {
		{
			m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NodeConfig.ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.NodeConfig.ServiceAccountRef,
				Selector:     mg.Spec.InitProvider.NodeConfig.ServiceAccountSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NodeConfig.ServiceAccount")
		}
		mg.Spec.InitProvider.NodeConfig.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.NodeConfig.ServiceAccountRef = rsp.ResolvedReference

	}

	return nil
}
