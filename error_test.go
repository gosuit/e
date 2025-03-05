package e

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNew(t *testing.T) {
	msg := "An error occurred"
	testErr := errors.New("some error")
	code := NotFound

	err := New(msg, code, testErr)

	if err.GetMessage() != msg {
		t.Errorf("Expected message %q, got %q", msg, err.GetMessage())
	}

	if errors.Is(err.GetError(), err) {
		t.Errorf("Expected message %s, got %s", err.Error(), err.GetError().Error())
	}

	if err.GetCode() != code {
		t.Errorf("Expected code %v, got %v", code, err.GetCode())
	}
}

func TestE(t *testing.T) {
	t.Parallel()

	testErr := errors.New("some error")
	customErr := New("", Internal, testErr)

	tests := []struct {
		name string
		err  error
		want Error
	}{
		{
			name: "Nil error",
			err:  nil,
			want: nil,
		},
		{
			name: "Not nil error",
			err:  customErr,
			want: customErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, E(tt.err))
		})
	}
}

func TestToGRPC(t *testing.T) {
	tests := []struct {
		msg      string
		grpcCode codes.Code
		status   Status
	}{
		{
			msg:      "some msg",
			grpcCode: codes.Internal,
			status:   Internal,
		},
		{
			msg:      "some msg",
			grpcCode: codes.AlreadyExists,
			status:   Conflict,
		},
		{
			msg:      "some msg",
			grpcCode: codes.InvalidArgument,
			status:   BadInput,
		},
		{
			msg:      "some msg",
			grpcCode: codes.NotFound,
			status:   NotFound,
		},
		{
			msg:      "some msg",
			grpcCode: codes.PermissionDenied,
			status:   Forbidden,
		},
		{
			msg:      "some msg",
			grpcCode: codes.Unauthenticated,
			status:   Unauthorize,
		},
	}

	for _, tt := range tests {
		err := status.Error(tt.grpcCode, tt.msg)

		custom := FromGRPC(err)

		assert.Equal(t, custom.GetCode(), tt.status)
		assert.Equal(t, custom.GetMessage(), tt.msg)
	}
}
