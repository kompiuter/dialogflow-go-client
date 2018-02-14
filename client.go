package dialogflow

// Client is a DialogFlow client
type Client struct {
	accessToken string
	protocol    string
	apiBaseURL  string
	apiLang     string
	sessionID   string
}

// GetProtocol returns client protocol
func (client *Client) GetProtocol() string {
	return client.protocol
}

// GetAPILanguage returns client API language
func (client *Client) GetAPILanguage() string {
	return client.apiLang
}

// GetBaseURL returns client base URL
func (client *Client) GetBaseURL() string {
	return client.apiBaseURL
}

// GetAccessToken returns client access token
func (client *Client) GetAccessToken() string {
	return client.accessToken
}
