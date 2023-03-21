package maximumflow

import (
	"fmt"
	"math"
)

// get paths from s(source) to t(sink)
func bfs(graph [][]int, s, t int) ([]int, bool) {
	numOfVertices := len(graph)

	paths := make([]int, numOfVertices)
	queue := make([]int, 0)

	visited := make([]bool, numOfVertices)

	queue = append(queue, s)
	visited[s] = true

	paths[s] = 0

	// standard bfs queue
	for len(queue) > 0 {
		source := queue[0]
		queue = queue[1:]

		for dest := 0; dest < numOfVertices; dest++ {
			edgeWeight := graph[source][dest]

			if !visited[dest] && edgeWeight > 0 {
				paths[dest] = source

				if dest == t {
					return paths, true
				}

				queue = append(queue, dest)
				visited[dest] = true
			}
		}
	}

	return paths, false
}

func fordFulkerson(graph [][]int, s, t int) (maxFlow int) {
	for {
		paths, pathExists := bfs(graph, s, t)

		if !pathExists {
			break
		}

		bottleNeck := math.MaxInt
		for dest := t; dest != s; dest = paths[dest] {
			src := paths[dest]
			edgeWeight := graph[src][dest]

			bottleNeck = int(math.Min(float64(bottleNeck), float64(edgeWeight)))
		}

		for dest := t; dest != s; dest = paths[dest] {
			src := paths[dest]
			graph[src][dest] -= bottleNeck
			graph[dest][src] += bottleNeck
		}

		maxFlow += bottleNeck
	}

	return maxFlow
}

func TestFordFulkerson() {
	graph := [][]int{
		{0, 8, 0, 0, 3, 0},
		{0, 0, 9, 0, 0, 0},
		{0, 0, 0, 0, 7, 2},
		{0, 0, 0, 0, 0, 5},
		{0, 0, 7, 4, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	fmt.Println(fordFulkerson(graph, 0, 5))
}
