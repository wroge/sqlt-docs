---
title: 2. Load from File
description: >
  Learn how to load SQL from a file and execute type-safe statements using sqlt.
weight: 2
drivers: [modernc.org/sqlite]
scanners: [ScanInt, ScanString]
executors: [Exec, First]
configs: [ParseFiles, Lookup]
---


{{< code language="go-template" source="tour/load_from_file/queries.sql" >}}{{< /code >}}  

{{< code language="go" source="tour/load_from_file/repository.go" >}}{{< /code >}}

