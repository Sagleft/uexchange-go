package uexchange

import (
	"errors"
	"io/ioutil"
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
	req, err := http.NewRequest("GET", urlWithParams, nil)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}

func (c *Client) sendPOSTRequest(requestURL string, params url.Values) ([]byte, error) {
	// declare http client
	httpClient := &http.Client{}

	// create request
	req, err := http.NewRequest("POST", requestURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, errors.New("failed to send POST request: " + err.Error())
	}

	// set cookie
	if c.AuthToken != "" {
		req.AddCookie(&http.Cookie{
			Name:  "auth_token",
			Value: c.AuthToken,
		})
	}

	// set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	// send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
