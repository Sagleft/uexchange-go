package uexchange

import (
	"encoding/json"
	"errors"
)

const (
	apiHost = "https://crp.is"
	apiPort = "8182"
)

// NewClient - ..
func NewClient() *Client {
	return &Client{}
}

func (c *Client) getAPIURL(endpoint string) string {
	return apiHost + ":" + apiPort + "/" + endpoint
}

// Auth client
func (c *Client) Auth(cred Credentials) error {
	c.APICredentials = cred
	body, err := sendRequest(c.getAPIURL("user/login"), map[string]interface{}{
		"PublicKey": cred.AccountPublicKey,
		"password":  cred.Password,
		"2fa_pin":   cred.TwoFACode,
	})
	if err != nil {
		return err
	}

	// decode response
	var response APIAuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.New("failed to decode request response: " + err.Error())
	}

	// set auth token
	c.AuthToken = response.Result.AuthToken
	return nil
}
