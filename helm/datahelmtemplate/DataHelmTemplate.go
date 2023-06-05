package datahelmtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-helm-go/helm/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-helm-go/helm/v6/datahelmtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template helm_template}.
type DataHelmTemplate interface {
	cdktf.TerraformDataSource
	ApiVersions() *[]*string
	SetApiVersions(val *[]*string)
	ApiVersionsInput() *[]*string
	Atomic() interface{}
	SetAtomic(val interface{})
	AtomicInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chart() *string
	SetChart(val *string)
	ChartInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Crds() *[]*string
	SetCrds(val *[]*string)
	CrdsInput() *[]*string
	CreateNamespace() interface{}
	SetCreateNamespace(val interface{})
	CreateNamespaceInput() interface{}
	DependencyUpdate() interface{}
	SetDependencyUpdate(val interface{})
	DependencyUpdateInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Devel() interface{}
	SetDevel(val interface{})
	DevelInput() interface{}
	DisableOpenapiValidation() interface{}
	SetDisableOpenapiValidation(val interface{})
	DisableOpenapiValidationInput() interface{}
	DisableWebhooks() interface{}
	SetDisableWebhooks(val interface{})
	DisableWebhooksInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeCrds() interface{}
	SetIncludeCrds(val interface{})
	IncludeCrdsInput() interface{}
	IsUpgrade() interface{}
	SetIsUpgrade(val interface{})
	IsUpgradeInput() interface{}
	Keyring() *string
	SetKeyring(val *string)
	KeyringInput() *string
	KubeVersion() *string
	SetKubeVersion(val *string)
	KubeVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Manifest() *string
	SetManifest(val *string)
	ManifestInput() *string
	Manifests() *map[string]*string
	SetManifests(val *map[string]*string)
	ManifestsInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	PassCredentials() interface{}
	SetPassCredentials(val interface{})
	PassCredentialsInput() interface{}
	Postrender() DataHelmTemplatePostrenderOutputReference
	PostrenderInput() *DataHelmTemplatePostrender
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RenderSubchartNotes() interface{}
	SetRenderSubchartNotes(val interface{})
	RenderSubchartNotesInput() interface{}
	Replace() interface{}
	SetReplace(val interface{})
	ReplaceInput() interface{}
	Repository() *string
	SetRepository(val *string)
	RepositoryCaFile() *string
	SetRepositoryCaFile(val *string)
	RepositoryCaFileInput() *string
	RepositoryCertFile() *string
	SetRepositoryCertFile(val *string)
	RepositoryCertFileInput() *string
	RepositoryInput() *string
	RepositoryKeyFile() *string
	SetRepositoryKeyFile(val *string)
	RepositoryKeyFileInput() *string
	RepositoryPassword() *string
	SetRepositoryPassword(val *string)
	RepositoryPasswordInput() *string
	RepositoryUsername() *string
	SetRepositoryUsername(val *string)
	RepositoryUsernameInput() *string
	ResetValues() interface{}
	SetResetValues(val interface{})
	ResetValuesInput() interface{}
	ReuseValues() interface{}
	SetReuseValues(val interface{})
	ReuseValuesInput() interface{}
	Set() DataHelmTemplateSetList
	SetInput() interface{}
	SetList() DataHelmTemplateSetListStructList
	SetListInput() interface{}
	SetSensitive() DataHelmTemplateSetSensitiveList
	SetSensitiveInput() interface{}
	SetString() DataHelmTemplateSetStringList
	SetStringInput() interface{}
	ShowOnly() *[]*string
	SetShowOnly(val *[]*string)
	ShowOnlyInput() *[]*string
	SkipCrds() interface{}
	SetSkipCrds(val interface{})
	SkipCrdsInput() interface{}
	SkipTests() interface{}
	SetSkipTests(val interface{})
	SkipTestsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Validate() interface{}
	SetValidate(val interface{})
	ValidateInput() interface{}
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	Verify() interface{}
	SetVerify(val interface{})
	VerifyInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Wait() interface{}
	SetWait(val interface{})
	WaitInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPostrender(value *DataHelmTemplatePostrender)
	PutSet(value interface{})
	PutSetList(value interface{})
	PutSetSensitive(value interface{})
	PutSetString(value interface{})
	ResetApiVersions()
	ResetAtomic()
	ResetCrds()
	ResetCreateNamespace()
	ResetDependencyUpdate()
	ResetDescription()
	ResetDevel()
	ResetDisableOpenapiValidation()
	ResetDisableWebhooks()
	ResetId()
	ResetIncludeCrds()
	ResetIsUpgrade()
	ResetKeyring()
	ResetKubeVersion()
	ResetManifest()
	ResetManifests()
	ResetNamespace()
	ResetNotes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassCredentials()
	ResetPostrender()
	ResetRenderSubchartNotes()
	ResetReplace()
	ResetRepository()
	ResetRepositoryCaFile()
	ResetRepositoryCertFile()
	ResetRepositoryKeyFile()
	ResetRepositoryPassword()
	ResetRepositoryUsername()
	ResetResetValues()
	ResetReuseValues()
	ResetSet()
	ResetSetList()
	ResetSetSensitive()
	ResetSetString()
	ResetShowOnly()
	ResetSkipCrds()
	ResetSkipTests()
	ResetTfValues()
	ResetTimeout()
	ResetValidate()
	ResetVerify()
	ResetVersion()
	ResetWait()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataHelmTemplate
type jsiiProxy_DataHelmTemplate struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataHelmTemplate) ApiVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ApiVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Atomic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) AtomicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Chart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ChartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Crds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) CrdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) CreateNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) CreateNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DependencyUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DependencyUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Devel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"develInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DisableOpenapiValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DisableOpenapiValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DisableWebhooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) DisableWebhooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) IncludeCrds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCrds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) IncludeCrdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCrdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) IsUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) IsUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Keyring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) KeyringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) KubeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) KubeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Manifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Manifests() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"manifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ManifestsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"manifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) PassCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) PassCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Postrender() DataHelmTemplatePostrenderOutputReference {
	var returns DataHelmTemplatePostrenderOutputReference
	_jsii_.Get(
		j,
		"postrender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) PostrenderInput() *DataHelmTemplatePostrender {
	var returns *DataHelmTemplatePostrender
	_jsii_.Get(
		j,
		"postrenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RenderSubchartNotes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RenderSubchartNotesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Replace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryCaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) RepositoryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ResetValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ResetValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ReuseValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ReuseValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Set() DataHelmTemplateSetList {
	var returns DataHelmTemplateSetList
	_jsii_.Get(
		j,
		"set",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetList() DataHelmTemplateSetListStructList {
	var returns DataHelmTemplateSetListStructList
	_jsii_.Get(
		j,
		"setList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetSensitive() DataHelmTemplateSetSensitiveList {
	var returns DataHelmTemplateSetSensitiveList
	_jsii_.Get(
		j,
		"setSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetString() DataHelmTemplateSetStringList {
	var returns DataHelmTemplateSetStringList
	_jsii_.Get(
		j,
		"setString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SetStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ShowOnly() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"showOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ShowOnlyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"showOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SkipCrds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SkipCrdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SkipTests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) SkipTestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Validate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ValidateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) Wait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplate) WaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template helm_template} Data Source.
func NewDataHelmTemplate(scope constructs.Construct, id *string, config *DataHelmTemplateConfig) DataHelmTemplate {
	_init_.Initialize()

	if err := validateNewDataHelmTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataHelmTemplate{}

	_jsii_.Create(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/helm/2.10.1/docs/data-sources/template helm_template} Data Source.
func NewDataHelmTemplate_Override(d DataHelmTemplate, scope constructs.Construct, id *string, config *DataHelmTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetApiVersions(val *[]*string) {
	if err := j.validateSetApiVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiVersions",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetAtomic(val interface{}) {
	if err := j.validateSetAtomicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atomic",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetChart(val *string) {
	if err := j.validateSetChartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chart",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetCrds(val *[]*string) {
	if err := j.validateSetCrdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crds",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetCreateNamespace(val interface{}) {
	if err := j.validateSetCreateNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createNamespace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDependencyUpdate(val interface{}) {
	if err := j.validateSetDependencyUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencyUpdate",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDevel(val interface{}) {
	if err := j.validateSetDevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devel",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDisableOpenapiValidation(val interface{}) {
	if err := j.validateSetDisableOpenapiValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableOpenapiValidation",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetDisableWebhooks(val interface{}) {
	if err := j.validateSetDisableWebhooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableWebhooks",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetIncludeCrds(val interface{}) {
	if err := j.validateSetIncludeCrdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeCrds",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetIsUpgrade(val interface{}) {
	if err := j.validateSetIsUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isUpgrade",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetKeyring(val *string) {
	if err := j.validateSetKeyringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetKubeVersion(val *string) {
	if err := j.validateSetKubeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubeVersion",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetManifest(val *string) {
	if err := j.validateSetManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifest",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetManifests(val *map[string]*string) {
	if err := j.validateSetManifestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifests",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetPassCredentials(val interface{}) {
	if err := j.validateSetPassCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passCredentials",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRenderSubchartNotes(val interface{}) {
	if err := j.validateSetRenderSubchartNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renderSubchartNotes",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetReplace(val interface{}) {
	if err := j.validateSetReplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepositoryCaFile(val *string) {
	if err := j.validateSetRepositoryCaFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryCaFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepositoryCertFile(val *string) {
	if err := j.validateSetRepositoryCertFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryCertFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepositoryKeyFile(val *string) {
	if err := j.validateSetRepositoryKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryKeyFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepositoryPassword(val *string) {
	if err := j.validateSetRepositoryPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryPassword",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetRepositoryUsername(val *string) {
	if err := j.validateSetRepositoryUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUsername",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetResetValues(val interface{}) {
	if err := j.validateSetResetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetValues",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetReuseValues(val interface{}) {
	if err := j.validateSetReuseValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reuseValues",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetShowOnly(val *[]*string) {
	if err := j.validateSetShowOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showOnly",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetSkipCrds(val interface{}) {
	if err := j.validateSetSkipCrdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipCrds",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetSkipTests(val interface{}) {
	if err := j.validateSetSkipTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipTests",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetValidate(val interface{}) {
	if err := j.validateSetValidateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetVerify(val interface{}) {
	if err := j.validateSetVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate)SetWait(val interface{}) {
	if err := j.validateSetWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wait",
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
func DataHelmTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataHelmTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataHelmTemplate_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataHelmTemplate_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataHelmTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataHelmTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataHelmTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.dataHelmTemplate.DataHelmTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataHelmTemplate) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutPostrender(value *DataHelmTemplatePostrender) {
	if err := d.validatePutPostrenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostrender",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSet(value interface{}) {
	if err := d.validatePutSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSetList(value interface{}) {
	if err := d.validatePutSetListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetList",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSetSensitive(value interface{}) {
	if err := d.validatePutSetSensitiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetSensitive",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSetString(value interface{}) {
	if err := d.validatePutSetStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetString",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetApiVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetApiVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetAtomic() {
	_jsii_.InvokeVoid(
		d,
		"resetAtomic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetCrds() {
	_jsii_.InvokeVoid(
		d,
		"resetCrds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetCreateNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetDependencyUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetDependencyUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetDevel() {
	_jsii_.InvokeVoid(
		d,
		"resetDevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetDisableOpenapiValidation() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableOpenapiValidation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetDisableWebhooks() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableWebhooks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetIncludeCrds() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeCrds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetIsUpgrade() {
	_jsii_.InvokeVoid(
		d,
		"resetIsUpgrade",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetKeyring() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetKubeVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetKubeVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetManifest() {
	_jsii_.InvokeVoid(
		d,
		"resetManifest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetManifests() {
	_jsii_.InvokeVoid(
		d,
		"resetManifests",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetNotes() {
	_jsii_.InvokeVoid(
		d,
		"resetNotes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetPassCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetPassCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetPostrender() {
	_jsii_.InvokeVoid(
		d,
		"resetPostrender",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRenderSubchartNotes() {
	_jsii_.InvokeVoid(
		d,
		"resetRenderSubchartNotes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetReplace() {
	_jsii_.InvokeVoid(
		d,
		"resetReplace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepository() {
	_jsii_.InvokeVoid(
		d,
		"resetRepository",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepositoryCaFile() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryCaFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepositoryCertFile() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryCertFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepositoryKeyFile() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryKeyFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepositoryPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetRepositoryUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetResetValues() {
	_jsii_.InvokeVoid(
		d,
		"resetResetValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetReuseValues() {
	_jsii_.InvokeVoid(
		d,
		"resetReuseValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSet() {
	_jsii_.InvokeVoid(
		d,
		"resetSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSetList() {
	_jsii_.InvokeVoid(
		d,
		"resetSetList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSetSensitive() {
	_jsii_.InvokeVoid(
		d,
		"resetSetSensitive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSetString() {
	_jsii_.InvokeVoid(
		d,
		"resetSetString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetShowOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetShowOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSkipCrds() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipCrds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetSkipTests() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipTests",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetTfValues() {
	_jsii_.InvokeVoid(
		d,
		"resetTfValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetValidate() {
	_jsii_.InvokeVoid(
		d,
		"resetValidate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetVerify() {
	_jsii_.InvokeVoid(
		d,
		"resetVerify",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) ResetWait() {
	_jsii_.InvokeVoid(
		d,
		"resetWait",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

