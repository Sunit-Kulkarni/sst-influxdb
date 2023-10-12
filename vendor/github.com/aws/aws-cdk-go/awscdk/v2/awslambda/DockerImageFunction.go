package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a lambda function where the handler is a docker image.
//
// Example:
//   lambda.NewDockerImageFunction(this, jsii.String("AssetFunction"), &DockerImageFunctionProps{
//   	Code: lambda.DockerImageCode_FromImageAsset(path.join(__dirname, jsii.String("docker-handler"))),
//   })
//
type DockerImageFunction interface {
	Function
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	Architecture() Architecture
	// Whether the addPermission() call adds any permissions.
	//
	// True for new Lambdas, false for version $LATEST and imported Lambdas
	// from different accounts.
	CanCreatePermissions() *bool
	// Access the Connections object.
	//
	// Will fail if not a VPC-enabled Lambda Function.
	Connections() awsec2.Connections
	// Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes.
	//
	// You can specify options for this version using the `currentVersionOptions`
	// prop when initializing the `lambda.Function`.
	CurrentVersion() Version
	// The DLQ (as queue) associated with this Lambda Function (this is an optional attribute).
	DeadLetterQueue() awssqs.IQueue
	// The DLQ (as topic) associated with this Lambda Function (this is an optional attribute).
	DeadLetterTopic() awssns.ITopic
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// ARN of this function.
	FunctionArn() *string
	// Name of this function.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() IVersion
	// The LogGroup where the Lambda function's logs are made available.
	//
	// If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
	// pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
	// period (never expire, by default).
	//
	// Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
	// to never expire even if it was configured with a different value.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	Node() constructs.Node
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	ResourceArnsForGrantInvoke() *[]*string
	// Execution role associated with this function.
	Role() awsiam.IRole
	// The runtime configured for this lambda.
	Runtime() Runtime
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The timeout configured for this lambda.
	Timeout() awscdk.Duration
	// Defines an alias for this function.
	//
	// The alias will automatically be updated to point to the latest version of
	// the function as it is being updated during a deployment.
	//
	// ```ts
	// declare const fn: lambda.Function;
	//
	// fn.addAlias('Live');
	//
	// // Is equivalent to
	//
	// new lambda.Alias(this, 'AliasLive', {
	//   aliasName: 'Live',
	//   version: fn.currentVersion,
	// });
	// ```.
	AddAlias(aliasName *string, options *AliasOptions) Alias
	// Adds an environment variable to this Lambda function.
	//
	// If this is a ref to a Lambda function, this operation results in a no-op.
	AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function
	// Adds an event source to this function.
	//
	// Event sources are implemented in the aws-cdk-lib/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	AddEventSource(source IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	// Adds a url to this lambda function.
	AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl
	// Adds one or more Lambda Layers to this Lambda function.
	AddLayers(layers ...ILayerVersion)
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	AddPermission(id *string, permission *Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	// A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.
	//
	// - function.currentVersion is invoked before or after the permission is created.
	//
	// This applies only to permissions on Lambda functions, not versions or aliases.
	// This function is overridden as a noOp for QualifiedFunctionBase.
	ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant
	// Mix additional information into the hash of the Version object.
	//
	// The Lambda Function construct does its best to automatically create a new
	// Version when anything about the Function changes (its code, its layers,
	// any of the other properties).
	//
	// However, you can sometimes source information from places that the CDK cannot
	// look into, like the deploy-time values of SSM parameters. In those cases,
	// the CDK would not force the creation of a new Version object when it actually
	// should.
	//
	// This method can be used to invalidate the current Version object. Pass in
	// any string into this method, and make sure the string changes when you know
	// a new Version needs to be created.
	//
	// This method may be called more than once.
	InvalidateVersionBasedOn(x *string)
	// Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	WarnInvokeFunctionPermissions(scope constructs.Construct)
}

// The jsii proxy struct for DockerImageFunction
type jsiiProxy_DockerImageFunction struct {
	jsiiProxy_Function
}

func (j *jsiiProxy_DockerImageFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) CurrentVersion() Version {
	var returns Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Runtime() Runtime {
	var returns Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewDockerImageFunction(scope constructs.Construct, id *string, props *DockerImageFunctionProps) DockerImageFunction {
	_init_.Initialize()

	if err := validateNewDockerImageFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerImageFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDockerImageFunction_Override(d DockerImageFunction, scope constructs.Construct, id *string, props *DockerImageFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		[]interface{}{scope, id, props},
		d,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func DockerImageFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	if err := validateDockerImageFunction_ClassifyVersionPropertyParameters(propertyName, locked); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
func DockerImageFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) IFunction {
	_init_.Initialize()

	if err := validateDockerImageFunction_FromFunctionArnParameters(scope, id, functionArn); err != nil {
		panic(err)
	}
	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
func DockerImageFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *FunctionAttributes) IFunction {
	_init_.Initialize()

	if err := validateDockerImageFunction_FromFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
func DockerImageFunction_FromFunctionName(scope constructs.Construct, id *string, functionName *string) IFunction {
	_init_.Initialize()

	if err := validateDockerImageFunction_FromFunctionNameParameters(scope, id, functionName); err != nil {
		panic(err)
	}
	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
		&returns,
	)

	return returns
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
func DockerImageFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDockerImageFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DockerImageFunction_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDockerImageFunction_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DockerImageFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDockerImageFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func DockerImageFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Default: max over 5 minutes.
//
func DockerImageFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Default: average over 5 minutes.
//
func DockerImageFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Default: sum over 5 minutes.
//
func DockerImageFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Default: sum over 5 minutes.
//
func DockerImageFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Default: sum over 5 minutes.
//
func DockerImageFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Default: max over 5 minutes.
//
func DockerImageFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateDockerImageFunction_MetricAllUnreservedConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) AddAlias(aliasName *string, options *AliasOptions) Alias {
	if err := d.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns Alias

	_jsii_.Invoke(
		d,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function {
	if err := d.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns Function

	_jsii_.Invoke(
		d,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) AddEventSource(source IEventSource) {
	if err := d.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addEventSource",
		[]interface{}{source},
	)
}

func (d *jsiiProxy_DockerImageFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	if err := d.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns EventSourceMapping

	_jsii_.Invoke(
		d,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
	if err := d.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionUrl

	_jsii_.Invoke(
		d,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) AddLayers(layers ...ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addLayers",
		args,
	)
}

func (d *jsiiProxy_DockerImageFunction) AddPermission(id *string, permission *Permission) {
	if err := d.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (d *jsiiProxy_DockerImageFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := d.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (d *jsiiProxy_DockerImageFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DockerImageFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := d.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (d *jsiiProxy_DockerImageFunction) ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string) {
	if err := d.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (d *jsiiProxy_DockerImageFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) InvalidateVersionBasedOn(x *string) {
	if err := d.validateInvalidateVersionBasedOnParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"invalidateVersionBasedOn",
		[]interface{}{x},
	)
}

func (d *jsiiProxy_DockerImageFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageFunction) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := d.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}

