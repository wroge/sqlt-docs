---
title: ScanString
date: 2025-04-06
description: >
  ScanString can be used to scan string and *string.
weight: 1
drivers: []
scanners: [Scan, ScanString]
executors: [One]
configs: []
---

{{% pageinfo %}}
{{ ScanString Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  Author   string
  CoAuthor *string
  Title    sql.Null[string]
}

var queryBook = sqlt.One[int64, Book](sqlt.Parse(`
  SELECT
    author       {{ ScanString "Author" }}
    , co_author  {{ ScanString "CoAuthor" }}
    , title      {{ Scan "Title" }}
  FROM books
  WHERE id = {{ . }}
`))
{{< /code >}}
