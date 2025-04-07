---
title: ScanBinary
date: 2025-04-06
description: >
  ScanBinary can be used to scan types that implements encoding.BinaryUnmarshaler.
categories: [Scanner]
weight: 7
---

{{% pageinfo %}}
{{ ScanBinary Field }}
{{% /pageinfo %}}

```go
type Book struct {
  ID        int64
  AmazonLink url.URL
}

var queryBook = sqlt.First[string, Book](sqlt.Parse(`
  SELECT
    books.id            {{ ScanInt64 "ID" }}
    , books.amazon_link {{ ScanBinary "AmazonLink" }}
  FROM books
  WHERE title = {{ . }}
`))
```
