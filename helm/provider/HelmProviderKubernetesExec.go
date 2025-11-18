// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderKubernetesExec struct {
	// API version for the exec plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.1/docs#api_version HelmProvider#api_version}
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Command to run for Kubernetes exec plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.1/docs#command HelmProvider#command}
	Command *string `field:"required" json:"command" yaml:"command"`
	// Arguments for the exec plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.1/docs#args HelmProvider#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Environment variables for the exec plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.1.1/docs#env HelmProvider#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

