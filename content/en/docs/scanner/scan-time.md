---
title: ScanTime
date: 2025-04-06
description: >
  ScanTime can be used to scan time.Time and *time.Time.
categories: [Scanner]
weight: 5
---

{{% pageinfo %}}
{{ ScanTime Field }}
{{% /pageinfo %}}

```go
type Book struct {
  Added     time.Time
  Updated   *time.Time
  Published sql.Null[time.Time]
}

var queryBooks = sqlt.All[string, Book](sqlt.Parse(`
  SELECT
    added_at        {{ ScanTime "Added" }}
    , updated_at    {{ ScanTime "Updated" }}
    , published_at  {{ Scan "Published" }}
  FROM books
  WHERE author = {{ . }}
`))
```
