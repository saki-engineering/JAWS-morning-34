package main

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"

	"github.com/aws/constructs-go/constructs/v10"
)

var (
	projectName       string
	projectNameAlias  string
	resourceStackName string
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	lambda.NewFunction(stack, jsii.String("JAWSMorning34Lambda"), &lambda.FunctionProps{
		FunctionName: jsii.String(strings.Join([]string{projectName, "func"}, "-")),
		Runtime:      lambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         lambda.AssetCode_FromAsset(jsii.String("../src"), &awss3assets.AssetOptions{}),
		LogRetention: awslogs.RetentionDays_ONE_WEEK,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	projectName = app.Node().TryGetContext(jsii.String("projectName")).(string)
	projectNameAlias = app.Node().TryGetContext(jsii.String("projectNameAlias")).(string)
	resourceStackName = app.Node().TryGetContext(jsii.String("resourceStackName")).(string)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			StackName: jsii.String(resourceStackName),
			Synthesizer: awscdk.NewDefaultStackSynthesizer(&awscdk.DefaultStackSynthesizerProps{
				Qualifier: jsii.String(projectNameAlias),
			}),
			Tags: &map[string]*string{
				"CreatedBy": jsii.String(resourceStackName),
			},
		},
	})

	app.Synth(nil)
}
