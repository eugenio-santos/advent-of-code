package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	datA := strings.Split(string(dat), "\n")
	max := make([]int, 4)
	for _, el := range datA {
		if el == "" {
			sort.Ints(max)
			max[0] = 0
		} else {
			it, err := strconv.Atoi(el)
			if err != nil {
				panic(err)
			}
			max[0] += it
		}
		fmt.Println(max)
	}
	fmt.Println(max[1] + max[2] + max[3])
}
