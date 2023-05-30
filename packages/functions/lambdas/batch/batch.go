package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Batch(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	if err := dbWrite(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return events.APIGatewayProxyResponse{
			Body:       "Your POST batch request is invalid:" + err.Error(),
			StatusCode: 400,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       "Your POST request was received at " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func dbWrite(ctx context.Context) error {
	// Create write client
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"
	token := os.Getenv("INFLUXDB_TOKEN")
	writeClient := influxdb2.NewClient(url, token)

	// Define write API
	org := "Dev Team"
	bucket := "sample-iot-data"
	writeAPI := writeClient.WriteAPIBlocking(org, bucket)

	data := map[string]map[string]interface{}{
		"point1": {
			"ipAddress":   "1.1.1.1",
			"temperature": 59.5625,
			"barrelId":    "1XD5",
			"locationId":  "Home",
		},
		"point2": {
			"ipAddress":   "1.1.1.1",
			"temperature": 59.00,
			"barrelId":    "1XD5",
			"locationId":  "Home",
		},
	}

	// Write data
	for key := range data {
		point := influxdb2.NewPointWithMeasurement("barrel").
			AddTag("ipAddress", data[key]["ipAddress"].(string)).
			AddTag("barrelId", data[key]["ipAddress"].(string)).
			AddTag("locationId", data[key]["locationId"].(string)).
			AddField("temperature", data[key]["temperature"])

		if err := writeAPI.WritePoint(ctx, point); err != nil {
			return fmt.Errorf("write API write point: %s", err)
		}
	}

	return nil
}

func main() {
	lambda.Start(Batch)
}
