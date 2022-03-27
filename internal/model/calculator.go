package model

import (
	"bufio"
	"encoding/json"
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
	pkgMap map[string]*config.Package
}

func NewCalculator() *calculator {
	return &calculator{
		pkgMap: make(map[string]*config.Package),
	}
}

func (c *calculator) PrintCovResult() {
	c.printIncrementCovRate()
	rate := c.totalWgCovRate()
	c.printPackages()
	printCoverageRate(rate)
}

func (c *calculator) printIncrementCovRate() {
	var err error
	_, err = core.ExecCommand(false, "/bin/bash", "-c", config.IncreasePrepareScript)
	if err != nil {
		log.Fatalf("error when download tool,error=%v", err)
	}
	_, err = core.ExecCommand(true, "/bin/bash", "-c", config.RunTest)
	if err != nil {
		log.Fatalf("error when running go test, your go test has failed, please check again")
	}
	_, err = core.ExecCommand(true, "/bin/bash", "-c", config.GenerateXml)
	if err != nil {
		log.Fatalf("error when transfer cover.out to xml, gocov or gocov-xml is missing or wrong")
	}
	_, err = core.ExecCommand(true, "/bin/bash", "-c", config.IncreaseDiffCalScript)
	if err != nil {
		log.Fatalf("error when running diff-cover, diff-cover missed or brach missed")
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
		pkg, statCnt, isCovered := c.cutCoverLine(line)
		totalStat += float64(statCnt)
		if _, ok := c.pkgMap[pkg]; !ok {
			c.pkgMap[pkg] = &config.Package{}
		}
		c.pkgMap[pkg].StatementsCnt += statCnt
		if isCovered {
			sum += float64(statCnt)
			c.pkgMap[pkg].CoveredStatementCnt += statCnt
		}
	}
	return sum * 100 / totalStat
}

func (c *calculator) cutCoverLine(line string) (string, int64, bool) {
	stat := strings.Split(line, " ")
	pkg := stat[0][:strings.LastIndex(stat[0], "/")]
	statCnt, err := strconv.ParseInt(stat[1], 10, 64)
	if err != nil {
		log.Fatalf("error when parse line statement count,err=%v", err)
	}
	cvr, err := strconv.ParseInt(strings.Trim(stat[2], "\n"), 10, 64)
	if err != nil {
		log.Fatalf("error when parse line covered flag,err=%v", err)
	}
	return pkg, statCnt, cvr > 0
}

func (c *calculator) printPackages() {
	if len(config.PackageMap) == 0 {
		config.PackageMap = c.pkgMap
		genConfFile()
	}
	var (
		pkg *config.Package
		ok  bool
	)
	sum := 0.0
	totalStat := 1e-11
	fmt.Println("----------------------------")
	fmt.Println("Packages involved in config:")
	for k := range config.PackageMap {
		if pkg, ok = c.pkgMap[k]; ok {
			sum += float64(pkg.CoveredStatementCnt * 100)
			totalStat += float64(pkg.StatementsCnt)
			fmt.Printf("    %s\n", k)
			continue
		}
		fmt.Printf("%s does not exits, skip this package\n", k)
	}
	fmt.Printf("verage coverage: %.1f%%\n", sum/totalStat)
	fmt.Println("----------------------------")
}

func printCoverageRate(rate float64) {
	fmt.Printf("The weighted average coverage: %.1f%% of statements\n", rate)
}

func genConfFile() {
	m := make(map[string][]string)
	packages := make([]string, 0, len(config.PackageMap))
	for k := range config.PackageMap {
		packages = append(packages, core.GetPackageName(k))
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
