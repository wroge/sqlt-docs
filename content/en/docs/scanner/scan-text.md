---
title: ScanText
date: 2025-04-06
description: >
  ScanText can be used to scan types that implements encoding.TextUnmarshaler.
categories: [Scanner]
weight: 8
---

{{% pageinfo %}}
{{ ScanText Field }}
{{% /pageinfo %}}

```go
type Book struct {
  ID    int64
  Sales big.Int
}

var queryBook = sqlt.First[string, Book](sqlt.Parse(`
  SELECT
    id         {{ ScanInt64 "ID" }}
    , sales    {{ ScanText "Sales" }}
  FROM books
  WHERE title = {{ . }}
`))
```
