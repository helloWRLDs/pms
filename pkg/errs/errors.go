package errs

import (
	"fmt"
	"strings"
)

type ErrInvalidInput struct {
	Object string
	Reason string
}

func (e ErrInvalidInput) Error() string {
	return strings.Trim(fmt.Sprintf("invalid input format for %s. %s", e.Object, e.Reason), ". ")
}

type ErrBadGateway struct {
	Object string
}

func (e ErrBadGateway) Error() string {
	return fmt.Sprintf("invalid format for %s", e.Object)
}

type ErrNotFound struct {
	Object string
	Field  string
	Value  string
}

func (e ErrNotFound) Error() string {
	var with string
	if e.Field != "" {
		with = fmt.Sprintf(" with %s = %s", e.Field, e.Value)
	}
	return fmt.Sprintf("%s%s not found", e.Object, with)
}

type ErrUnauthorized struct {
	Reason string
}

func (e ErrUnauthorized) Error() string {
	return fmt.Sprintf("authorization failed: %s", e.Reason)
}

type ErrForbidden struct {
	Reason string
}

func (e ErrForbidden) Error() string {
	return fmt.Sprintf("forbidden action: %s", e.Reason)
}

type ErrConflict struct {
	Reason string
}

func (e ErrConflict) Error() string {
	return fmt.Sprintf("conflict: %s", e.Reason)
}

type ErrTooManyRequests struct {
	Reason string
}

func (e ErrTooManyRequests) Error() string {
	return e.Reason
}

type ErrInternal struct {
	Reason string
}

func (e ErrInternal) Error() string {
	return e.Reason
}
