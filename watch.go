package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	fmt.Println("start time: "+time.Now().In(loc).Format(time.RFC3339))
	argsWithoutProg := os.Args[1:]
	actualArgsWithoutProg := argsWithoutProg[1:]
	cmd := exec.Command(argsWithoutProg[0], actualArgsWithoutProg...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	rd := bufio.NewReader(stdout)
	if err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		if runtime.GOOS == "windows" {
			str = str[:len(str)-1]
		}
		fmt.Println(str)
	}
	fmt.Println("end time: "+time.Now().In(loc).Format(time.RFC3339))
}
