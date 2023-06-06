package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"

	"github.com/Sunit-Kulkarni/sst-influxdb/v2/packages/functions/db"
	"github.com/jackc/pgx/v5"
)

type MemberRow struct {
	MemberId          string `db:"member_id"`
	ClubMemberAliasId string `db:"member_alias_id"`
	FamilyName        string `db:"familyName"`
}

type ResultBody struct {
	MemberId          string `json:"member_id"`
	ClubMemberAliasId string `json:"member_alias_id"`
	FamilyName        string `json:"familyName"`
}

func Get(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	memberAliasId := request.PathParameters["id"]

	dsn := os.Getenv("DATABASE_URL")
	conn, err := db.CreateCockroachDBConnection(dsn)
	defer conn.Close(context.Background())

	var Result ResultBody
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return getRow(context.Background(), tx, memberAliasId, &Result)
	})

	if err == pgx.ErrNoRows {
		return events.APIGatewayProxyResponse{
			Body:       "Membership not found:" + err.Error(),
			StatusCode: 404,
		}, nil
	}

	jsonBytes, err := json.Marshal(Result)
	JSON := string(jsonBytes)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       JSON,
	}, nil
}

func getRow(ctx context.Context, tx pgx.Tx, memberAliasId string, Result *ResultBody) error {
	// get row from "members" table.
	var rowResult MemberRow
	row := tx.QueryRow(ctx,
		`SELECT DISTINCT member_id,member_alias_id,family_name
		FROM country_club.public.members
		WHERE member_alias_id=($1)`,
		memberAliasId,
	)
	err := row.Scan(&rowResult.MemberId, &rowResult.ClubMemberAliasId, &rowResult.FamilyName)
	if err != nil {
		return err
	}

	Result.MemberId = rowResult.MemberId
	Result.ClubMemberAliasId = rowResult.ClubMemberAliasId
	Result.FamilyName = rowResult.FamilyName

	return nil
}

func main() {
	lambda.Start(Get)
}
