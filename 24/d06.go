package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func d06p1(file string) int {
	res := 0

	// angle := (3 * math.Pi) / 2
	angle := 270
	o := GetVec(angle)

	lab, g := getLab(file)

	L := len(lab)
	lab[g[0]][g[1]] = "."
	for {
		if !isPointInLab([2]int{g[0] + o[0], g[1] + o[1]}, L) {
			if lab[g[0]][g[1]] == "." {
				res++
			}
			lab[g[0]][g[1]] = "o"
			break
		}
		n := lab[g[0]+o[0]][g[1]+o[1]]

		//check obstacle
		if n == "#" {
			angle = (angle + 90) % 360
			o = GetVec(angle)
		} else {
			// not out of bounds not an obstacle ahaed
			// mark prev pos, advance
			if lab[g[0]][g[1]] == "." {
				res++
			}
			lab[g[0]][g[1]] = "o"
			g[0] = g[0] + o[0]
			g[1] = g[1] + o[1]
		}
	}

	for _, l := range lab {
		fmt.Println(l)
	}

	return res
}

func d06p2(file string) int {
	res := 0
	// visited := map[string]int{}
	// _ = visited

	angle := 270
	o := GetVec(angle)

	lab, g := getLab(file)
	for _, l := range lab {
		fmt.Println(l)
	}

	L := len(lab)
	lab[g[0]][g[1]] = "."

	GoLeft(lab, g, 90)

	for {
		if !isPointInLab([2]int{g[0] + o[0], g[1] + o[1]}, L) {
			lab[g[0]][g[1]] = "o"
			break
		}
		n := lab[g[0]+o[0]][g[1]+o[1]]

		//check obstacle
		if n == "#" {
			// go backwards
			GoLeft(lab, g, angle)
			// keep moving
			angle = (angle + 90) % 360

			o = GetVec(angle)
		} else {
			// not out of bounds not an obstacle ahaed,
			rVec := GetVec((angle - 90) % 360)
			// if visited and right turn is visited
			if lab[g[0]][g[1]] == "o" && lab[g[0]+rVec[0]][g[1]+rVec[1]] == "o" {
				fmt.Println("from", g[0], g[1], "can go in loop?", g[0]+o[0], g[1]+o[1])
				if checkIfLoop(lab, g, (angle+90)%360) {
					res++
				}
				// v := fs(g[0]) + fs(g[1]) + fs(g[0]+o[0]) + fs(g[1]+o[1])
				// visited[v] = 1
			}
			// mark curr pos, advance
			lab[g[0]][g[1]] = "o"
			g[0] = g[0] + o[0]
			g[1] = g[1] + o[1]
		}
	}

	for _, l := range lab {
		fmt.Println(l)
	}
	// fmt.Println(visited)
	return res
}

func checkIfLoop(lab [][]string, g [2]int, angle int) bool {
	L := len(lab)
	o := GetVec(angle)
	visited := map[string]int{}

	for {
		if _, ok := visited[fs(g[0])+fs(g[1])]; ok {
			fmt.Println("It is loop")
			return true
		}
		ni := g[0] + o[0]
		nj := g[1] + o[1]
		if !isPointInLab([2]int{ni, nj}, L) {
			break
		}

		n := lab[ni][nj]
		//check obstacle
		if n == "#" {
			angle = (angle + 90) % 360
			o = GetVec(angle)
		} else {
			visited[fs(g[0])+fs(g[1])] = 1
			g[0] = ni
			g[1] = nj
		}
	}
	return false
}

func GoLeft(lab [][]string, g [2]int, angle int) {
	L := len(lab)
	o := GetVec(angle)
	visited := map[string]int{}

	for {
		if _, ok := visited[fs(g[0])+fs(g[1])]; ok {
			fmt.Println("loop")
			break
		}
		fmt.Println("lef round", g, o, angle)
		ni := g[0] + o[0]
		nj := g[1] + o[1]
		if !isPointInLab([2]int{ni, nj}, L) {
			lab[g[0]][g[1]] = "o"
			break
		}

		n := lab[ni][nj]
		//check obstacle
		if n == "#" {
			angle = (angle - 90) % 360
			o = GetVec(angle)
		} else {
			visited[fs(g[0])+fs(g[1])] = 1
			lab[g[0]][g[1]] = "o"
			g[0] = ni
			g[1] = nj
		}
	}
}

func isPointInLab(point [2]int, l int) bool {
	if point[0] < 0 || point[0] >= l || point[1] < 0 || point[1] >= l {
		return false
	}
	return true
}

func GetVec(angle int) [2]int {
	a := float64(angle) * math.Pi / 180
	i := int(math.Sin(a))
	j := int(math.Cos(a))
	return [2]int{i, j}
}

func getLab(file string) ([][]string, [2]int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, [2]int{0, 0}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lab := [][]string{}
	var guard [2]int
	li := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i := strings.Index(line, "^"); i != -1 {
			guard[0] = li
			guard[1] = i
		}
		lab = append(lab, strings.Split(line, ""))
		li++
	}

	return lab, guard
}
