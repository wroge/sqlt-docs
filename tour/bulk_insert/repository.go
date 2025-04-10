package repository

import (
	"context"
	"database/sql"

	"github.com/wroge/sqlt"
	_ "modernc.org/sqlite"
)

type Book struct {
	ID     int64
	Title  string
	Author string
}

type Params struct {
	Title  string
	Author string
}

var (
	config = sqlt.Config{
		Templates: []sqlt.Template{
			sqlt.ParseFiles("bulk_insert/queries.sql"),
		},
	}

	schema = sqlt.Exec[any](config, sqlt.Lookup("schema"))

	create = sqlt.First[Params, int64](config, sqlt.Lookup("insert_book"))

	createMany = sqlt.All[[]Params, int64](config, sqlt.Lookup("insert_books"))

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

func (r Repository) Create(ctx context.Context, params Params) (int64, error) {
	return create.Exec(ctx, r.DB, params)
}

func (r Repository) CreateMany(ctx context.Context, params []Params) ([]int64, error) {
	return createMany.Exec(ctx, r.DB, params)
}

func (r Repository) Get(ctx context.Context, id int64) (Book, error) {
	return get.Exec(ctx, r.DB, id)
}
