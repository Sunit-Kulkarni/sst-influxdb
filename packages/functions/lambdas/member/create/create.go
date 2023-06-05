package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jackc/pgx/v5"
)

type CreateRequest struct {
	CountryClubMemberId string `json:"countryClubMemberId"`
	FamilyName          string `json:"familyName"`
}

func Create(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var newMemberRequest CreateRequest
	err := json.Unmarshal([]byte(request.Body), &newMemberRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Your POST request is invalid:" + err.Error(),
			StatusCode: 400,
		}, err
	}

	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			Body:       "Error parsing connection string:" + err.Error(),
			StatusCode: 400,
		}, err
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			Body:       "Error configuring runtime params:" + err.Error(),
			StatusCode: 400,
		}, err
	}
	defer conn.Close(context.Background())

	return events.APIGatewayProxyResponse{
		Body:       "New Member " + newMemberRequest.FamilyName + " was received at " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Create)
}
