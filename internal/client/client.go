package client

import (
    "github.com/jsirianni/coinbase-cli/internal/sdk"
)

type Client interface {
    GetAccounts() (sdk.Accounts, error)
}
