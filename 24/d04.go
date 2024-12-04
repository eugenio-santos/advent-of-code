package main

// 88 77 65 83
// X  M  A  S

type Cell struct {
	x   *Cell
	i   int
	j   int
	vec []int
}

func d04p1(m [][]rune) int {
	//  find all X
	xes := []Cell{}
	for i, row := range m {
		for j, ru := range row {
			if ru == 'X' {
				xes = append(xes, Cell{nil, i, j, nil})
			}
		}
	}

	// find all M next to an X and what is the direction(vector)
	emes := []Cell{}
	for _, x := range xes {
		emes = append(emes, getNeighboursMRune(m, x, 1, 'M')...)
	}
	// fmt.Println("emes", emes)

	as := []Cell{}
	for _, eme := range emes {
		newi := eme.i + eme.vec[0]
		newj := eme.j + eme.vec[1]
		if newi >= 0 && newi < len(m) && newj >= 0 && newj < len(m) &&
			m[newi][newj] == 'A' {
			newii := eme.i + (eme.vec[0] * 2)
			newjj := eme.j + (eme.vec[1] * 2)
			if newii >= 0 && newii < len(m) && newjj >= 0 && newjj < len(m) &&
				m[newii][newjj] == 'S' {
				as = append(as, eme)
			}
		}
	}

	// fmt.Println(as)

	return len(as)
}

func getNeighboursMRune(m [][]rune, c Cell, r int, ru rune) []Cell {
	startRow := max(c.i-r, 0)
	endRow := min(c.i+r+1, len(m))

	startCol := max(c.j-r, 0)
	endCol := min(c.j+r+1, len(m))

	// fmt.Println(c, "row", startRow, endRow, "col", startCol, endCol, "ru", ru)
	cells := []Cell{}

	for i := startRow; i < endRow; i++ {
		for j := startCol; j < endCol; j++ {
			// fmt.Println(i, j, m[i][j])
			if m[i][j] == ru {
				cells = append(cells, Cell{&c, i, j, []int{i - c.i, j - c.j}})
			}
		}
	}
	return cells
}

func d04p2(m [][]rune) int {
	A := []Cell{}
	for i, row := range m {
		for j, ru := range row {
			if ru == 'A' {
				A = append(A, Cell{nil, i, j, nil})
			}
		}
	}

	res := 0
	for _, a := range A {
		if evalMAS(m, a) {
			res++
		}
	}

	return res
}

// get the neighboring martix of a cell
func evalMAS(m [][]rune, c Cell) bool {
	if c.i-1 < 0 || c.i+1 >= len(m) || c.j-1 < 0 || c.j+1 >= len(m) {
		return false
	}

	if (m[c.i-1][c.j-1] == 'M' || m[c.i-1][c.j-1] == 'S') &&
		(m[c.i+1][c.j+1] == 'M' || m[c.i+1][c.j+1] == 'S') &&
		(m[c.i+1][c.j-1] == 'M' || m[c.i+1][c.j-1] == 'S') &&
		(m[c.i-1][c.j+1] == 'M' || m[c.i-1][c.j+1] == 'S') {
		if (m[c.i-1][c.j-1] != m[c.i+1][c.j+1]) && (m[c.i+1][c.j-1] != m[c.i-1][c.j+1]) {
			return true
		}
	}
	return false
}
