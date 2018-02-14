package dialogflow

// NewClient creates a new DialogFlow client
// You must provide a valid agent access token
func NewClient(token string) *Client {
	client := &Client{
		accessToken: token,
		apiBaseURL:  "https://api.dialogflow.com/v1/",
		apiLang:     "en",
		protocol:    "20150910",
	}

	return client
}

// SetProtocol sets the client's protocol
// There are two protocols available:
// 20150910 -> sys.number values are returned as strings (default)
// 20170712 -> sys.number values are returned as integers
func (c *Client) SetProtocol(s string) {
	c.protocol = s
}
