package uexchange

import (
	"encoding/json"
	"errors"
)

// GetTradeHistory - get trading history by pairs
func (c *Client) GetTradeHistory(pairSymbol string) (*TradeHistoryDataContainer, error) {
	body, err := c.sendRequest(c.getAPIURL("history/trade"), mapTable{
		"pair": pairSymbol,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APITradeHistoryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get trade history") // TODO
	}
	return &response.Result, nil
}
