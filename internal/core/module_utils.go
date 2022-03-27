package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/robertsong9972/utkit/internal/config"
)

func SysInit() {
	processGitIgnore()
	processConf()
}

func GetPackageName(pack string) string {
	return strings.ReplaceAll(pack, config.ModuleName+"/", "")
}

func processGitIgnore() {
	_, err := os.Stat(".gitignore")
	if os.IsNotExist(err) {
		_, err = ExecCommand(false, "/bin/bash", "-c", config.MakeGitIgnoreFile)
		if err != nil {
			log.Fatalf("error when create .gitignore, error=%v", err)
		}
		return
	}
	lines, err := ExecCommand(false, "cat", ".gitignore")
	if err != nil {
		log.Fatalf("error when get .gitignore info, error=%v", err)
	}
	reg := regexp.MustCompile("localfiles/.*")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		if reg.MatchString(line) {
			return
		}
	}
	_, err = ExecCommand(false, "/bin/bash", "-c", config.MakeGitIgnoreFile)
	if err != nil {
		log.Fatalf("error when add lines to .gitignore, error=%v", err)
	}
}

func processConf() {
	_, err := os.Stat(config.ConfPath)
	if os.IsNotExist(err) {
		return
	}
	file, err := os.Open(config.ConfPath)
	if err != nil {
		log.Fatalf("error when open testdata/ut_package.json, error=%v", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error when read testdata/ut_package.json, error=%v", err)
	}
	cf := make(map[string][]string)
	pkgMap := make(map[string]*config.Package)
	err = json.Unmarshal(data, &cf)
	if err != nil {
		log.Fatalf("error when unmarshal testdata/ut_package.json, error=%v", err)
	}
	for _, s := range cf["packages"] {
		pkg := fmt.Sprintf("%s/%s", config.ModuleName, s)
		pkgMap[pkg] = &config.Package{}
	}
	config.PackageMap = pkgMap
}
