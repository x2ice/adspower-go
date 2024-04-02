package adspower

import (
	"context"
	"fmt"
	"net/http"
)

func (a *AdsPower) QueryAllOpenedBrowsers(ctx context.Context) (Browsers, error) {
	url := fmt.Sprintf("%s/local-active", RootUrl)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	a.rl.Take()

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	decodedBody, err := decodeResponseBody[openedBrowsersResponse](resp)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return nil, err
	}

	list := decodedBody.Data.List
	len := len(list)
	openedBrowsers := make(Browsers, len)

	copy(openedBrowsers, list)
	return openedBrowsers, nil
}
