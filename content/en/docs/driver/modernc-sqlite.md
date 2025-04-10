---
title: modernc.org/sqlite
date: 2025-04-06
description: >
  Package sqlite is a sql/database driver using a CGo-free port of the C SQLite3 library.
drivers: [modernc.org/sqlite]
scanners: [ScanInt, ScanString]
executors: [One]
configs: [UnlimitedSizeCache]
---

{{< code language="go" title="Example" >}}
import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
	"github.com/wroge/sqlt"
)

type Book struct {
  ID    int64
  Title string
}

type Query struct {
  Title string
}

var queryBook = sqlt.One[Query, Book](sqlt.UnlimitedSizeCache(time.Second), sqlt.Parse(`
  SELECT
    id        {{ ScanInt "ID" }}
    , title   {{ ScanString "Title" }}
  FROM books
  WHERE title = {{ .Title }}
`))

db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)")
if err != nil {
  panic(err)
}

book, err := queryBook.Exec(ctx, db, Query{Title: "Moby-Dick"})
if err != nil {
  panic(err)
}
{{< /code >}}