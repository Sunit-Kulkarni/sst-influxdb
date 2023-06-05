package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	"github.com/jackc/pgx/v5"
)

type MemberRow struct {
	MemberId          string `db:"member_id"`
	ClubMemberAliasId string `db:"member_alias_id"`
	FamilyName        string `db:"familyName"`
}

func Get(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	memberAliasId := request.PathParameters["id"]

	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			Body:       "Error parsing connection string:" + err.Error(),
			StatusCode: 500,
		}, nil
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			Body:       "Error passing connection config:" + err.Error(),
			StatusCode: 500,
		}, nil
	}
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return getRow(context.Background(), tx, memberAliasId)
	})

	if err == pgx.ErrNoRows {
		return events.APIGatewayProxyResponse{
			Body:       "Membership not found:" + err.Error(),
			StatusCode: 404,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       memberAliasId + " has valid club membership",
		StatusCode: 200,
	}, nil
}

func getRow(ctx context.Context, tx pgx.Tx, memberAliasId string) error {
	// get row from "members" table.
	var result MemberRow

	row := tx.QueryRow(ctx,
		`SELECT DISTINCT member_id,member_alias_id,family_name
		FROM country_club.public.members
		WHERE member_alias_id=($1)`,
		memberAliasId,
	)
	err := row.Scan(&result.MemberId, &result.ClubMemberAliasId, &result.FamilyName)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(Get)
}
