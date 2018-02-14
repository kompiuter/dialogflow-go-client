package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"
)

const queryPath = "query"

// Query queries DialogFlow with a GET request with query encoded as query parameters
func (client *Client) Query(session string, query model.Query) (*model.QueryResponse, error) {
	return queryClient(client, session, query, false)
}

// QueryBody queries DialogFlow with a POST request with query in the body of the request
func (client *Client) QueryBody(session string, query model.Query) (*model.QueryResponse, error) {
	return queryClient(client, session, query, true)
}

func queryClient(client *Client, session string, query model.Query, body bool) (*model.QueryResponse, error) {
	if session == "" {
		return nil, errors.New("session cannot be empty")
	}
	query.SessionID = session

	if reflect.DeepEqual(query, model.Query{}) {
		return nil, errors.New("query cannot be empty")
	}

	if query.V == "" {
		query.V = client.GetProtocol()
	}

	if query.Lang == "" {
		query.Lang = client.GetAPILanguage()
	}

	opts := requestOptions{Path: queryPath}
	if body {
		opts.Method = http.MethodPost
		opts.Body = query
	} else {
		opts.Method = http.MethodGet
		opts.QueryParams = query.ToMap()
	}

	request := newRequest(client, opts)
	data, err := request.perform()
	if err != nil {
		return nil, err
	}

	var response model.QueryResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}
