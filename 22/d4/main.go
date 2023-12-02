package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world!")

	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(dat), "\n")
	sum := 0
	for _, pair := range pairs {
		fmt.Println(pair)
		p1min, p1max := getMinMax(strings.Split(string(pair), ",")[0])
		p2min, p2max := getMinMax(strings.Split(string(pair), ",")[1])
		fmt.Println(p1min, p1max, p2min, p2max)

		if (p2min >= p1min && p2min <= p1max) || (p2max >= p1min && p2max <= p1max) || (p1min >= p2min && p1min <= p2max) || (p1max >= p2min && p1max <= p2max) {
			fmt.Println("over")
			sum += 1
		}
	}
	fmt.Println(sum)
}

func getMinMax(p string) (int, int) {
	min, _ := strconv.Atoi(strings.Split(p, "-")[0])
	max, _ := strconv.Atoi(strings.Split(p, "-")[1])
	return min, max
}
