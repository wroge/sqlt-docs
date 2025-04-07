---
title: ScanInt64
date: 2025-04-06
description: >
  ScanInt64 can be used to scan int64 and *int64.
categories: [Scanner]
weight: 2
---

{{% pageinfo %}}
{{ ScanInt64 Field }}
{{% /pageinfo %}}

```go
type Book struct {
  ID      int64
  Pages   *int64
  Edition sql.Null[int64]
}

type Query struct {
  Title string
}

var queryBook = sqlt.First[Query, Book](sqlt.Parse(`
  SELECT
    id          {{ ScanInt64 "ID" }}
    , pages     {{ ScanInt64 "Pages" }}
    , edition   {{ Scan "Edition" }}
  FROM books
  WHERE title = {{ .Title }}
`))
```
