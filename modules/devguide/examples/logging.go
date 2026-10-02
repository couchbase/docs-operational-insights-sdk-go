package main

import (
	"context"
	"fmt"
	"os"

	"github.com/couchbase/gocbinsights"
	"github.com/sirupsen/logrus"
)

// #tag::loggerwrapper[]
type MyLogrusLogger struct {
	logrusLogger *logrus.Logger
}

func (logger *MyLogrusLogger) Error(format string, v ...interface{}) {
	logger.logrusLogger.Error(fmt.Sprintf(format, v...))
}

func (logger *MyLogrusLogger) Warn(format string, v ...interface{}) {
	logger.logrusLogger.Warn(fmt.Sprintf(format, v...))
}

func (logger *MyLogrusLogger) Info(format string, v ...interface{}) {
	logger.logrusLogger.Info(fmt.Sprintf(format, v...))
}

func (logger *MyLogrusLogger) Debug(format string, v ...interface{}) {
	logger.logrusLogger.Debug(fmt.Sprintf(format, v...))
}

func (logger *MyLogrusLogger) Trace(format string, v ...interface{}) {
	logger.logrusLogger.Trace(fmt.Sprintf(format, v...))
}

// #end::loggerwrapper[]

func logging() {
	// #tag::creation[]
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	opts := cbinsights.NewClusterOptions().SetLogger(&MyLogrusLogger{logger})
	// #end::creation[]

	connStr := "couchbases://..."
	username := "..."
	password := "..."

	cluster, err := cbinsights.NewCluster(
		connStr,
		cbinsights.NewBasicAuthCredential(username, password),
		opts,
	)
	handleErr(err)

	_, err = cluster.ExecuteQuery(context.Background(), "select 1")
	if err != nil {
		panic(err)
	}

	err = cluster.Close()
	if err != nil {
		panic(err)
	}
}

func builtInLogger() {
	// #tag::creationBuiltIn[]
	cbinsights.NewClusterOptions().SetLogger(cbinsights.NewInfoLogger())
	// #end::creationBuiltIn[]
}
