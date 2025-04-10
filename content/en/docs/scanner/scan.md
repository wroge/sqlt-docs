---
title: Scan
date: 2025-04-06
description: >
  Scan can be used to scan sql.Scanner, string, int, uint, float or time.Time types.
weight: 1
drivers: []
scanners: [Scan]
executors: [One]
configs: []
---

{{% pageinfo %}}
{{ Scan Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID        int64
  Title     string 
  Author    string
  CoAuthor  *string
  Published *time.Time
  Pages     sql.Null[int64]
}

var queryBook = sqlt.One[int64, Book](sqlt.Parse(`
  SELECT
    id            {{ Scan "ID" }}
    , title       {{ Scan "Title" }}
    , author      {{ Scan "Author" }}
    , co_author   {{ Scan "CoAuthor" }}
    , published   {{ Scan "Published" }}
    , pages       {{ Scan "Pages" }}
  FROM books
  WHERE id = {{ . }}
`))
{{< /code >}}
