package minimumspanningtree

import (
	"fmt"
	"sort"
)

type Kruskals struct {
	Vertices int
	Edges    Edges
}

func (k Kruskals) GetTree() {
	sort.Sort(k.Edges)

	subsets := make([]Subset, k.Vertices)

	for i := 0; i < k.Vertices; i++ {
		subsets[i] = Subset{Parent: noParent}
	}

	result := make(Edges, k.Vertices-1)
	savedVertices := 0
	i := 0
	

	for savedVertices < k.Vertices-1 {
		if i >= len(k.Edges) {
			break
		}

		next := k.Edges[i]
		x := Find(subsets, next.Src)
		y := Find(subsets, next.Dest)

		if x != y {
			result[savedVertices] = next
			Union(subsets, x, y)
			savedVertices++
		}

		i++
	}

	for _, anEdge := range result {
		fmt.Printf("%d - %d | %d\n", anEdge.Src, anEdge.Dest, anEdge.Cost)
	}
}
