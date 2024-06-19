package uexchange

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetTradeHistory - get trading history by pairs
func (c *Client) GetTradeHistory(pairSymbol string) (*TradeHistoryDataContainer, error) {
	reqFields := url.Values{}
	reqFields.Add("pair", pairSymbol)
	body, err := c.sendRequest(c.getAPIURL("history/trade"), requestTypeGET, reqFields)
	if err != nil {
		return nil, err
	}

	// decode response
	var response APITradeHistoryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("get trade history: %s", response.Error)
	}
	return &response.Result, nil
}

// GetAccountHistoryService - get operations history service
type GetAccountHistoryService struct {
	// required
	ExchangeClient *Client
	RequestType    string // history type: profile/trade/billing

	// optional
	FromID     string // pagination offset: uuid
	RecordType string // billing operation type (only for billing): payment/comission/withdraw or combined
	Currency   string // currency (only for billing type)
}

// NewGetAccountHistoryService - create new get account history service
func (c *Client) NewGetAccountHistoryService(requestType string) *GetAccountHistoryService {
	// TODO: validate request type
	return &GetAccountHistoryService{
		ExchangeClient: c,
		RequestType:    requestType,
	}
}

// SetRequestType - set request type for request
func (s *GetAccountHistoryService) SetRequestType(newRequestType string) *GetAccountHistoryService {
	s.RequestType = newRequestType
	return s
}

// SetFromID - set form ID for request: pagination offset: uuid
func (s *GetAccountHistoryService) SetFromID(newID string) *GetAccountHistoryService {
	s.FromID = newID
	return s
}

// SetRecordType - set record type for request:
// billing operation type (only for billing): payment/comission/withdraw or combined
func (s *GetAccountHistoryService) SetRecordType(newRecordType string) *GetAccountHistoryService {
	// TODO: validate record type
	s.RecordType = newRecordType
	return s
}

// SetCurrency - set currency for request: currency (only for billing type)
func (s *GetAccountHistoryService) SetCurrency(newCurrency string) *GetAccountHistoryService {
	s.Currency = newCurrency
	return s
}

// Do request
func (s *GetAccountHistoryService) Do() (*OperationsHistoryDataContainer, error) {
	requestFieldsMap := url.Values{}
	requestFieldsMap.Add("type", s.RequestType)
	if s.FromID != "" {
		requestFieldsMap.Add("from_id", s.FromID)
	}
	if s.RecordType != "" {
		requestFieldsMap.Add("record_type", s.RecordType)
	}
	if s.Currency != "" {
		requestFieldsMap.Add("currency", s.Currency)
	}

	body, err := s.ExchangeClient.sendRequest(
		s.ExchangeClient.getAPIURL("history"), // endpoint
		"POST",                                // request type
		requestFieldsMap,                      // request fields
	)
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIOperationsHistoryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("decode response: %s", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("get operations history: %s", response.Error)
	}
	return &response.Result, nil
}
