package code

import (
	"errors"
	"fmt"
)

var (
	ErrIdempotent = errors.New("idempotent failed")
)

func IsIdempotentError(err error) bool {
	return errors.Is(err, ErrIdempotent)
}

type BizError struct {
	msg string
}

func NewBizError(msg string) error {
	return &BizError{msg: msg}
}

func (e *BizError) Error() string {
	return e.msg
}

type ExtSrvError struct {
	cause   error
	service string
	method  string
}

func NewExtSrvError(cause error, service string, method string) error {
	return &ExtSrvError{
		cause:   cause,
		service: service,
		method:  method,
	}
}

func (e *ExtSrvError) Error() string {
	return fmt.Sprintf("%s-%s:%s", e.service, e.method, e.cause.Error())
}

func (e *ExtSrvError) Unwrap() error {

	return e.cause
}
