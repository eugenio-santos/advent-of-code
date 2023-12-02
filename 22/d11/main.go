package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

type Operation func(int) int

type monkey struct {
	items []int
	op    Operation
	test  int
	t     int
	f     int
}

var nums [][]int
var itemCounter int
var monkeys []monkey

func main() {
	dat, _ := os.ReadFile("in")

	datA := strings.Split(string(dat), "\n")
	re := regexp.MustCompile("[0-9]+")

	monkeyCount := 0

	for _, l := range datA {
		if strings.HasPrefix(l, "Monkey") {
			monkeyCount++
		}
	}

	for i, l := range datA {
		if strings.HasPrefix(l, "Monkey") {
			m := monkey{}

			// items
			items := re.FindAllString(strings.Split(datA[i+1], ":")[1], -1)
			for _, item := range items {
				it, _ := strconv.Atoi(item)
				m.items = append(m.items, itemCounter)
				n := []int{}

				for mc := 0; mc < monkeyCount; mc++ {
					n = append(n, it)
				}

				nums = append(nums, n)
				itemCounter++
			}

			// operation
			o := strings.Split(strings.Split(datA[i+2], "=")[1], " ")
			if o[1] == "old" && o[3] == "old" {
				m.op = func(old int) int {
					return ops[o[2]](old, old)
				}
			} else if o[1] == "old" && o[3] != "old" {
				o2, _ := strconv.Atoi(o[3])
				m.op = func(old int) int {
					return ops[o[2]](old, o2)
				}
			} else if o[1] != "old" && o[3] == "old" {
				o1, _ := strconv.Atoi(o[1])
				m.op = func(old int) int {
					return ops[o[2]](o1, old)
				}
			} else {
				o1, _ := strconv.Atoi(o[1])
				o2, _ := strconv.Atoi(o[3])
				m.op = func(old int) int {
					return ops[o[2]](o1, o2)
				}
			}

			m.test, _ = strconv.Atoi(re.FindAllString(strings.Split(datA[i+3], ":")[1], -1)[0])
			m.t, _ = strconv.Atoi(re.FindAllString(strings.Split(datA[i+4], ":")[1], -1)[0])
			m.f, _ = strconv.Atoi(re.FindAllString(strings.Split(datA[i+5], ":")[1], -1)[0])
			monkeys = append(monkeys, m)
		}
	}
	fmt.Println("nums", nums)
	var res [8]int
	fmt.Println(monkeys)
	for round := 1; round <= 10000; round++ {
		for m := 0; m < len(monkeys); m++ {
			for len(monkeys[m].items) > 0 {
				res[m]++
				nums[monkeys[m].items[0]] = updateItem(nums[monkeys[m].items[0]], monkeys[m].op)

				if nums[monkeys[m].items[0]][m] == 0 {
					monkeys[monkeys[m].t].items = append(monkeys[monkeys[m].t].items, monkeys[m].items[0])
				} else {
					monkeys[monkeys[m].f].items = append(monkeys[monkeys[m].f].items, monkeys[m].items[0])

				}
				monkeys[m].items = monkeys[m].items[1:]
			}
		}
		if round%1000 == 0 {
			fmt.Println("ROUND", round)
			fmt.Println("res", res)
			for _, v := range monkeys {
				fmt.Println(v.items)
			}
		}
	}
	fmt.Println("ITEMS")
	for m := 0; m < len(monkeys); m++ {
		fmt.Println(monkeys[m].items)
	}

	fmt.Println()
	fmt.Println(res)
}

func updateItem(item []int, op Operation) []int {
	for i := 0; i < len(monkeys); i++ {
		item[i] = op(item[i])
		item[i] = item[i] % monkeys[i].test
	}
	return item
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	s[len(s)-1] = 0
	return s[:len(s)-1]
}
