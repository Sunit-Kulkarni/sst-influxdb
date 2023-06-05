package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Sunit-Kulkarni/sst-influxdb/v2/packages/functions/db"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	"github.com/jackc/pgx/v5"
)

type CreateRequest struct {
	ClubMemberId string `json:"clubMemberId"`
	FamilyName   string `json:"familyName"`
}

func Create(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var newMemberRequest CreateRequest
	err := json.Unmarshal([]byte(request.Body), &newMemberRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Your POST request body is invalid:" + err.Error(),
			StatusCode: 400,
		}, err
	}

	dsn := os.Getenv("DATABASE_URL")
	conn, err := db.CreateCockroachDBConnection(dsn)
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return insertRow(context.Background(), tx, newMemberRequest)
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "DB Row insertion failed:" + err.Error(),
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "New Member " + newMemberRequest.FamilyName + " was received at " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func insertRow(ctx context.Context, tx pgx.Tx, newMemberRequest CreateRequest) error {
	// Insert row into the "members" table.
	log.Println("Creating new rows...")
	if _, err := tx.Exec(ctx,
		`INSERT INTO country_club.public.members (member_alias_id, family_name) 
		VALUES ($1, $2)`,
		newMemberRequest.ClubMemberId, newMemberRequest.FamilyName,
	); err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(Create)
}
