package uexchange

// NewClient - ..
func NewClient(c Credentials) Client {
	return Client{
		APICredentials: c,
	}
}
