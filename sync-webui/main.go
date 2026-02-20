package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	goWebuiRepo = "github.com/webui-dev/go-webui/v2"
	webuiRepo   = "github.com/webui-dev/webui"
)

func goWebuiVersion(version string) string {
	return fmt.Sprintf("%s@%s", goWebuiRepo, version)
}

func webuiVersion(version string) string {
	return fmt.Sprintf("%s@%s", webuiRepo, version)
}

func main() {
	// Ensure the current directory is part of a go module.
	if !fileExists("go.mod") {
		errExit(`error: failed to find go.mod file in current directory.
       To set up the go-webui module, use this script in a module directory.`)
	}
	goTools := GetGo()
	// Run go commands
	iferrExit(goTools.Cmd("mod", "tidy").Run())
	iferrExit(goTools.Cmd("get", goWebuiVersion("main")).Run())
	iferrExit(goTools.CmdSilent("get", webuiVersion("main")).Run())

	// Retrieve GOPATH (use environment variable if defined, otherwise go env)
	goPath, err := goTools.Gopath()
	iferrExit(err)

	// Parse the first matching version lines from go.sum
	// For go-webui/v2:
	goWebuiFullVersion := field(headN1(grep(goWebuiRepo, "go.sum")), 2)
	// For webui:
	webuiFullVersion := field(headN1(grep(webuiRepo, "go.sum")), 2)

	// Construct paths based on the parsed versions
	goWebuiPath := filepath.Join(goPath, "pkg", "mod", goWebuiVersion(goWebuiFullVersion))
	webuiPath := filepath.Join(goPath, "pkg", "mod", webuiVersion(webuiFullVersion))

	// Validate that these paths actually exist
	iferrExit(validatePaths(goWebuiPath, webuiPath))
	linkName := filepath.Join(goWebuiPath, "webui")
	// Exit if link already exists in the directory of the used go-webui version.
	if !dirExists(linkName) {
		return
	}

	// Store original permissions.
	// Not strictly necessary, yet we ensure end without changes to the original permissions.
	ogPerms, err := getPerms(goWebuiPath)
	iferrExit(err)
	iferrExit(os.Chmod(goWebuiPath, 0200+ogPerms))
	// Linking allows using WebUI C even in cases of multiple go-webui versions without creating bloat.
	if err := os.Symlink(webuiPath, linkName); err != nil {
		if runtime.GOOS == "windows" {
			// (Requires Administrator privileges or Developer Mode)
			errExit("mklink failed. Run as Administrator if needed.")
		}
		errExit(err.Error())
	}
	// Restore original permissions.
	iferrExit(os.Chmod(goWebuiPath, ogPerms))
	iferrExit(goTools.Cmd("mod", "tidy").Run())
}

func validatePaths(goWebuiPath, webuiPath string) error {
	var errs []error
	if !dirExists(goWebuiPath) {
		errs = append(errs, fmt.Errorf("failed to find go-webui in '%s'", goWebuiPath))
	}
	if !dirExists(webuiPath) {
		errs = append(errs, fmt.Errorf("failed to find webui in '%s'", webuiPath))
	}
	return errors.Join(errs...)
}

func iferrExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func errExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
