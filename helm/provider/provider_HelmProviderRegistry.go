package provider


type HelmProviderRegistry struct {
	// The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#password HelmProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// OCI URL in form of oci://host:port or oci://host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#url HelmProvider#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#username HelmProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

