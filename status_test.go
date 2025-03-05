package e

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestToString(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{status: Internal, expected: "Internal"},
		{status: NotFound, expected: "NotFound"},
		{status: BadInput, expected: "BadInput"},
		{status: Conflict, expected: "Conflict"},
		{status: Forbidden, expected: "Forbidden"},
		{status: Unauthorize, expected: "Unauthorize"},
		{status: 99, expected: "Internal"}, // Testing default case
	}

	for _, tt := range tests {
		value := tt.status.ToString()
		assert.Equal(t, value, tt.expected)
	}
}

func TestToHttp(t *testing.T) {
	tests := []struct {
		status   Status
		expected int
	}{
		{status: Internal, expected: http.StatusInternalServerError},
		{status: NotFound, expected: http.StatusNotFound},
		{status: BadInput, expected: http.StatusBadRequest},
		{status: Conflict, expected: http.StatusConflict},
		{status: Forbidden, expected: http.StatusForbidden},
		{status: Unauthorize, expected: http.StatusUnauthorized},
		{status: 99, expected: http.StatusInternalServerError}, // Testing default case
	}

	for _, tt := range tests {
		value := tt.status.ToHttp()
		assert.Equal(t, value, tt.expected)
	}
}

func TestToGRPCStatus(t *testing.T) {
	tests := []struct {
		status   Status
		expected codes.Code
	}{
		{status: Internal, expected: codes.Internal},
		{status: NotFound, expected: codes.NotFound},
		{status: BadInput, expected: codes.InvalidArgument},
		{status: Conflict, expected: codes.AlreadyExists},
		{status: Forbidden, expected: codes.PermissionDenied},
		{status: Unauthorize, expected: codes.Unauthenticated},
		{status: 99, expected: codes.Internal}, // Testing default case
	}

	for _, tt := range tests {
		value := tt.status.ToGRPC()
		assert.Equal(t, value, tt.expected)
	}
}
