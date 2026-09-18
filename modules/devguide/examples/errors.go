package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/couchbase/gocbanalytics"
)

func analyticsError(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::analyticsError[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select 1=1",
	)
	if err != nil {
		var analyticsErr cbanalytics.AnalyticsError
		if errors.As(err, &analyticsErr) {
			// Do something with this error.
		}

		// This error occurred out of a client-server interaction.
	}
	// end::analyticsError[]

	handleResult(result)
}

func queryError(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::queryError[]
	handleQueryError := func(err error) {
		if err != nil {
			var queryErr cbanalytics.QueryError
			if errors.As(err, &queryErr) {
				fmt.Printf("Error code: %d, error message: %s", queryErr.Code(), queryErr.Message())
				return
			}

			// This error isn't a result of query processing, possibly something like a connection error.
		}
	}

	result, err := cluster.ExecuteQuery(
		ctx,
		"selec 1=1", // Syntax error
	)
	handleQueryError(err)
	// end::queryError[]

	for row := result.NextRow(); row != nil; row = result.NextRow() {
	}

	err = result.Err()
	handleQueryError(err)

	handleResult(result)
}
