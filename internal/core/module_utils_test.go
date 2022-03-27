package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/robertsong9972/utkit/internal/config"
)

func TestGetPackageName(t *testing.T) {
	config.ModuleName = "github.com/robertsong9972/utkit"
	assert.Equal(t, "internal/core",
		GetPackageName("github.com/robertsong9972/utkit/internal/core"))
}
