package adspower

import (
	"context"
	"fmt"

	json "github.com/goccy/go-json"
)

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

type ProxyConfig struct {
	ProxySoft     string `json:"proxy_soft"`
	ProxyType     string `json:"proxy_type,omitempty"`
	ProxyHost     string `json:"proxy_host,omitempty"`
	ProxyPort     string `json:"proxy_port,omitempty"`
	ProxyUser     string `json:"proxy_user,omitempty"`
	ProxyPassword string `json:"proxy_password,omitempty"`
	ProxyURL      string `json:"proxy_url,omitempty"`
}

type FingerprintConfig struct {
	AutomaticTimezone   string              `json:"automatic_timezone"`
	Timezone            string              `json:"timezone,omitempty"`
	WebRTC              string              `json:"webrtc,omitempty"`
	Location            string              `json:"location,omitempty"`
	LocationSwitch      string              `json:"location_switch,omitempty"`
	Longitude           string              `json:"longitude,omitempty"`
	Latitude            string              `json:"latitude,omitempty"`
	Accuracy            string              `json:"accuracy,omitempty"`
	Language            []string            `json:"language,omitempty"`
	LanguageSwitch      string              `json:"language_switch,omitempty"`
	PageLanguageSwitch  string              `json:"page_language_switch,omitempty"`
	PageLanguage        string              `json:"page_language,omitempty"`
	UA                  string              `json:"ua,omitempty"`
	ScreenResolution    string              `json:"screen_resolution,omitempty"`
	Fonts               []string            `json:"fonts,omitempty"`
	Canvas              string              `json:"canvas,omitempty"`
	WebGLImage          string              `json:"webgl_image,omitempty"`
	WebGL               string              `json:"webgl,omitempty"`
	WebGLConfig         WebGLConfig         `json:"webgl_config,omitempty"`
	Audio               string              `json:"audio,omitempty"`
	DoNotTrack          string              `json:"do_not_track,omitempty"`
	HardwareConcurrency string              `json:"hardware_concurrency,omitempty"`
	DeviceMemory        string              `json:"device_memory,omitempty"`
	Flash               string              `json:"flash,omitempty"`
	ScanPortType        string              `json:"scan_port_type,omitempty"`
	AllowScanPorts      []string            `json:"allow_scan_ports,omitempty"`
	MediaDevices        string              `json:"media_devices,omitempty"`
	MediaDevicesNum     MediaDevicesNum     `json:"media_devices_num,omitempty"`
	AdsPowerRects       string              `json:"AdsPower_rects,omitempty"`
	DeviceNameSwitch    string              `json:"device_name_switch,omitempty"`
	DeviceName          string              `json:"device_name,omitempty"`
	RandomUA            RandomUA            `json:"random_ua,omitempty"`
	SpeechSwitch        string              `json:"speech_switch,omitempty"`
	MacAddressConfig    MacAddressConfig    `json:"mac_address_config,omitempty"`
	BrowserKernelConfig BrowserKernelConfig `json:"browser_kernel_config,omitempty"`
	GPU                 int                 `json:"gpu,omitempty"`
}

type WebGLConfig struct {
	UnmaskedVendor   string `json:"unmasked_vendor,omitempty"`
	UnmaskedRenderer string `json:"unmasked_renderer,omitempty"`
	WebGPU           WebGPU `json:"webgpu,omitempty"`
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

type CreateProfileOptions struct {
	Name              string   `json:"name"`
	DomainName        string   `json:"domain_name"`
	OpenURLs          []string `json:"open_urls"`
	RepeatConfig      []int    `json:"repeat_config"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
	Fakey             string   `json:"fakey"`
	Cookie            []Cookie `json:"cookie"`
	IgnoreCookieError int      `json:"ignore_cookie_error"`
	GroupID           string   `json:"group_id"`
	IP                string   `json:"ip"`
	Country           string   `json:"country"`
	Region            string   `json:"region"`
	City              string   `json:"city"`
	Remark            string   `json:"remark"`
	IPChecker         string   `json:"ipchecker"`
	SysAppCateID      string   `json:"sys_app_cate_id"`
}

type NewProfileResponse struct {
	Code int `json:"code"`
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func (c *AdsPower) CreateProfile(ctx context.Context, groupID string, proxyConfig *ProxyConfig, fingerprintConfig *FingerprintConfig, opts ...*CreateProfileOptions) (error, error) {
	// url_ := fmt.Sprintf("%s/create", UserApi)

	type Payload struct {
		GroupID           string             `json:"group_id"`
		ProxyConfig       *ProxyConfig       `json:"user_proxy_config"`
		FingerprintConfig *FingerprintConfig `json:"fingerprint_config"`
		*CreateProfileOptions
	}

	var opts_ *CreateProfileOptions
	if len(opts) != 0 {
		opts_ = opts[0]
	}

	payload := &Payload{groupID, proxyConfig, fingerprintConfig, opts_}

	fmt.Println(payload.GroupID)

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// buf := bytes.NewBuffer(b)

	// req, err := http.NewRequestWithContext(ctx, http.MethodPost, url_, buf)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = c.HTTPClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
