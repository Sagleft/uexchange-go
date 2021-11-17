package uexchange

import (
	"encoding/json"
	"errors"
)

// GetOrdersService - get orders service with optional params
type GetOrdersService struct {
	ExchangeClient *Client
	OrderType      string // order type: open/close/cancel/hold
	TaskType       string // task type: buy/sell
}

// NewGetOrdersService - create new get orders service
func (c *Client) NewGetOrdersService() *GetOrdersService {
	return &GetOrdersService{
		ExchangeClient: c,
	}
}

// SetOrderType - set order type: open/close/cancel/hold
func (s *GetOrdersService) SetOrderType(newType string) *GetOrdersService {
	s.OrderType = newType
	return s
}

// SetTaskType - set task type: buy/sell
func (s *GetOrdersService) SetTaskType(taskType string) *GetOrdersService {
	s.TaskType = taskType
	return s
}

// Do request
func (s *GetOrdersService) Do() (*OrdersDataContainer, error) {
	requestFieldsMap := mapTable{}
	if s.OrderType != "" {
		requestFieldsMap["status"] = s.OrderType
	}
	if s.TaskType != "" {
		requestFieldsMap["task"] = s.TaskType
	}

	body, err := s.ExchangeClient.sendRequest(
		s.ExchangeClient.getAPIURL("orders"), // endpoint
		requestFieldsMap,                     // reuqest fields
	)
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIOrdersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get orders list") // TODO
	}
	return &response.Result, nil
}

// GetOrderHistory - get orders history
func (c *Client) GetOrderHistory(orderID string) (*OrdersHistoryDataContainer, error) {
	body, err := c.sendRequest(c.getAPIURL("orders/history"), mapTable{
		"order_id": orderID,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIOrdersHistoryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, errors.New("failed to get order history") // TODO
	}
	return &response.Result, nil
}
