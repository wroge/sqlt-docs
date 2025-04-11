package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/wroge/sqlt"
	_ "modernc.org/sqlite"
)

type Book struct {
	ID      int64
	Title   string
	Author  string
	AddedAt time.Time
}

type Params struct {
	Title  string
	Author string
}

var (
	config = sqlt.Config{
		Templates: []sqlt.Template{
			sqlt.Funcs(sprig.TxtFuncMap()),
			sqlt.ParseFiles("transactions/queries.sql"),
		},
	}

	schema = sqlt.Exec[any](config, sqlt.Lookup("schema"))

	create = sqlt.Transaction(nil,
		sqlt.One[Params, int64](config, sqlt.Lookup("upsert_author")),
		sqlt.One[Params, int64](config, sqlt.Lookup("insert_book")),
	)

	createMany = sqlt.Transaction(nil,
		sqlt.All[[]Params, int64](config, sqlt.Lookup("upsert_authors")),
		sqlt.All[[]Params, int64](config, sqlt.Lookup("insert_books")),
	)

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
	ctx, err := create.Exec(ctx, r.DB, params)
	if err != nil {
		return 0, err
	}

	id, ok := ctx.Value(sqlt.ContextKey("insert_book")).(int64)
	if !ok {
		return 0, errors.New("internal error")
	}

	return id, nil
}

func (r Repository) CreateMany(ctx context.Context, params []Params) ([]int64, error) {
	ctx, err := createMany.Exec(ctx, r.DB, params)
	if err != nil {
		return nil, err
	}

	ids, ok := ctx.Value(sqlt.ContextKey("insert_books")).([]int64)
	if !ok {
		return nil, errors.New("internal error")
	}

	return ids, nil
}

func (r Repository) Get(ctx context.Context, id int64) (Book, error) {
	return get.Exec(ctx, r.DB, id)
}
