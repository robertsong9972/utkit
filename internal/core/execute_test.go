package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecCommand(t *testing.T) {
	res, err := ExecCommand(false, "ls", "-l")
	assert.Greater(t, len(res), 0)
	assert.Nil(t, err)
	res, err = ExecCommand(true, "ls", "-l")
	assert.Greater(t, len(res), 0)
	assert.Nil(t, err)
}
