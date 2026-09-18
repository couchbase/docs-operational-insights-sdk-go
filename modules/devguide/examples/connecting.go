package main

import (
	"time"

	"github.com/couchbase/gocbanalytics"
)

func connecting() {
	// tag::connecting[]
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
	// end::connecting[]

	err = cluster.Close()
	handleErr(err)
}
