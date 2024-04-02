package adspower

import (
	"net/http"
	"net/url"

	json "github.com/goccy/go-json"
)

func decodeResponseBody[V any](resp *http.Response) (*V, error) {
	var v V
	err := json.NewDecoder(resp.Body).Decode(&v)
	return &v, err
}

func DefaultProxyConfig() *ProxyConfig {
	return &ProxyConfig{Soft: "no_proxy"}
}

func NewProxyConfigFromUrl(proxySoft, proxyUrl string) (*ProxyConfig, error) {
	url, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, errInvalidProxyFormat
	}

	password, _ := url.User.Password()
	return &ProxyConfig{
		Soft:        proxySoft,
		Type:        url.Scheme,
		Host:        url.Hostname(),
		Port:        url.Port(),
		User:        url.User.Username(),
		Password:    password,
		ChangeIpUrl: url.Path,
	}, nil
}
