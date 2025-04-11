---
title: ScanInt
date: 2025-04-06
description: >
  ScanInt can be used to scan ints and pointer to ints.
drivers: []
scanners: [Scan, ScanInt]
executors: [First]
configs: [Parse]
weight: 2
---

{{% pageinfo %}}
{{ ScanInt Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID      int
  Pages   *int64
  Edition sql.Null[int64]
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    id          {{ ScanInt "ID" }}
    , pages     {{ ScanInt "Pages" }}
    , edition   {{ Scan "Edition" }}
  FROM books
  WHERE title = {{ .Title }}
`))
{{< /code >}}
