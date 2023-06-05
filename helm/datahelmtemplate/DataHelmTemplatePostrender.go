package datahelmtemplate


type DataHelmTemplatePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template#binary_path DataHelmTemplate#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
}

