// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datahelmtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHelmTemplateSetList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHelmTemplateSetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

