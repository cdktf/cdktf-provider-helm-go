package release


type ReleasePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#binary_path Release#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
	// an argument to the post-renderer (can specify multiple).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#args Release#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

