package datahelmtemplate


type DataHelmTemplateSetListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template#value DataHelmTemplate#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

