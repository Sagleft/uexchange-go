package uexchange

import (
	"encoding/json"
	"errors"
)

func (c *Client) sendTradeTask(orderType string, pairSymbol string, amount, price float64) (int64, error) {
	body, err := c.sendRequest(c.getAPIURL("market/"+orderType), mapTable{
		"pair":   pairSymbol,
		"amount": amount,
		"price":  price,
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
		return 0, errors.New("failed to send trade task") // TODO
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

// Hold or Unhold order
func (c *Client) Hold(orderID int64) error {
	body, err := c.sendRequest(c.getAPIURL("market/hold"), mapTable{
		"order_id": orderID,
	})
	if err != nil {
		return err
	}

	// decode response
	var response APITradeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return errors.New("failed to send trade task") // TODO
	}
	return nil
}

// Cancel the specified order
func (c *Client) Cancel(orderID int64) error {
	body, err := c.sendRequest(c.getAPIURL("market/cancel"), mapTable{
		"order_id": orderID,
	})
	if err != nil {
		return err
	}

	// decode response
	var response APITradeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return errors.New("failed to send trade task") // TODO
	}
	return nil
}

// GetPairs - get trading pairs list
func (c *Client) GetPairs() ([]PairsDataContainer, error) {
	body, err := c.sendRequest(c.getAPIURL("market/pairs"), mapTable{})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIPairsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get pairs list") // TODO
	}
	return response.Result, nil
}

// GetOrderBook by trade pair
func (c *Client) GetOrderBook(pairSymbol string) (*BookValueDataContainer, error) {
	body, err := c.sendRequest(c.getAPIURL("market/pairs"), mapTable{
		"pair": pairSymbol,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIBookValueResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get order book by pair") // TODO
	}
	return &response.Result, nil
}

// GetMarketCurrenciesList - get exchange currencies list
func (c *Client) GetMarketCurrenciesList(pairSymbol string) (*CurrenciesListData, error) {
	body, err := c.sendRequest(c.getAPIURL("market/pairs"), mapTable{
		"pair": pairSymbol,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APICurrenciesListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get currencies list") // TODO
	}
	return &response.Result, nil
}
