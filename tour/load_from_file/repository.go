package repository

import (
	"context"
	"database/sql"

	"github.com/wroge/sqlt"
	_ "modernc.org/sqlite"
)

type Book struct {
	ID    int64
	Title string
}

var (
	config = sqlt.Config{
		Templates: []sqlt.Template{
			sqlt.ParseFiles("load_from_file/queries.sql"),
		},
	}

	schema = sqlt.Exec[any](config, sqlt.Lookup("schema"))

	create = sqlt.First[string, int64](config, sqlt.Lookup("insert_book"))

	get = sqlt.First[int64, Book](config, sqlt.Lookup("get_book"))
)

func NewRepository() (Repository, error) {
	db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)")
	if err != nil {
		return Repository{}, err
	}

	_, err = schema.Exec(context.Background(), db, nil)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		DB: db,
	}, nil
}

type Repository struct {
	DB *sql.DB
}

func (r Repository) Create(ctx context.Context, title string) (int64, error) {
	return create.Exec(ctx, r.DB, title)
}

func (r Repository) Get(ctx context.Context, id int64) (Book, error) {
	return get.Exec(ctx, r.DB, id)
}
