package main

import (
	"context"
	"fmt"
	books "study/feature_library/books"
	feature_library "study/feature_library/sql"
)

func main() {
	ctx := context.Background()

	conn, err := feature_library.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	lib := feature_library.NewDbConnection(conn)
	if err := lib.CreateTable(ctx); err != nil {
		panic(err)
	}

	// cmpTime := time.Now()

	// newB := books.Book{
	// 	ID:              2,
	// 	Title:           "REDACTED_TEST!!!!",
	// 	Description:     "12",
	// 	Author:          "34",
	// 	Review:          "TEST1",
	// 	PublicationYear: 1111,
	// 	Completed:       true,
	// 	CreatedAt:       time.Now(),
	// 	CompletedAt:     &cmpTime,
	// }

	// books.AddNewBook(ctx, lib, newB)

	// if err := books.ListBooks(ctx, lib); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if err := books.RedactBookByStruct(ctx, lib, newB); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	slice := []int{3, 4}

	books.DeleteBooksBySlice(ctx, lib, slice)

	fmt.Println("succed!")
}
