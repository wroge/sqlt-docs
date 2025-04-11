---
title: 5. Complex Query
description: >
  This example shows how to build complex queries.
weight: 5
drivers: [modernc.org/sqlite]
scanners: [Scan, ScanStringSlice]
executors: [Exec, Transaction, All, One]
configs: [ParseFiles, Lookup, Masterminds/sprig, Log, NoExpirationCache]
---


{{< code language="go-template" source="tour/complex_query/queries.sql" >}}{{< /code >}}  

{{< code language="go" source="tour/complex_query/repository.go" >}}{{< /code >}}

