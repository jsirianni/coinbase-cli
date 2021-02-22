package env

import (
    "os"
    "fmt"
)

const envAPIKey = "COINBASE_API_KEY"
const envSecret = "COINBASE_SECRET"

type Env struct {

}

func (e Env) GetAPIKey() (string, error) {
    return getEnvNotEmpty(envAPIKey)
}

func (e Env) GetSecret() (string, error) {
    return getEnvNotEmpty(envSecret)
}

func getEnvNotEmpty(e string) (string, error) {
    x := os.Getenv(e)
    if x == "" {
        return "", fmt.Errorf(fmt.Sprintf("Expected environment variable %s to be set but it was empty", e))
    }
    return x, nil
}
