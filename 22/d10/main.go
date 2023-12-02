package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var datA []string

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	datA = strings.Split(string(dat), "\n")

	cycle := 1
	pc := 0
	inst, arg := getInst(pc)
	X := 1
	var crt []string
	crtRow := 0
	crt = append(crt, "")
	for {

		fmt.Print("cycle ", cycle, (cycle)%40, " X: ", X, " pc: ", pc, " inst: ", inst, " arg: ", arg, " ")
		if inst == "q" {
			break
		}

		if ((cycle)%40)-1 >= X-1 && ((cycle)%40)-1 <= X+1 {
			crt[crtRow] += "#"
		} else {
			crt[crtRow] += "."
		}

		if (cycle)%40 == 0 {
			crtRow++
			crt = append(crt, "")
		}

		if inst == "f" {
			inst = "s"
		} else {
			if inst == "n" {
				pc++
			} else if inst == "s" {
				pc++
				X += arg
			}
			inst, arg = getInst(pc)
		}

		fmt.Print("\n")
		cycle++
	}

	fmt.Println()
	for _, v := range crt {
		fmt.Println(v)
	}
}

func getInst(pc int) (string, int) {
	if pc >= len(datA) {
		return "q", 0
	} else {
		l := datA[pc]
		if l == "noop" {
			return "n", 0
		} else {
			a := strings.Split(l, " ")
			arg, _ := strconv.Atoi(a[1])
			return "f", arg
		}
	}
}
