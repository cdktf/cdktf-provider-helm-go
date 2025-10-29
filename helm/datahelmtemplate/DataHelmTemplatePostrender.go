// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahelmtemplate


type DataHelmTemplatePostrender struct {
	// The common binary path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.0/docs/data-sources/template#binary_path DataHelmTemplate#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
	// An argument to the post-renderer (can specify multiple).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.0/docs/data-sources/template#args DataHelmTemplate#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

