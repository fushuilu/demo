package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Name(t *testing.T) {
	assert.Equal(t, Name(), "demo")
}
