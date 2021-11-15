package uexchange

import (
	"encoding/json"
	"errors"
)

// Buy currency. returns order ID
func (c *Client) Buy(pairSymbol string, amount, price float64) (int64, error) {
	body, err := sendRequest(c.getAPIURL("market/buy"), mapTable{
		"pair":       pairSymbol,
		"amount":     amount,
		"price":      price,
		"auth_token": c.AuthToken,
	})
	if err != nil {
		return 0, err
	}

	// decode response
	var response APITradeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return 0, errors.New("failed to get balance") // TODO
	}
	return response.Result.OrderID, nil
}
