package dialogflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/kompiuter/go-dialogflow/model"

	uuid "github.com/satori/go.uuid"
)

const (
	apiBaseURL = "https://api.dialogflow.com/v1/"
	apiVersion = "20150910" //https://dialogflow.com/docs/reference/agent/#protocol_version
)

type DialogFlowClient struct {
	accessToken string
	apiVersion  string
	apiBaseURL  string
	apiLang     string
	sessionID   string
}

type Options struct {
	AccessToken string
	APIVersion  string
	SessionID   string
}

// Create API.AI instance
func NewDialogFlowClient(options Options) (*DialogFlowClient, error) {
	if (reflect.DeepEqual(options, Options{}) || options.AccessToken == "") {
		return nil, errors.New("access token is required for new dialogflow client")
	}

	client := &DialogFlowClient{
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

// Takes natural language text and information as query parameters and returns information as JSON
func (client *DialogFlowClient) QueryFindRequest(query model.Query) (model.QueryResponse, error) {
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
			URI:         client.GetBaseUrl() + "query",
			Method:      "GET",
			Body:        nil,
			QueryParams: query.ToMap(),
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Takes natural language text and information as JSON in the POST body and returns information as JSON
func (client *DialogFlowClient) QueryCreateRequest(query model.Query) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(query, model.Query{}) {
		return response, errors.New("query cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "query?v=" + client.GetApiVersion(),
			Method: "POST",
			Body:   query,
		},
	)

	data, err := request.Perform()
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	return response, err
}

// Retrieves a list of all entities for the agent
func (client *DialogFlowClient) EntitiesFindAllRequest() ([]model.Entity, error) {
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

// Retrieves the specified entity
func (client *DialogFlowClient) EntitiesFindByIdRequest(eid string) (model.Entity, error) {
	var response model.Entity

	if eid == "" {
		return response, errors.New("eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
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

// Creates a new entity
func (client *DialogFlowClient) EntitiesCreateRequest(entity model.Entity) (model.QueryResponse, error) {
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

// Adds entries to the specified entity.
func (client *DialogFlowClient) EntitiesAddEntryRequest(eid string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, []model.Entry{}) || eid == "" {
		return response, errors.New("entries and eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
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

// Creates or updates an array of entities
func (client *DialogFlowClient) EntitiesUpdateRequest(entities []model.Entity) (model.QueryResponse, error) {
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

// Updates the specified entity
func (client *DialogFlowClient) EntitiesUpdateEntityRequest(eid string, entity model.Entity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entity, model.Entity{}) || eid == "" {
		return response, errors.New("entity and eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
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

// Updates entity entries
func (client *DialogFlowClient) EntitiesUpdateEntityEntriesRequest(eid string, entries []model.Entry) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(entries, model.Entry{}) || eid == "" {
		return response, errors.New("entries and eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
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

// Deletes the specified entity
func (client *DialogFlowClient) EntitiesDeleteRequest(eid string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if eid == "" {
		return response, errors.New("eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "?v=" + client.GetApiVersion(),
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

// Deletes entity entries
func (client *DialogFlowClient) EntitiesDeleteEntriesRequest(eid string, values []string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if len(values) == 0 || eid == "" {
		return response, errors.New("values and eid cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "entities/" + eid + "/entries?v=" + client.GetApiVersion(),
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

// Adds one or multiple user entities for a session.
func (client *DialogFlowClient) UserEntitiesCreateRequest(userEntities []model.UserEntity) (model.QueryResponse, error) {
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

// Updates user entity specified by name
func (client *DialogFlowClient) UserEntitiesUpdateRequest(name string, userEntity model.UserEntity) (model.QueryResponse, error) {
	var response model.QueryResponse

	if reflect.DeepEqual(userEntity, model.UserEntity{}) || name == "" {
		return response, errors.New("user entity and name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion(),
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

// Gets a user entity object by name
func (client *DialogFlowClient) UserEntitiesFindByNameRequest(name string) (model.UserEntity, error) {
	var response model.UserEntity

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
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

// Deletes a user entity object with a specified name
func (client *DialogFlowClient) UserEntitiesDeleteByNameRequest(name string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "userEntities/" + name + "?v=" + client.GetApiVersion() + "&sessionId=" + client.GetSessionID(),
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

// Retrieves a list of all intents for the agent
func (client *DialogFlowClient) IntentsFindAllRequest() ([]model.IntentAgent, error) {
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

// Retrieves the specified intent
func (client *DialogFlowClient) IntentsFindByIdRequest(id string) (model.Intent, error) {
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

// Creates a new intent
func (client *DialogFlowClient) IntentsCreateRequest(intent model.Intent) (model.QueryResponse, error) {
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
func (client *DialogFlowClient) IntentsUpdateRequest(id string, intent model.Intent) (model.QueryResponse, error) {
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

// Deletes the specified intent
func (client *DialogFlowClient) IntentsDeleteRequest(id string) (model.QueryResponse, error) {
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

// retrieves the list of all currently active contexts for the specified session
func (client *DialogFlowClient) ContextsFindAllRequest() ([]model.Context, error) {
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

// Retrieves the specified context for the specified session
func (client *DialogFlowClient) ContextsFindByNameRequest(name string) (model.Context, error) {
	var response model.Context

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
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

// Adds new active contexts to the specified session
func (client *DialogFlowClient) ContextsCreateRequest(contexts []model.Context) (model.QueryResponse, error) {
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

// Deletes all contexts from the specified session
func (client *DialogFlowClient) ContextsDeleteRequest() (model.QueryResponse, error) {
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

// Deletes the specified context from the specified session
func (client *DialogFlowClient) ContextsDeleteByNameRequest(name string) (model.QueryResponse, error) {
	var response model.QueryResponse

	if name == "" {
		return response, errors.New("name cannot be empty")
	}

	request := NewRequest(
		client,
		model.RequestOptions{
			URI:    client.GetBaseUrl() + "contexts/" + name + "?sessionId=" + client.GetSessionID(),
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

// GET API.AI access token
func (client *DialogFlowClient) GetAccessToken() string {
	return client.accessToken
}

// GET API.AI version
func (client *DialogFlowClient) GetApiVersion() string {
	if client.apiVersion != "" {
		return client.apiVersion
	}
	return apiVersion
}

// GET API.AI language
func (client *DialogFlowClient) GetApiLang() string {
	return client.apiLang
}

// Get API.AI base url
func (client *DialogFlowClient) GetBaseUrl() string {
	return client.apiBaseURL
}

// Get current session ID
func (client *DialogFlowClient) GetSessionID() string {
	return client.sessionID
}

// Set a new seesion ID
func (client *DialogFlowClient) SetSessionID(sessionID string) {
	client.sessionID = sessionID
}
