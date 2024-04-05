package adspower

import "context"

type iResponseMessage interface {
	GetCode() int
	GetMsg() string
}

type IProxySeller interface {
	GetResIdentialProxy(ctx context.Context, options ...any) (string, error)
}
