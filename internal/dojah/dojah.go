package dojah

import (
	"errors"
	"net/http"
)

type DojahClient struct {
	secretKey    string
	appId        string
	httpclient   *http.Client
	baseURL      string
	isProduction bool
}

type DojahOption struct {
	// Secret key used for authentication
	SecretKey string
	// App ID gotten on the dashboard
	AppID string
	// Custom implementation of HTTP client to use
	HTTPClient *http.Client
}

func NewDojahClient(opt DojahOption) (*DojahClient, error) {
	if len(opt.SecretKey) == 0 {
		return nil, errors.New("dojah secret key is missing")
	}

	if len(opt.AppID) == 0 {
		return nil, errors.New("dojah app id is missing")
	}

	client := &DojahClient{
		secretKey:    opt.SecretKey,
		appId:        opt.AppID,
		baseURL:      baseURLSandbox,
		httpclient:   http.DefaultClient,
		isProduction: false,
	}
	if opt.HTTPClient != nil {
		client.httpclient = opt.HTTPClient
	}
	return client, nil
}
