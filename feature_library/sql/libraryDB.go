package feature_library

import (
	"context"

	"time"

	"github.com/jackc/pgx/v5"
)

type DataBase struct {
	Conn *pgx.Conn
}

func NewDbConnection(conn *pgx.Conn) *DataBase {
	return &DataBase{
		Conn: conn,
	}
}

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:4382@localhost:5432/postgres")
}

func (d *DataBase) CreateTable(ctx context.Context) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS library(
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		author VARCHAR(100) NOT NULL,
		review VARCHAR(5000),
		publication_year INTEGER NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,
		
		UNIQUE(title)
		);
		`

	_, err := d.Conn.Exec(ctx, sqlQuery)

	return err
}

func (d *DataBase) InsertRow(ctx context.Context,
	title string,
	description string,
	author string,
	review string,
	publication_year int,
	completed bool,
	createdAt time.Time,
) error {
	sqlQuery := `
			INSERT INTO library(
				title,
				description,
				author,
				review,
				publication_year,
				completed,
				created_at
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7);
				`

	_, err := d.Conn.Exec(ctx, sqlQuery, title, description, author, review, publication_year, completed, createdAt)

	return err

}

func (d *DataBase) DeleteRow(ctx context.Context, id int) error {
	sqlQuery := `
	DELETE FROM library
	WHERE id = $1;
	`

	_, err := d.Conn.Exec(ctx, sqlQuery, id)
	return err
}

func (d *DataBase) UpdateRow(ctx context.Context, id int, completeTime time.Time) error {
	sqlQuery := `
	UPDATE library 
	SET completed = TRUE, 
	completed_at = $1
	WHERE id = $2;
	`

	_, err := d.Conn.Exec(ctx, sqlQuery, time.Now(), completeTime, id)
	return err
}

func (d *DataBase) GetRows(ctx context.Context) (pgx.Rows, error) {
	return d.Conn.Query(ctx, "SELECT * FROM library ORDER BY id ASC")
}
