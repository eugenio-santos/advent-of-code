package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	LOSE string = "X"
	DRAW        = "Y"
	WIN         = "Z"
)

const (
	ROCK    string = "A"
	PAPER          = "B"
	SCISSOR        = "C"
)

var shapes = [3]string{ROCK, PAPER, SCISSOR}

func main() {
	dat, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	score := 0
	for _, g := range strings.Split(string(dat), "\n") {
		game := strings.Split(g, " ")
		score += getTotalPoints(decideChoices(game))
	}
	fmt.Println(score)
}

func getTotalPoints(choices []string) int {
	return getChoicePoints(choices[1]) + getGamePoints(choices[0], choices[1])
}

func decideChoices(game []string) []string {
	switch game[1] {
	case LOSE:
		game[1] = getShape(getShapeIndex(game[0]) - 1)
	case DRAW:
		game[1] = getShape(getShapeIndex(game[0]))
	case WIN:
		game[1] = getShape(getShapeIndex(game[0]) + 1)
	}
	return game
}

func getShapeIndex(s string) int {
	for i, shape := range shapes {
		if s == shape {
			return i
		}
	}
	return -10
}

func getShape(i int) string {
	i = i % 3
	if i < 0 {
		i = i + 3
	}
	return shapes[i]
}

func getChoicePoints(choice string) int {
	switch choice {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSOR:
		return 3
	default:
		return 0
	}
}

func getGamePoints(elf string, me string) int {
	if diff := getChoicePoints(elf) - getChoicePoints(me); diff == 0 {
		return 3
	} else if diff == -1 || diff == 2 {
		return 6
	} else {
		return 0
	}
}
