package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/wroge/sqlt"
	_ "modernc.org/sqlite"
)

type Book struct {
	ID      int64
	Title   string
	Authors []string
	AddedAt time.Time
}

type Query struct {
	Title       string
	Author      string
	AddedBefore time.Time
}

type Insert struct {
	Title   string
	Authors []string
}

var (
	config = sqlt.Config{
		Templates: []sqlt.Template{
			sqlt.Funcs(sprig.TxtFuncMap()),
			sqlt.ParseFiles("complex_query/queries.sql"),
		},
		Log: func(ctx context.Context, info sqlt.Info) {
			if info.Cached {
				fmt.Println(info.SQL, info.Args)
			}
		},
	}

	schema = sqlt.Exec[any](config, sqlt.Lookup("schema"))

	create = sqlt.Transaction(nil,
		sqlt.All[Insert, int64](config, sqlt.Lookup("upsert_authors")),
		sqlt.One[Insert, int64](config, sqlt.Lookup("insert_book")),
		sqlt.Exec[Insert](config, sqlt.Lookup("link_book_authors")),
	)

	query = sqlt.All[Query, Book](config, sqlt.NoExpirationCache(512), sqlt.Lookup("query_books"))
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

func (r Repository) Create(ctx context.Context, params Insert) (int64, error) {
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

func (r Repository) Query(ctx context.Context, params Query) ([]Book, error) {
	return query.Exec(ctx, r.DB, params)
}
