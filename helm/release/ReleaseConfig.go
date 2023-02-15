package release

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReleaseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#chart Release#chart}
	Chart *string `field:"required" json:"chart" yaml:"chart"`
	// Release name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#name Release#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#atomic Release#atomic}
	Atomic interface{} `field:"optional" json:"atomic" yaml:"atomic"`
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#cleanup_on_fail Release#cleanup_on_fail}
	CleanupOnFail interface{} `field:"optional" json:"cleanupOnFail" yaml:"cleanupOnFail"`
	// Create the namespace if it does not exist.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#create_namespace Release#create_namespace}
	CreateNamespace interface{} `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// Run helm dependency update before installing the chart.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#dependency_update Release#dependency_update}
	DependencyUpdate interface{} `field:"optional" json:"dependencyUpdate" yaml:"dependencyUpdate"`
	// Add a custom description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#description Release#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#devel Release#devel}
	Devel interface{} `field:"optional" json:"devel" yaml:"devel"`
	// Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#disable_crd_hooks Release#disable_crd_hooks}
	DisableCrdHooks interface{} `field:"optional" json:"disableCrdHooks" yaml:"disableCrdHooks"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#disable_openapi_validation Release#disable_openapi_validation}
	DisableOpenapiValidation interface{} `field:"optional" json:"disableOpenapiValidation" yaml:"disableOpenapiValidation"`
	// Prevent hooks from running.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#disable_webhooks Release#disable_webhooks}
	DisableWebhooks interface{} `field:"optional" json:"disableWebhooks" yaml:"disableWebhooks"`
	// Force resource update through delete/recreate if needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#force_update Release#force_update}
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#id Release#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location of public keys used for verification. Used only if `verify` is true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#keyring Release#keyring}
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// Run helm lint when planning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#lint Release#lint}
	Lint interface{} `field:"optional" json:"lint" yaml:"lint"`
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#max_history Release#max_history}
	MaxHistory *float64 `field:"optional" json:"maxHistory" yaml:"maxHistory"`
	// Namespace to install the release into.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#namespace Release#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Pass credentials to all domains.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#pass_credentials Release#pass_credentials}
	PassCredentials interface{} `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	// postrender block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#postrender Release#postrender}
	Postrender *ReleasePostrender `field:"optional" json:"postrender" yaml:"postrender"`
	// Perform pods restart during upgrade/rollback.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#recreate_pods Release#recreate_pods}
	RecreatePods interface{} `field:"optional" json:"recreatePods" yaml:"recreatePods"`
	// If set, render subchart notes along with the parent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#render_subchart_notes Release#render_subchart_notes}
	RenderSubchartNotes interface{} `field:"optional" json:"renderSubchartNotes" yaml:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#replace Release#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
	// Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository Release#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The Repositories CA File.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository_ca_file Release#repository_ca_file}
	RepositoryCaFile *string `field:"optional" json:"repositoryCaFile" yaml:"repositoryCaFile"`
	// The repositories cert file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository_cert_file Release#repository_cert_file}
	RepositoryCertFile *string `field:"optional" json:"repositoryCertFile" yaml:"repositoryCertFile"`
	// The repositories cert key file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository_key_file Release#repository_key_file}
	RepositoryKeyFile *string `field:"optional" json:"repositoryKeyFile" yaml:"repositoryKeyFile"`
	// Password for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository_password Release#repository_password}
	RepositoryPassword *string `field:"optional" json:"repositoryPassword" yaml:"repositoryPassword"`
	// Username for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#repository_username Release#repository_username}
	RepositoryUsername *string `field:"optional" json:"repositoryUsername" yaml:"repositoryUsername"`
	// When upgrading, reset the values to the ones built into the chart.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#reset_values Release#reset_values}
	ResetValues interface{} `field:"optional" json:"resetValues" yaml:"resetValues"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#reuse_values Release#reuse_values}
	ReuseValues interface{} `field:"optional" json:"reuseValues" yaml:"reuseValues"`
	// set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#set Release#set}
	Set interface{} `field:"optional" json:"set" yaml:"set"`
	// set_sensitive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#set_sensitive Release#set_sensitive}
	SetSensitive interface{} `field:"optional" json:"setSensitive" yaml:"setSensitive"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#skip_crds Release#skip_crds}
	SkipCrds interface{} `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// Time in seconds to wait for any individual kubernetes operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#timeout Release#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// List of values in raw yaml format to pass to helm.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#values Release#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
	// Verify the package before installing it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#verify Release#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#version Release#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Will wait until all resources are in a ready state before marking the release as successful.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#wait Release#wait}
	Wait interface{} `field:"optional" json:"wait" yaml:"wait"`
	// If wait is enabled, will wait until all Jobs have been completed before marking the release as successful.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#wait_for_jobs Release#wait_for_jobs}
	WaitForJobs interface{} `field:"optional" json:"waitForJobs" yaml:"waitForJobs"`
}

