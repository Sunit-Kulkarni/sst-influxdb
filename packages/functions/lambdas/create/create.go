package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
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
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})
	if err != nil {
		return fmt.Errorf("could not setup InfluxDB v3 client: %s", err)
	}

	// Write data point
	point := influxdb3.NewPointWithMeasurement("barrel").
		AddTag("ipAddress", measurement.IPAddress).
		AddTag("barrelId", measurement.BarrelId).
		AddTag("locationId", measurement.LocationId).
		AddField("temperature", measurement.Temperature)

	if err := client.WritePoints(ctx, point); err != nil {
		return fmt.Errorf("write API write point: %s", err)
	}

	return nil
}

func main() {
	lambda.Start(Create)
}
