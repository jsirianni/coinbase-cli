package main

import (
    "os"

    "github.com/jsirianni/coinbase-cli/internal/auth"
    "github.com/jsirianni/coinbase-cli/internal/client/coinbase"
    "github.com/jsirianni/coinbase-cli/internal/logger/standard"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "coinbase",
	Short: "Command line utility for interacting with the Coinbase API",
}

func init() {
	cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().BoolVar(&logJson, "json", false, "enable json output")
    rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level")

}

func initConfig() {
    // skip config init if any of these sub commands
    // were passed
	y := []string{"help", "version"}
	for _, subCmd := range y {
		if subCmd == os.Args[1] {
			return
		}
	}

    // Try call InitLog() before logging anything
    InitLog()

    if err := InitAuth(); err != nil {
        log.Error(err)
        os.Exit(1)
    }
}

// InitLog will initialize the log driver
func InitLog() {
    // TODO: implement json logger

    /*log.SetFormatter(&log.TextFormatter{
        DisableTimestamp: true,
    })
    if logJson {
        log.SetFormatter(&log.JSONFormatter{})
    }
    log.SetOutput(os.Stdout)

    level, err := log.ParseLevel(logLevel)
    if err != nil {
        log.Error(errors.Wrap(err, fmt.Sprintf("Invalid log level %s, defaulting to INFO", logLevel)))
        level = log.InfoLevel
    }
    log.SetLevel(level)*/


    log = standard.New(logLevel)

}

// InitAuth will authenticate the coinbase client
func InitAuth() error {
    a := auth.NewEnv()
    key, err := a.GetAPIKey()
    if err != nil {
        return err
    }
    secret, err := a.GetSecret()
    if err != nil {
        return err
    }

    // main.go/client
    c = coinbase.ApiKeyClient(key, secret)
    return nil
}
