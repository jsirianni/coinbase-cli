package main

import (
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Interact with accounts",
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
