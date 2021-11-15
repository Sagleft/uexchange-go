package uexchange

import (
	"encoding/json"
	"errors"
)

// GetBalance - get balance data for all coins
func (c *Client) GetBalance() ([]BalanceData, error) {
	body, err := sendRequest(c.getAPIURL("user/balance"), mapTable{
		"auth_token": c.AuthToken,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIBalanceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get balance") // TODO
	}
	return response.Result.AllBalance, nil
}
