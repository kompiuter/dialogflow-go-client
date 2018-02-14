package api

import (
	"encoding/json"
	"errors"
	"reflect"

	df "github.com/kompiuter/go-dialogflow"
	"github.com/kompiuter/go-dialogflow/model"
)

// AllEntities returns all of the agent's entities
func (client *df.Client) AllEntities() ([]model.Entity, error) {
	var response []model.Entity

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
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

// GetEntity returns the entity with ID id
func (client *df.Client) GetEntity(id string) (model.Entity, error) {
	var response model.Entity

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "?v=" + client.GetApiVersion(),
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

// CreateEntity creates a new entity
func (client *df.Client) CreateEntity(entity model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entity, model.Entity{}) {
		return response, errors.New("entity cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   entity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// AddEntityEntries adds entries to the entity with ID id
func (client *df.Client) AddEntityEntries(id string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, []model.Entry{}) || id == "" {
		return response, errors.New("entries and id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "/entries?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   entries,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntities creates or updates an array of entities
func (client *df.Client) UpdateEntities(entities []model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entities, []model.Entity{}) {
		return response, errors.New("entities cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entities,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntity updates the entity with ID id
func (client *df.Client) UpdateEntity(id string, entity model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entity, model.Entity{}) || id == "" {
		return response, errors.New("entity and id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateEntityEntries updates entries of entity with ID id
func (client *df.Client) UpdateEntityEntries(id string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, model.Entry{}) || id == "" {
		return response, errors.New("entries and id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "/entries?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   entries,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// DeleteEntity deletes the entity with ID id
func (client *df.Client) DeleteEntity(id string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if id == "" {
		return response, errors.New("id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "?v=" + client.GetApiVersion(),
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

// DeleteEntityEntries deletes entries of entity with ID id
func (client *df.Client) DeleteEntityEntries(id string, values []string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if len(values) == 0 || id == "" {
		return response, errors.New("values and id cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + id + "/entries?v=" + client.GetApiVersion(),
			Method: "DELETE",
			Body:   values,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// CreateUserEntities creates one or more user entities for a single session
func (client *df.Client) CreateUserEntities(userEntities []model.UserEntity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(userEntities, []model.UserEntity{}) {
		return response, errors.New("user entities cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities?v=" + client.GetApiVersion(),
			Method: "POST",
			Body: struct {
				SessionID string
				Entities  []model.UserEntity
			}{
				SessionID: client.GetSessionID(),
				Entities:  userEntities,
			},
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// UpdateUserEntity updates the user entity with name s
func (client *df.Client) UpdateUserEntity(s string, userEntity model.UserEntity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(userEntity, model.UserEntity{}) || s == "" {
		return response, errors.New("user entity and name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + s + "?v=" + client.GetApiVersion(),
			Method: "PUT",
			Body:   userEntity,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// GetUserEntity gets a user entity with name s
func (client *df.Client) GetUserEntity(s string) (model.UserEntity, error) {
	var response model.UserEntity

	if s == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + s + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
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

// DeleteUserEntity deletes a user entity with name s
func (client *df.Client) DeleteUserEntity(s string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if s == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + s + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
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
