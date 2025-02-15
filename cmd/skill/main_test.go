package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	assert.IsType(t, NewApp(), &App{})
}
