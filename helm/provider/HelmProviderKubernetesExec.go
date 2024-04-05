// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderKubernetesExec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#api_version HelmProvider#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#command HelmProvider#command}.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#args HelmProvider#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#env HelmProvider#env}.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

