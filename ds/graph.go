package ds

type VertexPair = [2]string

type Graph struct {
	vertices    *Set[string]
	edgeWeights map[VertexPair]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices:    NewSet[string](),
		edgeWeights: make(map[VertexPair]int),
	}
}

func (g Graph) Vertices() []string {
	return g.vertices.Items()
}

func (g Graph) EdgeWeight(pair VertexPair) int {
	return g.edgeWeights[pair]
}

func (g *Graph) AddVertex(v string) {
	g.vertices.Add(v)
}

func (g *Graph) AddEdgeWeight(pair VertexPair, weight int) {
	g.edgeWeights[pair] = weight
}
