// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderRegistry struct {
	// The password to use for the OCI HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#password HelmProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// OCI URL in form of oci://host:port or oci://host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#url HelmProvider#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The username to use for the OCI HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.13.0/docs#username HelmProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

