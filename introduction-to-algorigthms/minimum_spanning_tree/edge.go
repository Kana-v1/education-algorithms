package minimumspanningtree

type Edge struct {
	Src  int
	Dest int
	Cost int
}

type Edges []Edge

// we need these 3 methods below for sort.Sort
func (e Edges) Len() int { return len(e) }

func (e Edges) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func (e Edges) Less(i, j int) bool { return e[i].Cost < e[j].Cost }

func MakeEdged(graph [][]int) Edges {
	edges := make([]Edge, 0)

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			// 0 means there is no edges between vertices
			if graph[i][j] != 0 {
				edges = append(edges, Edge{Src: i, Dest: j, Cost: graph[i][j]})
			}
		}
	}

	return edges
}
