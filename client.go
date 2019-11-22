package goomg

import (
	"fmt"
	"github.com/onionltd/go-omg/http/request"
	"github.com/onionltd/go-omg/spec"
	"io/ioutil"
	"net/http"
)

type Client struct {
	cli *http.Client
}

func (c Client) do(req *http.Request) ([]byte, error) {
	res, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid HTTP Response code '%d'", res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)
}

// GetMirrorsMessage makes an HTTP request to a host and downloads contents of mirrors.txt.
func (c Client) GetMirrorsMessage(host string) (spec.Mirrors, error) {
	req, err := request.Host(host).NewRequestMirrors()
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// GetCanaryMessage makes an HTTP request to a host and downloads contents of canary.txt.
func (c Client) GetCanaryMessage(host string) (spec.Canary, error) {
	req, err := request.Host(host).NewRequestCanary()
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// GetRelatedMessage makes an HTTP request to a host and downloads contents of related.txt.
func (c Client) GetRelatedMessage(host string) (spec.Mirrors, error) {
	req, err := request.Host(host).NewRequestRelated()
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// NewClient returns an instance of Client, which wraps passed httpClient.
func NewClient(httpClient *http.Client) *Client {
	return &Client{cli: httpClient}
}
