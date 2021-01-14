package request

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type Host string

func (h Host) NewRequestMirrors(ctx context.Context) (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "mirrors.txt")
	req, err := newRequest(ctx, u)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (h Host) NewRequestCanary(ctx context.Context) (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "canary.txt")
	req, err := newRequest(ctx, u)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (h Host) NewRequestRelated(ctx context.Context) (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "related.txt")
	req, err := newRequest(ctx, u)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func newRequest(ctx context.Context, host *url.URL) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, host.String(), nil)
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
