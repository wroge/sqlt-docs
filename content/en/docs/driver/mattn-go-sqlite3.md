---
title: mattn/go-sqlite3
date: 2025-04-06
description: >
  sqlite3 driver for go using database/sql.
drivers: [mattn/go-sqlite3]
scanners: [ScanInt, ScanString]
executors: [First]
configs: []
---

{{% pageinfo %}}
go-sqlite3 is cgo package. If you want to build your app using go-sqlite3, you need gcc.
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wroge/sqlt"
)

type Book struct {
  ID    int64
  Title string
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    id        {{ ScanInt "ID" }}
    , title   {{ ScanString "Title" }}
  FROM books
  WHERE title = {{ .Title }}
`))

db, err := sql.Open("sqlite3", ":memory:?_fk=1")
if err != nil {
  panic(err)
}

book, err := queryBook.Exec(ctx, db, Query{Title: "Moby-Dick"})
if err != nil {
  panic(err)
}
{{< /code >}}