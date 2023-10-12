package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for a newly created singleton Lambda.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var adotLayerVersion adotLayerVersion
//   var architecture architecture
//   var code code
//   var codeSigningConfig codeSigningConfig
//   var destination iDestination
//   var eventSource iEventSource
//   var fileSystem fileSystem
//   var key key
//   var lambdaInsightsVersion lambdaInsightsVersion
//   var layerVersion layerVersion
//   var paramsAndSecretsLayerVersion paramsAndSecretsLayerVersion
//   var policyStatement policyStatement
//   var profilingGroup profilingGroup
//   var queue queue
//   var role role
//   var runtime runtime
//   var runtimeManagementMode runtimeManagementMode
//   var securityGroup securityGroup
//   var size size
//   var snapStartConf snapStartConf
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var topic topic
//   var vpc vpc
//
//   singletonFunctionProps := &SingletonFunctionProps{
//   	Code: code,
//   	Handler: jsii.String("handler"),
//   	Runtime: runtime,
//   	Uuid: jsii.String("uuid"),
//
//   	// the properties below are optional
//   	AdotInstrumentation: &AdotInstrumentationConfig{
//   		ExecWrapper: awscdk.Aws_lambda.AdotLambdaExecWrapper_REGULAR_HANDLER,
//   		LayerVersion: adotLayerVersion,
//   	},
//   	AllowAllOutbound: jsii.Boolean(false),
//   	AllowPublicSubnet: jsii.Boolean(false),
//   	Architecture: architecture,
//   	CodeSigningConfig: codeSigningConfig,
//   	CurrentVersionOptions: &VersionOptions{
//   		CodeSha256: jsii.String("codeSha256"),
//   		Description: jsii.String("description"),
//   		MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   		OnFailure: destination,
//   		OnSuccess: destination,
//   		ProvisionedConcurrentExecutions: jsii.Number(123),
//   		RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   		RetryAttempts: jsii.Number(123),
//   	},
//   	DeadLetterQueue: queue,
//   	DeadLetterQueueEnabled: jsii.Boolean(false),
//   	DeadLetterTopic: topic,
//   	Description: jsii.String("description"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	EnvironmentEncryption: key,
//   	EphemeralStorageSize: size,
//   	Events: []*iEventSource{
//   		eventSource,
//   	},
//   	Filesystem: fileSystem,
//   	FunctionName: jsii.String("functionName"),
//   	InitialPolicy: []*policyStatement{
//   		policyStatement,
//   	},
//   	InsightsVersion: lambdaInsightsVersion,
//   	LambdaPurpose: jsii.String("lambdaPurpose"),
//   	Layers: []iLayerVersion{
//   		layerVersion,
//   	},
//   	LogRetention: awscdk.Aws_logs.RetentionDays_ONE_DAY,
//   	LogRetentionRetryOptions: &LogRetentionRetryOptions{
//   		Base: cdk.Duration_*Minutes(jsii.Number(30)),
//   		MaxRetries: jsii.Number(123),
//   	},
//   	LogRetentionRole: role,
//   	MaxEventAge: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MemorySize: jsii.Number(123),
//   	OnFailure: destination,
//   	OnSuccess: destination,
//   	ParamsAndSecrets: paramsAndSecretsLayerVersion,
//   	Profiling: jsii.Boolean(false),
//   	ProfilingGroup: profilingGroup,
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	RetryAttempts: jsii.Number(123),
//   	Role: role,
//   	RuntimeManagementMode: runtimeManagementMode,
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	SnapStart: snapStartConf,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	Tracing: awscdk.*Aws_lambda.Tracing_ACTIVE,
//   	Vpc: vpc,
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type SingletonFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	// Default: Duration.hours(6)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	// Default: - no destination.
	//
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	// Default: - no destination.
	//
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	// Default: 2.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Specify the configuration of AWS Distro for OpenTelemetry (ADOT) instrumentation.
	// See: https://aws-otel.github.io/docs/getting-started/lambda
	//
	// Default: - No ADOT instrumentation.
	//
	AdotInstrumentation *AdotInstrumentationConfig `field:"optional" json:"adotInstrumentation" yaml:"adotInstrumentation"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	// Default: false.
	//
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	// Default: Architecture.X86_64
	//
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Code signing config associated with this function.
	// Default: - Not Sign the Code.
	//
	CodeSigningConfig ICodeSigningConfig `field:"optional" json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	// Default: - default options as described in `VersionOptions`.
	//
	CurrentVersionOptions *VersionOptions `field:"optional" json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	// Default: - SQS queue with 14 day retention period if `deadLetterQueueEnabled` is `true`.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	// Default: - false unless `deadLetterQueue` is set, which implies DLQ is enabled.
	//
	DeadLetterQueueEnabled *bool `field:"optional" json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	// Default: - no SNS topic.
	//
	DeadLetterTopic awssns.ITopic `field:"optional" json:"deadLetterTopic" yaml:"deadLetterTopic"`
	// A description of the function.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	// Default: - AWS Lambda creates and uses an AWS managed customer master key (CMK).
	//
	EnvironmentEncryption awskms.IKey `field:"optional" json:"environmentEncryption" yaml:"environmentEncryption"`
	// The size of the function’s /tmp directory in MiB.
	// Default: 512 MiB.
	//
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	// Default: - No event sources.
	//
	Events *[]IEventSource `field:"optional" json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	// Default: - will not mount any filesystem.
	//
	Filesystem FileSystem `field:"optional" json:"filesystem" yaml:"filesystem"`
	// A name for the function.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that
	// ID for the function's name. For more information, see Name Type.
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	// Default: - No policy statements are added to the created Lambda role.
	//
	InitialPolicy *[]awsiam.PolicyStatement `field:"optional" json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	// Default: - No Lambda Insights.
	//
	InsightsVersion LambdaInsightsVersion `field:"optional" json:"insightsVersion" yaml:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	// Default: - No layers.
	//
	Layers *[]ILayerVersion `field:"optional" json:"layers" yaml:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Default: - Default AWS SDK retry options.
	//
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Default: - A new role is created.
	//
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Default: 128.
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Specify the configuration of Parameters and Secrets Extension.
	// See: https://docs.aws.amazon.com/systems-manager/latest/userguide/ps-integration-lambda-extensions.html
	//
	// Default: - No Parameters and Secrets Extension.
	//
	ParamsAndSecrets ParamsAndSecretsLayerVersion `field:"optional" json:"paramsAndSecrets" yaml:"paramsAndSecrets"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Default: - No profiling.
	//
	Profiling *bool `field:"optional" json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Default: - A new profiling group will be created if `profiling` is set.
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `field:"optional" json:"profilingGroup" yaml:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	// Default: - No specific limit - account limit.
	//
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	// Default: - A unique role will be generated for this lambda function.
	// Both supplied and generated roles can always be changed by calling `addToRolePolicy`.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Sets the runtime management configuration for a function's version.
	// Default: Auto.
	//
	RuntimeManagementMode RuntimeManagementMode `field:"optional" json:"runtimeManagementMode" yaml:"runtimeManagementMode"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Default: - If the function is placed within a VPC and a security group is
	// not specified, either by this or securityGroup prop, a dedicated security
	// group will be created for this function.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Enable SnapStart for Lambda Function.
	//
	// SnapStart is currently supported only for Java 11, 17 runtime.
	// Default: - No snapstart.
	//
	SnapStart SnapStartConf `field:"optional" json:"snapStart" yaml:"snapStart"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Default: Duration.seconds(3)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	// Default: Tracing.Disabled
	//
	Tracing Tracing `field:"optional" json:"tracing" yaml:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	// This is required when `vpcSubnets` is specified.
	// Default: - Function is not placed within a VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// This requires `vpc` to be specified in order for interfaces to actually be
	// placed in the subnets. If `vpc` is not specify, this will raise an error.
	//
	// Note: Internet access for Lambda Functions requires a NAT Gateway, so picking
	// public subnets is not allowed (unless `allowPublicSubnet` is set to `true`).
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include
	// namespaces and other qualifiers, depending on the runtime.
	// For more information, see https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html.
	//
	// Use `Handler.FROM_IMAGE` when defining a function from a Docker image.
	//
	// NOTE: If you specify your source code as inline text by specifying the
	// ZipFile property within the Code property, specify index.function_name as
	// the handler.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when defining a function from a Docker image.
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// A unique identifier to identify this lambda.
	//
	// The identifier should be unique across all custom resource providers.
	// We recommend generating a UUID per provider.
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
	// A descriptive name for the purpose of this Lambda.
	//
	// If the Lambda does not have a physical name, this string will be
	// reflected its generated name. The combination of lambdaPurpose
	// and uuid must be unique.
	// Default: SingletonLambda.
	//
	LambdaPurpose *string `field:"optional" json:"lambdaPurpose" yaml:"lambdaPurpose"`
}

