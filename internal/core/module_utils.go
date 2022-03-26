package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/robertsong9972/utkit/internal/config"
	"github.com/robertsong9972/utkit/internal/util"
)

func LoadOrCreateConf() {
	checkGitIgnore()
	m := make(map[string][]string)
	file, err := os.Open(config.ConfPath)
	outLines := getOutLines()
	c := make(map[string]*config.Package)
	if os.IsNotExist(err) {
		config.PackageMap = getPackMap(c, outLines)
		log.Printf("do not find %s,will scan all packages and generate one", config.ConfFileName)
		genConfFile()
		return
	}
	data, errData := ioutil.ReadAll(file)
	util.AssertError(errData)
	errData = json.Unmarshal(data, &m)
	util.AssertError(errData)
	for _, s := range m["packages"] {
		c[s] = &config.Package{}
	}
	config.PackageMap = getPackMap(c, outLines)
}

func GetPackageName(pack string) string {
	return strings.ReplaceAll(pack, config.ModuleName+"/", "")
}

func InitModuleName() {
	lines, err := ExecCommand(false, "head", "-1", config.GoModFileName)
	if err != nil {
		panic(errors.New("error when get go.mod"))
	}
	if len(lines) == 0 {
		return
	}
	arr := config.ModuleReg.FindStringSubmatch(lines[0])
	if len(arr) < 2 {
		panic("go.mod file is invalid")
	}
	config.ModuleName = arr[1]
}

func getPackMap(packMap map[string]*config.Package, lines []string) map[string]*config.Package {
	isEmpty := len(packMap) == 0
	for _, line := range lines {
		res := config.TestLineReg.FindStringSubmatch(line)
		if len(res) == 0 {
			continue
		}
		name := GetPackageName(res[1])
		if _, ok := packMap[name]; !ok && !isEmpty {
			continue
		}
		packageCov, err := strconv.ParseFloat(res[3], 64)
		util.AssertError(err)
		packagePath := strings.ReplaceAll(res[1], config.ModuleName, config.RootPath)
		packMap[name] = &config.Package{
			CoverageRate: packageCov,
			PackagePath:  packagePath,
		}
	}
	return packMap
}

func genConfFile() {
	m := make(map[string][]string)
	packages := make([]string, 0, len(config.PackageMap))
	for k := range config.PackageMap {
		packages = append(packages, k)
	}
	m["packages"] = packages
	_, err := os.Stat(config.ConfDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(config.ConfDir, 0777)
		util.AssertError(err)
	}
	file, err := os.Create(config.ConfPath)
	util.AssertError(err)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	buf, err := json.Marshal(m)
	util.AssertError(err)
	_, err = file.Write(buf)
	util.AssertError(err)
}

func getOutLines() []string {
	lines, err := ExecCommand(true, "/bin/bash", "-c", config.WeightCalShellScript)
	if err != nil {
		panic(errors.New("error when running go test to get weighted result"))
	}
	return lines
}

func checkGitIgnore() {
	_, err := os.Stat(".gitignore")
	if os.IsNotExist(err) {
		_, err = ExecCommand(false, "/bin/bash", "-c", config.MakeGitIgnoreFile)
		if err != nil {
			panic(errors.New("error when create .gitignore"))
		}
		return
	}
	lines, err := ExecCommand(false, "cat", ".gitignore")
	if err != nil {
		panic(errors.New("error when get .gitignore info"))
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
		panic(errors.New("error when create .gitignore"))
	}
}
