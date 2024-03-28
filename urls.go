package adspower

import "fmt"

const BASE_HTTP_API = "http://local.adspower.net:50325/api/v1"

var (
	RootUrl  = fmt.Sprintf("%s/browser", BASE_HTTP_API)
	GroupApi = fmt.Sprintf("%s/group", BASE_HTTP_API)
	UserApi  = fmt.Sprintf("%s/user", BASE_HTTP_API)
)
