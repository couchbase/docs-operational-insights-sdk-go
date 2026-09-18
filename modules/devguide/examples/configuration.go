package main

import (
	"time"

	"github.com/couchbase/gocbanalytics"
)

func configuration() {
	// #tag::configuration[]
	cluster, err := cbanalytics.NewCluster(
		connStr,
		cbanalytics.NewBasicAuthCredential(username, password),
		cbanalytics.NewClusterOptions().
			SetTimeoutOptions(
				cbanalytics.NewTimeoutOptions().
					SetConnectTimeout(30*time.Second).
					SetQueryTimeout(2*time.Minute),
			),
	)
	handleErr(err)
	// #end::configuration[]

	err = cluster.Close()
	handleErr(err)
}
