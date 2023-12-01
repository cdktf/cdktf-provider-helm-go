// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderExperiments struct {
	// Enable full diff by storing the rendered manifest in the state.
	//
	// This has similar limitations as when using helm install --dry-run. See https://helm.sh/docs/chart_best_practices/custom_resource_definitions/#install-a-crd-declaration-before-using-the-resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.12.1/docs#manifest HelmProvider#manifest}
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

