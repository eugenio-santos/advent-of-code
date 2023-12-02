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

	datA := strings.Split(string(dat), "\n")
	// datAlen := len(datA)
	arr := make([]int, len(datA[0]))
	fmt.Println(arr)

	ox := getOx(datA, 0)
	c0 := getC0(datA, 0)
	fmt.Println(ox, c0)

	oxi, err := strconv.ParseInt(ox, 2, 64)
	c0i, err := strconv.ParseInt(c0, 2, 64)
	fmt.Println(oxi, c0i, oxi*c0i)

}

func getOx(dat []string, pos int) string {
	fmt.Println(dat, pos)
	if len(dat) == 1 {
		return dat[0]
	}
	res := make([]string, 0)
	sum := 0
	for _, d := range dat {
		if d[pos] == 49 {
			sum++
		}
	}

	mostComon := '0'
	fmt.Println("common", sum, float64(len(dat))/2.0, len(dat))
	if float64(sum) >= float64(len(dat))/2 {
		mostComon = '1'
	}

	for _, d := range dat {
		fmt.Println(d[pos], mostComon)
		if d[pos] == byte(mostComon) {
			res = append(res, d)
		}
	}
	return getOx(res, pos+1)
}

func getC0(dat []string, pos int) string {
	fmt.Println(dat, pos)
	if len(dat) == 1 {
		return dat[0]
	}
	res := make([]string, 0)
	sum := 0
	for _, d := range dat {
		if d[pos] == 49 {
			sum++
		}
	}

	leastsComon := '0'
	fmt.Println("common", sum, float64(len(dat))/2.0, len(dat))
	if float64(sum) < float64(len(dat))/2 {
		leastsComon = '1'
	}

	for _, d := range dat {
		fmt.Println(d[pos], leastsComon)
		if d[pos] == byte(leastsComon) {
			res = append(res, d)
		}
	}
	return getC0(res, pos+1)
}

// PART ONE
// for _, d := range datA {
// 	for i, c := range d {
// 		if c == 49 {
// 			arr[i]++
// 		}
// 	}
// }

// // get rates
// var gamma, epsilon string

// for _, a := range arr {
// 	if a > datAlen/2 {
// 		gamma += "1"
// 		epsilon += "0"
// 	} else {
// 		gamma += "0"
// 		epsilon += "1"
// 	}
// }
// fmt.Println(arr)
// fmt.Println(gamma, epsilon)
// gi, err := strconv.ParseInt(gamma, 2, 64)
// ep, err := strconv.ParseInt(epsilon, 2, 64)
// fmt.Println(gi, ep)
