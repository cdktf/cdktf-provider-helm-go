// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package release


type ReleaseSetSensitive struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs/resources/release#name Release#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs/resources/release#value Release#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs/resources/release#type Release#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

