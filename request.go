package goomg

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

func NewRequestMirrors(ctx context.Context, hostURL string) (*http.Request, error) {
	hostURL, err := appendFilename(hostURL, "mirrors.txt")
	if err != nil {
		return nil, err
	}
	return newRequest(ctx, hostURL)
}

func NewRequestCanary(ctx context.Context, hostURL string) (*http.Request, error) {
	hostURL, err := appendFilename(hostURL, "canary.txt")
	if err != nil {
		return nil, err
	}
	return newRequest(ctx, hostURL)
}

func NewRequestRelated(ctx context.Context, hostURL string) (*http.Request, error) {
	hostURL, err := appendFilename(hostURL, "related.txt")
	if err != nil {
		return nil, err
	}
	return newRequest(ctx, hostURL)
}

func appendFilename(hostURL, filename string) (string, error) {
	u, err := url.Parse(hostURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, filename)
	return u.String(), nil
}

func newRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	setUserAgent(req)
	return req, nil
}

func setUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", "go-omg-client/1.0")
}
