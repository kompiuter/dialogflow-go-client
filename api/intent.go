package api

import (
	"encoding/json"
	"errors"
	"reflect"

	df "github.com/kompiuter/go-dialogflow"
	"github.com/kompiuter/go-dialogflow/model"
)

// AllIntents returns all of the agent's intents
func (client *df.Client) AllIntents() ([]model.IntentAgent, error) {
	var response []model.IntentAgent

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
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

// GetIntent returns the intent with ID id
func (client *df.Client) GetIntent(id string) (model.Intent, error) {
	var response model.Intent

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
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

// CreateIntent creates a new intent
func (client *df.Client) CreateIntent(intent model.Intent) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(intent, model.Intent{}) {
		return response, errors.New("intent cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "intents?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   intent,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Updates the specified intent
func (client *df.Client) UpdateIntent(id string, intent model.Intent) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(intent, model.Intent{}) || id == "" {
		return response, errors.New("intent and id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   intent,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateIntent updates the intent with ID is
func (client *df.Client) IntentsDeleteRequest(id string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "intents/" + id + "?v=" + client.GetApiVersion(),
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
