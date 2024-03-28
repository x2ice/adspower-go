package adspower

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type QueryProfileOptions struct {
	ID           string
	GroupID      string
	SerialNumber string
	Offset       int
	Limit        int
}

func (c *AdsPower) QueryProfiles(ctx context.Context, opts ...*QueryProfileOptions) (Profiles, error) {
	url_ := fmt.Sprintf("%s/list", UserApi)
	if len(opts) != 0 {
		opts_ := opts[0]
		if opts_ != nil {
			query := url.Values{}

			if opts_.ID != "" {
				query.Set("user_id", opts_.ID)
			}

			if opts_.GroupID != "" {
				query.Set("group_id", opts_.GroupID)
			}

			if opts_.SerialNumber != "" {
				query.Set("serial_number", opts_.SerialNumber)
			}

			if opts_.Offset > 1 {
				query.Set("page", fmt.Sprintf("%d", opts_.Offset))
			}

			if opts_.Limit > 1 {
				query.Set("page_size", fmt.Sprintf("%d", opts_.Limit))
			}

			url_ = url_ + "?" + query.Encode()
		}
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url_, nil)
	c.rl.Take()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[queryProfilesResponse](resp)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return nil, err
	}

	profiles := decodedBody.Data.List
	return profiles, nil
}

func (c *AdsPower) QueryProfilesByGroupName(ctx context.Context, groupName string, offset, limit int) (Profiles, error) {
	_, err := c.QueryGroups(ctx, &QueryGroupOptions{Name: groupName})
	if err != nil {
		return nil, err
	}

	opts := &QueryProfileOptions{
		GroupID: "",
		Offset:  offset,
		Limit:   limit,
	}

	return c.QueryProfiles(ctx, opts)
}
