---
title: 3. Bulk Insert
description: >
  This example shows how to create statements for bulk inserts.
weight: 3
drivers: [modernc.org/sqlite]
scanners: [Scan]
executors: [Exec, First, All]
configs: [ParseFiles]
---


{{< code language="go-template" source="tour/bulk_insert/queries.sql" >}}{{< /code >}}  

{{< code language="go" source="tour/bulk_insert/repository.go" >}}{{< /code >}}

