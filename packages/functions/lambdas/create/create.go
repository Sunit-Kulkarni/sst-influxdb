package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type CreateRequest struct {
	IPAddress   string  `json:"ipAddress"`
	Temperature float32 `json:"temperature"`
	BarrelId    string  `json:"barrelId"`
	LocationId  string  `json:"locationId"`
}

func Create(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var measurement CreateRequest
	err := json.Unmarshal([]byte(request.Body), &measurement)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Your POST request is invalid:" + err.Error(),
			StatusCode: 400,
		}, err
	}

	if err := CreateWrite(context.Background(), measurement); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return events.APIGatewayProxyResponse{
			Body:       "Error writing to DB:" + err.Error(),
			StatusCode: 400,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       "Your POST request was received at " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func CreateWrite(ctx context.Context, measurement CreateRequest) error {
	// Create write client
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"
	token := os.Getenv("INFLUXDB_TOKEN")
	writeClient := influxdb2.NewClient(url, token)

	// Define write API
	org := "Dev Team"
	bucket := "sample-iot-data"
	writeAPI := writeClient.WriteAPIBlocking(org, bucket)

	// Write data point

	point := influxdb2.NewPointWithMeasurement("barrel").
		AddTag("ipAddress", measurement.IPAddress).
		AddTag("barrelId", measurement.BarrelId).
		AddTag("locationId", measurement.LocationId).
		AddField("temperature", measurement.Temperature)

	if err := writeAPI.WritePoint(ctx, point); err != nil {
		return fmt.Errorf("write API write point: %s", err)
	}

	return nil
}

func main() {
	lambda.Start(Create)
}
