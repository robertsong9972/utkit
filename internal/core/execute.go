package core

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func ExecCommand(shouldPrint bool, commandName string, params ...string) ([]string, error) {
	lines := make([]string, 0, 1024)
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		lines = append(lines, line)
	}
	if shouldPrint {
		for _, line := range lines {
			fmt.Print(line)
		}
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
