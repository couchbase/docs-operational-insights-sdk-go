package main

import (
	"context"
	"fmt"

	"github.com/couchbase/gocbanalytics"
)

func queries() {
	cluster, err := cbanalytics.NewCluster(
		connStr,
		cbanalytics.NewBasicAuthCredential(username, password),
	)
	handleErr(err)

	scopeLevelQuery(context.Background(), cluster)
	clusterLevelQuery(context.Background(), cluster)
	positionalParamQuery(context.Background(), cluster)
	namedParamQuery(context.Background(), cluster)

	err = cluster.Close()
	handleErr(err)
}

func scopeLevelQuery(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::scopeLevelQuery[]
	scope := cluster.Database("my_database").Scope("my_scope")
	result, err := scope.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::scopeLevelQuery[]
}

func clusterLevelQuery(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::clusterLevelQuery[]
	result, err := cluster.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::clusterLevelQuery[]
}

func positionalParamQuery(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::positionalParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select ?=1",
		cbanalytics.NewQueryOptions().SetPositionalParameters([]interface{}{1}),
	)
	handleErr(err)
	// end::positionalParamQuery[]

	handleResult(result)
}

func namedParamQuery(ctx context.Context, cluster *cbanalytics.Cluster) {
	// tag::namedParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select $foo=1",
		cbanalytics.NewQueryOptions().SetNamedParameters(map[string]interface{}{"foo": 1}),
	)
	handleErr(err)
	// end::namedParamQuery[]

	handleResult(result)
}

func handleResult(result *cbanalytics.QueryResult) {
	// tag::handleResults[]
	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err := row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}

	if err := result.Err(); err != nil {
		handleErr(err)
	}
	// end::handleResults[]
}

func metadata(result *cbanalytics.QueryResult) {
	// tag::metadata[]
	meta, err := result.MetaData()
	handleErr(err)

	fmt.Printf("Got meta: %v", meta)
	// end::metadata[]
}
