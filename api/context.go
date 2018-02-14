package api

import (
	"encoding/json"
	"errors"
	"reflect"

	df "github.com/kompiuter/go-dialogflow"
	"github.com/kompiuter/go-dialogflow/model"
)

// AllContexts returns all of the contexts for the specified session
func (client *df.Client) AllContexts() ([]model.Context, error) {
	var response []model.Context

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GetContext returns the context with name s
func (client *df.Client) GetContext(s string) (model.Context, error) {
	var response model.Context

	if s == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + s + "?sessionId=" + client.GetSessionID(),
			Method: "GET",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// AddContexts adds new active contexts to the specified session
func (client *df.Client) AddContexts(contexts []model.Context) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(contexts, []model.Context{}) {
		return response, errors.New("contexts cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "POST",
			Body:   contexts,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteAllContexts deletes all contexts from the specified session
func (client *df.Client) DeleteAllContexts() (model.QueryResponse, error) {
	var response model.QueryResponse

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts?sessionId=" + client.GetSessionID(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteContext deletes the context with name s from the specified session
func (client *df.Client) DeleteContext(s string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if s == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + s + "?sessionId=" + client.GetSessionID(),
			Method: "DELETE",
			Body:   nil,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
