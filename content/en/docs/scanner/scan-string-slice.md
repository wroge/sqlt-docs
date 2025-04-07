---
title: ScanStringSlice
date: 2025-04-06
description: >
  ScanStringSlice can be used to parse string values to []string.
categories: [Scanner]
---

{{% pageinfo %}}
{{ ScanStringSlice Field Sep }}
{{% /pageinfo %}}

```go
var queryTags = sqlt.All[any, []string](sqlt.Parse(`
  SELECT 'hello,world'; {{ ScanStringSlice "" "," }}
`))
```
