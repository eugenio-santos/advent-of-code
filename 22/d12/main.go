package main

import (
	"container/heap"
	"os"
	"strings"
)

// Vertex represents a single vertex in the graph.
type Vertex struct {
	// ID is the unique identifier for the vertex.
	ID int
	// Edges is a list of edges connected to this vertex.
	Edges []*Vertex
}

// Graph is a data structure for storing a collection of vertices
// and the edges connecting them.
type Graph struct {
	// Vertices is a map of vertex IDs to vertex objects.
	Vertices map[int]*Vertex
}

// AddVertex adds a new vertex to the graph.
func (g *Graph) AddVertex(id int) *Vertex {
	v := &Vertex{ID: id, Edges: []*Vertex{}}
	g.Vertices[id] = v
	return v
}

// AddEdge adds a new, unweighted, edge to the graph.
func (g *Graph) AddEdge(srcID, dstID int) {
	src := g.Vertices[srcID]
	dst := g.Vertices[dstID]
	src.Edges = append(src.Edges, dst)
}

// ###vertexHeap is a min-heap of vertices sorted by distance.
type vertexHeap struct {
	vertices []vertex
}

// A vertex represents a vertex in a graph with a pointer to the vertex object,
// a distance, a path to the vertex, and an index in the heap.
type vertex struct {
	v     *Vertex
	dist  int
	path  []*Vertex
	index int
}

// Implement the heap.Interface interface for the vertexHeap type.

func (h *vertexHeap) Len() int           { return len(h.vertices) }
func (h *vertexHeap) Less(i, j int) bool { return h.vertices[i].dist < h.vertices[j].dist }
func (h *vertexHeap) Swap(i, j int) {
	h.vertices[i], h.vertices[j] = h.vertices[j], h.vertices[i]
	h.vertices[i].index = i
	h.vertices[j].index = j
}

// Push adds a new vertex to the heap.
func (h *vertexHeap) Push(x interface{}) {
	n := h.Len()
	v := x.(*vertex)
	v.index = n
	h.vertices = append(h.vertices, *v)
}

// Pop removes and returns the vertex with the smallest distance from the heap.
func (h *vertexHeap) Pop() interface{} {
	old := h.vertices
	n := len(old)
	v := old[n-1]
	v.index = -1 // for safety
	h.vertices = old[0 : n-1]
	return &v
}

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	datA := strings.Split(string(dat), "\n")
}

func ShortestPath(g *Graph, startID, endID int) []*Vertex {
	// Create a priority queue of vertices and initialize it with the start vertex.
	vertices := &vertexHeap{
		vertices: []vertex{
			{
				v:     g.Vertices[startID],
				path:  []*Vertex{g.Vertices[startID]},
				dist:  0,
				index: 0,
			},
		},
	}
	heap.Init(vertices)

	// Create a set of visited vertices.
	visited := make(map[int]struct{})

	// Keep track of the current and previous vertices in the shortest path.
	var currVertex, prevVertex *vertex

	// Keep looping until the priority queue is empty or the end vertex is found.
	for vertices.Len() > 0 {
		// Get the vertex with the shortest distance from the priority queue.
		currVertex = heap.Pop(vertices).(*vertex)

		// Check if the end vertex has been reached.
		if currVertex.v.ID == endID {
			break
		}

		// Skip this vertex if it has already been visited.
		if _, ok := visited[currVertex.v.ID]; ok {
			continue
		}

		// Add this vertex to the set of visited vertices.
		visited[currVertex.v.ID] = struct{}{}

		// Add all of the unvisited neighbors of this vertex to the priority queue.
		for _, neighbor := range currVertex.v.Edges {
			if _, ok := visited[neighbor.ID]; !ok {
				prevVertex = &vertex{
					v:     neighbor,
					path:  append(currVertex.path, neighbor),
					dist:  currVertex.dist + 1,
					index: vertices.Len(),
				}
				heap.Push(vertices, prevVertex)
			}
		}
	}

	// Return the shortest path, or the empty slice if there is no path.
	if currVertex.v.ID == endID {
		return currVertex.path
	}
	return []*Vertex{}
}
