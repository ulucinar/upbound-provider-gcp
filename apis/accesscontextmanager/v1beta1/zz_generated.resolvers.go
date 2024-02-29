//go:build (accesscontextmanager || all) && !ignore_autogenerated

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

// ResolveReferences of this AccessLevelCondition.
func (mg *AccessLevelCondition) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessLevel),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.AccessLevelRef,
		Selector:     mg.Spec.ForProvider.AccessLevelSelector,
		To: reference.To{
			List:    &AccessLevelList{},
			Managed: &AccessLevel{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessLevel")
	}
	mg.Spec.ForProvider.AccessLevel = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessLevelRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccessLevel),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.AccessLevelRef,
		Selector:     mg.Spec.InitProvider.AccessLevelSelector,
		To: reference.To{
			List:    &AccessLevelList{},
			Managed: &AccessLevel{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccessLevel")
	}
	mg.Spec.InitProvider.AccessLevel = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccessLevelRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServicePerimeter.
func (mg *ServicePerimeter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Spec); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Spec[i3].AccessLevels),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Spec[i3].AccessLevelsRefs,
			Selector:      mg.Spec.ForProvider.Spec[i3].AccessLevelsSelector,
			To: reference.To{
				List:    &AccessLevelList{},
				Managed: &AccessLevel{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Spec[i3].AccessLevels")
		}
		mg.Spec.ForProvider.Spec[i3].AccessLevels = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Spec[i3].AccessLevelsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Status); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Status[i3].AccessLevels),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Status[i3].AccessLevelsRefs,
			Selector:      mg.Spec.ForProvider.Status[i3].AccessLevelsSelector,
			To: reference.To{
				List:    &AccessLevelList{},
				Managed: &AccessLevel{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Status[i3].AccessLevels")
		}
		mg.Spec.ForProvider.Status[i3].AccessLevels = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Status[i3].AccessLevelsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Status); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Status[i3].IngressPolicies); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelRef,
						Selector:     mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelSelector,
						To: reference.To{
							List:    &AccessLevelList{},
							Managed: &AccessLevel{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel")
					}
					mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Spec); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Spec[i3].AccessLevels),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.Spec[i3].AccessLevelsRefs,
			Selector:      mg.Spec.InitProvider.Spec[i3].AccessLevelsSelector,
			To: reference.To{
				List:    &AccessLevelList{},
				Managed: &AccessLevel{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Spec[i3].AccessLevels")
		}
		mg.Spec.InitProvider.Spec[i3].AccessLevels = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Spec[i3].AccessLevelsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Status); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Status[i3].AccessLevels),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.Status[i3].AccessLevelsRefs,
			Selector:      mg.Spec.InitProvider.Status[i3].AccessLevelsSelector,
			To: reference.To{
				List:    &AccessLevelList{},
				Managed: &AccessLevel{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Status[i3].AccessLevels")
		}
		mg.Spec.InitProvider.Status[i3].AccessLevels = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.Status[i3].AccessLevelsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Status); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Status[i3].IngressPolicies); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelRef,
						Selector:     mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelSelector,
						To: reference.To{
							List:    &AccessLevelList{},
							Managed: &AccessLevel{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel")
					}
					mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevel = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Status[i3].IngressPolicies[i4].IngressFrom[i5].Sources[i6].AccessLevelRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}

// ResolveReferences of this ServicePerimeterResource.
func (mg *ServicePerimeterResource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PerimeterName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.PerimeterNameRef,
		Selector:     mg.Spec.ForProvider.PerimeterNameSelector,
		To: reference.To{
			List:    &ServicePerimeterList{},
			Managed: &ServicePerimeter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PerimeterName")
	}
	mg.Spec.ForProvider.PerimeterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PerimeterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PerimeterName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.InitProvider.PerimeterNameRef,
		Selector:     mg.Spec.InitProvider.PerimeterNameSelector,
		To: reference.To{
			List:    &ServicePerimeterList{},
			Managed: &ServicePerimeter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PerimeterName")
	}
	mg.Spec.InitProvider.PerimeterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PerimeterNameRef = rsp.ResolvedReference

	return nil
}
