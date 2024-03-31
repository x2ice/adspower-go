package adspower

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type QueryGroupOptions struct {
	Name   string
	Offset int
	Limit  int
}

func (a *AdsPower) QueryGroups(ctx context.Context, opts ...*QueryGroupOptions) (Groups, error) {
	url_ := fmt.Sprintf("%s/list", GroupApi)
	if opts != nil {
		opts_ := opts[0]
		query := url.Values{}
		if opts_.Name != "" {
			query.Set("group_name", opts_.Name)
		}

		if opts_.Offset > 1 {
			query.Set("page", fmt.Sprintf("%d", opts_.Offset))
		}

		if opts_.Limit > 1 {
			query.Set("page_size", fmt.Sprintf("%d", opts_.Limit))
		}

		url_ = url_ + "?" + query.Encode()
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url_, nil)
	defer req.Body.Close()

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[queryGroupResponse](resp)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return nil, err
	}

	list := decodedBody.Data.List
	len := len(list)
	if len == 0 {
		return nil, ErrNoGroupsFound
	}

	groups := make(Groups, len)

	copy(groups, list)
	return groups, nil
}

func (a *AdsPower) QueryGroupByName(ctx context.Context, name string) (*Group, error) {
	opts := &QueryGroupOptions{Name: name}
	groups, err := a.QueryGroups(ctx, opts)
	if err != nil {
		_, ok := err.(errNoGroupsFound)
		if ok {
			ErrGroupNotFound = fmt.Errorf("group %s not found", name)
			return nil, ErrGroupNotFound
		}

		return nil, err
	}

	return groups[0], nil
}
