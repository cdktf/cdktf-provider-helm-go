package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-helm-go/helm/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-helm-go/helm/v6/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.0/docs helm}.
type HelmProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	BurstLimit() *float64
	SetBurstLimit(val *float64)
	BurstLimitInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Debug() interface{}
	SetDebug(val interface{})
	DebugInput() interface{}
	Experiments() *HelmProviderExperiments
	SetExperiments(val *HelmProviderExperiments)
	ExperimentsInput() *HelmProviderExperiments
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HelmDriver() *string
	SetHelmDriver(val *string)
	HelmDriverInput() *string
	Kubernetes() *HelmProviderKubernetes
	SetKubernetes(val *HelmProviderKubernetes)
	KubernetesInput() *HelmProviderKubernetes
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	PluginsPath() *string
	SetPluginsPath(val *string)
	PluginsPathInput() *string
	// Experimental.
	RawOverrides() interface{}
	Registry() interface{}
	SetRegistry(val interface{})
	RegistryConfigPath() *string
	SetRegistryConfigPath(val *string)
	RegistryConfigPathInput() *string
	RegistryInput() interface{}
	RepositoryCache() *string
	SetRepositoryCache(val *string)
	RepositoryCacheInput() *string
	RepositoryConfigPath() *string
	SetRepositoryConfigPath(val *string)
	RepositoryConfigPathInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetBurstLimit()
	ResetDebug()
	ResetExperiments()
	ResetHelmDriver()
	ResetKubernetes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginsPath()
	ResetRegistry()
	ResetRegistryConfigPath()
	ResetRepositoryCache()
	ResetRepositoryConfigPath()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HelmProvider
type jsiiProxy_HelmProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_HelmProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) BurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) BurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Debug() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) DebugInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Experiments() *HelmProviderExperiments {
	var returns *HelmProviderExperiments
	_jsii_.Get(
		j,
		"experiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) ExperimentsInput() *HelmProviderExperiments {
	var returns *HelmProviderExperiments
	_jsii_.Get(
		j,
		"experimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) HelmDriver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helmDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) HelmDriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helmDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Kubernetes() *HelmProviderKubernetes {
	var returns *HelmProviderKubernetes
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) KubernetesInput() *HelmProviderKubernetes {
	var returns *HelmProviderKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) PluginsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) PluginsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) Registry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RegistryConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryConfigPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RegistryConfigPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryConfigPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RepositoryCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RepositoryCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RepositoryConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryConfigPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) RepositoryConfigPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryConfigPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HelmProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.0/docs helm} Resource.
func NewHelmProvider(scope constructs.Construct, id *string, config *HelmProviderConfig) HelmProvider {
	_init_.Initialize()

	if err := validateNewHelmProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HelmProvider{}

	_jsii_.Create(
		"@cdktf/provider-helm.provider.HelmProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.0/docs helm} Resource.
func NewHelmProvider_Override(h HelmProvider, scope constructs.Construct, id *string, config *HelmProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.provider.HelmProvider",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HelmProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetBurstLimit(val *float64) {
	_jsii_.Set(
		j,
		"burstLimit",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetDebug(val interface{}) {
	if err := j.validateSetDebugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debug",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetExperiments(val *HelmProviderExperiments) {
	if err := j.validateSetExperimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experiments",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetHelmDriver(val *string) {
	_jsii_.Set(
		j,
		"helmDriver",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetKubernetes(val *HelmProviderKubernetes) {
	if err := j.validateSetKubernetesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetes",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetPluginsPath(val *string) {
	_jsii_.Set(
		j,
		"pluginsPath",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetRegistry(val interface{}) {
	if err := j.validateSetRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registry",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetRegistryConfigPath(val *string) {
	_jsii_.Set(
		j,
		"registryConfigPath",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetRepositoryCache(val *string) {
	_jsii_.Set(
		j,
		"repositoryCache",
		val,
	)
}

func (j *jsiiProxy_HelmProvider)SetRepositoryConfigPath(val *string) {
	_jsii_.Set(
		j,
		"repositoryConfigPath",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func HelmProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHelmProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.provider.HelmProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HelmProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHelmProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.provider.HelmProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HelmProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHelmProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.provider.HelmProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HelmProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.provider.HelmProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HelmProvider) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HelmProvider) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HelmProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		h,
		"resetAlias",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetBurstLimit() {
	_jsii_.InvokeVoid(
		h,
		"resetBurstLimit",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetDebug() {
	_jsii_.InvokeVoid(
		h,
		"resetDebug",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetExperiments() {
	_jsii_.InvokeVoid(
		h,
		"resetExperiments",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetHelmDriver() {
	_jsii_.InvokeVoid(
		h,
		"resetHelmDriver",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetKubernetes() {
	_jsii_.InvokeVoid(
		h,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetPluginsPath() {
	_jsii_.InvokeVoid(
		h,
		"resetPluginsPath",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetRegistry() {
	_jsii_.InvokeVoid(
		h,
		"resetRegistry",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetRegistryConfigPath() {
	_jsii_.InvokeVoid(
		h,
		"resetRegistryConfigPath",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetRepositoryCache() {
	_jsii_.InvokeVoid(
		h,
		"resetRepositoryCache",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) ResetRepositoryConfigPath() {
	_jsii_.InvokeVoid(
		h,
		"resetRepositoryConfigPath",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HelmProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelmProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelmProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelmProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

