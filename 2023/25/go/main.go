package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	input, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	g := graph.New(graph.StringHash)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)

		splt := strings.Split(s, ":")

		first := splt[0]
		rest := splt[1]
		g.AddVertex(first)

		for _, e := range strings.Split(strings.Trim(rest, " "), " ") {
			g.AddVertex(string(e))
			g.AddEdge(first, e)
		}
	}

	fullG, _ := os.Create("./full")
	_ = draw.DOT(g, fullG)
	fullG.Close()

	g.RemoveEdge("hfx", "pzl")
	g.RemoveEdge("bvb", "cmg")
	g.RemoveEdge("nvd", "jqt")

	cutG, _ := os.Create("./cut")
	_ = draw.DOT(g, cutG)
	cutG.Close()

}
