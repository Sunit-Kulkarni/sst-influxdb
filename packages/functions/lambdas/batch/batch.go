package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func Batch(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	if err := dbWrite(context.Background()); err != nil {
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
		point := influxdb3.NewPointWithMeasurement("barrel").
			AddTag("ipAddress", data[key]["ipAddress"].(string)).
			AddTag("barrelId", data[key]["ipAddress"].(string)).
			AddTag("locationId", data[key]["locationId"].(string)).
			AddField("temperature", data[key]["temperature"])

		if err := client.WritePoints(ctx, point); err != nil {
			return fmt.Errorf("write API write point: %s", err)
		}
	}

	return nil
}

func main() {
	lambda.Start(Batch)
}
