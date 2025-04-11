---
title: ScanTime
date: 2025-04-06
description: >
  ScanTime can be used to scan time.Time and *time.Time.
weight: 5
drivers: []
scanners: [Scan, ScanTime]
executors: [All]
configs: [Parse]
---

{{% pageinfo %}}
{{ ScanTime Field }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
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
{{< /code >}}
