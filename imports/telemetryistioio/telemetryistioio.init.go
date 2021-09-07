package telemetryistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"telemetryistioio.Telemetry",
		reflect.TypeOf((*Telemetry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Telemetry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetryProps",
		reflect.TypeOf((*TelemetryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpec",
		reflect.TypeOf((*TelemetrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecAccessLogging",
		reflect.TypeOf((*TelemetrySpecAccessLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecAccessLoggingProviders",
		reflect.TypeOf((*TelemetrySpecAccessLoggingProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecMetrics",
		reflect.TypeOf((*TelemetrySpecMetrics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecMetricsOverrides",
		reflect.TypeOf((*TelemetrySpecMetricsOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecMetricsOverridesMatch",
		reflect.TypeOf((*TelemetrySpecMetricsOverridesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"telemetryistioio.TelemetrySpecMetricsOverridesMatchMetric",
		reflect.TypeOf((*TelemetrySpecMetricsOverridesMatchMetric)(nil)).Elem(),
		map[string]interface{}{
			"ALL_METRICS": TelemetrySpecMetricsOverridesMatchMetric_ALL_METRICS,
			"REQUEST_COUNT": TelemetrySpecMetricsOverridesMatchMetric_REQUEST_COUNT,
			"REQUEST_DURATION": TelemetrySpecMetricsOverridesMatchMetric_REQUEST_DURATION,
			"REQUEST_SIZE": TelemetrySpecMetricsOverridesMatchMetric_REQUEST_SIZE,
			"RESPONSE_SIZE": TelemetrySpecMetricsOverridesMatchMetric_RESPONSE_SIZE,
			"TCP_OPENED_CONNECTIONS": TelemetrySpecMetricsOverridesMatchMetric_TCP_OPENED_CONNECTIONS,
			"TCP_CLOSED_CONNECTIONS": TelemetrySpecMetricsOverridesMatchMetric_TCP_CLOSED_CONNECTIONS,
			"TCP_SENT_BYTES": TelemetrySpecMetricsOverridesMatchMetric_TCP_SENT_BYTES,
			"TCP_RECEIVED_BYTES": TelemetrySpecMetricsOverridesMatchMetric_TCP_RECEIVED_BYTES,
			"GRPC_REQUEST_MESSAGES": TelemetrySpecMetricsOverridesMatchMetric_GRPC_REQUEST_MESSAGES,
			"GRPC_RESPONSE_MESSAGES": TelemetrySpecMetricsOverridesMatchMetric_GRPC_RESPONSE_MESSAGES,
		},
	)
	_jsii_.RegisterEnum(
		"telemetryistioio.TelemetrySpecMetricsOverridesMatchMode",
		reflect.TypeOf((*TelemetrySpecMetricsOverridesMatchMode)(nil)).Elem(),
		map[string]interface{}{
			"CLIENT_AND_SERVER": TelemetrySpecMetricsOverridesMatchMode_CLIENT_AND_SERVER,
			"CLIENT": TelemetrySpecMetricsOverridesMatchMode_CLIENT,
			"SERVER": TelemetrySpecMetricsOverridesMatchMode_SERVER,
		},
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecMetricsOverridesTagOverrides",
		reflect.TypeOf((*TelemetrySpecMetricsOverridesTagOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"telemetryistioio.TelemetrySpecMetricsOverridesTagOverridesOperation",
		reflect.TypeOf((*TelemetrySpecMetricsOverridesTagOverridesOperation)(nil)).Elem(),
		map[string]interface{}{
			"UPSERT": TelemetrySpecMetricsOverridesTagOverridesOperation_UPSERT,
			"REMOVE": TelemetrySpecMetricsOverridesTagOverridesOperation_REMOVE,
		},
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecMetricsProviders",
		reflect.TypeOf((*TelemetrySpecMetricsProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecSelector",
		reflect.TypeOf((*TelemetrySpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracing",
		reflect.TypeOf((*TelemetrySpecTracing)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTags",
		reflect.TypeOf((*TelemetrySpecTracingCustomTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsEnvironment",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsHeader",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingCustomTagsLiteral",
		reflect.TypeOf((*TelemetrySpecTracingCustomTagsLiteral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"telemetryistioio.TelemetrySpecTracingProviders",
		reflect.TypeOf((*TelemetrySpecTracingProviders)(nil)).Elem(),
	)
}
