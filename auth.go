package uexchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// Auth client
func (c *Client) Auth(cred Credentials) (*APIAuthResultContainer, error) {
	if cred.AccountPublicKey == "" {
		return nil, errors.New("account public key is not set")
	}
	if cred.Password == "" {
		return nil, errors.New("password is not set")
	}

	c.APICredentials = cred
	reqFields := url.Values{}
	reqFields.Add("PublicKey", cred.AccountPublicKey)
	reqFields.Add("password", cred.Password)
	reqFields.Add("2fa_pin", cred.TwoFACode)

	body, err := c.sendRequest(c.getAPIURL("user/login"), requestTypePOST, reqFields)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	// decode response
	var response APIAuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	if !response.Success {
		return nil, fmt.Errorf("auth: %s", response.Error)
	}

	// set auth token
	c.AuthToken = response.Result.AuthToken
	return &response.Result, nil
}

// Logout - close auth session
func (c *Client) Logout() error {
	body, err := c.sendRequest(c.getAPIURL("user/logout"), requestTypePOST, url.Values{})
	if err != nil {
		return err
	}

	// decode response
	var response APIPlainResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	if !response.Success {
		return fmt.Errorf("logout: %s", response.Error)
	}
	return nil
}
