package api

import (
	df "github.com/kompiuter/go-dialogflow"
)

// GetAPIVersion returns client API version
func (client *df.Client) GetAPIVersion() string {
	if client.apiVersion != "" {
		return client.apiVersion
	}
	return apiVersion
}

// GetAPILanguage returns client API language
func (client *df.Client) GetAPILanguage() string {
	return client.apiLang
}

// GetBaseURL returns client base URL
func (client *df.Client) GetBaseURL() string {
	return client.apiBaseURL
}

// GetAccessToken returns client access token
func (client *df.Client) GetAccessToken() string {
	return client.accessToken
}
