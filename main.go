package main

import (
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/gghcode/cdk8s-demo-lib/imports/k8s"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	k8s.NewKubeNamespace(chart, jsii.String("hello-k8s"), &k8s.KubeNamespaceProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("Hello-world"),
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	// lib.NewRestAppChart(app, "app", &lib.RestAppChartProps{
	// 	Name:            "nginx",
	// 	Image:           "nginx",
	// 	Replicas:        3,
	// 	Port:            80,
	// 	HealthCheckPath: "/",
	// })
	// NewMyChart(app, "go-cdk8s-demo", nil)
	app.Synth()
}
