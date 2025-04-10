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
	schema = sqlt.Exec[any](sqlt.Parse(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL
		)
	`))

	create = sqlt.First[string, int64](sqlt.Parse(`
		INSERT INTO books (title) VALUES ({{ . }}) RETURNING id;
	`))

	get = sqlt.First[int64, Book](sqlt.Parse(`
		SELECT 
			id              {{ Scan "ID" }}
			, title         {{ Scan "Title" }}
		FROM books
		WHERE id = {{ . }};
	`))
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
