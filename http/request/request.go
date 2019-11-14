package request

import (
	"net/http"
	"net/url"
	"path"
)

type Host string

func (h Host) NewRequestMirrors() (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "mirrors.txt")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (h Host) NewRequestCanary() (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "canary.txt")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (h Host) NewRequestRelated() (*http.Request, error) {
	u, err := url.Parse(string(h))
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "related.txt")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
