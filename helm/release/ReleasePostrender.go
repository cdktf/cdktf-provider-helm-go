// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package release


type ReleasePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.16.1/docs/resources/release#binary_path Release#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
	// an argument to the post-renderer (can specify multiple).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.16.1/docs/resources/release#args Release#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

