package main

import (
	"bufio"
	"os"
	"strings"
)

func fileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func grep(text, path string) []string {
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, text) {
			lines = append(lines, line)
		}
	}
	return lines
}

func headN1(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return lines[0]
}

func field(line string, field int) string {
	line = strings.TrimSpace(line)
	fields := strings.Split(line, " ")
	if len(fields) < field {
		return ""
	}
	return fields[field-1]
}

func getPerms(path string) (os.FileMode, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}
