package main

import (
	"context"
	"fmt"
	"time"

	"github.com/couchbase/gocbanalytics"
)

func helloWorld() {
	cluster, err := cbanalytics.NewCluster(
		connStr,
		cbanalytics.NewBasicAuthCredential(username, password),
		// The third parameter is optional.
		// This example sets the default server query timeout to 3 minutes,
		// that is the timeout value sent to the query server.
		cbanalytics.NewClusterOptions().SetTimeoutOptions(
			cbanalytics.NewTimeoutOptions().SetQueryTimeout(3*time.Minute),
		),
	)
	handleErr(err)

	// We create a new context with a timeout of 2 minutes.
	// This context will apply timeout/cancellation on the client side.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	printRows := func(result *cbanalytics.QueryResult) {
		for row := result.NextRow(); row != nil; row = result.NextRow() {
			var content map[string]interface{}

			err = row.ContentAs(&content)
			handleErr(err)

			fmt.Printf("Got row content: %v", content)
		}
	}

	// Execute a query and process rows as they arrive from server.
	// Results are always streamed from the server.
	result, err := cluster.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	printRows(result)

	// Execute a query with positional arguments.
	result, err = cluster.ExecuteQuery(
		ctx,
		"select ?=1",
		cbanalytics.NewQueryOptions().SetPositionalParameters([]interface{}{1}),
	)
	handleErr(err)

	printRows(result)

	// Execute a query with named arguments.
	result, err = cluster.ExecuteQuery(
		ctx,
		"select $foo=1",
		cbanalytics.NewQueryOptions().SetNamedParameters(map[string]interface{}{"foo": 1}),
	)
	handleErr(err)

	printRows(result)

	err = cluster.Close()
	handleErr(err)
}
