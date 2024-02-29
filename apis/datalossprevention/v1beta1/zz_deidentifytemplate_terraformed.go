//go:build (datalossprevention || all) && !ignore_autogenerated

// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this DeidentifyTemplate
func (mg *DeidentifyTemplate) GetTerraformResourceType() string {
	return "google_data_loss_prevention_deidentify_template"
}

// GetConnectionDetailsMapping for this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"deidentify_config[*].record_transformations[*].field_transformations[*].info_type_transformations[*].transformations[*].primitive_transformation[*].crypto_deterministic_config[*].crypto_key[*].unwrapped[*].key": "spec.forProvider.deidentifyConfig[*].recordTransformations[*].fieldTransformations[*].infoTypeTransformations[*].transformations[*].primitiveTransformation[*].cryptoDeterministicConfig[*].cryptoKey[*].unwrapped[*].keySecretRef", "deidentify_config[*].record_transformations[*].field_transformations[*].info_type_transformations[*].transformations[*].primitive_transformation[*].crypto_hash_config[*].crypto_key[*].unwrapped[*].key": "spec.forProvider.deidentifyConfig[*].recordTransformations[*].fieldTransformations[*].infoTypeTransformations[*].transformations[*].primitiveTransformation[*].cryptoHashConfig[*].cryptoKey[*].unwrapped[*].keySecretRef", "deidentify_config[*].record_transformations[*].field_transformations[*].info_type_transformations[*].transformations[*].primitive_transformation[*].crypto_replace_ffx_fpe_config[*].crypto_key[*].unwrapped[*].key": "spec.forProvider.deidentifyConfig[*].recordTransformations[*].fieldTransformations[*].infoTypeTransformations[*].transformations[*].primitiveTransformation[*].cryptoReplaceFfxFpeConfig[*].cryptoKey[*].unwrapped[*].keySecretRef", "deidentify_config[*].record_transformations[*].field_transformations[*].info_type_transformations[*].transformations[*].primitive_transformation[*].date_shift_config[*].crypto_key[*].unwrapped[*].key": "spec.forProvider.deidentifyConfig[*].recordTransformations[*].fieldTransformations[*].infoTypeTransformations[*].transformations[*].primitiveTransformation[*].dateShiftConfig[*].cryptoKey[*].unwrapped[*].keySecretRef"}
}

// GetObservation of this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DeidentifyTemplate
func (tr *DeidentifyTemplate) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DeidentifyTemplate
func (tr *DeidentifyTemplate) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this DeidentifyTemplate
func (tr *DeidentifyTemplate) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this DeidentifyTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DeidentifyTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &DeidentifyTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DeidentifyTemplate) GetTerraformSchemaVersion() int {
	return 0
}
