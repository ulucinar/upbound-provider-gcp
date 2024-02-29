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
	v1beta1 "github.com/upbound/provider-gcp/apis/servicenetworking/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizedNetwork),
		Extract:      resource.ExtractParamPath("network", false),
		Reference:    mg.Spec.ForProvider.AuthorizedNetworkRef,
		Selector:     mg.Spec.ForProvider.AuthorizedNetworkSelector,
		To: reference.To{
			List:    &v1beta1.ConnectionList{},
			Managed: &v1beta1.Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizedNetwork")
	}
	mg.Spec.ForProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizedNetworkRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizedNetwork),
		Extract:      resource.ExtractParamPath("network", false),
		Reference:    mg.Spec.InitProvider.AuthorizedNetworkRef,
		Selector:     mg.Spec.InitProvider.AuthorizedNetworkSelector,
		To: reference.To{
			List:    &v1beta1.ConnectionList{},
			Managed: &v1beta1.Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizedNetwork")
	}
	mg.Spec.InitProvider.AuthorizedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizedNetworkRef = rsp.ResolvedReference

	return nil
}
