// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm


type HelmProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#alias HelmProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Debug indicates whether or not Helm is running in Debug mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#debug HelmProvider#debug}
	Debug interface{} `field:"optional" json:"debug" yaml:"debug"`
	// experiments block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#experiments HelmProvider#experiments}
	Experiments *HelmProviderExperiments `field:"optional" json:"experiments" yaml:"experiments"`
	// The backend storage driver. Values are: configmap, secret, memory, sql.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#helm_driver HelmProvider#helm_driver}
	HelmDriver *string `field:"optional" json:"helmDriver" yaml:"helmDriver"`
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#kubernetes HelmProvider#kubernetes}
	Kubernetes *HelmProviderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// The path to the helm plugins directory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#plugins_path HelmProvider#plugins_path}
	PluginsPath *string `field:"optional" json:"pluginsPath" yaml:"pluginsPath"`
	// The path to the registry config file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#registry_config_path HelmProvider#registry_config_path}
	RegistryConfigPath *string `field:"optional" json:"registryConfigPath" yaml:"registryConfigPath"`
	// The path to the file containing cached repository indexes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#repository_cache HelmProvider#repository_cache}
	RepositoryCache *string `field:"optional" json:"repositoryCache" yaml:"repositoryCache"`
	// The path to the file containing repository names and URLs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#repository_config_path HelmProvider#repository_config_path}
	RepositoryConfigPath *string `field:"optional" json:"repositoryConfigPath" yaml:"repositoryConfigPath"`
}

