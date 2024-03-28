package adspower

import (
	"context"
	"fmt"
	"net/http"
)

func (c *AdsPower) QueryAllOpenedBrowsers(ctx context.Context) (Browsers, error) {
	url := fmt.Sprintf("%s/local-active", RootUrl)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	c.rl.Take()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
