// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahelmtemplate


type DataHelmTemplateSetSensitive struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.14.0/docs/data-sources/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.14.0/docs/data-sources/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.14.0/docs/data-sources/template#type DataHelmTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

