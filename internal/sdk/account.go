package sdk

import (
    "time"
    "encoding/json"
)

type Accounts struct {
	Pagination struct {
		EndingBefore         interface{} `json:"ending_before"`
		StartingAfter        interface{} `json:"starting_after"`
		PreviousEndingBefore interface{} `json:"previous_ending_before"`
		NextStartingAfter    string      `json:"next_starting_after"`
		Limit                int         `json:"limit"`
		Order                string      `json:"order"`
		PreviousURI          interface{} `json:"previous_uri"`
		NextURI              string      `json:"next_uri"`
	} `json:"pagination"`
	Data []Account `json:"data"`
    Warnings struct {
        ID      string `json:"id"`
        Message string `json:"message"`
        URL     string `json:"url"`
    } `json:"warnings"`
}

type Account struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Primary  bool   `json:"primary"`
    Type     string `json:"type"`
    Currency struct {
        Code         string `json:"code"`
        Name         string `json:"name"`
        Color        string `json:"color"`
        SortIndex    int    `json:"sort_index"`
        Exponent     int    `json:"exponent"`
        Type         string `json:"type"`
        AddressRegex string `json:"address_regex"`
        AssetID      string `json:"asset_id"`
        Slug         string `json:"slug"`
    } `json:"currency,omitempty"`
    Balance struct {
        Amount   string `json:"amount"`
        Currency string `json:"currency"`
    } `json:"balance"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
    Resource         string    `json:"resource"`
    ResourcePath     string    `json:"resource_path"`
    AllowDeposits    bool      `json:"allow_deposits"`
    AllowWithdrawals bool      `json:"allow_withdrawals"`
    Rewards          struct {
        Apy          string `json:"apy"`
        FormattedApy string `json:"formatted_apy"`
        Label        string `json:"label"`
    } `json:"rewards,omitempty"`
}


func (a Accounts) ToBytes() ([]byte, error) {
    return json.Marshal(a)
}
