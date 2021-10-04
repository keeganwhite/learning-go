package client

import "net/http"

type HTTPClient struct {
	client *http.Client
	backendURL string
}

func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		backendURL: uri,
		client: &http.Client{},
	}

}
