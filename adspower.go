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
	duration := time.Duration(float64(time.Second) * 1.5)
	ratelimitPer := ratelimit.Per(duration)
	return &AdsPower{
		HTTPClient: &http.Client{},
		rl:         ratelimit.New(1, ratelimitPer),
	}
}
