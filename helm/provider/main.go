// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-helm.provider.HelmProvider",
		reflect.TypeOf((*HelmProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "burstLimit", GoGetter: "BurstLimit"},
			_jsii_.MemberProperty{JsiiProperty: "burstLimitInput", GoGetter: "BurstLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "debug", GoGetter: "Debug"},
			_jsii_.MemberProperty{JsiiProperty: "debugInput", GoGetter: "DebugInput"},
			_jsii_.MemberProperty{JsiiProperty: "experiments", GoGetter: "Experiments"},
			_jsii_.MemberProperty{JsiiProperty: "experimentsInput", GoGetter: "ExperimentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "helmDriver", GoGetter: "HelmDriver"},
			_jsii_.MemberProperty{JsiiProperty: "helmDriverInput", GoGetter: "HelmDriverInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetes", GoGetter: "Kubernetes"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesInput", GoGetter: "KubernetesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsPath", GoGetter: "PluginsPath"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsPathInput", GoGetter: "PluginsPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "registry", GoGetter: "Registry"},
			_jsii_.MemberProperty{JsiiProperty: "registryConfigPath", GoGetter: "RegistryConfigPath"},
			_jsii_.MemberProperty{JsiiProperty: "registryConfigPathInput", GoGetter: "RegistryConfigPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "registryInput", GoGetter: "RegistryInput"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCache", GoGetter: "RepositoryCache"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCacheInput", GoGetter: "RepositoryCacheInput"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryConfigPath", GoGetter: "RepositoryConfigPath"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryConfigPathInput", GoGetter: "RepositoryConfigPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetBurstLimit", GoMethod: "ResetBurstLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetDebug", GoMethod: "ResetDebug"},
			_jsii_.MemberMethod{JsiiMethod: "resetExperiments", GoMethod: "ResetExperiments"},
			_jsii_.MemberMethod{JsiiMethod: "resetHelmDriver", GoMethod: "ResetHelmDriver"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetes", GoMethod: "ResetKubernetes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginsPath", GoMethod: "ResetPluginsPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegistry", GoMethod: "ResetRegistry"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegistryConfigPath", GoMethod: "ResetRegistryConfigPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepositoryCache", GoMethod: "ResetRepositoryCache"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepositoryConfigPath", GoMethod: "ResetRepositoryConfigPath"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_HelmProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-helm.provider.HelmProviderConfig",
		reflect.TypeOf((*HelmProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-helm.provider.HelmProviderExperiments",
		reflect.TypeOf((*HelmProviderExperiments)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-helm.provider.HelmProviderKubernetes",
		reflect.TypeOf((*HelmProviderKubernetes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-helm.provider.HelmProviderKubernetesExec",
		reflect.TypeOf((*HelmProviderKubernetesExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-helm.provider.HelmProviderRegistry",
		reflect.TypeOf((*HelmProviderRegistry)(nil)).Elem(),
	)
}
