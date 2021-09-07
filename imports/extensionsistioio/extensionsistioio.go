// extensionsistioio
package extensionsistioio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/gghcode/cdk8s-demo-lib/imports/extensionsistioio/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/gghcode/cdk8s-demo-lib/imports/extensionsistioio/internal"
)

type WasmPlugin interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for WasmPlugin
type jsiiProxy_WasmPlugin struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_WasmPlugin) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WasmPlugin) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WasmPlugin) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WasmPlugin) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WasmPlugin) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WasmPlugin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "WasmPlugin" API object.
func NewWasmPlugin(scope constructs.Construct, id *string, props *WasmPluginProps) WasmPlugin {
	_init_.Initialize()

	j := jsiiProxy_WasmPlugin{}

	_jsii_.Create(
		"extensionsistioio.WasmPlugin",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "WasmPlugin" API object.
func NewWasmPlugin_Override(w WasmPlugin, scope constructs.Construct, id *string, props *WasmPluginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"extensionsistioio.WasmPlugin",
		[]interface{}{scope, id, props},
		w,
	)
}

// Renders a Kubernetes manifest for "WasmPlugin".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func WasmPlugin_Manifest(props *WasmPluginProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"extensionsistioio.WasmPlugin",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func WasmPlugin_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"extensionsistioio.WasmPlugin",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func WasmPlugin_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"extensionsistioio.WasmPlugin",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (w *jsiiProxy_WasmPlugin) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (w *jsiiProxy_WasmPlugin) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
func (w *jsiiProxy_WasmPlugin) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (w *jsiiProxy_WasmPlugin) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (w *jsiiProxy_WasmPlugin) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (w *jsiiProxy_WasmPlugin) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WasmPlugin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type WasmPluginProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec *WasmPluginSpec `json:"spec"`
}

type WasmPluginSpec struct {
	// The pull behaviour to be applied when fetching an OCI image.
	ImagePullPolicy WasmPluginSpecImagePullPolicy `json:"imagePullPolicy"`
	// Credentials to use for OCI image pulling.
	ImagePullSecret *string `json:"imagePullSecret"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase WasmPluginSpecPhase `json:"phase"`
	// The configuration that will be passed on to the plugin.
	PluginConfig interface{} `json:"pluginConfig"`
	PluginName *string `json:"pluginName"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	Priority *float64 `json:"priority"`
	Selector *WasmPluginSpecSelector `json:"selector"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	Sha256 *string `json:"sha256"`
	// URL of a Wasm module or OCI container.
	Url *string `json:"url"`
	VerificationKey *string `json:"verificationKey"`
}

// The pull behaviour to be applied when fetching an OCI image.
type WasmPluginSpecImagePullPolicy string

const (
	WasmPluginSpecImagePullPolicy_UNSPECIFIED_POLICY WasmPluginSpecImagePullPolicy = "UNSPECIFIED_POLICY"
	WasmPluginSpecImagePullPolicy_IF_NOT_PRESENT WasmPluginSpecImagePullPolicy = "IF_NOT_PRESENT"
	WasmPluginSpecImagePullPolicy_ALWAYS WasmPluginSpecImagePullPolicy = "ALWAYS"
)

// Determines where in the filter chain this `WasmPlugin` is to be injected.
type WasmPluginSpecPhase string

const (
	WasmPluginSpecPhase_UNSPECIFIED_PHASE WasmPluginSpecPhase = "UNSPECIFIED_PHASE"
	WasmPluginSpecPhase_AUTHN WasmPluginSpecPhase = "AUTHN"
	WasmPluginSpecPhase_AUTHZ WasmPluginSpecPhase = "AUTHZ"
	WasmPluginSpecPhase_STATS WasmPluginSpecPhase = "STATS"
)

type WasmPluginSpecSelector struct {
	MatchLabels *map[string]*string `json:"matchLabels"`
}

