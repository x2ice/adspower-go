package adspower

import (
	"errors"
	"fmt"
	"strings"
)

type errNoGroupsFound error

var (
	ErrProfileLimitReached error
	ErrProxyFailure        error
	ErrGroupNotFound       error
	ErrNoGroupsFound       errNoGroupsFound = errors.New("no groups found")
	ErrOpenBrowserFailure  error
	ErrStopBrowserFailure  error
	ErrUnknownError        error
)

func handleResponseError(r iResponseMessage) error {
	if r.GetCode() == 0 {
		return nil
	}

	msg := r.GetMsg()
	switch {
	case strings.Contains(msg, "proxy"):
		ErrProxyFailure = fmt.Errorf(msg)
		return ErrProxyFailure

	case strings.Contains(msg, "not open"):
		ErrStopBrowserFailure = fmt.Errorf(msg)
		return ErrStopBrowserFailure

	case strings.Contains(msg, "exceeds"):
		ErrProfileLimitReached = fmt.Errorf(msg)
		return ErrProfileLimitReached

	default:
		ErrUnknownError = fmt.Errorf(msg)
		return ErrUnknownError
	}
}
