package e

import (
	"errors"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestToJson(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  Error
		want jsonError
	}{
		{
			name: "Default",
			err:  New("some error", Internal),
			want: jsonError{Error: "some error"},
		},
		{
			name: "With err",
			err:  New("some error", Internal, errors.New("invalid data")),
			want: jsonError{Error: "some error"},
		},
		{
			name: "Empty message",
			err:  New("", Internal),
			want: jsonError{Error: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.err.ToJson())
		})
	}
}

func TestToHttpCode(t *testing.T) {
	tests := []struct {
		code     Status
		expected int
	}{
		{Internal, http.StatusInternalServerError},
		{NotFound, http.StatusNotFound},
		{BadInput, http.StatusBadRequest},
		{Unauthorize, http.StatusUnauthorized},
		{Forbidden, http.StatusForbidden},
		{Conflict, http.StatusConflict},
		{99, http.StatusInternalServerError}, // Testing default case
	}

	for _, tt := range tests {
		err := New("", tt.code, nil)
		assert.Equal(t, tt.expected, err.ToHttpCode())
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		err    []error
		msg    string
		expect string
	}{{
		name:   "AllFields",
		err:    []error{errors.New("some error")},
		msg:    "some message",
		expect: "some message: some error",
	},
		{
			name:   "Only Message",
			err:    []error{},
			msg:    "some message",
			expect: "some message",
		},
		{
			name:   "Only error",
			err:    []error{errors.New("some error")},
			msg:    "",
			expect: "some error",
		},
		{
			name:   "Nil fileds",
			err:    []error{},
			msg:    "",
			expect: "<nil>",
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := New(tc.msg, Internal, tc.err...)

			assert.Equal(t, tc.expect, err.Error())
		})
	}
}

func TestToGRPCErr(t *testing.T) {
	msg := "This is a gRPC error message"
	testErr := errors.New("some grpc error")
	err := New(msg, Internal, testErr)

	grpcErr := err.ToGRPC()
	if grpcErr == nil {
		t.Fatal("Expected non-nil gRPC error")
	}

	// Check if the gRPC error message is as expected
	assert.ErrorIs(t, grpcErr, status.Error(codes.Internal, msg))

	// Check if the gRPC error code is as expected
	grpcStatus, ok := status.FromError(grpcErr)
	if !ok {
		t.Fatal("Expected gRPC status from error")
	}

	if grpcStatus.Code() != codes.Internal {
		t.Errorf("Expected gRPC code %v; got %v", codes.Internal, grpcStatus.Code())
	}
}

func TestToGRPCCode(t *testing.T) {
	tests := []struct {
		code     Status
		expected codes.Code
	}{
		{code: Internal, expected: codes.Internal},
		{code: NotFound, expected: codes.NotFound},
		{code: BadInput, expected: codes.InvalidArgument},
		{code: Conflict, expected: codes.AlreadyExists},
		{code: 99, expected: codes.Internal},
	}

	for _, tc := range tests {
		err := New("", tc.code, nil)
		assert.Equal(t, tc.expected, err.ToGRPCCode())
	}
}

func TestSlErr(t *testing.T) {
	msg := "test error for slog"
	testErr := errors.New("some error")
	code := Internal
	err := New(msg, code, testErr)

	assert.Equal(t, slog.String("error", err.Error()), err.SlErr())
}
