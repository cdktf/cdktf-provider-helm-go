package provider


type HelmProviderKubernetesExec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#api_version HelmProvider#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#command HelmProvider#command}.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#args HelmProvider#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#env HelmProvider#env}.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

