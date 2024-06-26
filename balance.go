package uexchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// GetBalance - get balance data for all coins
func (c *Client) GetBalance() ([]BalanceData, error) {
	body, err := c.sendRequest(c.getAPIURL("user/balance"), "GET", url.Values{})
	if err != nil {
		return nil, err
	}

	// validate response
	responseStr := string(body)
	firstResponseChar := responseStr[0:1]
	if firstResponseChar != "{" && firstResponseChar != "[" {
		return nil, errors.New(responseStr)
	}

	// decode response
	var response APIBalanceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("decode response: " + err.Error())
	}

	if !response.Success {
		return nil, fmt.Errorf("get balance: %s", response.Error)
	}
	return response.Result.AllBalance, nil
}
