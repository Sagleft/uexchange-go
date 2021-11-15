package uexchange

// NewClient - ..
func NewClient() *Client {
	return &Client{}
}

func (c *Client) getAPIURL(endpoint string) string {
	return apiHost + ":" + apiPort + "/" + endpoint
}
