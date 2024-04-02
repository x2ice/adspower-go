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

type ProxyConfig struct {
	Soft        string `json:"proxy_soft"`
	Type        string `json:"proxy_type,omitempty"`
	Host        string `json:"proxy_host,omitempty"`
	Port        string `json:"proxy_port,omitempty"`
	User        string `json:"proxy_user,omitempty"`
	Password    string `json:"proxy_password,omitempty"`
	ChangeIpUrl string `json:"proxy_url,omitempty"`
}

type FingerprintConfig struct {
	AutomaticTimezone   string               `json:"automatic_timezone,omitempty"`
	Timezone            string               `json:"timezone,omitempty"`
	WebRTC              string               `json:"webrtc,omitempty"`
	Location            string               `json:"location,omitempty"`
	LocationSwitch      string               `json:"location_switch,omitempty"`
	Longitude           string               `json:"longitude,omitempty"`
	Latitude            string               `json:"latitude,omitempty"`
	Accuracy            string               `json:"accuracy,omitempty"`
	Language            []string             `json:"language,omitempty"`
	LanguageSwitch      string               `json:"language_switch,omitempty"`
	PageLanguageSwitch  string               `json:"page_language_switch,omitempty"`
	PageLanguage        string               `json:"page_language,omitempty"`
	UA                  string               `json:"ua,omitempty"`
	ScreenResolution    string               `json:"screen_resolution,omitempty"`
	Fonts               []string             `json:"fonts,omitempty"`
	Canvas              string               `json:"canvas,omitempty"`
	WebGLImage          string               `json:"webgl_image,omitempty"`
	WebGL               string               `json:"webgl,omitempty"`
	WebGLConfig         *WebGLConfig         `json:"webgl_config,omitempty"`
	Audio               string               `json:"audio,omitempty"`
	DoNotTrack          string               `json:"do_not_track,omitempty"`
	HardwareConcurrency string               `json:"hardware_concurrency,omitempty"`
	DeviceMemory        string               `json:"device_memory,omitempty"`
	Flash               string               `json:"flash,omitempty"`
	ScanPortType        string               `json:"scan_port_type,omitempty"`
	AllowScanPorts      []string             `json:"allow_scan_ports,omitempty"`
	MediaDevices        string               `json:"media_devices,omitempty"`
	MediaDevicesNum     *MediaDevicesNum     `json:"media_devices_num,omitempty"`
	AdsPowerRects       string               `json:"AdsPower_rects,omitempty"`
	DeviceNameSwitch    string               `json:"device_name_switch,omitempty"`
	DeviceName          string               `json:"device_name,omitempty"`
	RandomUA            *RandomUA            `json:"random_ua,omitempty"`
	SpeechSwitch        string               `json:"speech_switch,omitempty"`
	MacAddressConfig    *MacAddressConfig    `json:"mac_address_config,omitempty"`
	BrowserKernelConfig *BrowserKernelConfig `json:"browser_kernel_config,omitempty"`
	GPU                 int                  `json:"gpu,omitempty"`
}

type WebGLConfig struct {
	UnmaskedVendor   string  `json:"unmasked_vendor,omitempty"`
	UnmaskedRenderer string  `json:"unmasked_renderer,omitempty"`
	WebGPU           *WebGPU `json:"webgpu,omitempty"`
}

type WebGPU struct {
	Switch string `json:"webgpu_switch,omitempty"`
}

type MediaDevicesNum struct {
	AudioInputNum  string `json:"audioinput_num,omitempty"`
	VideoInputNum  string `json:"videoinput_num,omitempty"`
	AudioOutputNum string `json:"audiooutput_num,omitempty"`
}

type RandomUA struct {
	UABrowser       []string `json:"ua_browser,omitempty"`
	UAVersion       []string `json:"ua_version,omitempty"`
	UASystemVersion []string `json:"ua_system_version,omitempty"`
}

type MacAddressConfig struct {
	Model   string `json:"model,omitempty"`
	Address string `json:"address,omitempty"`
}

type BrowserKernelConfig struct {
	Version string `json:"version,omitempty"`
	Type    string `json:"type,omitempty"`
}

type Cookie struct {
	Domain         string `json:"domain"`
	ExpirationDate string `json:"expirationDate"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	SameSite       string `json:"sameSite"`
	Secure         bool   `json:"secure"`
	Value          string `json:"value"`
	ID             int    `json:"id"`
}

type Cookies []*Cookie

type createProfileRequest struct {
	GroupId           string             `json:"group_id"`
	ProxyConfig       *ProxyConfig       `json:"user_proxy_config"`
	FingerprintConfig *FingerprintConfig `json:"fingerprint_config"`
	*CreateProfileOptions
}

type createProfileResponse struct {
	*response
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

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

type updateProfileRequest struct {
	Id string `json:"user_id"`
	*UpdateProfileOptions
}

type updateProfileResponse struct {
	response
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

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
	GroupId      string `json:"group_id"`
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
