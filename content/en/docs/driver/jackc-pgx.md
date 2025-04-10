---
title: jackc/pgx
date: 2025-04-06
description: >
  pgx is a pure Go driver and toolkit for PostgreSQL.
drivers: [jackc/pgx]
scanners: [ScanInt, ScanString]
executors: [First]
configs: [Dollar, NoExpirationCache]
---

{{< code language="go" title="Example" >}}
import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wroge/sqlt"
)

type Book struct {
  ID    int64
  Title string
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Dollar, sqlt.NoExpirationCache(100), sqlt.Parse(`
  SELECT
    id        {{ ScanInt "ID" }}
    , title   {{ ScanString "Title" }}
  FROM books
  WHERE title = {{ .Title }}
`))

db, err := sql.Open("pgx", "user=postgres password=secret host=localhost port=5432 database=pgx_test sslmode=disable")
if err != nil {
  return err
}

book, err := queryBook.Exec(ctx, db, Query{Title: "Moby-Dick"})
if err != nil {
  panic(err)
}
{{< /code >}}