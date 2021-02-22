package coinbase

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/jsirianni/coinbase-cli/internal/sdk"

    "github.com/pkg/errors"
)

type Coinbase struct {
    auth Auth
}

type Auth struct {
    Key string
    Secret string
}

func ApiKeyClient(key, secret string) Coinbase {
    return Coinbase{
        auth: Auth{
            key,
            secret,
        },
    }
}

func (c Coinbase) GetAccounts() (sdk.Accounts, error) {
    // TOOD: use limit and starting_after parameters to ensure
    // all accounts are accounted for
    path := "/v2/accounts?limit=100"
    a, err := c.getAccounts(path)
    if err != nil {
        return sdk.Accounts{}, err
    }
    return a, nil
}

func (c Coinbase) getAccounts(path string) (sdk.Accounts, error) {
    client := &http.Client{
        // TODO: set values like timeout
    }

    key := c.auth.Key
    secret := c.auth.Secret
    method := "GET"
    baseURI := "https://api.coinbase.com"
    uri     := baseURI + path
    reqBody := []byte{}

    req, err := http.NewRequest(method, uri, nil)
    if err != nil {
        return sdk.Accounts{}, err
    }

    if err := sdk.Authenticate(req, key, secret, method, path, reqBody); err != nil {
        return sdk.Accounts{}, errors.Wrap(err, "failed to add authentication headers to request")
    }

    resp, err := client.Do(req)
    if err != nil {
        return sdk.Accounts{}, errors.Wrap(err, "request to api failed")
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return sdk.Accounts{}, errors.Wrap(err, "failed to read response body")
    }

    if resp.StatusCode != http.StatusOK {
        return sdk.Accounts{}, fmt.Errorf("expected status ok" + string(bodyBytes))
    }

    account := sdk.Accounts{}
    err = json.Unmarshal(bodyBytes, &account)
    return account, err
}
