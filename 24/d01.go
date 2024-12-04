package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func d01p1() int {
	file, err := os.Open("inputs/d01")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		lr := strings.Split(line, "   ")
		lv, _ := strconv.Atoi(lr[0])
		rv, _ := strconv.Atoi(lr[1])

		left = sortedInsert(lv, left)
		right = sortedInsert(rv, right)
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += absDiff(left[i], right[i])
	}

	return sum
}

func sortedInsert(v int, s []int) []int {
	i, _ := sort.Find(len(s), func(i int) int {
		return v - s[i]
	})

	return append(s[:i], append([]int{v}, s[i:]...)...)
}

func absDiff(x int, y int) int {
	s := x - y

	if s < 0 {
		return -s
	}

	return s
}

func d01p2() int {
	file, err := os.Open("inputs/d01")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := make(map[int]int)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		lr := strings.Split(line, "   ")
		lv, _ := strconv.Atoi(lr[0])
		rv, _ := strconv.Atoi(lr[1])

		left = append(left, lv)
		right[rv] = right[rv] + 1
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += left[i] * right[left[i]]
	}

	return sum
}
