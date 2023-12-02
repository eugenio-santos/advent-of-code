package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var prio = map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := 0
	sacks := strings.Split(string(dat), "\n")
	for i := 0; i < len(sacks); i += 3 {
		first := make(map[rune]int)
		for _, s := range sacks[i] {
			first[s] = 0
		}

		second := make(map[rune]int)
		for _, s := range sacks[i+1] {
			if _, ok := first[s]; ok {
				second[s] = 0
			}
		}

		for _, s := range sacks[i+2] {
			if _, ok := second[s]; ok {
				sum += prio[s]
				break
			}
		}
	}

	fmt.Println(sum)
}

// part one
// sum := 0
// for _, sack := range strings.Split(string(dat), "\n") {
// 	// iterate 1st half
// 	first := make(map[rune]int)
// 	for _, s := range sack[:len(sack)/2] {
// 		first[s] = 0
// 	}

// 	// 2nd half
// 	for _, s := range sack[len(sack)/2:] {
// 		if _, ok := first[s]; ok {
// 			sum += prio[s]
// 			break
// 		}
// 	}
// }
// fmt.Println(sum)
