package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	cell struct {
		num    int
		col    int
		row    int
		marked bool
	}

	card struct {
		cols  []int
		rows  []int
		cells map[int]*cell
	}
)

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	datA := strings.Split(string(dat), "\n")

	re := regexp.MustCompile("[0-9]+")

	var cards []card
	currRow := 0
	currCard := card{make([]int, 5), make([]int, 5), make(map[int]*cell)}
	cards = append(cards, currCard)
	fmt.Println(currCard)

	for _, l := range datA[2:] { // create cards
		if l == "" {
			currCard = card{make([]int, 5), make([]int, 5), make(map[int]*cell)}
			cards = append(cards, currCard)
			currRow = 0
		} else {
			nums := re.FindAllString(l, -1)
			for j, ns := range nums {
				n, _ := strconv.Atoi(ns)
				cell := &cell{num: n, col: j, row: currRow, marked: false}
				currCard.cells[n] = cell
			}
			currRow++
		}
	}

	numbers := re.FindAllString(datA[0], -1)

	winner, lastn := getWinner(cards, numbers)

	fmt.Println(winner)

	sum := 0
	for _, cel := range winner.cells {
		if !cel.marked {
			sum += cel.num
		}
	}

	fmt.Println(sum * lastn)

}

func getWinner(cards []card, numbers []string) (card, int) {
	for _, ns := range numbers {
		// fmt.Println(ns)
		n, _ := strconv.Atoi(ns)

		for _, c := range cards {
			if cel, ok := c.cells[n]; ok {
				// fmt.Println(c, cel)
				c.rows[cel.row]++
				c.cols[cel.col]++
				c.cells[n].marked = true

				for _, col := range c.cols {
					if col > 4 {
						return c, n
					}
				}
				for _, row := range c.rows {
					if row > 4 {
						return c, n
					}
				}
			}
		}
	}
	return card{}, 0
}
