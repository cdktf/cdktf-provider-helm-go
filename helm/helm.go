// Prebuilt helm Provider for Terraform CDK (cdktf)
package helm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-helm-go/helm/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-helm-go/helm/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/helm/d/template helm_template}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	PutSetSensitive(value interface{})
	PutSetString(value interface{})
	ResetApiVersions()
	ResetAtomic()
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

func (j *jsiiProxy_DataHelmTemplate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
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


// Create a new {@link https://www.terraform.io/docs/providers/helm/d/template helm_template} Data Source.
func NewDataHelmTemplate(scope constructs.Construct, id *string, config *DataHelmTemplateConfig) DataHelmTemplate {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplate{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/helm/d/template helm_template} Data Source.
func NewDataHelmTemplate_Override(d DataHelmTemplate, scope constructs.Construct, id *string, config *DataHelmTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetApiVersions(val *[]*string) {
	_jsii_.Set(
		j,
		"apiVersions",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetAtomic(val interface{}) {
	_jsii_.Set(
		j,
		"atomic",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetChart(val *string) {
	_jsii_.Set(
		j,
		"chart",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetCreateNamespace(val interface{}) {
	_jsii_.Set(
		j,
		"createNamespace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDependencyUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"dependencyUpdate",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDevel(val interface{}) {
	_jsii_.Set(
		j,
		"devel",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDisableOpenapiValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableOpenapiValidation",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetDisableWebhooks(val interface{}) {
	_jsii_.Set(
		j,
		"disableWebhooks",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetIncludeCrds(val interface{}) {
	_jsii_.Set(
		j,
		"includeCrds",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetIsUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"isUpgrade",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetKeyring(val *string) {
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetManifest(val *string) {
	_jsii_.Set(
		j,
		"manifest",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetManifests(val *map[string]*string) {
	_jsii_.Set(
		j,
		"manifests",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetNotes(val *string) {
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetPassCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"passCredentials",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRenderSubchartNotes(val interface{}) {
	_jsii_.Set(
		j,
		"renderSubchartNotes",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetReplace(val interface{}) {
	_jsii_.Set(
		j,
		"replace",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepositoryCaFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCaFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepositoryCertFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCertFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepositoryKeyFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryKeyFile",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepositoryPassword(val *string) {
	_jsii_.Set(
		j,
		"repositoryPassword",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetRepositoryUsername(val *string) {
	_jsii_.Set(
		j,
		"repositoryUsername",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetResetValues(val interface{}) {
	_jsii_.Set(
		j,
		"resetValues",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetReuseValues(val interface{}) {
	_jsii_.Set(
		j,
		"reuseValues",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetShowOnly(val *[]*string) {
	_jsii_.Set(
		j,
		"showOnly",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetSkipCrds(val interface{}) {
	_jsii_.Set(
		j,
		"skipCrds",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetSkipTests(val interface{}) {
	_jsii_.Set(
		j,
		"skipTests",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetValidate(val interface{}) {
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetVerify(val interface{}) {
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplate) SetWait(val interface{}) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.DataHelmTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataHelmTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.DataHelmTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataHelmTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutPostrender(value *DataHelmTemplatePostrender) {
	_jsii_.InvokeVoid(
		d,
		"putPostrender",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSet(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putSet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSetSensitive(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putSetSensitive",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataHelmTemplate) PutSetString(value interface{}) {
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

type DataHelmTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#chart DataHelmTemplate#chart}
	Chart *string `field:"required" json:"chart" yaml:"chart"`
	// Release name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Kubernetes api versions used for Capabilities.APIVersions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#api_versions DataHelmTemplate#api_versions}
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	// If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#atomic DataHelmTemplate#atomic}
	Atomic interface{} `field:"optional" json:"atomic" yaml:"atomic"`
	// Create the namespace if it does not exist.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#create_namespace DataHelmTemplate#create_namespace}
	CreateNamespace interface{} `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// Run helm dependency update before installing the chart.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#dependency_update DataHelmTemplate#dependency_update}
	DependencyUpdate interface{} `field:"optional" json:"dependencyUpdate" yaml:"dependencyUpdate"`
	// Add a custom description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#description DataHelmTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#devel DataHelmTemplate#devel}
	Devel interface{} `field:"optional" json:"devel" yaml:"devel"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#disable_openapi_validation DataHelmTemplate#disable_openapi_validation}
	DisableOpenapiValidation interface{} `field:"optional" json:"disableOpenapiValidation" yaml:"disableOpenapiValidation"`
	// Prevent hooks from running.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#disable_webhooks DataHelmTemplate#disable_webhooks}
	DisableWebhooks interface{} `field:"optional" json:"disableWebhooks" yaml:"disableWebhooks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#id DataHelmTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Include CRDs in the templated output.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#include_crds DataHelmTemplate#include_crds}
	IncludeCrds interface{} `field:"optional" json:"includeCrds" yaml:"includeCrds"`
	// Set .Release.IsUpgrade instead of .Release.IsInstall.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#is_upgrade DataHelmTemplate#is_upgrade}
	IsUpgrade interface{} `field:"optional" json:"isUpgrade" yaml:"isUpgrade"`
	// Location of public keys used for verification. Used only if `verify` is true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#keyring DataHelmTemplate#keyring}
	Keyring *string `field:"optional" json:"keyring" yaml:"keyring"`
	// Concatenated rendered chart templates. This corresponds to the output of the `helm template` command.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#manifest DataHelmTemplate#manifest}
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// Map of rendered chart templates indexed by the template name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#manifests DataHelmTemplate#manifests}
	Manifests *map[string]*string `field:"optional" json:"manifests" yaml:"manifests"`
	// Namespace to install the release into.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#namespace DataHelmTemplate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Rendered notes if the chart contains a `NOTES.txt`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#notes DataHelmTemplate#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Pass credentials to all domains.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#pass_credentials DataHelmTemplate#pass_credentials}
	PassCredentials interface{} `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	// postrender block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#postrender DataHelmTemplate#postrender}
	Postrender *DataHelmTemplatePostrender `field:"optional" json:"postrender" yaml:"postrender"`
	// If set, render subchart notes along with the parent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#render_subchart_notes DataHelmTemplate#render_subchart_notes}
	RenderSubchartNotes interface{} `field:"optional" json:"renderSubchartNotes" yaml:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#replace DataHelmTemplate#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
	// Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository DataHelmTemplate#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The Repositories CA File.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository_ca_file DataHelmTemplate#repository_ca_file}
	RepositoryCaFile *string `field:"optional" json:"repositoryCaFile" yaml:"repositoryCaFile"`
	// The repositories cert file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository_cert_file DataHelmTemplate#repository_cert_file}
	RepositoryCertFile *string `field:"optional" json:"repositoryCertFile" yaml:"repositoryCertFile"`
	// The repositories cert key file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository_key_file DataHelmTemplate#repository_key_file}
	RepositoryKeyFile *string `field:"optional" json:"repositoryKeyFile" yaml:"repositoryKeyFile"`
	// Password for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository_password DataHelmTemplate#repository_password}
	RepositoryPassword *string `field:"optional" json:"repositoryPassword" yaml:"repositoryPassword"`
	// Username for HTTP basic authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#repository_username DataHelmTemplate#repository_username}
	RepositoryUsername *string `field:"optional" json:"repositoryUsername" yaml:"repositoryUsername"`
	// When upgrading, reset the values to the ones built into the chart.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#reset_values DataHelmTemplate#reset_values}
	ResetValues interface{} `field:"optional" json:"resetValues" yaml:"resetValues"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'reset_values' is specified, this is ignored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#reuse_values DataHelmTemplate#reuse_values}
	ReuseValues interface{} `field:"optional" json:"reuseValues" yaml:"reuseValues"`
	// set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#set DataHelmTemplate#set}
	Set interface{} `field:"optional" json:"set" yaml:"set"`
	// set_sensitive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#set_sensitive DataHelmTemplate#set_sensitive}
	SetSensitive interface{} `field:"optional" json:"setSensitive" yaml:"setSensitive"`
	// set_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#set_string DataHelmTemplate#set_string}
	SetString interface{} `field:"optional" json:"setString" yaml:"setString"`
	// Only show manifests rendered from the given templates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#show_only DataHelmTemplate#show_only}
	ShowOnly *[]*string `field:"optional" json:"showOnly" yaml:"showOnly"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#skip_crds DataHelmTemplate#skip_crds}
	SkipCrds interface{} `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// If set, tests will not be rendered. By default, tests are rendered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#skip_tests DataHelmTemplate#skip_tests}
	SkipTests interface{} `field:"optional" json:"skipTests" yaml:"skipTests"`
	// Time in seconds to wait for any individual kubernetes operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#timeout DataHelmTemplate#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Validate your manifests against the Kubernetes cluster you are currently pointing at.
	//
	// This is the same validation performed on an install
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#validate DataHelmTemplate#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
	// List of values in raw yaml format to pass to helm.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#values DataHelmTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
	// Verify the package before installing it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#verify DataHelmTemplate#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#version DataHelmTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Will wait until all resources are in a ready state before marking the release as successful.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#wait DataHelmTemplate#wait}
	Wait interface{} `field:"optional" json:"wait" yaml:"wait"`
}

type DataHelmTemplatePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#binary_path DataHelmTemplate#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
}

type DataHelmTemplatePostrenderOutputReference interface {
	cdktf.ComplexObject
	BinaryPath() *string
	SetBinaryPath(val *string)
	BinaryPathInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataHelmTemplatePostrender
	SetInternalValue(val *DataHelmTemplatePostrender)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplatePostrenderOutputReference
type jsiiProxy_DataHelmTemplatePostrenderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) BinaryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) BinaryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) InternalValue() *DataHelmTemplatePostrender {
	var returns *DataHelmTemplatePostrender
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataHelmTemplatePostrenderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataHelmTemplatePostrenderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplatePostrenderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplatePostrenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataHelmTemplatePostrenderOutputReference_Override(d DataHelmTemplatePostrenderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplatePostrenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetBinaryPath(val *string) {
	_jsii_.Set(
		j,
		"binaryPath",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetInternalValue(val *DataHelmTemplatePostrender) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplatePostrenderOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplatePostrenderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#type DataHelmTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type DataHelmTemplateSetList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataHelmTemplateSetOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetList
type jsiiProxy_DataHelmTemplateSetList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataHelmTemplateSetList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataHelmTemplateSetList {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetList{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetList_Override(d DataHelmTemplateSetList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetList) Get(index *float64) DataHelmTemplateSetOutputReference {
	var returns DataHelmTemplateSetOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSetOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetOutputReference
type jsiiProxy_DataHelmTemplateSetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataHelmTemplateSetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetOutputReference_Override(d DataHelmTemplateSetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSetSensitive struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#type DataHelmTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type DataHelmTemplateSetSensitiveList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataHelmTemplateSetSensitiveOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetSensitiveList
type jsiiProxy_DataHelmTemplateSetSensitiveList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetSensitiveList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataHelmTemplateSetSensitiveList {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetSensitiveList{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetSensitiveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetSensitiveList_Override(d DataHelmTemplateSetSensitiveList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetSensitiveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) Get(index *float64) DataHelmTemplateSetSensitiveOutputReference {
	var returns DataHelmTemplateSetSensitiveOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSetSensitiveOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetSensitiveOutputReference
type jsiiProxy_DataHelmTemplateSetSensitiveOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetSensitiveOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataHelmTemplateSetSensitiveOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetSensitiveOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetSensitiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetSensitiveOutputReference_Override(d DataHelmTemplateSetSensitiveOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetSensitiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetSensitiveOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSetString struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#name DataHelmTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/d/template#value DataHelmTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

type DataHelmTemplateSetStringList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataHelmTemplateSetStringOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetStringList
type jsiiProxy_DataHelmTemplateSetStringList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetStringList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataHelmTemplateSetStringList {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetStringList{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetStringList_Override(d DataHelmTemplateSetStringList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetStringList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetStringList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringList) Get(index *float64) DataHelmTemplateSetStringOutputReference {
	var returns DataHelmTemplateSetStringOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataHelmTemplateSetStringOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataHelmTemplateSetStringOutputReference
type jsiiProxy_DataHelmTemplateSetStringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDataHelmTemplateSetStringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataHelmTemplateSetStringOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataHelmTemplateSetStringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataHelmTemplateSetStringOutputReference_Override(d DataHelmTemplateSetStringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.DataHelmTemplateSetStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataHelmTemplateSetStringOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataHelmTemplateSetStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/helm helm}.
type HelmProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
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
	RegistryConfigPath() *string
	SetRegistryConfigPath(val *string)
	RegistryConfigPathInput() *string
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
	ResetDebug()
	ResetExperiments()
	ResetHelmDriver()
	ResetKubernetes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginsPath()
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


// Create a new {@link https://www.terraform.io/docs/providers/helm helm} Resource.
func NewHelmProvider(scope constructs.Construct, id *string, config *HelmProviderConfig) HelmProvider {
	_init_.Initialize()

	j := jsiiProxy_HelmProvider{}

	_jsii_.Create(
		"@cdktf/provider-helm.HelmProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/helm helm} Resource.
func NewHelmProvider_Override(h HelmProvider, scope constructs.Construct, id *string, config *HelmProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.HelmProvider",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HelmProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetDebug(val interface{}) {
	_jsii_.Set(
		j,
		"debug",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetExperiments(val *HelmProviderExperiments) {
	_jsii_.Set(
		j,
		"experiments",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetHelmDriver(val *string) {
	_jsii_.Set(
		j,
		"helmDriver",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetKubernetes(val *HelmProviderKubernetes) {
	_jsii_.Set(
		j,
		"kubernetes",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetPluginsPath(val *string) {
	_jsii_.Set(
		j,
		"pluginsPath",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetRegistryConfigPath(val *string) {
	_jsii_.Set(
		j,
		"registryConfigPath",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetRepositoryCache(val *string) {
	_jsii_.Set(
		j,
		"repositoryCache",
		val,
	)
}

func (j *jsiiProxy_HelmProvider) SetRepositoryConfigPath(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.HelmProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HelmProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.HelmProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HelmProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HelmProvider) OverrideLogicalId(newLogicalId *string) {
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

type HelmProviderExperiments struct {
	// Enable full diff by storing the rendered manifest in the state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#manifest HelmProvider#manifest}
	Manifest interface{} `field:"optional" json:"manifest" yaml:"manifest"`
}

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

type HelmProviderKubernetesExec struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#api_version HelmProvider#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#command HelmProvider#command}.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#args HelmProvider#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm#env HelmProvider#env}.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

// Represents a {@link https://www.terraform.io/docs/providers/helm/r/release helm_release}.
type Release interface {
	cdktf.TerraformResource
	Atomic() interface{}
	SetAtomic(val interface{})
	AtomicInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chart() *string
	SetChart(val *string)
	ChartInput() *string
	CleanupOnFail() interface{}
	SetCleanupOnFail(val interface{})
	CleanupOnFailInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	DisableCrdHooks() interface{}
	SetDisableCrdHooks(val interface{})
	DisableCrdHooksInput() interface{}
	DisableOpenapiValidation() interface{}
	SetDisableOpenapiValidation(val interface{})
	DisableOpenapiValidationInput() interface{}
	DisableWebhooks() interface{}
	SetDisableWebhooks(val interface{})
	DisableWebhooksInput() interface{}
	ForceUpdate() interface{}
	SetForceUpdate(val interface{})
	ForceUpdateInput() interface{}
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
	Keyring() *string
	SetKeyring(val *string)
	KeyringInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Lint() interface{}
	SetLint(val interface{})
	LintInput() interface{}
	Manifest() *string
	MaxHistory() *float64
	SetMaxHistory(val *float64)
	MaxHistoryInput() *float64
	Metadata() ReleaseMetadataList
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PassCredentials() interface{}
	SetPassCredentials(val interface{})
	PassCredentialsInput() interface{}
	Postrender() ReleasePostrenderOutputReference
	PostrenderInput() *ReleasePostrender
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecreatePods() interface{}
	SetRecreatePods(val interface{})
	RecreatePodsInput() interface{}
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
	Set() ReleaseSetList
	SetInput() interface{}
	SetSensitive() ReleaseSetSensitiveList
	SetSensitiveInput() interface{}
	SkipCrds() interface{}
	SetSkipCrds(val interface{})
	SkipCrdsInput() interface{}
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	WaitForJobs() interface{}
	SetWaitForJobs(val interface{})
	WaitForJobsInput() interface{}
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
	PutPostrender(value *ReleasePostrender)
	PutSet(value interface{})
	PutSetSensitive(value interface{})
	ResetAtomic()
	ResetCleanupOnFail()
	ResetCreateNamespace()
	ResetDependencyUpdate()
	ResetDescription()
	ResetDevel()
	ResetDisableCrdHooks()
	ResetDisableOpenapiValidation()
	ResetDisableWebhooks()
	ResetForceUpdate()
	ResetId()
	ResetKeyring()
	ResetLint()
	ResetMaxHistory()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassCredentials()
	ResetPostrender()
	ResetRecreatePods()
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
	ResetSetSensitive()
	ResetSkipCrds()
	ResetTfValues()
	ResetTimeout()
	ResetVerify()
	ResetVersion()
	ResetWait()
	ResetWaitForJobs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Release
type jsiiProxy_Release struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Release) Atomic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) AtomicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atomicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Chart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ChartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CleanupOnFail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CleanupOnFailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupOnFailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CreateNamespace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) CreateNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependencyUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependencyUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Devel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"develInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableCrdHooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCrdHooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableCrdHooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCrdHooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableOpenapiValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableOpenapiValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableOpenapiValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableWebhooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) DisableWebhooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableWebhooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForceUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForceUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Keyring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) KeyringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Lint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) LintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Manifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) MaxHistory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) MaxHistoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Metadata() ReleaseMetadataList {
	var returns ReleaseMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PassCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PassCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Postrender() ReleasePostrenderOutputReference {
	var returns ReleasePostrenderOutputReference
	_jsii_.Get(
		j,
		"postrender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) PostrenderInput() *ReleasePostrender {
	var returns *ReleasePostrender
	_jsii_.Get(
		j,
		"postrenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RecreatePods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatePods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RecreatePodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recreatePodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RenderSubchartNotes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RenderSubchartNotesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renderSubchartNotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Replace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCaFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) RepositoryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ResetValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ResetValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReuseValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ReuseValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Set() ReleaseSetList {
	var returns ReleaseSetList
	_jsii_.Get(
		j,
		"set",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetSensitive() ReleaseSetSensitiveList {
	var returns ReleaseSetSensitiveList
	_jsii_.Get(
		j,
		"setSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SetSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SkipCrds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) SkipCrdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCrdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) Wait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitForJobs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitForJobsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForJobsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Release) WaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/helm/r/release helm_release} Resource.
func NewRelease(scope constructs.Construct, id *string, config *ReleaseConfig) Release {
	_init_.Initialize()

	j := jsiiProxy_Release{}

	_jsii_.Create(
		"@cdktf/provider-helm.Release",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/helm/r/release helm_release} Resource.
func NewRelease_Override(r Release, scope constructs.Construct, id *string, config *ReleaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.Release",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Release) SetAtomic(val interface{}) {
	_jsii_.Set(
		j,
		"atomic",
		val,
	)
}

func (j *jsiiProxy_Release) SetChart(val *string) {
	_jsii_.Set(
		j,
		"chart",
		val,
	)
}

func (j *jsiiProxy_Release) SetCleanupOnFail(val interface{}) {
	_jsii_.Set(
		j,
		"cleanupOnFail",
		val,
	)
}

func (j *jsiiProxy_Release) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Release) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Release) SetCreateNamespace(val interface{}) {
	_jsii_.Set(
		j,
		"createNamespace",
		val,
	)
}

func (j *jsiiProxy_Release) SetDependencyUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"dependencyUpdate",
		val,
	)
}

func (j *jsiiProxy_Release) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Release) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Release) SetDevel(val interface{}) {
	_jsii_.Set(
		j,
		"devel",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableCrdHooks(val interface{}) {
	_jsii_.Set(
		j,
		"disableCrdHooks",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableOpenapiValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableOpenapiValidation",
		val,
	)
}

func (j *jsiiProxy_Release) SetDisableWebhooks(val interface{}) {
	_jsii_.Set(
		j,
		"disableWebhooks",
		val,
	)
}

func (j *jsiiProxy_Release) SetForceUpdate(val interface{}) {
	_jsii_.Set(
		j,
		"forceUpdate",
		val,
	)
}

func (j *jsiiProxy_Release) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Release) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Release) SetKeyring(val *string) {
	_jsii_.Set(
		j,
		"keyring",
		val,
	)
}

func (j *jsiiProxy_Release) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Release) SetLint(val interface{}) {
	_jsii_.Set(
		j,
		"lint",
		val,
	)
}

func (j *jsiiProxy_Release) SetMaxHistory(val *float64) {
	_jsii_.Set(
		j,
		"maxHistory",
		val,
	)
}

func (j *jsiiProxy_Release) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Release) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_Release) SetPassCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"passCredentials",
		val,
	)
}

func (j *jsiiProxy_Release) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Release) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Release) SetRecreatePods(val interface{}) {
	_jsii_.Set(
		j,
		"recreatePods",
		val,
	)
}

func (j *jsiiProxy_Release) SetRenderSubchartNotes(val interface{}) {
	_jsii_.Set(
		j,
		"renderSubchartNotes",
		val,
	)
}

func (j *jsiiProxy_Release) SetReplace(val interface{}) {
	_jsii_.Set(
		j,
		"replace",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryCaFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCaFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryCertFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryCertFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryKeyFile(val *string) {
	_jsii_.Set(
		j,
		"repositoryKeyFile",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryPassword(val *string) {
	_jsii_.Set(
		j,
		"repositoryPassword",
		val,
	)
}

func (j *jsiiProxy_Release) SetRepositoryUsername(val *string) {
	_jsii_.Set(
		j,
		"repositoryUsername",
		val,
	)
}

func (j *jsiiProxy_Release) SetResetValues(val interface{}) {
	_jsii_.Set(
		j,
		"resetValues",
		val,
	)
}

func (j *jsiiProxy_Release) SetReuseValues(val interface{}) {
	_jsii_.Set(
		j,
		"reuseValues",
		val,
	)
}

func (j *jsiiProxy_Release) SetSkipCrds(val interface{}) {
	_jsii_.Set(
		j,
		"skipCrds",
		val,
	)
}

func (j *jsiiProxy_Release) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_Release) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (j *jsiiProxy_Release) SetVerify(val interface{}) {
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_Release) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_Release) SetWait(val interface{}) {
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (j *jsiiProxy_Release) SetWaitForJobs(val interface{}) {
	_jsii_.Set(
		j,
		"waitForJobs",
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
func Release_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-helm.Release",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Release_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-helm.Release",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Release) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Release) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Release) PutPostrender(value *ReleasePostrender) {
	_jsii_.InvokeVoid(
		r,
		"putPostrender",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) PutSet(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putSet",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) PutSetSensitive(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putSetSensitive",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Release) ResetAtomic() {
	_jsii_.InvokeVoid(
		r,
		"resetAtomic",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetCleanupOnFail() {
	_jsii_.InvokeVoid(
		r,
		"resetCleanupOnFail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetCreateNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetCreateNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDependencyUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetDependencyUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDevel() {
	_jsii_.InvokeVoid(
		r,
		"resetDevel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableCrdHooks() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableCrdHooks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableOpenapiValidation() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableOpenapiValidation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetDisableWebhooks() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableWebhooks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetForceUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetForceUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetKeyring() {
	_jsii_.InvokeVoid(
		r,
		"resetKeyring",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetLint() {
	_jsii_.InvokeVoid(
		r,
		"resetLint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetMaxHistory() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxHistory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetPassCredentials() {
	_jsii_.InvokeVoid(
		r,
		"resetPassCredentials",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetPostrender() {
	_jsii_.InvokeVoid(
		r,
		"resetPostrender",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRecreatePods() {
	_jsii_.InvokeVoid(
		r,
		"resetRecreatePods",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRenderSubchartNotes() {
	_jsii_.InvokeVoid(
		r,
		"resetRenderSubchartNotes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetReplace() {
	_jsii_.InvokeVoid(
		r,
		"resetReplace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepository() {
	_jsii_.InvokeVoid(
		r,
		"resetRepository",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryCaFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryCaFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryCertFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryCertFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryKeyFile() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryKeyFile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetRepositoryUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetRepositoryUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetResetValues() {
	_jsii_.InvokeVoid(
		r,
		"resetResetValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetReuseValues() {
	_jsii_.InvokeVoid(
		r,
		"resetReuseValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSet() {
	_jsii_.InvokeVoid(
		r,
		"resetSet",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSetSensitive() {
	_jsii_.InvokeVoid(
		r,
		"resetSetSensitive",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetSkipCrds() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipCrds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetTfValues() {
	_jsii_.InvokeVoid(
		r,
		"resetTfValues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetVerify() {
	_jsii_.InvokeVoid(
		r,
		"resetVerify",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetWait() {
	_jsii_.InvokeVoid(
		r,
		"resetWait",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) ResetWaitForJobs() {
	_jsii_.InvokeVoid(
		r,
		"resetWaitForJobs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Release) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Release) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

type ReleaseMetadata struct {
}

type ReleaseMetadataList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ReleaseMetadataOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseMetadataList
type jsiiProxy_ReleaseMetadataList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ReleaseMetadataList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewReleaseMetadataList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ReleaseMetadataList {
	_init_.Initialize()

	j := jsiiProxy_ReleaseMetadataList{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseMetadataList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewReleaseMetadataList_Override(r ReleaseMetadataList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseMetadataList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseMetadataList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_ReleaseMetadataList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataList) Get(index *float64) ReleaseMetadataOutputReference {
	var returns ReleaseMetadataOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleaseMetadataOutputReference interface {
	cdktf.ComplexObject
	AppVersion() *string
	Chart() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ReleaseMetadata
	SetInternalValue(val *ReleaseMetadata)
	Name() *string
	Namespace() *string
	Revision() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *string
	Version() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseMetadataOutputReference
type jsiiProxy_ReleaseMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) AppVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Chart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) InternalValue() *ReleaseMetadata {
	var returns *ReleaseMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Values() *string {
	var returns *string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewReleaseMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReleaseMetadataOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ReleaseMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReleaseMetadataOutputReference_Override(r ReleaseMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) SetInternalValue(val *ReleaseMetadata) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseMetadataOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleasePostrender struct {
	// The command binary path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#binary_path Release#binary_path}
	BinaryPath *string `field:"required" json:"binaryPath" yaml:"binaryPath"`
	// an argument to the post-renderer (can specify multiple).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#args Release#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
}

type ReleasePostrenderOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	BinaryPath() *string
	SetBinaryPath(val *string)
	BinaryPathInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ReleasePostrender
	SetInternalValue(val *ReleasePostrender)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetArgs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleasePostrenderOutputReference
type jsiiProxy_ReleasePostrenderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) BinaryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) BinaryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) InternalValue() *ReleasePostrender {
	var returns *ReleasePostrender
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewReleasePostrenderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReleasePostrenderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ReleasePostrenderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleasePostrenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReleasePostrenderOutputReference_Override(r ReleasePostrenderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleasePostrenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetArgs(val *[]*string) {
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetBinaryPath(val *string) {
	_jsii_.Set(
		j,
		"binaryPath",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetInternalValue(val *ReleasePostrender) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleasePostrenderOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		r,
		"resetArgs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleasePostrenderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleaseSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#name Release#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#value Release#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#type Release#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type ReleaseSetList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ReleaseSetOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseSetList
type jsiiProxy_ReleaseSetList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ReleaseSetList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewReleaseSetList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ReleaseSetList {
	_init_.Initialize()

	j := jsiiProxy_ReleaseSetList{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewReleaseSetList_Override(r ReleaseSetList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseSetList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_ReleaseSetList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetList) Get(index *float64) ReleaseSetOutputReference {
	var returns ReleaseSetOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleaseSetOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseSetOutputReference
type jsiiProxy_ReleaseSetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReleaseSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewReleaseSetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReleaseSetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ReleaseSetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReleaseSetOutputReference_Override(r ReleaseSetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (r *jsiiProxy_ReleaseSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		r,
		"resetType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReleaseSetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleaseSetSensitive struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#name Release#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#value Release#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/helm/r/release#type Release#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type ReleaseSetSensitiveList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ReleaseSetSensitiveOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseSetSensitiveList
type jsiiProxy_ReleaseSetSensitiveList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ReleaseSetSensitiveList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewReleaseSetSensitiveList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ReleaseSetSensitiveList {
	_init_.Initialize()

	j := jsiiProxy_ReleaseSetSensitiveList{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetSensitiveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewReleaseSetSensitiveList_Override(r ReleaseSetSensitiveList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetSensitiveList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_ReleaseSetSensitiveList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveList) Get(index *float64) ReleaseSetSensitiveOutputReference {
	var returns ReleaseSetSensitiveOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ReleaseSetSensitiveOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReleaseSetSensitiveOutputReference
type jsiiProxy_ReleaseSetSensitiveOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewReleaseSetSensitiveOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ReleaseSetSensitiveOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ReleaseSetSensitiveOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetSensitiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewReleaseSetSensitiveOutputReference_Override(r ReleaseSetSensitiveOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-helm.ReleaseSetSensitiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ReleaseSetSensitiveOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		r,
		"resetType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReleaseSetSensitiveOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

