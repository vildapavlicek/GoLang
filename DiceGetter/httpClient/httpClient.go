package httpClient

import (
	"fmt"
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

// GetResponse gets response
func (c *customClient) GetResponse(method, baseURL, paramsURL string, body io.Reader) (*http.Response, error) {
	uri := baseURL + paramsURL
	request, err := getRequest(method, uri, body)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Response %v\n", response.Status)
	return response, nil
}
