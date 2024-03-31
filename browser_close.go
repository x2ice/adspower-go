package adspower

import (
	"context"
	"fmt"
	"net/http"
)

func (a *AdsPower) CloseBrowser(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/stop?user_id=%s", RootUrl, id)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	defer req.Body.Close()

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	decodedBody, err := decodeResponseBody[closeBrowserResponse](resp)
	if err != nil {
		return err
	}

	return handleResponseError(decodedBody)
}
