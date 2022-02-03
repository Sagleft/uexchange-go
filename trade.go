package uexchange

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

func (c *Client) sendTradeTask(orderType string, pairSymbol string, amount, price float64) (int64, error) {
	reqFields := url.Values{}
	reqFields.Add("pair", pairSymbol)
	reqFields.Add("amount", strconv.FormatFloat(amount, 'f', 8, 64))
	reqFields.Add("price", strconv.FormatFloat(price, 'f', 8, 64))

	body, err := c.sendRequest(c.getAPIURL("market/"+orderType), "POST", reqFields)
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
	reqFields := url.Values{}
	reqFields.Add("order_id", strconv.FormatInt(orderID, 10))
	body, err := c.sendRequest(c.getAPIURL("market/hold"), "POST", reqFields)
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
	reqFields := url.Values{}
	reqFields.Add("order_id", strconv.FormatInt(orderID, 10))
	body, err := c.sendRequest(c.getAPIURL("market/cancel"), "POST", reqFields)
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
	body, err := c.sendRequest(c.getAPIURL("market/pairs"), "GET", url.Values{})
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
	return response.Result.Pairs, nil
}

// GetOrderBook by trade pair
func (c *Client) GetOrderBook(pairSymbol string) (*BookValueDataContainer, error) {
	reqFields := url.Values{}
	reqFields.Add("pair", pairSymbol)
	body, err := c.sendRequest(c.getAPIURL("market/panel"), "POST", reqFields)
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
	reqFields := url.Values{}
	reqFields.Add("pair", pairSymbol)
	body, err := c.sendRequest(c.getAPIURL("market/curlist"), "GET", reqFields)
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
