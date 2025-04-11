package main

import (
	"context"
	"fmt"

	bulk_insert "github.com/wroge/sqlt-docs/tour/bulk_insert"
	complex_query "github.com/wroge/sqlt-docs/tour/complex_query"
	create_statements "github.com/wroge/sqlt-docs/tour/create_statements"
	load_from_file "github.com/wroge/sqlt-docs/tour/load_from_file"
	transactions "github.com/wroge/sqlt-docs/tour/transactions"
)

func main() {
	create_statements_example()
	load_from_file_example()
	bulk_insert_example()
	transactions_example()
	complex_query_example()
}

func complex_query_example() {
	r, err := complex_query.NewRepository()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	_, err = r.Create(ctx, complex_query.Insert{
		Title:   "Moby-Dick",
		Authors: []string{"Herman Melville"},
	})
	if err != nil {
		panic(err)
	}

	_, err = r.Create(ctx, complex_query.Insert{
		Title:   "Good Omens",
		Authors: []string{"Neil Gaiman", "Terry Pratchett"},
	})
	if err != nil {
		panic(err)
	}

	_, err = r.Create(ctx, complex_query.Insert{
		Title:   "Discworld",
		Authors: []string{"Terry Pratchett"},
	})
	if err != nil {
		panic(err)
	}

	for range 2 {
		books, err := r.Query(ctx, complex_query.Query{
			Author: "Terry Pratchett",
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(books)
	}

	fmt.Println("complex_query_example ✅")
}

func transactions_example() {
	r, err := transactions.NewRepository()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	id, err := r.Create(ctx, transactions.Params{
		Title:  "Moby-Dick",
		Author: "Herman Melville",
	})
	if err != nil {
		panic(err)
	}

	book, err := r.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)

	ids, err := r.CreateMany(ctx, []transactions.Params{
		{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
		},
		{
			Title:  "Lord of the Flies",
			Author: "William Golding",
		},
	})
	if err != nil {
		panic(err)
	}

	for _, id := range ids {
		book, err := r.Get(ctx, id)
		if err != nil {
			panic(err)
		}

		fmt.Println(book)
	}

	fmt.Println("transactions_example ✅")
}

func bulk_insert_example() {
	r, err := bulk_insert.NewRepository()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	id, err := r.Create(ctx, bulk_insert.Params{
		Title:  "Moby-Dick",
		Author: "Herman Melville",
	})
	if err != nil {
		panic(err)
	}

	book, err := r.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)

	ids, err := r.CreateMany(ctx, []bulk_insert.Params{
		{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
		},
		{
			Title:  "Lord of the Flies",
			Author: "William Golding",
		},
	})
	if err != nil {
		panic(err)
	}

	for _, id := range ids {
		book, err := r.Get(ctx, id)
		if err != nil {
			panic(err)
		}

		fmt.Println(book)
	}

	fmt.Println("bulk_insert_example ✅")
}

func load_from_file_example() {
	r, err := load_from_file.NewRepository()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	id, err := r.Create(ctx, "Moby-Dick")
	if err != nil {
		panic(err)
	}

	book, err := r.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
	fmt.Println("load_from_file_example ✅")
}

func create_statements_example() {
	r, err := create_statements.NewRepository()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	id, err := r.Create(ctx, "Moby-Dick")
	if err != nil {
		panic(err)
	}

	book, err := r.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(book)
	fmt.Println("create_statements_example ✅")
}
