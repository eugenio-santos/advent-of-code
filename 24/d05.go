package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func d05p1(file string) (int, map[int][]int, [][]int) {
	prio, pages := getData(file)
	res := 0
	unsafe := [][]int{}
a:
	for _, page := range pages {
		for i := len(page) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				// if something is out of order go to the next page
				if slices.Contains(prio[page[i]], page[j]) {
					unsafe = append(unsafe, page)
					continue a
				}
			}
		}

		res += page[int(math.Floor(float64(len(page))/2.0))]
	}

	return res, prio, unsafe
}

func d05p2(file string) int {
	_, prio, pages := d05p1(file)
	res := 0

	for _, page := range pages {
	a:
		for i := len(page) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				// if something is out of order swap places
				if slices.Contains(prio[page[i]], page[j]) {
					aux := page[i]
					page[i] = page[j]
					page[j] = aux
					i++
					continue a
				}
			}
		}
		res += page[int(math.Floor(float64(len(page))/2.0))]
	}

	return res
}

func getData(file string) (map[int][]int, [][]int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer f.Close()

	prio := map[int][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		p := strings.Split(line, "|")
		i, _ := strconv.Atoi(p[0])
		j, _ := strconv.Atoi(p[1])
		prio[i] = append(prio[i], j)
	}

	pages := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		Spage := strings.Split(line, ",")

		page := []int{}
		for _, p := range Spage {
			pi, _ := strconv.Atoi(p)
			page = append(page, pi)
		}

		pages = append(pages, page)
	}

	return prio, pages
}
