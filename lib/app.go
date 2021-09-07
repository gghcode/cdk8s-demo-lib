package lib

import (
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
	"github.com/gghcode/cdk8s-demo-lib/imports/k8s"
)

type RestAppChartProps struct {
	cdk8s.ChartProps

	Image    string
	Name     string
	Port     float64
	Replicas float64

	HealthCheckPath string
}

func NewRestAppChart(scope constructs.Construct, id string, props *RestAppChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	metadataLabels := &map[string]*string{
		"app": jsii.String(props.Name),
	}

	// define resources here
	k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
		Metadata: &k8s.ObjectMeta{
			Name:   jsii.String(props.Name),
			Labels: metadataLabels,
		},
		Spec: &k8s.DeploymentSpec{
			Replicas: jsii.Number(props.Replicas),
			Selector: &k8s.LabelSelector{
				MatchLabels: metadataLabels,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: metadataLabels,
				},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{
						{
							Name:  jsii.String("main"),
							Image: jsii.String(props.Image),
							Ports: &[]*k8s.ContainerPort{
								{
									ContainerPort: &props.Port,
								},
							},
							LivenessProbe: &k8s.Probe{
								HttpGet: &k8s.HttpGetAction{
									Port: k8s.IntOrString_FromNumber(jsii.Number(props.Port)),
									Path: jsii.String(props.HealthCheckPath),
								},
							},
							ReadinessProbe: &k8s.Probe{
								HttpGet: &k8s.HttpGetAction{
									Port: k8s.IntOrString_FromNumber(jsii.Number(props.Port)),
									Path: jsii.String(props.HealthCheckPath),
								},
							},
							Resources: &k8s.ResourceRequirements{
								Limits: &map[string]k8s.Quantity{
									"cpu":    k8s.Quantity_FromNumber(jsii.Number(1)),
									"memory": k8s.Quantity_FromString(jsii.String("1Gi")),
								},
							},
						},
					},
				},
			},
		},
	})

	k8s.NewKubeService(chart, jsii.String(props.Name), &k8s.KubeServiceProps{
		Metadata: &k8s.ObjectMeta{
			Name: &props.Name,
		},
		Spec: &k8s.ServiceSpec{
			Selector: metadataLabels,
			Ports: &[]*k8s.ServicePort{
				{
					Port:       &props.Port,
					TargetPort: k8s.IntOrString_FromNumber(&props.Port),
				},
			},
		},
	})

	return chart
}
