package adspower

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
)

type CreateProfileOptions struct {
	Name              string   `json:"name,omitempty"`
	DomainName        string   `json:"domain_name,omitempty"`
	OpenURLs          []string `json:"open_urls,omitempty"`
	RepeatConfig      []int    `json:"repeat_config,omitempty"`
	Username          string   `json:"username,omitempty"`
	Password          string   `json:"password,omitempty"`
	Fakey             string   `json:"fakey,omitempty"`
	Cookie            Cookies  `json:"cookie,omitempty"`
	IgnoreCookieError int      `json:"ignore_cookie_error,omitempty"`
	GroupId           string   `json:"group_Id,omitempty"`
	IP                string   `json:"ip,omitempty"`
	Country           string   `json:"country,omitempty"`
	Region            string   `json:"region,omitempty"`
	City              string   `json:"city,omitempty"`
	Remark            string   `json:"remark,omitempty"`
	IPChecker         string   `json:"ipchecker,omitempty"`
	SysAppCateId      string   `json:"sys_app_cate_Id,omitempty"`
}

func (a *AdsPower) CreateProfile(ctx context.Context, groupId string, proxyConfig *ProxyConfig, fingerprintConfig *FingerprintConfig, opts ...*CreateProfileOptions) (string, error) {
	var opts_ *CreateProfileOptions
	if len(opts) != 0 {
		opts_ = opts[0]
	}

	payload := &createProfileRequest{groupId, proxyConfig, fingerprintConfig, opts_}

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(b)

	url := fmt.Sprintf("%s/create", UserApi)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")

	a.rl.Take()
	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	decodedBody, err := decodeResponseBody[createProfileResponse](resp)
	if err != nil {
		return "", err
	}

	err = handleResponseError(decodedBody)
	if err != nil {
		return "", err
	}

	return decodedBody.Data.Id, nil
}
