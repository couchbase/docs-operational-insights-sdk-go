// tag::server-async[]
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	cbanalytics "github.com/couchbase/gocbanalytics"
)

func waitForQueryResults(ctx context.Context, handle *cbanalytics.QueryHandle, delay time.Duration) (*cbanalytics.QueryResultHandle, error) {
	var lastStatus *cbanalytics.QueryStatus
	deadline, _ := ctx.Deadline()

	for {
		status, err := handle.FetchStatus(ctx)
		if err != nil {
			fmt.Printf("Error fetching query results: %v\n", err)
		} else {
			if status.ResultsReady() {
				return status.ResultHandle()
			}
			lastStatus = status
		}

		nextPoll := time.Now().Add(delay)
		if deadline.Before(nextPoll) {
			return nil, fmt.Errorf("query results not ready")
		}

		if lastStatus != nil {
			fmt.Printf("Query status: %s\n", lastStatus)
		}
		fmt.Printf("Query results not ready yet, sleeping for %s...\n", delay)
		time.Sleep(delay)
	}
}

func main() {
	endpoint := "https://localhost"
	username := "Administrator"
	password := "password"

	cred := cbanalytics.NewBasicAuthCredential(username, password)
	cluster, err := cbanalytics.NewCluster(endpoint, cred)
	if err != nil {
		log.Fatalf("failed to create cluster: %v", err)
	}
	defer func() {
		if err := cluster.Close(); err != nil {
			log.Printf("failed to close cluster: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	statement := `SELECT VALUE SLEEP("x", 100) FROM RANGE(1, 100) AS id;`

	handle, err := cluster.StartQuery(ctx, statement)
	if err != nil {
		log.Fatalf("failed to start query: %v", err)
	}

	resultHandle, err := waitForQueryResults(ctx, handle, 2500*time.Millisecond)
	if err != nil {
		log.Fatalf("waiting for results: %v", err)
	}

	res, err := resultHandle.FetchResults(ctx)
	if err != nil {
		log.Fatalf("failed to fetch results: %v", err)
	}

	for row := res.NextRow(); row != nil; row = res.NextRow() {
		var val interface{}
		if err := row.ContentAs(&val); err != nil {
			log.Printf("failed to decode row: %v", err)
			continue
		}
		fmt.Printf("Found row: %v\n", val)
	}

	if err := res.Err(); err != nil {
		log.Fatalf("result error: %v", err)
	}

	metadata, err := res.MetaData()
	if err != nil {
		log.Fatalf("failed to get metadata: %v", err)
	}
	fmt.Printf("metadata=%+v\n", metadata)

	if err := resultHandle.DiscardResults(ctx); err != nil {
		log.Printf("failed to discard results: %v", err)
	}
}
// end::server-async[]
