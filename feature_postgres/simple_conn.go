package featurepostgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SimpleConn() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:4382@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подлкючение к БД успешно!")

}
