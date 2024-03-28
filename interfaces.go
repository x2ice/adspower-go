package adspower

type iResponseMessage interface {
	GetCode() int
	GetMsg() string
}
