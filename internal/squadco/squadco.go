package squadco

import (
	"errors"
	"net/http"
)

type Client struct {
	apiKey     string
	httpclient *http.Client
	baseURL    string
}

type SquadOption struct {
	// API key used for authentication
	ApiKey string
	// Custom implementation of HTTP client to use
	HTTPClient *http.Client
}

func NewSquadClient(opt SquadOption) (*Client, error) {
	if len(opt.ApiKey) == 0 {
		return nil, errors.New("squadco api key is missing")
	}

	client := &Client{
		apiKey:     opt.ApiKey,
		baseURL:    baseURLSandbox,
		httpclient: http.DefaultClient,
	}
	if opt.HTTPClient != nil {
		client.httpclient = opt.HTTPClient
	}
	return client, nil
}
