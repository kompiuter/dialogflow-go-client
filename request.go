package dialogflow

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type requestOptions struct {
	Path        string
	SessionID   string
	Method      string
	Body        interface{}
	QueryParams map[string]string
}

type request struct {
	URI         string
	Method      string
	Headers     map[string]string
	Body        interface{}
	QueryParams map[string]string
}

func newRequest(client *Client, options requestOptions) *request {
	headers := map[string]string{
		"Authorization": "Bearer " + client.GetAccessToken(),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	req := &request{
		URI:         prepare(client, options.Path, options.SessionID),
		Method:      options.Method,
		Headers:     headers,
		QueryParams: options.QueryParams,
		Body:        options.Body,
	}

	return req
}

// Perform executes the HTTP request
func (r *request) perform() ([]byte, error) {
	var data []byte
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URI, nil)

	if r.Method != http.MethodGet {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(r.Body)
		req, err = http.NewRequest(r.Method, r.URI, b)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range r.QueryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	if err != nil {
		return data, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func prepare(client *Client, path, session string) string {
	m := url.Values{}
	if session != "" {
		m.Add("sessionId", session)
	}
	m.Add("v", client.GetProtocol())

	return client.GetBaseURL() + path + "?" + m.Encode()
}
