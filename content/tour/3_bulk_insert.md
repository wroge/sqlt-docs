---
title: 3. Bulk insert
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

<div style="padding-top: 2em; text-align: right"><a href="/sqlt-docs/tour/4_transactions/">>> 4. Transactions</a></div>