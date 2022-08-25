// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm


type HelmProviderKubernetes struct {
	// PEM-encoded client certificate for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#client_certificate HelmProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded client certificate key for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#client_key HelmProvider#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// PEM-encoded root certificates bundle for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#cluster_ca_certificate HelmProvider#cluster_ca_certificate}
	ClusterCaCertificate *string `field:"optional" json:"clusterCaCertificate" yaml:"clusterCaCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#config_context HelmProvider#config_context}.
	ConfigContext *string `field:"optional" json:"configContext" yaml:"configContext"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#config_context_auth_info HelmProvider#config_context_auth_info}.
	ConfigContextAuthInfo *string `field:"optional" json:"configContextAuthInfo" yaml:"configContextAuthInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#config_context_cluster HelmProvider#config_context_cluster}.
	ConfigContextCluster *string `field:"optional" json:"configContextCluster" yaml:"configContextCluster"`
	// Path to the kube config file. Can be set with KUBE_CONFIG_PATH.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#config_path HelmProvider#config_path}
	ConfigPath *string `field:"optional" json:"configPath" yaml:"configPath"`
	// A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#config_paths HelmProvider#config_paths}
	ConfigPaths *[]*string `field:"optional" json:"configPaths" yaml:"configPaths"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#exec HelmProvider#exec}
	Exec *HelmProviderKubernetesExec `field:"optional" json:"exec" yaml:"exec"`
	// The hostname (in form of URI) of Kubernetes master.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#host HelmProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Whether server should be accessed without verifying the TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#insecure HelmProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#password HelmProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// URL to the proxy to be used for all API requests.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#proxy_url HelmProvider#proxy_url}
	ProxyUrl *string `field:"optional" json:"proxyUrl" yaml:"proxyUrl"`
	// Token to authenticate an service account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#token HelmProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#username HelmProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}
