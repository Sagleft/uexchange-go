package uexchange

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// NewClient - ..
func NewClient() *Client {
	return &Client{}
}

func (c *Client) getAPIURL(endpoint string) string {
	return apiHost + ":" + apiPort + "/" + endpoint
}

func (c *Client) sendRequest(requestURL string, requestType string, params url.Values) ([]byte, error) {
	switch requestType {
	default:
		return nil, errors.New("invalid request type given: " + requestType)
	case "POST":
		return c.sendPOSTRequest(requestURL, params)
	case "GET":
		return c.sendGETRequest(requestURL, params)
	}
}

func (c *Client) sendGETRequest(requestURL string, params url.Values) ([]byte, error) {
	// declare http client
	httpClient := &http.Client{}

	// create request
	urlWithParams := requestURL + "?" + params.Encode()
	req, err := http.NewRequest(requestTypeGET, urlWithParams, nil)
	if err != nil {
		return nil, err
	}

	// set cookie
	if c.AuthToken != "" {
		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: c.AuthToken,
		})
	}

	// send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %s", err)
	}

	defer resp.Body.Close()
	return body, nil
}

func (c *Client) sendPOSTRequest(requestURL string, params url.Values) ([]byte, error) {
	// declare http client
	httpClient := &http.Client{}

	// create request
	req, err := http.NewRequest(requestTypePOST, requestURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	// set cookie
	if c.AuthToken != "" {
		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: c.AuthToken,
		})
	}

	// set headers
	req.Header.Set(headerContentType, apiRequestContentType)

	// send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %s", err)
	}

	defer resp.Body.Close()
	return body, nil
}
