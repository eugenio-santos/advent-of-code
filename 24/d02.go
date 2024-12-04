package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d02p1() int {
	file, err := os.Open("inputs/d02")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safe_lvls := 0
LineLoop:
	for scanner.Scan() {
		line := scanner.Text()
		Slvls := strings.Split(line, " ")

		lvls := []int{}
		for _, l := range Slvls {
			il, _ := strconv.Atoi(l)
			lvls = append(lvls, il)
		}

		prev := lvls[1]
		prevDiff := lvls[0] - lvls[1]

		if !(abs(prevDiff) >= 1 && abs(prevDiff) <= 3) {
			continue
		}

		for _, lvl := range lvls[2:] {
			diff := prev - lvl

			if diff*prevDiff < 0 {
				continue LineLoop
			}
			if !(abs(diff) >= 1 && abs(diff) <= 3) {
				continue LineLoop
			}

			prev = lvl
			prevDiff = diff
		}

		safe_lvls++
	}

	return safe_lvls
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func d02p2() int {
	file, err := os.Open("inputs/d02")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	safe_lvls := 0
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		Slvls := strings.Split(line, " ")

		lvls := []int{}
		for _, l := range Slvls {
			il, _ := strconv.Atoi(l)
			lvls = append(lvls, il)
		}

		r := evalL(lvls)

		if len(r) == 0 {
			safe_lvls++
		} else {
			for k := range r {
				aux := k
				subS := make([]int, len(lvls))
				copy(subS, lvls)
				subS = append(subS[:aux], subS[aux+1:]...)
				if isSafe(subS) {
					safe_lvls++
					break
				}
			}
		}
	}

	return safe_lvls
}

func evalL(lvls []int) map[int]int {
	res := make(map[int]int)
	prevDiff := lvls[0] - lvls[1]

	if !(abs(prevDiff) >= 1 && abs(prevDiff) <= 3) {
		res[0] = 1
		res[1] = 1
	}
	if prevDiff*(lvls[1]-lvls[2]) < 0 {
		res[0] = 1
		res[1] = 1
	}

	for i, lvl := range lvls[2:] {
		diff := lvls[i+1] - lvl
		if diff*prevDiff < 0 || !(abs(diff) >= 1 && abs(diff) <= 3) {
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
