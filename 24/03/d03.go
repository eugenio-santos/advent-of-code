package main

import (
	"regexp"
	"strconv"
)

func d03p1(hay string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ms := re.FindAllStringSubmatch(hay, -1)

	sum := 0
	for _, m := range ms {
		m1, _ := strconv.Atoi(m[1])
		m2, _ := strconv.Atoi(m[2])
		sum += m1 * m2
	}

	return sum
}

func d03p2(hay string) int {
	re := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)
	ms := re.FindAllStringSubmatch(hay, -1)

	enalber := true

	sum := 0
	for _, m := range ms {
		if m[0] == "don't()" {
			enalber = false
		} else if m[0] == "do()" {
			enalber = true
		} else if enalber {
			m1, _ := strconv.Atoi(m[1])
			m2, _ := strconv.Atoi(m[2])
			sum += m1 * m2
		}
	}
	return sum
}
