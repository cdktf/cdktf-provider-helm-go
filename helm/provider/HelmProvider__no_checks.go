// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HelmProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HelmProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHelmProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHelmProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHelmProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HelmProvider) validateSetDebugParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HelmProvider) validateSetExperimentsParameters(val *HelmProviderExperiments) error {
	return nil
}

func (j *jsiiProxy_HelmProvider) validateSetKubernetesParameters(val *HelmProviderKubernetes) error {
	return nil
}

func (j *jsiiProxy_HelmProvider) validateSetRegistryParameters(val interface{}) error {
	return nil
}

func validateNewHelmProviderParameters(scope constructs.Construct, id *string, config *HelmProviderConfig) error {
	return nil
}

