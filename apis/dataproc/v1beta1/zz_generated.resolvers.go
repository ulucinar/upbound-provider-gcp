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
	v1beta1 "github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1"
	v1beta12 "github.com/upbound/provider-gcp/apis/compute/v1beta1"
	v1beta11 "github.com/upbound/provider-gcp/apis/kms/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ClusterConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountRef,
				Selector:     mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountSelector,
				To: reference.To{
					List:    &v1beta1.ServiceAccountList{},
					Managed: &v1beta1.ServiceAccount{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount")
			}
			mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ClusterConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount),
				Extract:      resource.ExtractParamPath("email", true),
				Reference:    mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountRef,
				Selector:     mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountSelector,
				To: reference.To{
					List:    &v1beta1.ServiceAccountList{},
					Managed: &v1beta1.ServiceAccount{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount")
			}
			mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ClusterConfig[i3].GceClusterConfig[i4].ServiceAccountRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Placement); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Placement[i3].ClusterName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.Placement[i3].ClusterNameRef,
			Selector:     mg.Spec.ForProvider.Placement[i3].ClusterNameSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Placement[i3].ClusterName")
		}
		mg.Spec.ForProvider.Placement[i3].ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Placement[i3].ClusterNameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Region),
		Extract:      resource.ExtractParamPath("region", false),
		Reference:    mg.Spec.ForProvider.RegionRef,
		Selector:     mg.Spec.ForProvider.RegionSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Region")
	}
	mg.Spec.ForProvider.Region = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegionRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Placement); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Placement[i3].ClusterName),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.InitProvider.Placement[i3].ClusterNameRef,
			Selector:     mg.Spec.InitProvider.Placement[i3].ClusterNameSelector,
			To: reference.To{
				List:    &ClusterList{},
				Managed: &Cluster{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Placement[i3].ClusterName")
		}
		mg.Spec.InitProvider.Placement[i3].ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Placement[i3].ClusterNameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Region),
		Extract:      resource.ExtractParamPath("region", false),
		Reference:    mg.Spec.InitProvider.RegionRef,
		Selector:     mg.Spec.InitProvider.RegionSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Region")
	}
	mg.Spec.InitProvider.Region = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RegionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MetastoreService.
func (mg *MetastoreService) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionConfig[i3].KMSKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EncryptionConfig[i3].KMSKeyRef,
			Selector:     mg.Spec.ForProvider.EncryptionConfig[i3].KMSKeySelector,
			To: reference.To{
				List:    &v1beta11.CryptoKeyList{},
				Managed: &v1beta11.CryptoKey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionConfig[i3].KMSKey")
		}
		mg.Spec.ForProvider.EncryptionConfig[i3].KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionConfig[i3].KMSKeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkConfig[i3].Consumers); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].Subnetwork),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].SubnetworkRef,
				Selector:     mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].SubnetworkSelector,
				To: reference.To{
					List:    &v1beta12.SubnetworkList{},
					Managed: &v1beta12.Subnetwork{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].Subnetwork")
			}
			mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkConfig[i3].Consumers[i4].SubnetworkRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EncryptionConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionConfig[i3].KMSKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EncryptionConfig[i3].KMSKeyRef,
			Selector:     mg.Spec.InitProvider.EncryptionConfig[i3].KMSKeySelector,
			To: reference.To{
				List:    &v1beta11.CryptoKeyList{},
				Managed: &v1beta11.CryptoKey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionConfig[i3].KMSKey")
		}
		mg.Spec.InitProvider.EncryptionConfig[i3].KMSKey = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EncryptionConfig[i3].KMSKeyRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.NetworkConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.NetworkConfig[i3].Consumers); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].Subnetwork),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].SubnetworkRef,
				Selector:     mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].SubnetworkSelector,
				To: reference.To{
					List:    &v1beta12.SubnetworkList{},
					Managed: &v1beta12.Subnetwork{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].Subnetwork")
			}
			mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.NetworkConfig[i3].Consumers[i4].SubnetworkRef = rsp.ResolvedReference

		}
	}

	return nil
}
