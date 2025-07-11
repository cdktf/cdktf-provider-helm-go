// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type HelmProviderKubernetes struct {
	// PEM-encoded client certificate for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#client_certificate HelmProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded client certificate key for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#client_key HelmProvider#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// PEM-encoded root certificates bundle for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#cluster_ca_certificate HelmProvider#cluster_ca_certificate}
	ClusterCaCertificate *string `field:"optional" json:"clusterCaCertificate" yaml:"clusterCaCertificate"`
	// Context to choose from the config file. Can be sourced from KUBE_CTX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#config_context HelmProvider#config_context}
	ConfigContext *string `field:"optional" json:"configContext" yaml:"configContext"`
	// Authentication info context of the kube config (name of the kubeconfig user, --user flag in kubectl).
	//
	// Can be sourced from KUBE_CTX_AUTH_INFO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#config_context_auth_info HelmProvider#config_context_auth_info}
	ConfigContextAuthInfo *string `field:"optional" json:"configContextAuthInfo" yaml:"configContextAuthInfo"`
	// Cluster context of the kube config (name of the kubeconfig cluster, --cluster flag in kubectl).
	//
	// Can be sourced from KUBE_CTX_CLUSTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#config_context_cluster HelmProvider#config_context_cluster}
	ConfigContextCluster *string `field:"optional" json:"configContextCluster" yaml:"configContextCluster"`
	// Path to the kube config file. Can be set with KUBE_CONFIG_PATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#config_path HelmProvider#config_path}
	ConfigPath *string `field:"optional" json:"configPath" yaml:"configPath"`
	// A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#config_paths HelmProvider#config_paths}
	ConfigPaths *[]*string `field:"optional" json:"configPaths" yaml:"configPaths"`
	// Exec configuration for Kubernetes authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#exec HelmProvider#exec}
	Exec *HelmProviderKubernetesExec `field:"optional" json:"exec" yaml:"exec"`
	// The hostname (in form of URI) of kubernetes master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#host HelmProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Whether server should be accessed without verifying the TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#insecure HelmProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#password HelmProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// URL to the proxy to be used for all API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#proxy_url HelmProvider#proxy_url}
	ProxyUrl *string `field:"optional" json:"proxyUrl" yaml:"proxyUrl"`
	// Server name passed to the server for SNI and is used in the client to check server certificates against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#tls_server_name HelmProvider#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Token to authenticate a service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#token HelmProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/helm/3.0.2/docs#username HelmProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

