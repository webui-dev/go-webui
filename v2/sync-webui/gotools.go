package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Gotools struct {
	exe string
}

func GetGo() *Gotools {
	exe, err := exec.LookPath("go")
	if err != nil {
		panic(fmt.Errorf("failed to find go executable on PATH: %w", err))
	}
	return &Gotools{
		exe: exe,
	}
}

func (gt *Gotools) Cmd(args ...string) *exec.Cmd {
	cmd := exec.Command(gt.exe, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func (gt *Gotools) CmdSilent(args ...string) *exec.Cmd {
	cmd := gt.Cmd(args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd
}

func (gt *Gotools) Gopath() (string, error) {
	gopath, ok := os.LookupEnv("GOPATH")
	if ok {
		return gopath, nil
	}
	output, err := gt.Cmd("env", "GOPATH").Output()
	if err != nil {
		return "", fmt.Errorf("failed to check GOPATH from go tools: %w", err)
	}
	gopath = strings.TrimSpace(string(output))
	return gopath, nil
}
