package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	return events.APIGatewayProxyResponse{
		Body:       "New Member " + newMemberRequest.FamilyName + " was received at " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Create)
}
