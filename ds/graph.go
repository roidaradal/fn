package ds

import "github.com/roidaradal/fn/dict"

type VertexPair = [2]string

type Graph struct {
	vertices    *Set[string]
	edges       map[string][]string
	edgeWeights map[VertexPair]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices:    NewSet[string](),
		edges:       make(map[string][]string),
		edgeWeights: make(map[VertexPair]int),
	}
}

func (g Graph) Vertices() []string {
	return g.vertices.Items()
}

func (g Graph) EdgeWeight(pair VertexPair) int {
	return g.edgeWeights[pair]
}

func (g Graph) Edges(v string) []string {
	return g.edges[v]
}

func (g *Graph) AddVertex(v string) {
	g.vertices.Add(v)
}

func (g *Graph) AddEdgeWeight(pair VertexPair, weight int) {
	g.edgeWeights[pair] = weight
}

func (g *Graph) AddEdge(v1, v2 string) {
	g.edges[v1] = append(g.edges[v1], v2)
}

func (g *Graph) InitEdges() {
	for _, v := range g.Vertices() {
		if !dict.HasKey(g.edges, v) {
			g.edges[v] = make([]string, 0)
		}
	}
}
