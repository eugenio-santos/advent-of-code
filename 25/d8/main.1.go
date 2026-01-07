package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fn string
var maxPair int

func main() {
	if os.Args[1] == "test" {
		fn = "test"
		maxPair = 20
	} else {
		fn = "input"
		maxPair = 1000
	}

	pointss := lines(fn)
	var points []Point
	for _, p := range pointss {
		points = append(points, psToP(p))
	}

	h := &DistHeap{}
	heap.Init(h)

	for i, pi := range points {
		for j := i + 1; j < len(points); j++ {
			// fmt.Println(j, points[j])
			d := sqrtDist(pi, points[j])
			dist := &Dist{i, j, d}
			heap.Push(h, dist)
		}
	}

	// for i, e := range *h {
	// 	_ = i
	// 	fmt.Println(e.d)
	// }

	circ := []map[string]struct{}{}

	var winPair Dist

	for len(*h) > 0 {
		pair := heap.Pop(h).(*Dist)
		tomerge := []int{}

		for i, c := range circ {
			if InSet(pointss, c, *pair) {
				c[pointss[pair.i]] = struct{}{}
				c[pointss[pair.j]] = struct{}{}
				tomerge = append(tomerge, i)
			}
		}

		if len(tomerge) == 0 {
			newm := map[string]struct{}{pointss[pair.i]: {}, pointss[pair.j]: {}}
			circ = append(circ, newm)
		} else {
			for _, i := range tomerge[1:] {
				for k, v := range circ[i] {
					circ[tomerge[0]][k] = v
				}
				circ[i] = circ[len(circ)-1]
				circ = circ[:len(circ)-1]
			}
		}

		if len(circ) == 1 && len(circ[0]) == maxPair {
			winPair = *pair
			break
		}
	}

	fmt.Println(circ)
	fmt.Println(points[winPair.i], points[winPair.j])
	fmt.Println(points[winPair.i].x * points[winPair.j].x)

	// maxC := []int{0, 0, 0}

	// for _, c := range circ {
	// 	lc := len(c)
	// 	if lc > maxC[0] {
	// 		maxC[0] = lc
	// 		if lc > maxC[1] {
	// 			maxC[0] = maxC[1]
	// 			maxC[1] = lc
	// 			if lc > maxC[2] {
	// 				maxC[1] = maxC[2]
	// 				maxC[2] = lc
	// 			}
	// 		}
	// 	}
	// }

	// fmt.Println(maxC[0], maxC[1], maxC[2])
	// fmt.Println(maxC[0] * maxC[1] * maxC[2])
}

func InSet(pointss []string, set map[string]struct{}, dist Dist) bool {
	if _, ok := set[pointss[dist.i]]; ok {
		return true
	}
	if _, ok := set[pointss[dist.j]]; ok {
		return true
	}
	return false
}

func sqrtDist(p1, p2 Point) int {
	return pow2(p1.x-p2.x) + pow2(p1.y-p2.y) + pow2(p1.z-p2.z)
}

func pow2(i int) int {
	return i * i
}

type Point struct {
	x int
	y int
	z int
}

func psToP(ps string) Point {
	nums := strings.Split(ps, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	z, _ := strconv.Atoi(nums[2])

	return Point{x, y, z}
}
func lines(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}

type Dist struct {
	i int
	j int
	d int
}

type DistHeap []*Dist

func (d DistHeap) Len() int           { return len(d) }
func (d DistHeap) Less(i, j int) bool { return d[i].d < d[j].d }
func (d DistHeap) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func (d *DistHeap) PushBounded(item *Dist) {
	if len(*d) < 10 {
		heap.Push(d, item)
		return
	}

	if item.d < (*d)[0].d {
		(*d)[0] = item // Replace root
		heap.Fix(d, 0) // Sift down - O(log maxSize)
	}
}

func (d *DistHeap) Push(x any) {
	*d = append(*d, x.(*Dist))
}

func (d *DistHeap) Pop() any {
	old := *d
	n := len(old)
	x := old[n-1]
	*d = old[0 : n-1]
	return x
}
