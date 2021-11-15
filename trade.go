package uexchange

import (
	"encoding/json"
	"errors"
)

func (c *Client) sendTradeTask(orderType string, pairSymbol string, amount, price float64) (int64, error) {
	body, err := c.sendRequest(c.getAPIURL("market/"+orderType), mapTable{
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

// Buy currency. returns order ID
func (c *Client) Buy(pairSymbol string, amount, price float64) (int64, error) {
	return c.sendTradeTask("buy", pairSymbol, amount, price)
}

// Sell currency. returns order ID
func (c *Client) Sell(pairSymbol string, amount, price float64) (int64, error) {
	return c.sendTradeTask("sell", pairSymbol, amount, price)
}
