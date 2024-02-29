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
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ConsentStore.
func (mg *ConsentStore) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Dataset),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DatasetRef,
		Selector:     mg.Spec.ForProvider.DatasetSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Dataset")
	}
	mg.Spec.ForProvider.Dataset = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Dataset),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.DatasetRef,
		Selector:     mg.Spec.InitProvider.DatasetSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Dataset")
	}
	mg.Spec.InitProvider.Dataset = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatasetRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatasetIAMMember.
func (mg *DatasetIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatasetIDRef,
		Selector:     mg.Spec.ForProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatasetID")
	}
	mg.Spec.ForProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatasetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatasetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DatasetIDRef,
		Selector:     mg.Spec.InitProvider.DatasetIDSelector,
		To: reference.To{
			List:    &DatasetList{},
			Managed: &Dataset{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatasetID")
	}
	mg.Spec.InitProvider.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatasetIDRef = rsp.ResolvedReference

	return nil
}
