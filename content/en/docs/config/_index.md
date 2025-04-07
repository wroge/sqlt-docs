---
title: Config
description: Config options.
weight: 2
---

```go
import "github.com/Masterminds/sprig/v3"

// sqlt.Config implements sqlt.Option
config := sqlt.Config{
	Placeholder: sqlt.Dollar,
	Templates: []sqlt.Template{
		sqlt.Funcs(sprig.TxtFuncMap()),
	},
	Log: func(ctx context.Context, info sqlt.Info) {
		fmt.Println(info.SQL)
	},
}

// sqlt.First(opts ...sqlt.Option)
var queryBook = sqlt.First[Query, Book](config, sqlt.NoExpirationCache(100), sqlt.Parse(`
  SELECT
    id        {{ ScanInt64 "ID" }}
    , title   {{ ScanString "Title" }}
  FROM books
  WHERE title = {{ .Title }}
`))
```
