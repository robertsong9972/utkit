package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/robertsong9972/utkit/internal/config"
)

func TestGetPackageName(t *testing.T) {
	config.ModuleName = "utcal"
	assert.Equal(t, "internal/core", GetPackageName("utcal/internal/core"))
}

func Test_getPackMap(t *testing.T) {
	config.ModuleName = "utcal"
	lines := []string{
		"?       utcal   [no test files]",
		"?       utcal/internal/config   [no test files]",
		"ok      utcal/internal/core     0.567s  coverage: 44.6% of statements",
		"?       utcal/internal/model    [no test files]",
		"?       utcal/internal/util     [no test files]",
	}
	arrMap := map[string]*config.Package{
		"internal/core": {
			CoverageRate: 44.6,
			PackagePath:  fmt.Sprintf("%s/%s", config.RootPath, "internal/core"),
		},
	}
	m := getPackMap(make(map[string]*config.Package), lines)
	assert.Equal(t, arrMap, m)
	m = getPackMap(map[string]*config.Package{
		"internal/core": {},
	}, lines)
	assert.Equal(t, arrMap, m)
}
