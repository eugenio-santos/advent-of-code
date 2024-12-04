package main

import (
	"os"
	"strings"
)

func main() {
	_ = d03p1("")
}

func lines(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
