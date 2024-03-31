package adspower

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type OpenBrowserOptions struct {
	OpenTabs               bool
	IpTab                  bool
	NewIpTab               bool
	LaunchArgs             []string
	Headless               bool
	DisablePasswordFilling bool
	ClearCacheAfterClosing bool
	EnablePasswordSaving   bool
}

func (a *AdsPower) OpenBrowser(ctx context.Context, id string, opts ...*OpenBrowserOptions) (*Browser, error) {
	url_ := fmt.Sprintf("%s/start?user_id=%s", RootUrl, id)

	if len(opts) != 0 {
		opts_ := opts[0]
		if opts_ != nil {
			query := url.Values{}

			if opts_.OpenTabs {
				query.Set("open_tabs", "1")
			}

			if !opts_.IpTab {
				query.Set("ip_tab", "0")
			}

			if opts_.NewIpTab {
				query.Set("new_first_tab", "1")
			}

			if opts_.Headless {
				query.Set("headless", "1")
			}

			if opts_.DisablePasswordFilling {
				query.Set("disable_password_filling", "1")
			}

			if opts_.ClearCacheAfterClosing {
				query.Set("clear_cache_after_closing", "1")
			}

			if opts_.EnablePasswordSaving {
				query.Set("enable_password_saving", "1")
			}

			url_ += "&" + query.Encode()
		}
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url_, nil)
	defer req.Body.Close()

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[openBrowserResponse](resp)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return nil, err
	}

	openBrowser := decodedBody.Data
	return openBrowser, nil
}
