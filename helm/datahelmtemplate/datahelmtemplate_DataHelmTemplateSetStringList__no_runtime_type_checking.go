//go:build no_runtime_type_checking

package datahelmtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataHelmTemplateSetStringList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataHelmTemplateSetStringList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataHelmTemplateSetStringListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

