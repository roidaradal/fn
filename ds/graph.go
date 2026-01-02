package ds

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn/dict"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

type (
	Edge      [2]Vertex
	Vertex    = string
	VertexSet = *Set[Vertex]
	EdgeSet   = *Set[Edge]
)

type Graph struct {
	Vertices     []Vertex
	Edges        []Edge
	IndexOf      map[Vertex]int
	EdgesOf      map[Vertex][]Edge
	NeighborsOf  map[Vertex]VertexSet
	EdgeWeightOf map[Edge]float64
}

// Create new Edge from v1-v2 string
func NewEdge(edge string) Edge {
	parts := str.CleanSplit(edge, "-")
	return Edge{parts[0], parts[1]}
}

// Create new directed Edge from v1->v2 string
func NewDirectedEdge(edge string) Edge {
	parts := str.CleanSplit(edge, "->")
	return Edge{parts[0], parts[1]}
}

// Edge string representation
func (e Edge) String() string {
	return fmt.Sprintf("%s-%s", e[0], e[1])
}

// Get edge endpoints
func (e Edge) Tuple() (Vertex, Vertex) {
	return e[0], e[1]
}

// Create new empty graph
func NewGraph() *Graph {
	return &Graph{
		Vertices:     make([]Vertex, 0),
		Edges:        make([]Edge, 0),
		IndexOf:      make(map[Vertex]int),
		EdgesOf:      make(map[Vertex][]Edge),
		NeighborsOf:  make(map[Vertex]VertexSet),
		EdgeWeightOf: make(map[Edge]float64),
	}
}

// Create undirected graph from vertices (separated by whitespace),
// and edgePairs (v1-v2 edges separated by whitespace)
func GraphFrom(vertices, edgePairs string) *Graph {
	g := NewGraph()
	g.Vertices = strings.Fields(vertices)
	for i, vertex := range g.Vertices {
		g.IndexOf[vertex] = i
	}
	for edgePair := range strings.FieldsSeq(edgePairs) {
		v1, v2 := NewEdge(edgePair).Tuple()
		g.AddUndirectedEdge(v1, v2)
	}
	return g
}

// Create directed graph from vertices (separated by whitespace),
// and edgePairs (v1->v2 edges separated by whitespace)
func DirectedGraphFrom(vertices, edgePairs string) *Graph {
	g := NewGraph()
	g.Vertices = strings.Fields(vertices)
	for i, vertex := range g.Vertices {
		g.IndexOf[vertex] = i
	}
	for edgePair := range strings.FieldsSeq(edgePairs) {
		v1, v2 := NewDirectedEdge(edgePair).Tuple()
		g.AddDirectedEdge(v1, v2)
	}
	return g
}

// Add vertex to graph
func (g *Graph) AddVertex(vertex Vertex) {
	g.Vertices = append(g.Vertices, vertex)
	g.IndexOf[vertex] = len(g.Vertices) - 1
}

// Add undirected edge to graph
func (g *Graph) AddUndirectedEdge(vertex1, vertex2 Vertex) {
	g.addEdge(vertex1, vertex2, true)
}

// Add directed edge to graph
func (g *Graph) AddDirectedEdge(vertex1, vertex2 Vertex) {
	g.addEdge(vertex1, vertex2, false)
}

// Common: add undirected / directed edge to graph
func (g *Graph) addEdge(vertex1, vertex2 Vertex, undirected bool) {
	_, ok1 := g.IndexOf[vertex1]
	_, ok2 := g.IndexOf[vertex2]
	if !ok1 || !ok2 {
		return // skip if one vertex is not part of the graph
	}
	edge := Edge{vertex1, vertex2}
	g.Edges = append(g.Edges, edge)
	g.EdgesOf[vertex1] = append(g.EdgesOf[vertex1], edge)
	g.EdgesOf[vertex2] = append(g.EdgesOf[vertex2], edge)
	vertices := []Vertex{vertex1, vertex2}
	if !undirected {
		vertices = []Vertex{vertex1}
	}
	for _, vertex := range vertices {
		if _, ok := g.NeighborsOf[vertex]; !ok {
			g.NeighborsOf[vertex] = NewSet[Vertex]()
		}
	}
	g.NeighborsOf[vertex1].Add(vertex2)
	if undirected {
		g.NeighborsOf[vertex2].Add(vertex1)
	}
}

// Get list of vertex neighbors
func (g Graph) Neighbors(vertex Vertex) []Vertex {
	neighbors, ok := g.NeighborsOf[vertex]
	if !ok {
		return []Vertex{}
	}
	return neighbors.Items()
}

// Get list of vertex neighbors, considering the active edge set
func (g Graph) ActiveNeighbors(vertex Vertex, activeEdges EdgeSet) []Vertex {
	neighbors := NewSet[Vertex]()
	for _, edge := range g.EdgesOf[vertex] {
		if activeEdges != nil && activeEdges.HasNo(edge) {
			continue
		}
		neighbors.Add(edge[0])
		neighbors.Add(edge[1])
	}
	neighbors.Delete(vertex)
	return neighbors.Items()
}

// Get edge weight
func (g Graph) EdgeWeight(edge Edge) (float64, bool) {
	weight, ok := g.EdgeWeightOf[edge]
	return weight, ok
}

// Check if list of vertices forms a clique
func (g Graph) IsClique(vertices []Vertex) bool {
	vertexSet := SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := SetFrom(g.Neighbors(vertex))
		adjacent.Add(vertex)
		if vertexSet.Difference(adjacent).NotEmpty() {
			return false
		}
	}
	return true
}

// Check if list of vertices forms an independent set
func (g Graph) IsIndependentSet(vertices []Vertex) bool {
	vertexSet := SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := SetFrom(g.Neighbors(vertex))
		if vertexSet.Intersection(adjacent).NotEmpty() {
			return false
		}
	}
	return true
}

// Check if list of vertices forms a dominating set
func (g Graph) IsDominatingSet(vertices []Vertex) bool {
	if len(vertices) == 0 {
		return false
	}
	vertexSet := SetFrom(vertices)
	for _, vertex := range g.Vertices {
		adjacent := SetFrom(g.Neighbors(vertex))
		adjacent.Add(vertex)
		if vertexSet.Intersection(adjacent).IsEmpty() {
			return false
		}
	}
	return true
}

// Check if vertex path is a valid Hamiltonian path
func (g Graph) IsHamiltonianPath(vertices []Vertex) bool {
	numVertices := len(vertices)
	if numVertices == 0 {
		return false
	}
	visitCount := dict.NewCounter(g.Vertices)
	for i := range numVertices - 1 {
		curr, next := vertices[i], vertices[i+1]
		if g.NeighborsOf[curr].HasNo(next) {
			return false // invalid path if no edge from curr => next
		}
		visitCount[curr] += 1
	}
	last := vertices[numVertices-1]
	visitCount[last] += 1
	// Check that all vertices visited exactly once
	return list.AllEqual(dict.Values(visitCount), 1)
}

// Check if vertex path is a valid Hamiltonian cycle
func (g Graph) IsHamiltonianCycle(vertices []Vertex) bool {
	// Check if vertices form a Hamiltonian path
	if !g.IsHamiltonianPath(vertices) {
		return false
	}
	// Check if there is an edge to connect last vertex and first vertex
	first, last := vertices[0], list.Last(vertices, 1)
	return g.NeighborsOf[last].Has(first)
}

// Check if edge sequence is a valid Eulerian path
func (g Graph) IsEulerianPath(edges []Edge) (bool, [2]Vertex) {
	var pair [2]Vertex
	numEdges := len(edges)
	if numEdges < 2 {
		return false, pair
	}
	visitCount := dict.NewCounter(g.Edges)
	a1, b1 := edges[0].Tuple()
	a2, b2 := edges[1].Tuple()
	var head, tail Vertex
	if a1 == a2 {
		head = b1
		tail = b2
	} else if b1 == a2 {
		head = a1
		tail = b2
	} else if a1 == b2 {
		head = b1
		tail = a2
	} else if b1 == b2 {
		head = a1
		tail = a2
	} else {
		return false, pair
	}
	visitCount[edges[0]] += 1
	visitCount[edges[1]] += 1
	for _, edge := range edges[2:] {
		visitCount[edge] += 1
		a, b := edge.Tuple()
		switch tail {
		case a:
			tail = b
		case b:
			tail = a
		default:
			return false, pair
		}
	}
	// Check that all edges visited exactly once
	return list.AllEqual(dict.Values(visitCount), 1), [2]Vertex{head, tail}
}

// Check if edge sequence is a valid Eulerian cycle
func (g Graph) IsEulerianCycle(edges []Edge) bool {
	// Check if vertices form an Eulerian path
	ok, pair := g.IsEulerianPath(edges)
	if !ok {
		return false
	}
	head, tail := pair[0], pair[1]
	return head == tail
}

// Perform BFS traversal on the graph, starting at given vertex,
// considering the active edge set, return list of vertices visited
func (g Graph) BFSTraversal(start Vertex, activeEdges EdgeSet) []Vertex {
	q := NewQueue[Vertex]()
	q.Enqueue(start)
	visited := NewSet[Vertex]()
	for q.NotEmpty() {
		current, _ := q.Dequeue()
		if visited.Has(current) {
			continue
		}
		visited.Add(current)
		for _, neighbor := range g.ActiveNeighbors(current, activeEdges) {
			if visited.Has(neighbor) {
				continue
			}
			q.Enqueue(neighbor)
		}
	}
	return visited.Items()
}
