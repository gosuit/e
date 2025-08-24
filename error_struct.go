package e

import (
	"errors"
	"runtime"

	"github.com/gosuit/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (e *errorStruct) GetStatus() Status {
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

func (e *errorStruct) WithStatus(status Status) Error {
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

type jsonError struct {
	Error string `json:"error"`
}

func (e *errorStruct) ToJson() jsonError {
	return jsonError{
		Error: e.message,
	}
}

func (e *errorStruct) Error() string {
	if e.message == "" && (e.errs == nil || (e.errs != nil && len(e.errs) == 0)) {
		return "<nil>"
	}

	if e.errs == nil || (e.errs != nil && len(e.errs) == 0) {
		return e.message
	}

	if e.message == "" {
		return errors.Join(e.errs...).Error()
	}

	return e.message + ": " + errors.Join(e.errs...).Error()
}

func (e *errorStruct) GetHttpCode() int {
	return e.code.ToHttp()
}

func (e *errorStruct) GetGrpcCode() codes.Code {
	return e.code.ToGRPC()
}

func (e *errorStruct) ToGRPC() error {
	return status.Error(e.GetGrpcCode(), e.message)
}

func (e *errorStruct) SlErr() sl.Attr {
	return sl.StringAttr("error", e.Error())
}
