package adspower

import (
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

type AdsPower struct {
	HTTPClient *http.Client
	rl         ratelimit.Limiter
}

func NewAdsPower() *AdsPower {
	return &AdsPower{
		HTTPClient: &http.Client{},
		rl:         ratelimit.New(1, ratelimit.Per(time.Second*1)),
	}
}
