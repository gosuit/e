package e

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Status represents different error types that can occur in the system.
type Status int

const (
	// Internal means an internal server error.
	Internal Status = iota

	// NotFound means that the requested resource was not found.
	NotFound

	// BadInput means that the request input is invalid.
	BadInput

	// Conflict means a conflict when performing an action, for example, trying to create an already existing resource.
	Conflict

	// Forbidden means that the action is forbidden for the user.
	Forbidden

	// Unauthorize means that the user is not authorized to perform the action.
	Unauthorize
)

// ToString converts the Status value to its string representation.
func (s Status) ToString() string {
	switch s {

	case Internal:
		return "Internal"

	case NotFound:
		return "NotFound"

	case BadInput:
		return "BadInput"

	case Unauthorize:
		return "Unauthorize"

	case Forbidden:
		return "Forbidden"

	case Conflict:
		return "Conflict"

	default:
		return "Internal"
	}
}

// ToHttp converts the Status value to the corresponding HTTP status code.
func (s Status) ToHttp() int {
	switch s {

	case Internal:
		return http.StatusInternalServerError

	case NotFound:
		return http.StatusNotFound

	case BadInput:
		return http.StatusBadRequest

	case Unauthorize:
		return http.StatusUnauthorized

	case Forbidden:
		return http.StatusForbidden

	case Conflict:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}

// ToGRPC converts the Status value to the corresponding gRPC status code.
func (s Status) ToGRPC() codes.Code {
	switch s {

	case Internal:
		return codes.Internal

	case NotFound:
		return codes.NotFound

	case BadInput:
		return codes.InvalidArgument

	case Unauthorize:
		return codes.Unauthenticated

	case Conflict:
		return codes.AlreadyExists

	case Forbidden:
		return codes.PermissionDenied

	default:
		return codes.Internal

	}
}
