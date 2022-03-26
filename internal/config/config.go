package config

import (
	"fmt"
	"os"
	"regexp"
)

const (
	GoModFileName = "go.mod"
	ConfDir      = "testdata"
	ConfFileName = "testdata/ut_package.json"
)

var (
	// ConfPath is the path of the package json file which concludes
	// the packages needed to be scanned
	ConfPath string
	// RootPath is the root path of your project
	RootPath string
	// ModuleName is your project module name
	ModuleName string
	// PackageMap is to store the packages which need to be scanned
	PackageMap map[string]*Package
	// TestLineReg is a regexp to match available go test output file lines
	TestLineReg = regexp.MustCompile("ok\\s*(.*?)[\\s\\t](.*?)[\\s\\t]coverage:[\\s\\t](.*?)%")
	// TestFileReg is a regexp to match go test file such as xx_test.go
	TestFileReg = regexp.MustCompile("^(.*?)_test\\.go")
	// FunctionReg is a regexp to find the first effective line of go file
	FunctionReg = regexp.MustCompile("func\\s*(.*?)\\(.*?\\)")
	// ModuleReg is a regexp to get your project module name
	ModuleReg = regexp.MustCompile("^module\\s*\\t*(.*?)\\n")
	// EmptyReg is a regexp to match comment lines
	EmptyReg = regexp.MustCompile("//.*")
)

func InitConfig() {
	initRootPath()
	initConfPath()
}

func initRootPath() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	RootPath = rootPath
}

func initConfPath() {
	ConfPath = fmt.Sprintf("%s/%s", RootPath, ConfFileName)
}
