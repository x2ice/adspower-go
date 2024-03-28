package adspower

import (
	"context"
	"fmt"
	"net/http"
)

func (c *AdsPower) CloseBrowser(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/stop?user_id=%s", RootUrl, id)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	c.rl.Take()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[closeBrowserResponse](resp)
	if err != nil {
		return err
	}

	return handleResponseError(decodedBody)
}
