package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("d01")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := make(map[int]int)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		lr := strings.Split(line, "   ")
		lv, _ := strconv.Atoi(lr[0])
		rv, _ := strconv.Atoi(lr[1])

		left = append(left, lv)
		right[rv] = right[rv] + 1
	}

	fmt.Println(left)
	fmt.Println(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += left[i] * right[left[i]]
	}

	println(sum)
}
