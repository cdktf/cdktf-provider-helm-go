// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package release

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReleaseSetSensitiveList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReleaseSetSensitiveList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReleaseSetSensitiveList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetSensitiveList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetSensitiveList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetSensitiveList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReleaseSetSensitiveList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReleaseSetSensitiveListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

