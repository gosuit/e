package e

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/gosuit/lec"
)

type errorStruct struct {
	message string
	errs    []error
	tags    map[string]interface{}
	code    Status
	log     *slog.Logger
}

func (e *errorStruct) GetMessage() string {
	return e.message
}

func (e *errorStruct) GetCode() Status {
	return e.code
}

func (e *errorStruct) GetError() error {
	return errors.Join(e.errs...)
}

func (e *errorStruct) GetTag(key string) interface{} {
	return e.tags[key]
}

func (e *errorStruct) WithMessage(msg string) Error {
	return New(msg, e.code, e.errs...)
}

func (e *errorStruct) WithCode(status Status) Error {
	return New(e.message, status, e.errs...)
}

func (e *errorStruct) WithErr(err error) Error {
	return New(e.message, e.code, append(e.errs, err)...)
}

func (e *errorStruct) WithTag(key string, value interface{}) Error {
	err := New(e.message, e.code, e.errs...).(*errorStruct)

	for key, value := range e.tags {
		err.tags[key] = value
	}

	err.tags[key] = value

	return err
}

func (e *errorStruct) WithCtx(c lec.Context) Error {
	err := New(e.message, e.code, e.errs...).(*errorStruct)

	ctxErr := c.Err()
	if ctxErr != nil {
		err.errs = append(err.errs, c.Err())
	}

	for key, value := range e.tags {
		err.tags[key] = value
	}

	for key, value := range c.GetValues() {
		if value.Share {
			err.tags[key] = value.Val
		}
	}

	err.log = slog.New(c.SlHandler())

	c.AddErr(err)

	return err
}

func (e *errorStruct) Log(msg ...string) {
	l := e.log.With(e.SlErr())

	for key, value := range e.tags {
		l = l.With(key, value)
	}

	message := ""

	if len(msg) != 0 {
		message = strings.Join(msg, " ")
	}

	l.Error(message)
}
