package datahelmtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHelmTemplateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Chart name to be installed. A path may be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#chart DataHelmTemplate#chart}
	Chart *string `field:"required" json:"chart" yaml:"chart"`
	// Release name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#name DataHelmTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Kubernetes api versions used for Capabilities.APIVersions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#api_versions DataHelmTemplate#api_versions}
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	// If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#atomic DataHelmTemplate#atomic}
	Atomic interface{} `field:"optional" json:"atomic" yaml:"atomic"`
	// List of rendered CRDs from the chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#crds DataHelmTemplate#crds}
	Crds *[]*string `field:"optional" json:"crds" yaml:"crds"`
	// Create the namespace if it does not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#create_namespace DataHelmTemplate#create_namespace}
	CreateNamespace interface{} `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// Run helm dependency update before installing the chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#dependency_update DataHelmTemplate#dependency_update}
	DependencyUpdate interface{} `field:"optional" json:"dependencyUpdate" yaml:"dependencyUpdate"`
	// Add a custom description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#description DataHelmTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#devel DataHelmTemplate#devel}
	Devel interface{} `field:"optional" json:"devel" yaml:"devel"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#disable_openapi_validation DataHelmTemplate#disable_openapi_validation}
	DisableOpenapiValidation interface{} `field:"optional" json:"disableOpenapiValidation" yaml:"disableOpenapiValidation"`
	// Prevent hooks from running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#disable_webhooks DataHelmTemplate#disable_webhooks}
	DisableWebhooks interface{} `field:"optional" json:"disableWebhooks" yaml:"disableWebhooks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#id DataHelmTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Include CRDs in the templated output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#include_crds DataHelmTemplate#include_crds}
	IncludeCrds interface{} `field:"optional" json:"includeCrds" yaml:"includeCrds"`
	// Set .Release.IsUpgrade instead of .Release.IsInstall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#is_upgrade DataHelmTemplate#is_upgrade}
	IsUpgrade interface{} `field:"optional" json:"isUpgrade" yaml:"isUpgrade"`
	// Location of public keys used for verification. Used only if `verify` is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#keyring DataHelmTemplate#keyring}
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// Kubernetes version used for Capabilities.KubeVersion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#kube_version DataHelmTemplate#kube_version}
	KubeVersion *string `field:"optional" json:"kubeVersion" yaml:"kubeVersion"`
	// Concatenated rendered chart templates. This corresponds to the output of the `helm template` command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#manifest DataHelmTemplate#manifest}
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// Map of rendered chart templates indexed by the template name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#manifests DataHelmTemplate#manifests}
	Manifests *map[string]*string `field:"optional" json:"manifests" yaml:"manifests"`
	// Namespace to install the release into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#namespace DataHelmTemplate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Rendered notes if the chart contains a `NOTES.txt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#notes DataHelmTemplate#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Pass credentials to all domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#pass_credentials DataHelmTemplate#pass_credentials}
	PassCredentials interface{} `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	// postrender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#postrender DataHelmTemplate#postrender}
	Postrender *DataHelmTemplatePostrender `field:"optional" json:"postrender" yaml:"postrender"`
	// If set, render subchart notes along with the parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#render_subchart_notes DataHelmTemplate#render_subchart_notes}
	RenderSubchartNotes interface{} `field:"optional" json:"renderSubchartNotes" yaml:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#replace DataHelmTemplate#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
	// Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository DataHelmTemplate#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The Repositories CA File.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository_ca_file DataHelmTemplate#repository_ca_file}
	RepositoryCaFile *string `field:"optional" json:"repositoryCaFile" yaml:"repositoryCaFile"`
	// The repositories cert file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository_cert_file DataHelmTemplate#repository_cert_file}
	RepositoryCertFile *string `field:"optional" json:"repositoryCertFile" yaml:"repositoryCertFile"`
	// The repositories cert key file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository_key_file DataHelmTemplate#repository_key_file}
	RepositoryKeyFile *string `field:"optional" json:"repositoryKeyFile" yaml:"repositoryKeyFile"`
	// Password for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository_password DataHelmTemplate#repository_password}
	RepositoryPassword *string `field:"optional" json:"repositoryPassword" yaml:"repositoryPassword"`
	// Username for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#repository_username DataHelmTemplate#repository_username}
	RepositoryUsername *string `field:"optional" json:"repositoryUsername" yaml:"repositoryUsername"`
	// When upgrading, reset the values to the ones built into the chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#reset_values DataHelmTemplate#reset_values}
	ResetValues interface{} `field:"optional" json:"resetValues" yaml:"resetValues"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#reuse_values DataHelmTemplate#reuse_values}
	ReuseValues interface{} `field:"optional" json:"reuseValues" yaml:"reuseValues"`
	// set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#set DataHelmTemplate#set}
	Set interface{} `field:"optional" json:"set" yaml:"set"`
	// set_sensitive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#set_sensitive DataHelmTemplate#set_sensitive}
	SetSensitive interface{} `field:"optional" json:"setSensitive" yaml:"setSensitive"`
	// set_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#set_string DataHelmTemplate#set_string}
	SetString interface{} `field:"optional" json:"setString" yaml:"setString"`
	// Only show manifests rendered from the given templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#show_only DataHelmTemplate#show_only}
	ShowOnly *[]*string `field:"optional" json:"showOnly" yaml:"showOnly"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#skip_crds DataHelmTemplate#skip_crds}
	SkipCrds interface{} `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// If set, tests will not be rendered. By default, tests are rendered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#skip_tests DataHelmTemplate#skip_tests}
	SkipTests interface{} `field:"optional" json:"skipTests" yaml:"skipTests"`
	// Time in seconds to wait for any individual kubernetes operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#timeout DataHelmTemplate#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Validate your manifests against the Kubernetes cluster you are currently pointing at.
	//
	// This is the same validation performed on an install
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#validate DataHelmTemplate#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
	// List of values in raw yaml format to pass to helm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#values DataHelmTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
	// Verify the package before installing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#verify DataHelmTemplate#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#version DataHelmTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Will wait until all resources are in a ready state before marking the release as successful.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/2.9.0/docs/data-sources/template#wait DataHelmTemplate#wait}
	Wait interface{} `field:"optional" json:"wait" yaml:"wait"`
}

