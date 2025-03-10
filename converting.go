package e

import (
	"errors"

	"github.com/gosuit/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type jsonError struct {
	Error string `json:"error"`
}

func (e *errorStruct) ToJson() jsonError {
	return jsonError{
		Error: e.message,
	}
}

func (e *errorStruct) ToHttpCode() int {
	return e.code.ToHttp()
}

func (e *errorStruct) ToGRPCCode() codes.Code {
	return e.code.ToGRPC()
}

func (e *errorStruct) ToGRPC() error {
	return status.Error(e.ToGRPCCode(), e.message)
}

func (e *errorStruct) Error() string {
	if e.message == "" && (e.errs == nil || (len(e.errs) == 0)) {
		return "<nil>"
	}

	if e.errs == nil || (len(e.errs) == 0) {
		return e.message
	}

	if e.message == "" {
		return errors.Join(e.errs...).Error()
	}

	return e.message + ": " + errors.Join(e.errs...).Error()
}

func (e *errorStruct) SlErr() sl.Attr {
	return sl.StringAttr("error", e.Error())
}
