package graphql

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pinpt/agent.next/sdk"
)

type client struct {
	url     string
	headers map[string]string
}

var _ sdk.GraphQLClient = (*client)(nil)

func (g *client) Query(query string, variables map[string]interface{}, out interface{}, options ...sdk.WithGraphQLOption) error {
	payload := struct {
		Variables map[string]interface{} `json:"variables"`
		Query     string                 `json:"query"`
	}{variables, query}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, g.url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range g.headers {
		req.Header.Set(k, v)
	}
	for _, opt := range options {
		if err := opt(req); err != nil {
			return err
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		var datares struct {
			Data   interface{} `json:"data"`
			Errors []struct {
				Message string `json:"message"`
			} `json:"errors"`
		}
		err = json.Unmarshal(body, &datares)
		if err != nil {
			return err
		}
		if datares.Errors != nil {
			b, err := json.Marshal(datares.Errors)
			if err != nil {
				return err
			}
			return errors.New(string(b))
		}
		b, err := json.Marshal(datares.Data)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, out)
	}
	return fmt.Errorf("err: %s. status code: %s", string(body), resp.Status)
}

type manager struct {
}

var _ sdk.GraphQLClientManager = (*manager)(nil)

// New is for creating a new graphql client instance that can be reused
func (m *manager) New(url string, headers map[string]string) sdk.GraphQLClient {
	return &client{url, headers}
}

// New returns a new GraphQLClientManager
func New() sdk.GraphQLClientManager {
	return &manager{}
}