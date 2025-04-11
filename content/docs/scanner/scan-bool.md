---
title: ScanBool
date: 2025-04-06
description: >
  ScanBool can be used to scan bool and *bool.
drivers: []
scanners: [ScanBool]
executors: [First]
configs: [Parse]
weight: 4
---

{{% pageinfo %}}
{{ ScanBool Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Favorite bool

type Book struct {
  Favorite Favorite
  Read     bool
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    true    {{ ScanBool "Favorite" }}
    , read  {{ ScanBool "Read" }}
  FROM books
  WHERE title = {{ .Title }}
`))
{{< /code >}}
