package e

import (
	"errors"
	"log/slog"

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

// ToHttpCode convert Error to http status code.
func (e *errorStruct) ToHttpCode() int {
	return e.code.ToHttp()
}

// ToGRPCCode convert Error to grpc status code.
func (e *errorStruct) ToGRPCCode() codes.Code {
	return e.code.ToGRPC()
}

func (e *errorStruct) ToGRPC() error {
	return status.Error(e.ToGRPCCode(), e.message)
}

func (e *errorStruct) Error() string {
	if e.message == "" && (e.errs == nil || (len(e.errs) == 0)) {
		return "nil"
	}
	if e.errs == nil || (len(e.errs) == 0) {
		return e.message
	}
	if e.message == "" {
		return errors.Join(e.errs...).Error()
	}
	return e.message + ": " + errors.Join(e.errs...).Error()
}

func (e *errorStruct) SlErr() slog.Attr {
	return slog.String("error", e.Error())
}
