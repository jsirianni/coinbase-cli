package auth

import (
    "github.com/jsirianni/coinbase-cli/internal/auth/env"
)

type Auth interface{
    GetAPIKey() (string, error)
    GetSecret() (string, error)
}

func NewEnv() Auth {
    return env.Env{}
}
