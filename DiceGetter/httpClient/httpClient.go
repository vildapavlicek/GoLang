package httpClient

import (
	"io"
	"net/http"
	"time"
)

type customClient struct {
	client *http.Client
}

func New(timeout time.Duration) *customClient {
	return &customClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func getRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return request, err
}

func (c *customClient) getResponse(method, url string, body io.Reader) (*http.Response, error) {
	request, err := getRequest(method, url, body)
	defer request.Body.Close()
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
