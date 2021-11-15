package uexchange

import (
	"encoding/json"
	"errors"
)

// Auth client
func (c *Client) Auth(cred Credentials) (*APIAuthResultContainer, error) {
	c.APICredentials = cred
	body, err := c.sendRequest(c.getAPIURL("user/login"), map[string]interface{}{
		"PublicKey": cred.AccountPublicKey,
		"password":  cred.Password,
		"2fa_pin":   cred.TwoFACode,
	})
	if err != nil {
		return nil, err
	}

	// decode response
	var response APIAuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("failed to decode request response: " + err.Error())
	}

	// set auth token
	c.AuthToken = response.Result.AuthToken
	return &response.Result, nil
}

// Logout - close auth session
func (c *Client) Logout() error {
	body, err := c.sendRequest(c.getAPIURL("user/logout"), mapTable{
		"auth_token": c.AuthToken,
	})
	if err != nil {
		return err
	}

	// decode response
	var response APIPlainResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return errors.New("failed to decode request response: " + err.Error())
	}

	if !response.Success {
		return errors.New("failed to logout") // TODO
	}
	return nil
}
