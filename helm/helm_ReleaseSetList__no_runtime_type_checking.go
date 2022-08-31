//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReleaseSetList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReleaseSetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReleaseSetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

