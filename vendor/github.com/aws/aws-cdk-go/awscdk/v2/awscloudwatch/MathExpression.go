package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A math expression built with metric(s) emitted by a service.
//
// The math expression is a combination of an expression (x+y) and metrics to apply expression on.
// It also contains metadata which is used only in graphs, such as color and label.
// It makes sense to embed this in here, so that compound constructs can attach
// that metadata to metrics they expose.
//
// MathExpression can also be used for search expressions. In this case,
// it also optionally accepts a searchRegion and searchAccount property for cross-environment
// search expressions.
//
// This class does not represent a resource, so hence is not a construct. Instead,
// MathExpression is an abstraction that makes it easy to specify metrics for use in both
// alarms and graphs.
//
// Example:
//   var fn function
//
//
//   allProblems := cloudwatch.NewMathExpression(&MathExpressionProps{
//   	Expression: jsii.String("errors + throttles"),
//   	UsingMetrics: map[string]iMetric{
//   		"errors": fn.metricErrors(),
//   		"throttles": fn.metricThrottles(),
//   	},
//   })
//
type MathExpression interface {
	IMetric
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	Color() *string
	// The expression defining the metric.
	Expression() *string
	// Label for this metric when added to a Graph.
	Label() *string
	// Aggregation period of this metric.
	Period() awscdk.Duration
	// Account to evaluate search expressions within.
	SearchAccount() *string
	// Region to evaluate search expressions within.
	SearchRegion() *string
	// The metrics used in the expression as KeyValuePair <id, metric>.
	UsingMetrics() *map[string]IMetric
	// Warnings generated by this math expression.
	// Deprecated: - use warningsV2.
	Warnings() *[]*string
	// Warnings generated by this math expression.
	WarningsV2() *map[string]*string
	// Make a new Alarm for this metric.
	//
	// Combines both properties that may adjust the metric (aggregation) as well
	// as alarm properties.
	CreateAlarm(scope constructs.Construct, id *string, props *CreateAlarmOptions) Alarm
	// Inspect the details of the metric object.
	ToMetricConfig() *MetricConfig
	// Returns a string representation of an object.
	ToString() *string
	// Return a copy of Metric with properties changed.
	//
	// All properties except namespace and metricName can be changed.
	With(props *MathExpressionOptions) MathExpression
}

// The jsii proxy struct for MathExpression
type jsiiProxy_MathExpression struct {
	jsiiProxy_IMetric
}

func (j *jsiiProxy_MathExpression) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) SearchAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) SearchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) UsingMetrics() *map[string]IMetric {
	var returns *map[string]IMetric
	_jsii_.Get(
		j,
		"usingMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MathExpression) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}


func NewMathExpression(props *MathExpressionProps) MathExpression {
	_init_.Initialize()

	if err := validateNewMathExpressionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MathExpression{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.MathExpression",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewMathExpression_Override(m MathExpression, props *MathExpressionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.MathExpression",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MathExpression) CreateAlarm(scope constructs.Construct, id *string, props *CreateAlarmOptions) Alarm {
	if err := m.validateCreateAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns Alarm

	_jsii_.Invoke(
		m,
		"createAlarm",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MathExpression) ToMetricConfig() *MetricConfig {
	var returns *MetricConfig

	_jsii_.Invoke(
		m,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MathExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MathExpression) With(props *MathExpressionOptions) MathExpression {
	if err := m.validateWithParameters(props); err != nil {
		panic(err)
	}
	var returns MathExpression

	_jsii_.Invoke(
		m,
		"with",
		[]interface{}{props},
		&returns,
	)

	return returns
}

