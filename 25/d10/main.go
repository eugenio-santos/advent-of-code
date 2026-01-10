package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var fn string

func main() {
	if os.Args[1] == "test" {
		fn = "test"
	} else {
		fn = "input"
	}

	lines := lines(fn)
	sum := 0
	for _, l := range lines {
		fmt.Println(l)
		g, buttons := parseLine(l)
		allButXor := 0
		for _, b := range buttons {
			allButXor = allButXor ^ b
		}
		fmt.Println(g, buttons, allButXor)

		n := len(buttons)
		var i uint
		for opCount := 1; opCount < n; opCount++ {
			for i = 1; i < (1 << n); i++ {
				// fmt.Println(s, i)
				var combo []int
				if bits.OnesCount(i) == opCount {
					combo = []int{}
					for j := range n {
						// fmt.Println(j, i, j>>i, (j>>i)&1)
						if ((i >> j) & 1) == 1 {
							combo = append(combo, buttons[j])
						}
					}
					// fmt.Println(combo)
					// LOGIC HAPPENS HERE
					r := 0
					br := allButXor
					for _, c := range combo {
						r = r ^ c
						br = br ^ c
					}
					if r == g {
						sum += opCount
						goto nl
					}
				}
			}
		}
	nl:
	}

	fmt.Println("res: ", sum)
}

func parseLine(l string) (int, []int) {
	return getGoalR(l), getButtons(l)
}

func getButtons(l string) []int {
	lb := strings.IndexRune(l, ']') + 2
	b := strings.Split(l[lb:], " ")
	buttonsS := b[:len(b)-1]

	buttons := []int{}
	for _, but := range buttonsS {
		byts := []byte("0000000000000000000000000000000000000000000000000000000000000000")
		for s := range strings.SplitSeq(but[1:len(but)-1], ",") {
			i, _ := strconv.Atoi(s)
			byts[63-i] = '1'
		}
		nb, _ := strconv.ParseInt(string(byts), 2, 64)
		buttons = append(buttons, int(nb))
	}

	return buttons
}

func getGoal(l string) int {
	byts := ""

	for i := strings.IndexRune(l, ']') - 1; i > 0; i-- {
		if l[i] == '.' {
			byts = "0" + byts
		} else {
			byts = "1" + byts
		}
	}

	fmt.Println(byts)
	r, _ := strconv.ParseInt(byts, 2, 64)

	return int(r)
}

func getGoalR(l string) int {
	byts := ""

	for i := strings.IndexRune(l, ']') - 1; i > 0; i-- {
		if l[i] == '.' {
			byts += "0"
		} else {
			byts += "1"
		}
	}
	// fmt.Println(byts)
	r, _ := strconv.ParseInt(byts, 2, 64)

	return int(r)
}

func lines(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
