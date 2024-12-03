package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		// fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	safe_lvls := 0
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		Slvls := strings.Split(line, " ")

		lvls := []int{}
		for _, l := range Slvls {
			il, _ := strconv.Atoi(l)
			lvls = append(lvls, il)
		}

		r := evalL(lvls)

		// fmt.Println("L", lvls, "R", r)
		if len(r) == 0 {
			// fmt.Println(line, r)
			safe_lvls++
		} else {
			fmt.Println("L", lvls, "R", r)
			for k := range r {
				aux := k
				subS := make([]int, len(lvls))
				copy(subS, lvls)
				subS = append(subS[:aux], subS[aux+1:]...)
				fmt.Println("T", subS)
				if isSafe(subS) {
					fmt.Println("S", subS)
					safe_lvls++
					break
				}
			}
			fmt.Println()
			// fmt.Println(lvls)
		}
	}

	fmt.Println("coito", safe_lvls)
	// fmt.Println("lol", isSafe([]int{1, 2, 4, 5}))
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func evalL(lvls []int) map[int]int {
	res := make(map[int]int)
	prevDiff := lvls[0] - lvls[1]

	if !(abs(prevDiff) >= 1 && abs(prevDiff) <= 3) {
		// fmt.Println("adding to map 0 1")
		res[0] = 1
		res[1] = 1
	}
	if prevDiff*(lvls[1]-lvls[2]) < 0 {
		res[0] = 1
		res[1] = 1
	}

	for i, lvl := range lvls[2:] {
		// fmt.Println("lvl", lvl)
		diff := lvls[i+1] - lvl
		if diff*prevDiff < 0 || !(abs(diff) >= 1 && abs(diff) <= 3) {
			// fmt.Println("adding to map", i+1, i+2)
			res[i+1] = 1
			res[i+2] = 1
		}

		prevDiff = diff
	}

	return res
}

func isSafe(lvls []int) bool {
	prev := lvls[1]
	prevDiff := lvls[0] - lvls[1]

	if !(abs(prevDiff) >= 1 && abs(prevDiff) <= 3) {
		return false
	}

	for _, lvl := range lvls[2:] {
		diff := prev - lvl
		if diff*prevDiff < 0 || !(abs(diff) >= 1 && abs(diff) <= 3) {
			return false
		}

		prev = lvl
		prevDiff = diff
	}

	return true
}
