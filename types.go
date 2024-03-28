package adspower

type response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (o *response) GetCode() int {
	return o.Code
}

func (o *response) GetMsg() string {
	return o.Msg
}

// ------------------------------ //

type closeBrowserResponse struct {
	response
}

// ------------------------------ //

type openBrowserResponse struct {
	response
	Data *Browser `json:"data"`
}

type openedBrowsersResponse struct {
	response
	Data struct {
		List Browsers `json:"list"`
	} `json:"data"`
}

type Browser struct {
	ID string `json:"user_id"`
	Ws struct {
		Selenium  string `json:"selenium"`
		Puppeteer string `json:"puppeteer"`
	} `json:"ws"`
	DebugPort string `json:"debug_port"`
	WebDriver string `json:"webdriver"`
}

type Browsers []*Browser

// ------------------------------ //

type queryGroupResponse struct {
	response
	Data struct {
		List   []*Group `json:"list"`
		Offset int      `json:"page"`
		Limit  int      `json:"page_size"`
	} `json:"data"`
}

type Group struct {
	ID     string `json:"group_id"`
	Name   string `json:"group_name"`
	Remark string `json:"remark,omitempty"`
}

type Groups []*Group

// ------------------------------ //

type queryProfilesResponse struct {
	response
	Data struct {
		List   Profiles `json:"list"`
		Offset int      `json:"page"`
		Limit  int      `json:"page_size"`
	} `json:"data"`
}

type Profile struct {
	SerialNumber string `json:"serial_number"`
	ID           string `json:"user_id"`
	Name         string `json:"name"`
	GroupID      string `json:"group_id"`
	GroupName    string `json:"group_name"`
	DomainName   string `json:"domain_name"`
	Username     string `json:"username"`
	Remark       string `json:"remark"`
	SysAppCateID string `json:"sys_app_cate_id"`
	CreatedTime  string `json:"created_time"`
	IP           string `json:"ip"`
	IPCountry    string `json:"ip_country"`
	Password     string `json:"password"`
	LastOpenTime string `json:"last_open_time"`
}

type Profiles []*Profile
