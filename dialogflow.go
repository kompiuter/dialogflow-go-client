package dialogflow

import (
	"errors"
	"fmt"
	"reflect"

	uuid "github.com/satori/go.uuid"
)

const (
	apiBaseURL = "https://api.dialogflow.com/v1/"
	apiVersion = "20150910" // https://dialogflow.com/docs/reference/agent/#protocol_version
)

// Client is a DialogFlow client
type Client struct {
	accessToken string
	apiVersion  string
	apiBaseURL  string
	apiLang     string
	sessionID   string
}

// Options can be used to modify the DialogFlow client upon creation
type Options struct {
	AccessToken string
	APIVersion  string
	SessionID   string
}

// NewClient creates a new DialogFlow client
func NewClient(options Options) (*Client, error) {
	if (reflect.DeepEqual(options, Options{}) || options.AccessToken == "") {
		return nil, errors.New("access token is required for new dialogflow client")
	}

	client := &Client{
		accessToken: options.AccessToken,
		apiBaseURL:  apiBaseURL,
		apiLang:     "en",
		apiVersion:  apiVersion,
	}

	client.sessionID = options.SessionID
	if client.sessionID == "" {
		u, err := uuid.NewV4()
		if err != nil {
			return nil, fmt.Errorf("could not generate a session id: %v", err)
		}

		client.sessionID = u.String()
	}

	return client, nil
}
