package main

import (
	"context"
	"fmt"
	featurepostgres "study/feature_postgres"
	simplesql "study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := featurepostgres.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simplesql.InsertRow(ctx, conn, "TEST_3", "TEST_DESC3", true, time.Now()); err != nil {
		panic(err)
	}

	fmt.Println("succed")
}
