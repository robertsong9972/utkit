package model

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robertsong9972/utkit/internal/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/robertsong9972/utkit/internal/config"
	"github.com/robertsong9972/utkit/internal/core"
)

type calculator struct {
}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c *calculator) PrintCovResult() {
	c.printIncrementCovRate()
	rate := c.totalWgCovRate()
	printPackages()
	printCoverageRate(rate)
}

func (c *calculator) printIncrementCovRate() {
	var err error
	_, err = core.ExecCommand(true, "/bin/bash", "-c", config.IncreasePrepareScript)
	if err != nil {
		panic(errors.New("error when download tools"))
	}
	_, err = core.ExecCommand(false, "/bin/bash", "-c", config.RunTest)
	if err != nil {
		panic(errors.New("error when running go test scripts"))
	}
	_, err = core.ExecCommand(true, "/bin/bash", "-c", config.IncreaseDiffCalScript)
	if err != nil {
		panic(errors.New("error when running diff-cover"))
	}
}

func (c *calculator) totalWgCovRate() float64 {
	sum := 0.0
	totalStat := 1e-11
	cvrPath := "./localfiles/cover.out"
	file, err := os.Open(cvrPath)
	if err != nil {
		log.Fatalf("failed to load cover.out, are you sure ./localfiles/cover.out exit?")
	}
	rd := bufio.NewReader(file)
	var lineIdx int
	for ; ; lineIdx++ {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if lineIdx == 0 {
			continue
		}
		statCnt, isCovered := c.cutCoverLine(line)
		totalStat += float64(statCnt)
		if isCovered {
			sum += float64(statCnt)
		}
	}
	return sum * 100 / totalStat
}

func (c *calculator) cutCoverLine(line string) (int64, bool) {
	stat := strings.Split(line, " ")
	statCnt, err := strconv.ParseInt(stat[1], 10, 64)
	if err != nil {
		log.Fatalf("error when parse line statement count,err=%v", err)
	}
	cvr, err := strconv.ParseInt(strings.Trim(stat[2], "\n"), 10, 64)
	if err != nil {
		log.Fatalf("error when parse line covered flag,err=%v", err)
	}
	return statCnt, cvr > 0
}

func printPackages() {
	fmt.Println("The packages involved in the statistics are as follows:")
	for k := range config.PackageMap {
		fmt.Printf("package:%s/%s\n", config.ModuleName, k)
	}
}

func printCoverageRate(rate float64) {
	fmt.Printf("The weighted average coverage: %.2f%% of statements\n", rate)
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
