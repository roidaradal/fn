package ds

import "github.com/roidaradal/fn/dict"

type Graph struct {
	vertices    *Set[string]
	edges       map[string]*Set[string]
	edgeWeights map[VertexPair]float64
}

// Pair of vertices
type VertexPair = [2]string

// Creates a new, empty graph
func NewGraph() *Graph {
	return &Graph{
		vertices:    NewSet[string](),
		edges:       make(map[string]*Set[string]),
		edgeWeights: make(map[VertexPair]float64),
	}
}

// Initialize edge sets, from the current set of vertices,
// Assumes that vertex set has already been added
func (g *Graph) InitializeEdges() {
	for _, vertex := range g.Vertices() {
		dict.SetDefault(g.edges, vertex, NewSet[string]())
	}
}

// List of graph vertices
func (g Graph) Vertices() []string {
	return g.vertices.Items()
}

// Map of graph edges
func (g Graph) Edges() map[string][]string {
	edges := make(map[string][]string)
	for vertex, edgeSet := range g.edges {
		edges[vertex] = edgeSet.Items()
	}
	return edges
}

// Get edge weight of vertex pair
func (g Graph) EdgeWeight(pair VertexPair) (float64, bool) {
	weight, ok := g.edgeWeights[pair]
	return weight, ok
}

// List of vertices connected to given vertex
func (g Graph) Neighbors(vertex string) []string {
	return g.edges[vertex].Items()
}

// Add graph vertex
func (g *Graph) AddVertex(vertex string) {
	g.vertices.Add(vertex)
}

// Add undirected edge for the two given vertices
func (g *Graph) AddUndirectedEdge(vertex1, vertex2 string) {
	if !g.vertices.Contains(vertex1) || !g.vertices.Contains(vertex2) {
		return // skip if one of the vertices are not part of the graph
	}
	g.edges[vertex1].Add(vertex2)
	g.edges[vertex2].Add(vertex1)
}

// Add directed edge from vertex1 to vertex2
func (g *Graph) AddDirectedEdge(vertex1, vertex2 string) {
	if !g.vertices.Contains(vertex1) || !g.vertices.Contains(vertex2) {
		return // skip if one of the vertices are not part of the graph
	}
	g.edges[vertex1].Add(vertex2)
}

// Add graph edge weight
func (g *Graph) AddEdgeWeight(pair VertexPair, weight float64) {
	g.edgeWeights[pair] = weight
}
