package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	datA := strings.Split(string(dat), "\n")
	res := []string{}

	for i := 0; i < len(datA[0]); i++ {
		res = append(res, string(datA[0][len(datA[0])-i-1]))
	}
	fmt.Println(datA[0], res)
	for i := 0; i < 100; i++ {
		res = append(res, "0")
	}

	r := 0
	carryOver := 0
	last := 0
	for _, num := range datA[1:] {
		// fmt.Println(num)
		for i := 0; i < len(num); i++ {
			last = i
			s := SnToInt(string(num[len(num)-i-1]))

			r, carryOver = sum(SnToInt(res[i]), s+carryOver)
			res[i] = IntToSn(r)
		}
		for i := last + 1; carryOver != 0; i++ {
			r, carryOver = sum(SnToInt(res[i]), +carryOver)
			res[i] = IntToSn(r)
		}
		// printSol(res)
	}
	printSol(res)

}

func printSol(res []string) {
	sol := ""
	for _, v := range res {
		sol = v + sol
	}
	fmt.Println(sol)
}

func SnToInt(s string) int {
	switch s {
	case "2":
		return 2
	case "1":
		return 1
	case "0":
		return 0
	case "-":
		return -1
	case "=":
		return -2
	}
	return 0
}

func IntToSn(s int) string {
	switch s {
	case 2:
		return "2"
	case 1:
		return "1"
	case 0:
		return "0"
	case -1:
		return "-"
	case -2:
		return "="
	}
	return "0"
}

func sum(i, j int) (int, int) {
	s := i + j
	carryOver := 0
	if s > 2 { // positive overflow
		carryOver = 1
		s = s - 5
	} else if s < -2 { // negative overflow
		carryOver = -1
		s = s + 5
	}
	return s, carryOver
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
