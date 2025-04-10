---
title: ScanJSON
date: 2025-04-06
description: >
  ScanJSON can be used to scan any type as JSON.
drivers: []
scanners: [ScanInt, ScanJSON]
executors: [One]
configs: []
weight: 6
---

{{% pageinfo %}}
{{ ScanJSON Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID      int64
  Authors []string
}

var queryBook = sqlt.One[string, Book](sqlt.Parse(`
  SELECT
    books.id                 {{ ScanInt "ID" }}
    , JSON_AGG(authors.name) {{ ScanJSON "Authors" }}
  FROM books
  LEFT JOIN book_authors ON books.id = book_authors.book_id
  LEFT JOIN authors ON authors.id = book_authors.author_id
  WHERE books.title = {{ . }}
`))
{{< /code >}}
