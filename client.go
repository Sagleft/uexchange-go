package uexchange

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// NewClient - ..
func NewClient() *Client {
	return &Client{}
}

func (c *Client) getAPIURL(endpoint string) string {
	return apiHost + ":" + apiPort + "/" + endpoint
}

func (c *Client) sendRequest(url string, requestType string, data map[string]interface{}) ([]byte, error) {
	// declare http client
	httpClient := &http.Client{}

	// encode data fields to json
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("failed to encode request data to json: " + err.Error())
	}

	// declare HTTP Method and Url
	switch requestType {
	default:
		return nil, errors.New("invalid request type given: " + requestType)
	case "POST":
		break
	case "GET":
		break
	}
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(dataBytes))
	if err != nil {
		return nil, errors.New("failed to send " + requestType + " request: " + err.Error())
	}

	// set cookie
	if c.AuthToken != "" {
		req.Header.Set("Cookie", "auth_token="+c.AuthToken)
	}

	// send request
	resp, err := httpClient.Do(req)

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
