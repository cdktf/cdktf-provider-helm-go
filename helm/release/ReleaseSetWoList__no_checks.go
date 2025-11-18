// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReleaseSetWoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReleaseSetWoList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReleaseSetWoList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetWoList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetWoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetWoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetWoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReleaseSetWoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

