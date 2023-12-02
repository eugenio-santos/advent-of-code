package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x int
	y int
}

// Rope is a struct that represents a conceptual rope with a head and a tail.
type Rope struct {
	knots [10]Knot
}

// MoveHead moves the head of the rope in the specified direction.
func (r *Rope) MoveHead(dx, dy int) {
	r.knots[0].x += dx
	r.knots[0].y += dy
}

// MoveKnot moves the tail of the rope to keep up with the head.
func (r *Rope) MoveKnot(i int) {
	dx := r.knots[i-1].x - r.knots[i].x
	dy := r.knots[i-1].y - r.knots[i].y
	norm := int(math.Sqrt(float64(dx*dx + dy*dy)))
	// println("norm", dx, dy, norm)
	// println("tailb", r.knots[i].x, r.knots[i].y)

	// Check if the head is two steps directly up, down, left, or right from the tail.
	if norm >= 2 && r.knots[i-1].x != r.knots[i].x && r.knots[i-1].y != r.knots[i].y {
		r.knots[i].x = r.knots[i-1].x - dx/norm
		r.knots[i].y = r.knots[i-1].y - dy/norm
	} else if norm >= 2 && r.knots[i-1].x == r.knots[i].x {
		r.knots[i].y = r.knots[i-1].y - dy/norm
	} else if norm >= 2 && r.knots[i-1].y == r.knots[i].y {
		r.knots[i].x = r.knots[i-1].x - dx/norm
	}

	// println("taila", r.tailX, r.tailY)

}

func Norm(headX, headY, tailX, tailY int) int {
	// Calculate the difference between the X and Y coordinates of the two points
	dx := headX - tailX
	dy := headY - tailY

	// Use the Pythagorean theorem to calculate the norm
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	// Create a new Rope.
	var knots [10]Knot
	r := Rope{knots: knots}
	fmt.Println(r)
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	totalSteps := 0
	// Parse the input commands.
	commands := strings.Split(string(dat), "\n")
	visited := make(map[[2]int]bool)
	for _, command := range commands {
		direction := string(command[0])
		steps, _ := strconv.Atoi(string(command[2:]))

		fmt.Println("#############MOVE", direction, steps)
		totalSteps += steps
		for s := 0; s < steps; s++ {
			// Update the position of the head based on the command.
			// fmt.Println("Before Move")
			// fmt.Println("HEAD", r.headX, r.headY)
			// fmt.Println("TAIL", r.tailX, r.tailY)

			switch direction {
			case "U":
				r.MoveHead(0, 1)
			case "D":
				r.MoveHead(0, -1)
			case "L":
				r.MoveHead(-1, 0)
			case "R":
				r.MoveHead(1, 0)
			}

			// Update the position of the tail.
			for i := 1; i < 10; i++ {
				r.MoveKnot(i)
			}
			visited[[2]int{r.knots[9].x, r.knots[9].y}] = true
			fmt.Println("After Move")
			fmt.Println("HEAD", r.knots[0].x, r.knots[0].y)
			fmt.Println("TAIL", r.knots[9].x, r.knots[9].y)
			fmt.Println("V C", len(visited))
		}
	}

	// Print the result.
	fmt.Println(visited)
	fmt.Println(len(visited))
	fmt.Println(totalSteps)
}
