package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"
)

const contextEndpoint = "contexts"

// GetAllContexts returns all of the contexts for the specified session
func (client *Client) GetAllContexts(session string) ([]model.Context, error) {
	var response []model.Context

	request := newRequest(
		client,
		requestOptions{
			Path:      contextEndpoint,
			SessionID: session,
			Method:    http.MethodGet,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GetContext returns the context with name ctx for the specified session
func (client *Client) GetContext(session, ctx string) (model.Context, error) {
	var response model.Context

	if ctx == "" {
		return response, errors.New("ctx cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      fmt.Sprintf("%s/%s", contextEndpoint, ctx),
			SessionID: session,
			Method:    http.MethodGet,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// AddContexts adds new active contexts to the specified session
func (client *Client) AddContexts(session string, contexts []model.Context) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(contexts, []model.Context{}) {
		return response, errors.New("contexts cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      contextEndpoint,
			SessionID: session,
			Method:    http.MethodPost,
			Body:      contexts,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteAllContexts deletes all contexts from the specified session
func (client *Client) DeleteAllContexts(session string) (model.QueryResponse, error) {
	var response model.QueryResponse
	request := newRequest(
		client,
		requestOptions{
			Path:      contextEndpoint,
			SessionID: session,
			Method:    http.MethodDelete,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteContext deletes the context with name ctx from the specified session
func (client *Client) DeleteContext(session, ctx string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if ctx == "" {
		return response, errors.New("ctx cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      fmt.Sprintf("%s/%s", contextEndpoint, ctx),
			SessionID: session,
			Method:    http.MethodDelete,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
