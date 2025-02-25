package e

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type Status int

const (
	Internal Status = iota
	NotFound
	BadInput
	Conflict
	Forbidden
	Unauthorize
)

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

	default:
		return codes.Internal

	}
}
