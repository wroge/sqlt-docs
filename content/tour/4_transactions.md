---
title: 4. Transactions
description: >
  This example shows how to chain statements in a transaction.
weight: 4
drivers: [modernc.org/sqlite]
scanners: [Scan]
executors: [Exec, First, Transaction]
configs: [ParseFiles, Masterminds/sprig]
---


{{< code language="go-template" source="tour/transactions/queries.sql" >}}{{< /code >}}  

{{< code language="go" source="tour/transactions/repository.go" >}}{{< /code >}}

