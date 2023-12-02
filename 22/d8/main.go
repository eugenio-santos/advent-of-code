package main

import (
	"fmt"
	"os"
	"strings"
)

type tree struct {
	row int
	col int
	val byte
}

func main() {
	dat, _ := os.ReadFile("in")

	datA := strings.Split(string(dat), "\n")

	for _, v := range datA {
		for _, c := range v {
			fmt.Print(c, " ")
		}
		fmt.Println()
	}
	rows := len(datA)
	cols := len(datA[0])

	// mc := make([]int, cols)
	m := make([]([]int), rows)

	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
		m[i][0] = 1
		m[i][cols-1] = 1
	}
	for i := 0; i < cols; i++ {
		m[0][i] = 1
		m[rows-1][i] = 1
	}

	max := 0
	for row, line := range datA {
		fmt.Println(line)
		for col, char := range line {
			// check left
			east := 0
			for i := col - 1; i >= 0; i-- {
				east++
				if char <= rune(datA[row][i]) {
					break
				}
			}

			// check right
			west := 0
			for i := col + 1; i < cols; i++ {
				west++
				if char <= rune(datA[row][i]) {
					break
				}
			}

			// check up
			north := 0
			for i := row - 1; i >= 0; i-- {
				north++
				if char <= rune(datA[i][col]) {
					break
				}
			}

			// check down
			south := 0
			for i := row + 1; i < rows; i++ {
				south++
				if char <= rune(datA[i][col]) {
					break
				}
			}

			calc := east * west * north * south

			if calc > max {
				max = calc
			}
		}
	}

	fmt.Println(max)
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// countOnes(m)
}

func countOnes(m [][]int) {
	smu := 0
	for _, r := range m {
		for _, c := range r {
			if c == 1 {
				smu++
			}
		}
	}
	fmt.Println(smu)
}

// // mc := make([]int, cols)
// m := make([]([]int), rows)

// for i := 0; i < rows; i++ {
// 	m[i] = make([]int, cols)
// 	m[i][0] = 1
// 	m[i][cols-1] = 1
// }
// for i := 0; i < cols; i++ {
// 	m[0][i] = 1
// 	m[rows-1][i] = 1
// }

// for i := 1; i < rows-1; i++ { // row search
// 	// fmt.Println(i)
// 	h := tree{row: i, col: 0, val: datA[i][0]}
// 	h2 := tree{row: i, col: rows - 1, val: datA[i][rows-1]}

// 	// fmt.Println("j")
// 	for j := 1; j < cols-1; j++ {
// 		// fmt.Print(j, " ", cols-j-1, " ")

// 		if h.val < datA[i][j] {
// 			h.row = i
// 			h.col = j
// 			h.val = datA[i][j]
// 			m[h.row][h.col] = 1
// 		} else if h2.val < datA[i][cols-j-1] {
// 			h2.row = i
// 			h2.col = cols - j - 1
// 			h2.val = datA[i][cols-j-1]
// 			m[h2.row][h2.col] = 1
// 		}
// 	}
// 	// fmt.Print("\n")
// }

// for _, v := range m {
// 	fmt.Println(v)
// }
// countOnes(m)
// fmt.Println()
// for j := 1; j < cols-1; j++ { // column search
// 	// fmt.Println(j)

// 	h := tree{row: 0, col: j, val: datA[0][j]}
// 	h2 := tree{row: rows - 1, col: j, val: datA[rows-1][j]}

// 	// fmt.Println("i")

// 	for i := 1; i < rows-1; i++ {
// 		// fmt.Print(i, " ", rows-1-i, " ")
// 		// fmt.Println(string(datA[i][j]))
// 		if h.val < datA[i][j] {
// 			h.row = i
// 			h.col = j
// 			h.val = datA[i][j]
// 			m[h.row][h.col] = 1
// 		} else if h2.val < datA[rows-1-i][j] {
// 			h2.row = rows - 1 - i
// 			h2.col = j
// 			h2.val = datA[rows-1-i][j]
// 			m[h2.row][h2.col] = 1
// 		}
// 	}
// 	// fmt.Print("\n")
// }

// for _, v := range m {
// 	fmt.Println(v)
// }
// countOnes(m)
