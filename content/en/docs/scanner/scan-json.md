---
title: ScanJSON
date: 2025-04-06
description: >
  ScanJSON can be used to scan any type as JSON.
categories: [Scanner]
weight: 6
---

{{% pageinfo %}}
{{ ScanJSON Field }}
{{% /pageinfo %}}

```go
type Book struct {
  ID      int64
  Authors []string
}

var queryBook = sqlt.First[string, Book](sqlt.Parse(`
  SELECT
    books.id                 {{ ScanInt64 "ID" }}
    , JSON_AGG(authors.name) {{ ScanJSON "Authors" }}
  FROM books
  LEFT JOIN book_authors ON books.id = book_authors.book_id
  LEFT JOIN authors ON authors.id = book_authors.author_id
  WHERE books.title = {{ . }}
`))
```
