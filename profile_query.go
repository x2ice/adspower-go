package adspower

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type QueryProfileOptions struct {
	ID           string
	GroupId      string
	SerialNumber string
	Offset       int
	Limit        int
}

func (a *AdsPower) QueryProfiles(ctx context.Context, opts ...*QueryProfileOptions) (Profiles, error) {
	url_ := fmt.Sprintf("%s/list", UserApi)
	if len(opts) != 0 {
		opts_ := opts[0]
		if opts_ != nil {
			query := url.Values{}

			if opts_.ID != "" {
				query.Set("user_id", opts_.ID)
			}

			if opts_.GroupId != "" {
				query.Set("group_id", opts_.GroupId)
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

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
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

func (a *AdsPower) QueryProfilesByGroupName(ctx context.Context, groupName string, offset, limit int) (Profiles, error) {
	groups, err := a.QueryGroups(ctx, &QueryGroupOptions{Name: groupName})
	if err != nil {
		return nil, err
	}

	group := groups[0]
	opts := &QueryProfileOptions{
		GroupId: group.ID,
		Offset:  offset,
		Limit:   limit,
	}

	return a.QueryProfiles(ctx, opts)
}
