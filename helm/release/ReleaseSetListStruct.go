package release


type ReleaseSetListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.0/docs/resources/release#name Release#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.0/docs/resources/release#value Release#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

