package adspower

import (
	"errors"
	"fmt"
	"strings"
)

type (
	ErrProfileLimitExceeded error
	ErrProxyFailure         error
	ErrGroupNotFound        error
	ErrNoGroupsFound        error
	ErrOpenBrowserFailure   error
	ErrCloseBrowserFailure  error
	ErrInvalidProxyFormat   error
)

var (
	errProfileLimitReached ErrProfileLimitExceeded = errors.New("profile limit reached")
	errProxyFailure        ErrProxyFailure         = errors.New("proxy failure")
	errInvalidProxyFormat  ErrInvalidProxyFormat   = errors.New("invalid proxy format")
	errGroupNotFound       ErrGroupNotFound        = errors.New("group not found")
	errNoGroupsFound       ErrNoGroupsFound        = errors.New("no groups found")
	errOpenBrowserFailure  ErrOpenBrowserFailure   = errors.New("open browser failure")
	errStopBrowserFailure  ErrCloseBrowserFailure  = errors.New("close browser failure")
	ErrUnknownError        error
)

func handleResponseError(r iResponseMessage) error {
	if r.GetCode() == 0 {
		return nil
	}

	msg := r.GetMsg()
	switch {
	case strings.Contains(msg, "proxy"):
		return errProxyFailure

	case strings.Contains(msg, "not open"):
		return errStopBrowserFailure

	case strings.Contains(msg, "exceeds"):
		return errProfileLimitReached

	case strings.Contains(msg, "does not exist"):
		return errOpenBrowserFailure

	default:
		ErrUnknownError = fmt.Errorf(msg)
		return ErrUnknownError
	}
}
