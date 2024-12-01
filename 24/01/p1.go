package main

import (
	"os"
	"strings"
)

func main() {

}

func lines() []string {
	f, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}
