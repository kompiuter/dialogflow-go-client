package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"
)

const intentPath = "intents"

// GetAllIntents returns all of the agent's intents
func (client *Client) GetAllIntents() ([]model.IntentAgent, error) {
	var response []model.IntentAgent

	request := newRequest(
		client,
		requestOptions{
			Path:   intentPath,
			Method: http.MethodGet,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GetIntent returns the intent with ID id
func (client *Client) GetIntent(id string) (model.Intent, error) {
	var response model.Intent

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", intentPath, id),
			Method: http.MethodGet,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	log.Println(string(data))
	err = json.Unmarshal(data, &response)
	return response, err
}

// CreateIntent creates a new intent
func (client *Client) CreateIntent(intent model.Intent) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(intent, model.Intent{}) {
		return response, errors.New("intent cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   intentPath,
			Method: http.MethodPost,
			Body:   intent,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateIntent updates the intent with ID id
func (client *Client) UpdateIntent(id string, intent model.Intent) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(intent, model.Intent{}) || id == "" {
		return response, errors.New("intent and id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", intentPath, id),
			Method: http.MethodPut,
			Body:   intent,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteIntent deletes the intent with ID id
func (client *Client) DeleteIntent(id string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", intentPath, id),
			Method: http.MethodDelete,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}
