package model

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/robertsong9972/utkit/internal/core"
	"io"
	"os"
	"strings"

	"github.com/robertsong9972/utkit/internal/config"
	"github.com/robertsong9972/utkit/internal/util"
)

type calculator struct {
}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c *calculator) PrintCovResult() {
	c.printIncrementCovRate()
	rate := c.getWeightedCovRate()
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

func (c *calculator) getWeightedCovRate() float64 {
	sum := 0.0
	lineCount := 1e-11
	for _, v := range config.PackageMap {
		packLineCount := c.getPackageLines(v.PackagePath)
		sum += v.CoverageRate * packLineCount
		lineCount += packLineCount
	}
	return sum / lineCount
}

func (c *calculator) getPackageLines(path string) float64 {
	lineCount := 1e-11
	files, err := os.ReadDir(path)
	util.AssertError(err)
	return lineCount + c.traversePackage(files, path)
}

func (c *calculator) traversePackage(files []os.DirEntry, packagePath string) float64 {
	count := 0.0
	for _, file := range files {
		if config.TestFileReg.MatchString(file.Name()) || file.IsDir() {
			continue
		}
		goFile, err := os.Open(fmt.Sprintf("%s/%s", packagePath, file.Name()))
		util.AssertError(err)
		rd := bufio.NewReader(goFile)
		count += c.traverseFile(rd)
	}
	return count
}

func (c *calculator) traverseFile(reader *bufio.Reader) float64 {
	count := 0.0
	start := false
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if !start {
			start = config.FunctionReg.MatchString(line)
			continue
		}
		line = strings.Trim(line, " ")
		if line == "\n" || config.EmptyReg.MatchString(line) {
			continue
		}
		count++
	}
	if start {
		count++
	}
	return count
}

func printPackages() {
	fmt.Println("The packages involved in the statistics are as follows:")
	for k := range config.PackageMap {
		fmt.Printf("package:%s/%s\n", config.ModuleName, k)
	}
}
func printCoverageRate(rate float64) {
	fmt.Printf("The weighted average coverage: %.1f%% of statements\n", rate)
}
