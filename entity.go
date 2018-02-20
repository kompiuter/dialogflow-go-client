package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"
)

const (
	entityPath     = "entities"
	entryPath      = "entries"
	userEntityPath = "userEntities"
)

// GetAllEntities returns all of the agent's entities
func (client *Client) GetAllEntities() ([]model.Entity, error) {
	var response []model.Entity

	request := newRequest(
		client,
		requestOptions{
			Path:   entityPath,
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

// GetEntity returns the entity with ID id
func (client *Client) GetEntity(id string) (model.Entity, error) {
	var response model.Entity

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", entityPath, id),
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

// CreateEntity creates a new entity
func (client *Client) CreateEntity(entity model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entity, model.Entity{}) {
		return response, errors.New("entity cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   entityPath,
			Method: http.MethodPost,
			Body:   entity,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// AddEntityEntries adds entries to the entity with ID id
func (client *Client) AddEntityEntries(id string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, []model.Entry{}) || id == "" {
		return response, errors.New("entries and id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s/%s", entityPath, id, entryPath),
			Method: http.MethodPost,
			Body:   entries,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntities creates or updates an array of entities
func (client *Client) UpdateEntities(entities []model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entities, []model.Entity{}) {
		return response, errors.New("entities cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   entityPath,
			Method: http.MethodPut,
			Body:   entities,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntity updates the entity with ID id
func (client *Client) UpdateEntity(id string, entity model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entity, model.Entity{}) || id == "" {
		return response, errors.New("entity and id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", entityPath, id),
			Method: http.MethodPut,
			Body:   entity,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntityEntries updates entries of entity with ID id
func (client *Client) UpdateEntityEntries(id string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, model.Entry{}) || id == "" {
		return response, errors.New("entries and id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s/%s", entityPath, id, entryPath),
			Method: http.MethodPut,
			Body:   entries,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteEntity deletes the entity with ID id
func (client *Client) DeleteEntity(id string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s", entityPath, id),
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

// DeleteEntityEntries deletes entries of entity with ID id
func (client *Client) DeleteEntityEntries(id string, values []string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if len(values) == 0 || id == "" {
		return response, errors.New("values and id cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:   fmt.Sprintf("%s/%s/%s", entityPath, id, entryPath),
			Method: http.MethodDelete,
			Body:   values,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// CreateUserEntities creates one or more user entities for the specified session
func (client *Client) CreateUserEntities(session string, userEntities []model.UserEntity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(userEntities, []model.UserEntity{}) {
		return response, errors.New("user entities cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      userEntityPath,
			SessionID: session,
			Method:    http.MethodPost,
			Body: struct {
				SessionID string             `json:"sessionId"`
				Entities  []model.UserEntity `json:"entities"`
			}{
				SessionID: session,
				Entities:  userEntities,
			},
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateUserEntity updates the user entity with name name for the specified session
func (client *Client) UpdateUserEntity(session, name string, userEntity model.UserEntity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(userEntity, model.UserEntity{}) || name == "" {
		return response, errors.New("user entity and name cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      fmt.Sprintf("%s/%s", userEntityPath, name),
			SessionID: session,
			Method:    http.MethodPut,
			Body:      userEntity,
		},
	)

	data, err := request.perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GetUserEntity gets a user entity with name name for the specified session
func (client *Client) GetUserEntity(session, name string) (model.UserEntity, error) {
	var response model.UserEntity

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      fmt.Sprintf("%s/%s", userEntityPath, name),
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

// DeleteUserEntity deletes a user entity name name for the specified session
func (client *Client) DeleteUserEntity(session, name string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := newRequest(
		client,
		requestOptions{
			Path:      fmt.Sprintf("%s/%s", userEntityPath, name),
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
