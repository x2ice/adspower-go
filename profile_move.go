package adspower

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
)

type moveProfileRequest struct {
	UserIds []string `json:"user_ids"`
	GroupId string   `json:"group_id"`
}

func (a *AdsPower) MoveProfiles(ctx context.Context, ids []string, groupId string) error {

	payload := &moveProfileRequest{
		UserIds: ids,
		GroupId: groupId,
	}

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	buf := bytes.NewBuffer(b)

	url := UserApi + "/regroup"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[response](resp)
	if err != nil {
		return err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return err
	}

	return nil
}

func (a *AdsPower) MoveProfile(ctx context.Context, id, targetGroupId string) error {
	return a.MoveProfiles(ctx, []string{id}, targetGroupId)
}
