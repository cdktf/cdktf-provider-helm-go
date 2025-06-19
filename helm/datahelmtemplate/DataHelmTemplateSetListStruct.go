// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahelmtemplate


type DataHelmTemplateSetListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.1/docs/data-sources/template#value DataHelmTemplate#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.1/docs/data-sources/template#name DataHelmTemplate#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

