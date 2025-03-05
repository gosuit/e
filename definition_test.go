package e

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/gosuit/lec"
	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	msg := "Error"
	testErr := errors.New("some error")
	code := Internal
	err := New(msg, code, testErr)

	assert.Equal(t, msg, err.GetMessage())
}

func TestGetError(t *testing.T) {
	msg := "error"
	testErr1 := errors.New("some error 1")
	testErr2 := errors.New("some error 2")
	code := Internal

	joinedErr := errors.Join(testErr1, testErr2)
	err := New(msg, code, testErr1, testErr2)

	assert.Equal(t, joinedErr, err.GetError())
}

func TestGetTag(t *testing.T) {
	msg := "error"
	code := Internal
	err := New(msg, code)

	key := "key"
	value := "value"

	err = err.WithTag(key, value)

	assert.Equal(t, value, err.GetTag(key))
}

func TestGetCode(t *testing.T) {
	msg := "error"
	code := Internal
	err := New(msg, code)

	assert.Equal(t, code, err.GetCode())
}

func TestWithMessage(t *testing.T) {
	initialMsg := "Initial error"
	testErr := errors.New("some error")
	code := Internal
	err := New(initialMsg, code, testErr)

	newMsg := "Updated error"
	err = err.WithMessage(newMsg)

	if err.GetMessage() != newMsg {
		t.Errorf("Expected message %q, got %q", newMsg, err.GetMessage())
	}
}

func TestWithErr(t *testing.T) {
	msg := "error"
	code := Internal
	err := New(msg, code)

	testErr := errors.New("some error")
	joined := errors.Join(testErr)

	err = err.WithErr(testErr)

	assert.Equal(t, joined, err.GetError())
}

func TestWithTag(t *testing.T) {
	msg := "error"
	code := Internal
	err := New(msg, code)

	key := "key"
	value := "value"

	err = err.WithTag(key, value)

	assert.Equal(t, value, err.GetTag(key))
}

func TestWithCtx(t *testing.T) {
	c := lec.New(slog.Default())

	key := "key"
	value := "value"

	c.AddValue(key, value, true)

	msg := "error"
	code := Internal

	err := New(msg, code)

	err = err.WithCtx(c)

	assert.Equal(t, value, err.GetTag(key))
}

func TestWithCode(t *testing.T) {
	msg := "Some msg"
	testErr := errors.New("Some error")
	initialCode := Forbidden
	err := New(msg, initialCode, testErr)

	newCode := Internal
	err = err.WithCode(newCode)

	if err.GetCode() != newCode {
		t.Errorf("Expected code %v, got %v", newCode, err.GetCode())
	}
}
