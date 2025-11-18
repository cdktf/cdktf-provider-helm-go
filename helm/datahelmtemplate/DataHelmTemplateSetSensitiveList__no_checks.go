// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datahelmtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHelmTemplateSetSensitiveListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

