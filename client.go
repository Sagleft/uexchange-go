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

func (c *Client) sendRequest(url string, data map[string]interface{}) ([]byte, error) {
	// encode data fields to json
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("failed to encode request data to json: " + err.Error())
	}

	// send request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return nil, errors.New("failed to send POST request: " + err.Error())
	}
	defer resp.Body.Close()

	// read response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}
	return responseBody, nil
}
