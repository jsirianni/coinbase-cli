package main

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/jedib0t/go-pretty/v6/table"
)

var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "List accounts",
	Run: func(cmd *cobra.Command, args []string) {
		accountList()
	},
}

func init() {
	accountCmd.AddCommand(accountListCmd)
	accountListCmd.PersistentFlags().BoolVar(&includeEmpty, "include-empty", false, "show empty accounts")

}

func accountList() {
	accounts, err := c.GetAccounts()
	if err != nil {
		log.Error(err)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Type", "balance", "ID"})
	for _, a := range accounts.Data {
		balance, err := strconv.ParseFloat(a.Balance.Amount, 64)
		if err != nil {
			log.Error(err)
			continue
		}
		if balance == 0 && ! includeEmpty {
			continue
		}
		t.AppendRow([]interface{}{a.Currency.Name, a.Type, a.Balance.Amount, a.ID})
	}
	t.Render()
}
