---
title: ScanText
date: 2025-04-06
description: >
  ScanText can be used to scan types that implements encoding.TextUnmarshaler.
weight: 8
drivers: []
scanners: [ScanInt, ScanText]
executors: [First]
configs: [Parse]
---

{{% pageinfo %}}
{{ ScanText Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
type Book struct {
  ID    int64
  Sales big.Int
}

var queryBook = sqlt.First[string, Book](sqlt.Parse(`
  SELECT
    id         {{ ScanInt "ID" }}
    , sales    {{ ScanText "Sales" }}
  FROM books
  WHERE title = {{ . }}
`))
{{< /code >}}
