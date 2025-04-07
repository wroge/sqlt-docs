---
title: mattn/go-sqlite3
date: 2025-04-06
description: >
  sqlite3 driver for go using database/sql.
categories: [Driver, Sqlite]
---

{{% pageinfo %}}
go-sqlite3 is cgo package. If you want to build your app using go-sqlite3, you need gcc.
{{% /pageinfo %}}

```go
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
    id        {{ ScanInt64 "ID" }}
    , title   {{ ScanString "Title" }}
  FROM books
  WHERE title = {{ .Title }}
`))

db, err := sql.Open("sqlite3", "test.db?_fk=1")
if err != nil {
  panic(err)
}

book, err := queryBook.Exec(ctx, db, Query{Title: "Moby-Dick"})
if err != nil {
  panic(err)
}
```