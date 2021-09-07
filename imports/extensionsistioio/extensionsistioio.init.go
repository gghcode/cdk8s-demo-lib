package extensionsistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"extensionsistioio.WasmPlugin",
		reflect.TypeOf((*WasmPlugin)(nil)).Elem(),
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
			j := jsiiProxy_WasmPlugin{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"extensionsistioio.WasmPluginProps",
		reflect.TypeOf((*WasmPluginProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"extensionsistioio.WasmPluginSpec",
		reflect.TypeOf((*WasmPluginSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"extensionsistioio.WasmPluginSpecImagePullPolicy",
		reflect.TypeOf((*WasmPluginSpecImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED_POLICY": WasmPluginSpecImagePullPolicy_UNSPECIFIED_POLICY,
			"IF_NOT_PRESENT": WasmPluginSpecImagePullPolicy_IF_NOT_PRESENT,
			"ALWAYS": WasmPluginSpecImagePullPolicy_ALWAYS,
		},
	)
	_jsii_.RegisterEnum(
		"extensionsistioio.WasmPluginSpecPhase",
		reflect.TypeOf((*WasmPluginSpecPhase)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED_PHASE": WasmPluginSpecPhase_UNSPECIFIED_PHASE,
			"AUTHN": WasmPluginSpecPhase_AUTHN,
			"AUTHZ": WasmPluginSpecPhase_AUTHZ,
			"STATS": WasmPluginSpecPhase_STATS,
		},
	)
	_jsii_.RegisterStruct(
		"extensionsistioio.WasmPluginSpecSelector",
		reflect.TypeOf((*WasmPluginSpecSelector)(nil)).Elem(),
	)
}
