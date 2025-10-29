// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datahelmtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHelmTemplateSetWoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetWoList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetWoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetWoList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetWoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetWoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetWoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHelmTemplateSetWoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

