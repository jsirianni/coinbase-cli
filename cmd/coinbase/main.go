package main

import (
    "os"
    "fmt"

    "github.com/jsirianni/coinbase-cli/internal/client"
    "github.com/jsirianni/coinbase-cli/internal/logger"
)

var c client.Client
var log logger.Logger

// global flags, root.go
var (
    logJson bool
    logLevel string
)

// account_list.go
var (
    includeEmpty bool
)

func main() {
    if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
