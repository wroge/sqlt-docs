---
title: ScanUint
date: 2025-04-06
description: >
  ScanUint can be used to scan uints and pointer to uints.
weight: 3
drivers: []
scanners: [Scan, ScanUint]
executors: [First]
configs: []
---

{{% pageinfo %}}
{{ ScanUint Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID      uint
  Pages   *uint64
  Edition sql.Null[uint64]
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    id          {{ ScanUint "ID" }}
    , pages     {{ ScanUint "Pages" }}
    , edition   {{ Scan "Edition" }}
  FROM books
  WHERE title = {{ .Title }}
`))
{{< /code >}}
