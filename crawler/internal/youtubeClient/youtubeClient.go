package youtubeclient

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	client *http.Client
	addr   string
}

func New(timeout, idleConnections, idleTimeout int) *Client {

	timeoutDur := time.Duration(timeout)
	idleTimeoutDur := time.Duration(idleTimeout)
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		//log report error
	}

	transport := &http.Transport{
		MaxIdleConns:    idleConnections,
		IdleConnTimeout: idleTimeoutDur,
	}

	myClient := &http.Client{
		Timeout:   timeoutDur,
		Jar:       jar,
		Transport: transport,
	}

	return &Client{
		client: myClient,
		addr:   "https://youtube.com",
	}
}

// GetHTTPRequest returns *Request to do Do method with
func (c *Client) getHTTPRequest(suffix string) (*http.Request, error) {
	uri := c.addr + suffix
	req, err := http.NewRequest("GET", uri, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	return req, nil
}

func (c *Client) GetResponse(suffix string) (res *http.Response, err error) {
	req, err := c.getHTTPRequest(suffix)
	if err != nil {
		return nil, err
	}

	youTube, err := url.Parse(c.addr)
	if err != nil {
		//just log message
	}

	for _, v := range c.client.Jar.Cookies(youTube) {
		req.AddCookie(v)
	}

	res, err = c.client.Do(req)
	if err != nil {
		//log message
		return nil, err
	}

	return res, nil
}
