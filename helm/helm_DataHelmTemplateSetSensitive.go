// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm


type DataHelmTemplateSetSensitive struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#type DataHelmTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

