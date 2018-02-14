package api

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"
	df "github.com/kompiuter/go-dialogflow/model"
)

// Query queries DialogFlow with a GET request with query encoded as query parameters
func (client *df.Client) Query(query model.Query) (model.QueryResponse, error) {
	return queryClient(client, query, false)
}

// QueryBody queries DialogFlow with a POST request with query in the body of the request
func (client *df.Client) QueryBody(query model.Query) (model.QueryResponse, error) {
	return queryClient(client, query, true)
}

func queryClient(client *df.Client, query model.Query, body bool) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(query, model.Query{}) {
		return response, errors.New("query cannot be empty")
	}

	if query.V == "" {
		query.V = client.GetApiVersion()
	}

	if query.Lang == "" {
		query.Lang = client.GetApiLang()
	}

	if query.SessionID == "" {
		query.SessionID = client.GetSessionID()
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "query",
			Method: "GET",
		},
	)

	if body {
		request.RequestOptions.Body = query
	} else {
		request.RequestOptions.QueryParams = query.ToMap()
	}

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
