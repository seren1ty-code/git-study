package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	INSERT INTO tasks(
	title,
	description,
	completed,
	created_at
	)
	VALUES ('TEST_TASK2', 'TEST_DESCRIPTION 12345', FALSE, '2026-02-02 18:01:00');
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
