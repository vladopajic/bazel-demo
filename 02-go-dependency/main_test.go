package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	assert.Nil(t, t, "Sometimes it's okay to Fail")
}
