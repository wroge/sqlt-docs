---
title: ScanFloat
date: 2025-04-06
description: >
  ScanFloat can be used to scan float64 and *float64.
weight: 3
drivers: []
scanners: [Scan, ScanFloat]
executors: [All]
configs: []
---

{{% pageinfo %}}
{{ ScanUint Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID      uint64
  Rating  float64
}

var queryBook = sqlt.All[any, Book](sqlt.Parse(`
  SELECT
    id          {{ Scan "ID" }}
    , rating    {{ ScanFloat "Rating" }}
  FROM books;
`))
{{< /code >}}
