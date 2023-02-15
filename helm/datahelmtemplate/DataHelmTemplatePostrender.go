package datahelmtemplate


type DataHelmTemplatePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#binary_path DataHelmTemplate#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
}

