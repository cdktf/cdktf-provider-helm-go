// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#alias HelmProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Helm burst limit. Increase this if you have a cluster with many CRDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#burst_limit HelmProvider#burst_limit}
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Debug indicates whether or not Helm is running in Debug mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#debug HelmProvider#debug}
	Debug interface{} `field:"optional" json:"debug" yaml:"debug"`
	// Enable and disable experimental features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#experiments HelmProvider#experiments}
	Experiments *HelmProviderExperiments `field:"optional" json:"experiments" yaml:"experiments"`
	// The backend storage driver. Values are: configmap, secret, memory, sql.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#helm_driver HelmProvider#helm_driver}
	HelmDriver *string `field:"optional" json:"helmDriver" yaml:"helmDriver"`
	// Kubernetes Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#kubernetes HelmProvider#kubernetes}
	Kubernetes *HelmProviderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// The path to the helm plugins directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#plugins_path HelmProvider#plugins_path}
	PluginsPath *string `field:"optional" json:"pluginsPath" yaml:"pluginsPath"`
	// RegistryClient configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#registries HelmProvider#registries}
	Registries interface{} `field:"optional" json:"registries" yaml:"registries"`
	// The path to the registry config file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#registry_config_path HelmProvider#registry_config_path}
	RegistryConfigPath *string `field:"optional" json:"registryConfigPath" yaml:"registryConfigPath"`
	// The path to the file containing cached repository indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#repository_cache HelmProvider#repository_cache}
	RepositoryCache *string `field:"optional" json:"repositoryCache" yaml:"repositoryCache"`
	// The path to the file containing repository names and URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#repository_config_path HelmProvider#repository_config_path}
	RepositoryConfigPath *string `field:"optional" json:"repositoryConfigPath" yaml:"repositoryConfigPath"`
}

