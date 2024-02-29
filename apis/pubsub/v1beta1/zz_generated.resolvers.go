//go:build (pubsub || all) && !ignore_autogenerated

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
	v1beta1 "github.com/upbound/provider-gcp/apis/kms/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LiteSubscription.
func (mg *LiteSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicRef,
		Selector:     mg.Spec.ForProvider.TopicSelector,
		To: reference.To{
			List:    &LiteTopicList{},
			Managed: &LiteTopic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TopicRef,
		Selector:     mg.Spec.InitProvider.TopicSelector,
		To: reference.To{
			List:    &LiteTopicList{},
			Managed: &LiteTopic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LiteTopic.
func (mg *LiteTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ReservationConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservation),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservationRef,
			Selector:     mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservationSelector,
			To: reference.To{
				List:    &LiteReservationList{},
				Managed: &LiteReservation{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservation")
		}
		mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservation = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ReservationConfig[i3].ThroughputReservationRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ReservationConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservation),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservationRef,
			Selector:     mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservationSelector,
			To: reference.To{
				List:    &LiteReservationList{},
				Managed: &LiteReservation{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservation")
		}
		mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservation = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ReservationConfig[i3].ThroughputReservationRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Subscription.
func (mg *Subscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DeadLetterPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopic),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopicRef,
			Selector:     mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopicSelector,
			To: reference.To{
				List:    &TopicList{},
				Managed: &Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopic")
		}
		mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DeadLetterPolicy[i3].DeadLetterTopicRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicRef,
		Selector:     mg.Spec.ForProvider.TopicSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DeadLetterPolicy); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopic),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopicRef,
			Selector:     mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopicSelector,
			To: reference.To{
				List:    &TopicList{},
				Managed: &Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopic")
		}
		mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DeadLetterPolicy[i3].DeadLetterTopicRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TopicRef,
		Selector:     mg.Spec.InitProvider.TopicSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionIAMMember.
func (mg *SubscriptionIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subscription),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubscriptionRef,
		Selector:     mg.Spec.ForProvider.SubscriptionSelector,
		To: reference.To{
			List:    &SubscriptionList{},
			Managed: &Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subscription")
	}
	mg.Spec.ForProvider.Subscription = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubscriptionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Subscription),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SubscriptionRef,
		Selector:     mg.Spec.InitProvider.SubscriptionSelector,
		To: reference.To{
			List:    &SubscriptionList{},
			Managed: &Subscription{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subscription")
	}
	mg.Spec.InitProvider.Subscription = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubscriptionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Topic.
func (mg *Topic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.KMSKeyNameRef,
		Selector:     mg.Spec.ForProvider.KMSKeyNameSelector,
		To: reference.To{
			List:    &v1beta1.CryptoKeyList{},
			Managed: &v1beta1.CryptoKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyName")
	}
	mg.Spec.ForProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.KMSKeyNameRef,
		Selector:     mg.Spec.InitProvider.KMSKeyNameSelector,
		To: reference.To{
			List:    &v1beta1.CryptoKeyList{},
			Managed: &v1beta1.CryptoKey{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyName")
	}
	mg.Spec.InitProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TopicIAMMember.
func (mg *TopicIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicRef,
		Selector:     mg.Spec.ForProvider.TopicSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Topic")
	}
	mg.Spec.ForProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Topic),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TopicRef,
		Selector:     mg.Spec.InitProvider.TopicSelector,
		To: reference.To{
			List:    &TopicList{},
			Managed: &Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Topic")
	}
	mg.Spec.InitProvider.Topic = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicRef = rsp.ResolvedReference

	return nil
}
