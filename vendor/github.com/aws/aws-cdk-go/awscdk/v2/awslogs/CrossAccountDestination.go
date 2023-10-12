package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new CloudWatch Logs Destination for use in cross-account scenarios.
//
// CrossAccountDestinations are used to subscribe a Kinesis stream in a
// different account to a CloudWatch Subscription.
//
// Consumers will hardly ever need to use this class. Instead, directly
// subscribe a Kinesis stream using the integration class in the
// `aws-cdk-lib/aws-logs-destinations` package; if necessary, a
// `CrossAccountDestination` will be created automatically.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   crossAccountDestination := awscdk.Aws_logs.NewCrossAccountDestination(this, jsii.String("MyCrossAccountDestination"), &CrossAccountDestinationProps{
//   	Role: role,
//   	TargetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	DestinationName: jsii.String("destinationName"),
//   })
//
type CrossAccountDestination interface {
	awscdk.Resource
	ILogSubscriptionDestination
	// The ARN of this CrossAccountDestination object.
	DestinationArn() *string
	// The name of this CrossAccountDestination object.
	DestinationName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// Policy object of this CrossAccountDestination object.
	PolicyDocument() awsiam.PolicyDocument
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	AddToPolicy(statement awsiam.PolicyStatement)
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
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	Bind(_scope constructs.Construct, _sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CrossAccountDestination
type jsiiProxy_CrossAccountDestination struct {
	internal.Type__awscdkResource
	jsiiProxy_ILogSubscriptionDestination
}

func (j *jsiiProxy_CrossAccountDestination) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) DestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossAccountDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCrossAccountDestination(scope constructs.Construct, id *string, props *CrossAccountDestinationProps) CrossAccountDestination {
	_init_.Initialize()

	if err := validateNewCrossAccountDestinationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CrossAccountDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CrossAccountDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCrossAccountDestination_Override(c CrossAccountDestination, scope constructs.Construct, id *string, props *CrossAccountDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.CrossAccountDestination",
		[]interface{}{scope, id, props},
		c,
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
func CrossAccountDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCrossAccountDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CrossAccountDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CrossAccountDestination_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCrossAccountDestination_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CrossAccountDestination",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CrossAccountDestination_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCrossAccountDestination_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.CrossAccountDestination",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) AddToPolicy(statement awsiam.PolicyStatement) {
	if err := c.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToPolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CrossAccountDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CrossAccountDestination) Bind(_scope constructs.Construct, _sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig {
	if err := c.validateBindParameters(_scope, _sourceLogGroup); err != nil {
		panic(err)
	}
	var returns *LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossAccountDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
