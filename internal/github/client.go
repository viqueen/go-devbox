package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type client struct {
	baseURL    string
	HTTPClient *http.Client
}

func NewClient() Api {
	return &client{
		baseURL:    "https://api.github.com",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c client) SearchRepositories(query string) ([]Repository, error) {
	url := fmt.Sprintf("%s/search/repositories?q=language:go+%s", c.baseURL, query)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var searchResponse SearchRepositoriesResponse
	if err = json.NewDecoder(response.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	return searchResponse.Items, nil
}
