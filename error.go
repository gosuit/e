package e

import (
	"runtime"

	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error interface defines a custom error type that provides additional context
// and functionality beyond the standard error interface in Go. It is designed
// to encapsulate error messages, status codes, and conversion methods for gRPC
// and HTTP responses.
type Error interface {
	// GetMessage returns the error message as a string.
	// This method allows users to retrieve a human-readable description of the error.
	GetMessage() string

	// GetError returns the underlying error.
	// This method provides access to the original error that may contain more details.
	GetError() error

	// GetTag returns the tag value associated with the given key.
	// This method allows users to retrieve additional context information attached to the error.
	GetTag(key string) any

	// GetCode returns the status code associated with the error.
	// The StatusType can be a custom type that represents various error codes.
	GetCode() Status

	// GetSource returns the source file and line where error was created.
	GetSource() (string, int)

	// Log logs the error with an optional message, error creation source and other metadata.
	// This method provides a way to log error events with custom messages for debugging and monitoring.
	Log(msg ...string)

	// WithMessage sets a new error message for the error instance.
	// This method allows users to update the error message dynamically.
	WithMessage(string) Error

	// WithErr sets a new underlying error for the error instance.
	// This method allows users to associate a different error with this custom error type.
	WithErr(error) Error

	// WithTag sets a new tag (key-value pair) to the error instance.
	// This method allows users to add contextual information to the error.
	WithTag(key string, value any) Error

	// WithCtx sets a context to the error instance.
	// This method allows users to associate a context with the error.
	// It change error logger with logger from context, add err to context.
	WithCtx(c lec.Context) Error

	// WithCode sets a new status code for the error instance.
	// This method allows users to update the error code dynamically.
	WithCode(Status) Error

	// ToJson() returns error struct with json tags.
	ToJson() jsonError

	// ToGRPCCode converts the error's status code to a gRPC error code.
	// This method facilitates interoperability with gRPC services by providing
	// an appropriate error code representation.
	ToGRPCCode() codes.Code

	// ToHttpCode converts the error's status code to an HTTP status code.
	// This method helps in mapping application-specific errors to standard HTTP responses.
	ToHttpCode() int

	// Error returns the string representation of the error.
	// This method implements the standard error interface, allowing the error
	// to be used in contexts where a simple error message is required.
	Error() string

	// SlErr returns structured logging attributes for the error.
	// This method provides a way to log the error with additional context,
	// making it easier to analyze issues in logs.
	SlErr() sl.Attr

	// ToGRPCErr converts the custom error into a standard Go error type suitable for gRPC.
	// This method allows seamless integration with gRPC error handling mechanisms.
	ToGRPC() error
}

// New returns type Error with message.
func New(msg string, status Status, errs ...error) Error {
	if errs == nil {
		errs = []error{}
	}

	_, file, line, _ := runtime.Caller(1)

	return &errorStruct{
		message:     msg,
		errs:        errs,
		tags:        make(map[string]any),
		code:        status,
		log:         sl.Default(),
		source_file: file,
		source_line: line,
	}
}

// E creates a new custom error instance if the provided error is not nil.
// It initializes the custom error with an empty message and associates the given error
// with an internal status code by default. If the provided error is nil, it returns nil.
func E(err error) Error {
	if err != nil {
		if _, ok := err.(Error); ok {
			return err.(Error)
		}

		_, file, line, _ := runtime.Caller(1)

		return &errorStruct{
			message:     "",
			errs:        []error{err},
			tags:        make(map[string]any),
			code:        Internal,
			log:         sl.Default(),
			source_file: file,
			source_line: line,
		}
	}

	return nil
}

// FromGRPC converts a gRPC error to an Error type.
// It extracts the gRPC status code and message, maps the gRPC code to a predefined internal Status,
// and creates a new Error instance with the extracted message and mapped Status.
func FromGRPC(err error) Error {
	if err == nil {
		return nil
	}

	_, file, line, _ := runtime.Caller(1)

	stat, ok := status.FromError(err)
	if !ok {
		return &errorStruct{
			message:     "",
			errs:        []error{err},
			tags:        make(map[string]any),
			code:        Internal,
			log:         sl.Default(),
			source_file: file,
			source_line: line,
		}
	}

	var code Status

	switch stat.Code() {

	case codes.Internal:
		code = Internal

	case codes.NotFound:
		code = NotFound

	case codes.InvalidArgument:
		code = BadInput

	case codes.Unauthenticated:
		code = Unauthorize

	case codes.AlreadyExists:
		code = Conflict

	case codes.PermissionDenied:
		code = Forbidden

	default:
		code = Internal

	}

	return &errorStruct{
		message:     stat.Message(),
		errs:        make([]error, 0),
		tags:        make(map[string]any),
		code:        code,
		log:         sl.Default(),
		source_file: file,
		source_line: line,
	}
}
