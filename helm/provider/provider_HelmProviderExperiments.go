package provider


type HelmProviderExperiments struct {
	// Enable full diff by storing the rendered manifest in the state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#manifest HelmProvider#manifest}
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

