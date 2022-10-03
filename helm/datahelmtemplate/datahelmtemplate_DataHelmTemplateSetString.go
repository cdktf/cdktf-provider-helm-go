package datahelmtemplate


type DataHelmTemplateSetString struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

