package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("int")
	if err != nil {
		panic(err)
	}

	datA := strings.Split(string(dat), "\n")
	f := 0
	d := 0
	aim := 0
	for _, move := range datA {
		m, a := getM(move)

		switch m {
		case "forward":
			f += a
			d += aim * a
		case "down":
			aim += a
		case "up":
			aim -= a
		}
	}

	fmt.Println(f * d)
}

func getM(move string) (string, int) {
	m := strings.Split(move, " ")
	a, _ := strconv.Atoi(m[1])
	return m[0], a
}
