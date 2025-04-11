---
title: 1. Create statements
description: >
  This simple example demonstrates how to execute SQL statements and map the results to Go structs using sqlt.
weight: 1
drivers: [modernc.org/sqlite]
scanners: [Scan]
executors: [Exec, First]
configs: [Parse]
---

{{< code language="go" source="tour/create_statements/repository.go" >}}{{< /code >}}  

<div style="padding-top: 2em; text-align: right"><a href="/sqlt-docs/tour/2_load_from_file/">>> 2. Load from file</a></div>