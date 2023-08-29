// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderExperiments struct {
	// Enable full diff by storing the rendered manifest in the state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs#manifest HelmProvider#manifest}
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

