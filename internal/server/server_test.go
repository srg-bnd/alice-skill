package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	assert.IsType(t, NewServer(), &Server{})
}

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestConsts(t *testing.T) {
	assert.Equal(t, defaultHost, ":8080")
}
