package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

type QualifiedFunctionBase interface {
	FunctionBase
	// The architecture of this Lambda Function.
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
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN fo the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The underlying `IFunction`.
	Lambda() IFunction
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() IVersion
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
	// The qualifier of the version or alias of this function.
	//
	// A qualifier is the identifier that's appended to a version or alias ARN.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunctionConfiguration.html#API_GetFunctionConfiguration_RequestParameters
	//
	Qualifier() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	ResourceArnsForGrantInvoke() *[]*string
	// The IAM role associated with this function.
	//
	// Undefined if the function was imported without a role.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
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
	ConsiderWarningOnInvokeFunctionPermissions(_scope constructs.Construct, _action *string)
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

// The jsii proxy struct for QualifiedFunctionBase
type jsiiProxy_QualifiedFunctionBase struct {
	jsiiProxy_FunctionBase
}

func (j *jsiiProxy_QualifiedFunctionBase) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewQualifiedFunctionBase_Override(q QualifiedFunctionBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		[]interface{}{scope, id, props},
		q,
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
func QualifiedFunctionBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQualifiedFunctionBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func QualifiedFunctionBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQualifiedFunctionBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func QualifiedFunctionBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQualifiedFunctionBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) AddEventSource(source IEventSource) {
	if err := q.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addEventSource",
		[]interface{}{source},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	if err := q.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns EventSourceMapping

	_jsii_.Invoke(
		q,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
	if err := q.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionUrl

	_jsii_.Invoke(
		q,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) AddPermission(id *string, permission *Permission) {
	if err := q.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := q.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := q.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := q.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) ConsiderWarningOnInvokeFunctionPermissions(_scope constructs.Construct, _action *string) {
	if err := q.validateConsiderWarningOnInvokeFunctionPermissionsParameters(_scope, _action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{_scope, _action},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := q.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := q.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QualifiedFunctionBase) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := q.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}

