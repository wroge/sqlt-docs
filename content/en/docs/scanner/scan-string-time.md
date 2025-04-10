---
title: ScanStringTime
date: 2025-04-06
description: >
  ScanStringTime can be used to parse string values to time.Time.
weight: 10
drivers: []
scanners: [ScanStringTime]
executors: [One]
configs: []
---

{{% pageinfo %}}
{{ ScanStringTime Field Layout Location }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
var queryDate = sqlt.One[string, time.Time](sqlt.Parse(`
  SELECT
    '2025-04-06' {{ ScanStringTime "" "DateOnly" "UTC" }}
  FROM books
  WHERE title = {{ . }}
`))

// Predefined layouts:
var layoutMap = map[string]string{
	"DateTime":    time.DateTime,
	"DateOnly":    time.DateOnly,
	"TimeOnly":    time.TimeOnly,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Layout":      time.Layout,
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
}
{{< /code >}}
