package main

import (
	"context"
	"fmt"
	feature_library "study/feature_library/sql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := feature_library.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	library := feature_library.NewDbConnection(conn)
	if err := library.CreateTable(ctx); err != nil {
		panic(err)
	}
	if err := library.InsertRow(ctx, "TEST3", "TEST", "TEST", "TEST", 1111, false, time.Now()); err != nil {
		panic(err)
	}
	if err := library.UpdateRow(ctx, 5, time.Now()); err != nil {
		panic(err)
	}
	if err := library.DeleteRow(ctx, 3); err != nil {
		panic(err)
	}

	fmt.Println("succed!")
}
