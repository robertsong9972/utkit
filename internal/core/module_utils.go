package core

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/robertsong9972/utkit/internal/config"
)

func GetPackageName(pack string) string {
	return strings.ReplaceAll(pack, config.ModuleName+"/", "")
}

func CheckGitIgnore() {
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
