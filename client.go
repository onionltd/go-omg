package goomg

import (
	"context"
	"fmt"
	"github.com/onionltd/go-omg/spec"
	"io/ioutil"
	"net/http"
)

type Client struct {
	cli       *http.Client
	userAgent string
}

const DefaultUserAgent = "go-omg/1.0"

func (c Client) setUserAgent(req *http.Request) {
	ua := c.userAgent
	if ua == "" {
		ua = DefaultUserAgent
	}
	req.Header.Set("User-Agent", ua)
}

func (c Client) do(req *http.Request) ([]byte, error) {
	c.setUserAgent(req)
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
func (c Client) GetMirrorsMessage(ctx context.Context, hostURL string) (spec.Mirrors, error) {
	req, err := NewRequestMirrors(ctx, hostURL)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// GetCanaryMessage makes an HTTP request to a host and downloads contents of canary.txt.
func (c Client) GetCanaryMessage(ctx context.Context, hostURL string) (spec.Canary, error) {
	req, err := NewRequestCanary(ctx, hostURL)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

// GetRelatedMessage makes an HTTP request to a host and downloads contents of related.txt.
func (c Client) GetRelatedMessage(ctx context.Context, hostURL string) (spec.Mirrors, error) {
	req, err := NewRequestRelated(ctx, hostURL)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) SetUserAgent(ua string) {
	c.userAgent = ua
}

// NewClient returns an instance of Client, which wraps passed httpClient.
func NewClient(httpClient *http.Client) *Client {
	return &Client{cli: httpClient}
}
