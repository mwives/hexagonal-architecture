package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "test message"
	expected := []byte(`{"message":"test message"}`)
	result := jsonError(msg)

	assert.Equal(t, expected, result)
}
