---
title: ScanStringSlice
date: 2025-04-06
description: >
  ScanStringSlice can be used to parse string values to []string.
weight: 9
drivers: []
scanners: [ScanStringSlice]
executors: [First]
configs: [Parse]
---

{{% pageinfo %}}
{{ ScanStringSlice Field Sep }}
{{% /pageinfo %}}

{{< code language="go" title="Example" >}}
var queryTags = sqlt.First[any, []string](sqlt.Parse(`
  SELECT 'hello,world'; {{ ScanStringSlice "" "," }}
`))
{{< /code >}}
