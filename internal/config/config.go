package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	GoModFileName = "go.mod"
	ConfDir       = "testdata"
	ConfFileName  = "testdata/ut_package.json"
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
)

func InitConf() {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	RootPath = rootPath
	ConfPath = ConfFileName
	loadModuleName()
}

func loadModuleName() {
	file, err := os.Open(GoModFileName)
	if err != nil {
		log.Fatal("error when open go.mod, are you sure it exits?")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	rd := bufio.NewReader(file)
	line, err := rd.ReadString('\n')
	line = strings.Trim(line, "\n")
	ModuleName = strings.Split(line, " ")[1]
}
