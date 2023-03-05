package shortestpaths


type Edge struct {
	src int
	dest int
	weight int
}

type Graph struct {
	numOfVertices int
	numOfEdges int
	edges []Edge
}

func createGraph(numOfVertices, numOfEdges int) *Graph {
	return &Graph{
		numOfVertices: numOfVertices,
		numOfEdges: numOfEdges,
		edges: make([]Edge, 0),
	}
}

