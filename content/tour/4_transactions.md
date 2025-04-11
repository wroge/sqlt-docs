---
title: 4. Transactions
description: >
  This example shows how to chain statements in a transaction.
weight: 4
drivers: [modernc.org/sqlite]
scanners: [Scan]
executors: [Exec, First, One, All, Transaction]
configs: [ParseFiles, Lookup, Masterminds/sprig]
---

{{% pageinfo %}}
Alternatively, you can always create a *sql.Tx yourself and use it with each statement.
{{% /pageinfo %}}

{{< code language="go-template" source="tour/transactions/queries.sql" >}}{{< /code >}}  

{{< code language="go" source="tour/transactions/repository.go" >}}{{< /code >}}

<div style="padding-top: 2em; text-align: right"><a href="/sqlt-docs/tour/5_complex_query/">>> 5. Complex Query</a></div>