package uexchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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
	requestFieldsMap := url.Values{}
	if s.OrderType != "" {
		requestFieldsMap.Add("status", s.OrderType)
	}
	if s.TaskType != "" {
		requestFieldsMap.Add("task", s.TaskType)
	}

	body, err := s.ExchangeClient.sendRequest(
		s.ExchangeClient.getAPIURL("orders"), // endpoint
		"GET",                                // request type
		requestFieldsMap,                     // reuqest fields
	)
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIOrdersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("decode request response: " + err.Error())
	}

	if !response.Success {
		return nil, fmt.Errorf("get orders list: %s", response.Error)
	}
	return &response.Result, nil
}

// GetOrderHistory - get orders history
func (c *Client) GetOrderHistory(orderID string) (*OrdersHistoryDataContainer, error) {
	reqFields := url.Values{}
	reqFields.Add("order_id", orderID)
	body, err := c.sendRequest(c.getAPIURL("orders/history"), "POST", reqFields)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	// decode response
	var response APIOrdersHistoryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("get order history: %s", response.Error)
	}
	return &response.Result, nil
}
