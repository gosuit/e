package e

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
)

type errorStruct struct {
	message     string
	errs        []error
	tags        map[string]any
	code        Status
	source_file string
	source_line int
	log         sl.Logger
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

func (e *errorStruct) GetTag(key string) any {
	return e.tags[key]
}

func (e *errorStruct) GetSource() (string, int) {
	return e.source_file, e.source_line
}

func (e *errorStruct) WithMessage(msg string) Error {
	_, file, line, _ := runtime.Caller(1)

	return &errorStruct{
		message:     msg,
		errs:        e.errs,
		tags:        e.tags,
		code:        e.code,
		log:         e.log,
		source_file: file,
		source_line: line,
	}
}

func (e *errorStruct) WithCode(status Status) Error {
	_, file, line, _ := runtime.Caller(1)

	return &errorStruct{
		message:     e.message,
		errs:        e.errs,
		tags:        e.tags,
		code:        status,
		log:         e.log,
		source_file: file,
		source_line: line,
	}
}

func (e *errorStruct) WithErr(err error) Error {
	_, file, line, _ := runtime.Caller(1)

	return &errorStruct{
		message:     e.message,
		errs:        append(e.errs, err),
		tags:        e.tags,
		code:        e.code,
		log:         e.log,
		source_file: file,
		source_line: line,
	}
}

func (e *errorStruct) WithTag(key string, value any) Error {
	e.tags[key] = value

	_, file, line, _ := runtime.Caller(1)

	return &errorStruct{
		message:     e.message,
		errs:        e.errs,
		tags:        e.tags,
		code:        e.code,
		log:         e.log,
		source_file: file,
		source_line: line,
	}
}

func (e *errorStruct) WithCtx(c lec.Context) Error {
	_, file, line, _ := runtime.Caller(1)

	err := &errorStruct{
		message:     e.message,
		errs:        e.errs,
		tags:        e.tags,
		code:        e.code,
		log:         sl.New(c.Logger().Config()),
		source_file: file,
		source_line: line,
	}

	ctxErr := c.Err()
	if ctxErr != nil {
		err.errs = append(err.errs, c.Err())
	}

	for key, value := range c.GetValues() {
		if value.Share {
			err.tags[key] = value.Val
		}
	}

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

	l = l.With(
		sl.StringAttr("error_code", e.code.ToString()),
		sl.StringAttr("error_source", fmt.Sprintf("file: %s line: %d", e.source_file, e.source_line)),
	)

	l.Error(message)
}
