package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	mes := strings.Split(string(dat), "\n")
	sum := 0
	for i := 3; i < len(mes); i++ {
		prev := sumMes(mes[i-3], mes[i-2], mes[i-1])
		curr := sumMes(mes[i-2], mes[i-1], mes[i])
		if curr > prev {
			sum++
		}

	}
	fmt.Println(sum)
}

func sumMes(a, b, c string) int {
	m1, _ := strconv.Atoi(a)
	m2, _ := strconv.Atoi(b)
	m3, _ := strconv.Atoi(c)
	return m1 + m2 + m3
}
func getInts(c string, p string) (int, int) {
	curr, _ := strconv.Atoi(c)
	prev, _ := strconv.Atoi(p)
	return curr, prev
}
