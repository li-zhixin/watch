package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGINT,
	)
	loc, _ := time.LoadLocation("UTC")
	fmt.Println("start time: " + time.Now().In(loc).Format(time.RFC3339))
	argsWithoutProg := os.Args[1:]
	actualArgsWithoutProg := argsWithoutProg[1:]
	cmd := exec.Command(argsWithoutProg[0], actualArgsWithoutProg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Start()
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				_ = cmd.Process.Kill()
			default:
				fmt.Println("Unknown signal.")
			}
		}
	}()
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-done:
		fmt.Println("end time: " + time.Now().In(loc).Format(time.RFC3339))
	}
}
