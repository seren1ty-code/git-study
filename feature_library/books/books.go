package feature_library

import (
	"context"
	"fmt"
	feature_library "study/feature_library/sql"
	"time"

	"github.com/k0kubun/pp"
)

type Book struct {
	ID              int
	Title           string
	Description     string
	Author          string
	Review          string
	PublicationYear int
	Completed       bool
	CreatedAt       time.Time
	CompletedAt     *time.Time
}

func AddNewBook(ctx context.Context, dB *feature_library.DataBase, book Book) error {
	err := dB.InsertRow(ctx, book.Title,
		book.Description,
		book.Author,
		book.Review,
		book.PublicationYear,
		book.Completed,
		book.CreatedAt)

	return err
}

func ListBooks(ctx context.Context, dB *feature_library.DataBase) error {
	rows, err := dB.GetRows(ctx)
	if err != nil {
		return err
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.ID,
			&b.Title,
			&b.Description,
			&b.Author,
			&b.Review,
			&b.PublicationYear,
			&b.Completed,
			&b.CreatedAt,
			&b.CompletedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}

	pp.Println(books)
	return nil
}

func DeleteBooksBySlice(ctx context.Context, dB *feature_library.DataBase, idSlice []int) {

	for _, v := range idSlice {
		dB.DeleteRow(ctx, v)
	}
}

func RedactBookByStruct(ctx context.Context, dB *feature_library.DataBase, b Book) error {
	sqlQuery := `
	UPDATE library
	SET 
	title = $1,
	description = $2,
	author = $3, 
	review = $4, 
	publication_year = $5,
	completed = $6,
	created_at = $7,
	completed_at = $8
	WHERE id = $9;
	`

	_, err := dB.Conn.Exec(ctx, sqlQuery,
		b.Title,
		b.Description,
		b.Author,
		b.Review,
		b.PublicationYear,
		b.Completed,
		b.CreatedAt,
		b.CompletedAt,
		b.ID)
	return err
}
