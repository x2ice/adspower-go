package adspower

import (
	"net/http"

	json "github.com/goccy/go-json"
)

func decodeResponseBody[V any](resp *http.Response) (*V, error) {
	var v V
	err := json.NewDecoder(resp.Body).Decode(&v)
	return &v, err
}
