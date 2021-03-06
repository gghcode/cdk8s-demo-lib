// telemetryistioio
package telemetryistioio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/gghcode/cdk8s-demo-lib/imports/telemetryistioio/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/gghcode/cdk8s-demo-lib/imports/telemetryistioio/internal"
)

type Telemetry interface {
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

// The jsii proxy struct for Telemetry
type jsiiProxy_Telemetry struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Telemetry) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Telemetry) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Telemetry) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Telemetry) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Telemetry) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Telemetry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines a "Telemetry" API object.
func NewTelemetry(scope constructs.Construct, id *string, props *TelemetryProps) Telemetry {
	_init_.Initialize()

	j := jsiiProxy_Telemetry{}

	_jsii_.Create(
		"telemetryistioio.Telemetry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Telemetry" API object.
func NewTelemetry_Override(t Telemetry, scope constructs.Construct, id *string, props *TelemetryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"telemetryistioio.Telemetry",
		[]interface{}{scope, id, props},
		t,
	)
}

// Renders a Kubernetes manifest for "Telemetry".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Telemetry_Manifest(props *TelemetryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"telemetryistioio.Telemetry",
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
func Telemetry_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"telemetryistioio.Telemetry",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Telemetry_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"telemetryistioio.Telemetry",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_Telemetry) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_Telemetry) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
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
func (t *jsiiProxy_Telemetry) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (t *jsiiProxy_Telemetry) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
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
func (t *jsiiProxy_Telemetry) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
func (t *jsiiProxy_Telemetry) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_Telemetry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TelemetryProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	// Telemetry configuration for workloads.
	//
	// See more details at: https://istio.io/docs/reference/config/telemetry/telemetry.html
	Spec *TelemetrySpec `json:"spec"`
}

// Telemetry configuration for workloads.
//
// See more details at: https://istio.io/docs/reference/config/telemetry/telemetry.html
type TelemetrySpec struct {
	// Optional.
	AccessLogging *[]*TelemetrySpecAccessLogging `json:"accessLogging"`
	// Optional.
	Metrics *[]*TelemetrySpecMetrics `json:"metrics"`
	// Optional.
	Selector *TelemetrySpecSelector `json:"selector"`
	// Optional.
	Tracing *[]*TelemetrySpecTracing `json:"tracing"`
}

type TelemetrySpecAccessLogging struct {
	// Controls logging.
	Disabled *bool `json:"disabled"`
	// Optional.
	Providers *[]*TelemetrySpecAccessLoggingProviders `json:"providers"`
}

type TelemetrySpecAccessLoggingProviders struct {
	// Required.
	Name *string `json:"name"`
}

type TelemetrySpecMetrics struct {
	// Optional.
	Overrides *[]*TelemetrySpecMetricsOverrides `json:"overrides"`
	// Optional.
	Providers *[]*TelemetrySpecMetricsProviders `json:"providers"`
}

type TelemetrySpecMetricsOverrides struct {
	// Optional.
	Disabled *bool `json:"disabled"`
	// Match allows provides the scope of the override.
	Match *TelemetrySpecMetricsOverridesMatch `json:"match"`
	// Optional.
	TagOverrides *map[string]*TelemetrySpecMetricsOverridesTagOverrides `json:"tagOverrides"`
}

// Match allows provides the scope of the override.
type TelemetrySpecMetricsOverridesMatch struct {
	// Allows free-form specification of a metric.
	CustomMetric *string `json:"customMetric"`
	// One of the well-known Istio Standard Metrics.
	Metric TelemetrySpecMetricsOverridesMatchMetric `json:"metric"`
	// Controls which mode of metrics generation is selected: CLIENT and/or SERVER.
	Mode TelemetrySpecMetricsOverridesMatchMode `json:"mode"`
}

// One of the well-known Istio Standard Metrics.
type TelemetrySpecMetricsOverridesMatchMetric string

const (
	TelemetrySpecMetricsOverridesMatchMetric_ALL_METRICS TelemetrySpecMetricsOverridesMatchMetric = "ALL_METRICS"
	TelemetrySpecMetricsOverridesMatchMetric_REQUEST_COUNT TelemetrySpecMetricsOverridesMatchMetric = "REQUEST_COUNT"
	TelemetrySpecMetricsOverridesMatchMetric_REQUEST_DURATION TelemetrySpecMetricsOverridesMatchMetric = "REQUEST_DURATION"
	TelemetrySpecMetricsOverridesMatchMetric_REQUEST_SIZE TelemetrySpecMetricsOverridesMatchMetric = "REQUEST_SIZE"
	TelemetrySpecMetricsOverridesMatchMetric_RESPONSE_SIZE TelemetrySpecMetricsOverridesMatchMetric = "RESPONSE_SIZE"
	TelemetrySpecMetricsOverridesMatchMetric_TCP_OPENED_CONNECTIONS TelemetrySpecMetricsOverridesMatchMetric = "TCP_OPENED_CONNECTIONS"
	TelemetrySpecMetricsOverridesMatchMetric_TCP_CLOSED_CONNECTIONS TelemetrySpecMetricsOverridesMatchMetric = "TCP_CLOSED_CONNECTIONS"
	TelemetrySpecMetricsOverridesMatchMetric_TCP_SENT_BYTES TelemetrySpecMetricsOverridesMatchMetric = "TCP_SENT_BYTES"
	TelemetrySpecMetricsOverridesMatchMetric_TCP_RECEIVED_BYTES TelemetrySpecMetricsOverridesMatchMetric = "TCP_RECEIVED_BYTES"
	TelemetrySpecMetricsOverridesMatchMetric_GRPC_REQUEST_MESSAGES TelemetrySpecMetricsOverridesMatchMetric = "GRPC_REQUEST_MESSAGES"
	TelemetrySpecMetricsOverridesMatchMetric_GRPC_RESPONSE_MESSAGES TelemetrySpecMetricsOverridesMatchMetric = "GRPC_RESPONSE_MESSAGES"
)

// Controls which mode of metrics generation is selected: CLIENT and/or SERVER.
type TelemetrySpecMetricsOverridesMatchMode string

const (
	TelemetrySpecMetricsOverridesMatchMode_CLIENT_AND_SERVER TelemetrySpecMetricsOverridesMatchMode = "CLIENT_AND_SERVER"
	TelemetrySpecMetricsOverridesMatchMode_CLIENT TelemetrySpecMetricsOverridesMatchMode = "CLIENT"
	TelemetrySpecMetricsOverridesMatchMode_SERVER TelemetrySpecMetricsOverridesMatchMode = "SERVER"
)

type TelemetrySpecMetricsOverridesTagOverrides struct {
	// Operation controls whether or not to update/add a tag, or to remove it.
	Operation TelemetrySpecMetricsOverridesTagOverridesOperation `json:"operation"`
	// Value is only considered if the operation is `UPSERT`.
	Value *string `json:"value"`
}

// Operation controls whether or not to update/add a tag, or to remove it.
type TelemetrySpecMetricsOverridesTagOverridesOperation string

const (
	TelemetrySpecMetricsOverridesTagOverridesOperation_UPSERT TelemetrySpecMetricsOverridesTagOverridesOperation = "UPSERT"
	TelemetrySpecMetricsOverridesTagOverridesOperation_REMOVE TelemetrySpecMetricsOverridesTagOverridesOperation = "REMOVE"
)

type TelemetrySpecMetricsProviders struct {
	// Required.
	Name *string `json:"name"`
}

// Optional.
type TelemetrySpecSelector struct {
	MatchLabels *map[string]*string `json:"matchLabels"`
}

type TelemetrySpecTracing struct {
	// Optional.
	CustomTags *map[string]*TelemetrySpecTracingCustomTags `json:"customTags"`
	// Controls span reporting.
	DisableSpanReporting *bool `json:"disableSpanReporting"`
	// Optional.
	Providers *[]*TelemetrySpecTracingProviders `json:"providers"`
	RandomSamplingPercentage *float64 `json:"randomSamplingPercentage"`
}

type TelemetrySpecTracingCustomTags struct {
	// Environment adds the value of an environment variable to each span.
	Environment *TelemetrySpecTracingCustomTagsEnvironment `json:"environment"`
	// RequestHeader adds the value of an header from the request to each span.
	Header *TelemetrySpecTracingCustomTagsHeader `json:"header"`
	// Literal adds the same, hard-coded value to each span.
	Literal *TelemetrySpecTracingCustomTagsLiteral `json:"literal"`
}

// Environment adds the value of an environment variable to each span.
type TelemetrySpecTracingCustomTagsEnvironment struct {
	// Optional.
	DefaultValue *string `json:"defaultValue"`
	// Name of the environment variable from which to extract the tag value.
	Name *string `json:"name"`
}

// RequestHeader adds the value of an header from the request to each span.
type TelemetrySpecTracingCustomTagsHeader struct {
	// Optional.
	DefaultValue *string `json:"defaultValue"`
	// Name of the header from which to extract the tag value.
	Name *string `json:"name"`
}

// Literal adds the same, hard-coded value to each span.
type TelemetrySpecTracingCustomTagsLiteral struct {
	// The tag value to use.
	Value *string `json:"value"`
}

type TelemetrySpecTracingProviders struct {
	// Required.
	Name *string `json:"name"`
}

