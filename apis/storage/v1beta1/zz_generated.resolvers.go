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
	v1beta1 "github.com/upbound/provider-gcp/apis/pubsub/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BucketACL.
func (mg *BucketACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketAccessControl.
func (mg *BucketAccessControl) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketIAMMember.
func (mg *BucketIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketObject.
func (mg *BucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultObjectACL.
func (mg *DefaultObjectACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultObjectAccessControl.
func (mg *DefaultObjectAccessControl) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Notification.
func (mg *Notification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TopicRef,
		Selector:     mg.Spec.ForProvider.TopicSelector,
		To: reference.To{
			List:    &v1beta1.TopicList{},
			Managed: &v1beta1.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.TopicRef,
		Selector:     mg.Spec.InitProvider.TopicSelector,
		To: reference.To{
			List:    &v1beta1.TopicList{},
			Managed: &v1beta1.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ObjectACL.
func (mg *ObjectACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Object),
		Extract:      resource.ExtractParamPath("output_name", true),
		Reference:    mg.Spec.ForProvider.ObjectRef,
		Selector:     mg.Spec.ForProvider.ObjectSelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Object")
	}
	mg.Spec.ForProvider.Object = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ObjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Object),
		Extract:      resource.ExtractParamPath("output_name", true),
		Reference:    mg.Spec.InitProvider.ObjectRef,
		Selector:     mg.Spec.InitProvider.ObjectSelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Object")
	}
	mg.Spec.InitProvider.Object = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ObjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ObjectAccessControl.
func (mg *ObjectAccessControl) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Object),
		Extract:      resource.ExtractParamPath("output_name", true),
		Reference:    mg.Spec.ForProvider.ObjectRef,
		Selector:     mg.Spec.ForProvider.ObjectSelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Object")
	}
	mg.Spec.ForProvider.Object = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ObjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Object),
		Extract:      resource.ExtractParamPath("output_name", true),
		Reference:    mg.Spec.InitProvider.ObjectRef,
		Selector:     mg.Spec.InitProvider.ObjectSelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Object")
	}
	mg.Spec.InitProvider.Object = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ObjectRef = rsp.ResolvedReference

	return nil
}
