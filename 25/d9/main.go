package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fn string

type Point struct {
	x int
	y int
}

func main() {
	if os.Args[1] == "test" {
		fn = "test"
	} else {
		fn = "input"
	}

	redTiles := lines(fn)
	fmt.Println(redTiles)

	var vertices []Point
	for _, ti := range redTiles {
		vertices = append(vertices, tileToPoint(ti))
	}

	h := &RecHeap{}
	heap.Init(h)

	for i, ti := range vertices {
		for j := i + 1; j < len(vertices); j++ {
			tj := vertices[j]
			rec := &Rec{ti, tj, calcA(ti, tj)}
			heap.Push(h, rec)
		}
	}

	for i, e := range *h {
		if i > 100 {
			break
		}
		fmt.Println(i, e)
	}

a:
	for len(*h) > 0 {
		rec := heap.Pop(h).(*Rec)

		pa := rec.pa
		pc := rec.pc
		pb := Point{pa.x, pc.y}
		pd := Point{pc.x, pa.y}

		// test if pb and pd are on polygon (pa, pc we know they are)
		if isPointInPolygon(pb, vertices) && isPointInPolygon(pd, vertices) {
			// test if all edges are onpolygon a->b, b->c, c->d, d->a
			recVert := []Point{pa, pb, pc, pd}
			j := len(recVert) - 1
			for i, vi := range recVert {
				vj := recVert[j]

				diffx := vi.x - vj.x
				diffy := vi.y - vj.y

				diffx, diffy = decreaseDiffs(diffx, diffy)

				for diffx != 0 || diffy != 0 {
					if !isPointInPolygon(Point{vi.x - diffx, vi.y - diffy}, vertices) {
						goto a
					}

					diffx, diffy = decreaseDiffs(diffx, diffy)
				}

				j = i
			}

			fmt.Println("res: ", rec)
			break
		}
	}
}

func decreaseDiffs(x, y int) (int, int) {
	if x != 0 {
		if x < 0 {
			x++
		} else {
			x--
		}
	}
	if y != 0 {
		if y < 0 {
			y++
		} else {
			y--
		}
	}

	return x, y
}

func isPointInPolygon(p Point, vertices []Point) bool {
	j := len(vertices) - 1
	inside := false
	for i, va := range vertices {
		vb := vertices[j]

		// is point on edge
		if va.x == vb.x && va.x == p.x && min(va.y, vb.y) <= p.y && p.y <= max(va.y, vb.y) { // vertical
			return true
		} else if p.y == va.y && min(va.x, vb.x) <= p.x && p.x <= max(va.x, vb.x) { // horizontal
			return true
		}

		// is point left from edge
		if ((va.y > p.y) != (vb.y > p.y)) &&
			(p.x < (va.x + ((vb.x - va.x) * (p.y - va.y) / (vb.y - va.y)))) {
			inside = !inside
		}

		j = i
	}

	return inside
}

func lines(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

func calcA(pi, pj Point) int {
	// fmt.Println("points", pi, pj)
	l := abs(pi.x-pj.x) + 1
	a := abs(pi.y-pj.y) + 1
	// fmt.Println(l, a)
	return l * a
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func tileToPoint(t string) Point {
	nums := strings.Split(t, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])

	return Point{x, y}
}

type Rec struct {
	pa   Point
	pc   Point
	area int
}

type RecHeap []*Rec

func (r RecHeap) Len() int           { return len(r) }
func (r RecHeap) Less(i, j int) bool { return r[i].area > r[j].area }
func (r RecHeap) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r *RecHeap) Push(x any) {
	*r = append(*r, x.(*Rec))
}

func (r *RecHeap) Pop() any {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}
