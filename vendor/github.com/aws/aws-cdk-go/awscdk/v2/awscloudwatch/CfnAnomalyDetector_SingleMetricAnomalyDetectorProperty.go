package awscloudwatch


// Designates the CloudWatch metric and statistic that provides the time series the anomaly detector uses as input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   singleMetricAnomalyDetectorProperty := &SingleMetricAnomalyDetectorProperty{
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	Stat: jsii.String("stat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html
//
type CfnAnomalyDetector_SingleMetricAnomalyDetectorProperty struct {
	// The metric dimensions to create the anomaly detection model for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric to create the anomaly detection model for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric to create the anomaly detection model for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic to use for the metric and anomaly detection model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-singlemetricanomalydetector.html#cfn-cloudwatch-anomalydetector-singlemetricanomalydetector-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

