package shortestpaths

import (
	"fmt"
	"math"
)

const infinity = math.MaxInt

func BellmanFordAlg(graph *Graph, initVertice int) {
	distances := make([]int, graph.numOfVertices)

	for i := 0; i < graph.numOfVertices; i++ {
		distances[i] = infinity
	}

	distances[initVertice] = 0

	// relax n - 1 times
	for i := 0; i < graph.numOfVertices-1; i++ {
		for j := 0; j < graph.numOfEdges; j++ {
			src := graph.edges[j].src
			dest := graph.edges[j].dest
			weight := graph.edges[j].weight

			if distances[src] != infinity &&
				distances[src]+weight < distances[dest] {
				distances[dest] = distances[src] + weight
			}
		}
	}

	if graphContainsNegativeWeightCycle(graph, distances) {
		println("graph contains negative weight cycle")
		return
	}

	printRes(distances, initVertice)
}

func graphContainsNegativeWeightCycle(graph *Graph, distances []int) bool {
	for i := 0; i < graph.numOfEdges; i++ {
		src := graph.edges[i].src
		dest := graph.edges[i].dest
		weight := graph.edges[i].weight

		if distances[src] != infinity &&
			distances[src]+weight < distances[dest] {
			return true
		}
	}

	return false
}

func printRes(distances []int, initVertice int) {
	fmt.Printf("distances from init vertice %v\n", initVertice)

	for i := range distances {
		fmt.Printf("%v\t\t%v\n", i, distances[i])
	}

}

func TestBellmanFordAlg() {
	numOfVertices := 5
	numOfEdges := 8

	graph := createGraph(numOfVertices, numOfEdges)

	// https://media.geeksforgeeks.org/wp-content/uploads/bellmanford1.png graph
	graph.edges = append(graph.edges, Edge{src: 0, dest: 1, weight: -1})
	graph.edges = append(graph.edges, Edge{src: 0, dest: 2, weight: 4})
	graph.edges = append(graph.edges, Edge{src: 1, dest: 2, weight: 3})
	graph.edges = append(graph.edges, Edge{src: 1, dest: 3, weight: 2})
	graph.edges = append(graph.edges, Edge{src: 1, dest: 4, weight: 2})
	graph.edges = append(graph.edges, Edge{src: 3, dest: 2, weight: 5})
	graph.edges = append(graph.edges, Edge{src: 3, dest: 1, weight: 1})
	graph.edges = append(graph.edges, Edge{src: 4, dest: 3, weight: -3})

	BellmanFordAlg(graph, 0)
}
