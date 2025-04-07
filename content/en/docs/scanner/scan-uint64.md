---
title: ScanUint64
date: 2025-04-06
description: >
  ScanUint64 can be used to scan uint64 and *uint64.
categories: [Scanner]
weight: 3
---

{{% pageinfo %}}
{{ ScanUint64 Field }}
{{% /pageinfo %}}

```go
type Book struct {
  ID      uint64
  Pages   *uint64
  Edition sql.Null[uint64]
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    id          {{ ScanUint64 "ID" }}
    , pages     {{ ScanUint64 "Pages" }}
    , edition   {{ Scan "Edition" }}
  FROM books
  WHERE title = {{ .Title }}
`))
```
