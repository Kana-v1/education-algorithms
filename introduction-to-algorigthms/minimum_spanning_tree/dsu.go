package minimumspanningtree

// https://cp-algorithms.com/data_structures/disjoint_set_union.html

const noParent = -1

type Subset struct {
	Parent int
}

// get nodes parent or source node
func Find(subsets []Subset, value int) int {
	if subsets[value].Parent == noParent {
		return value
	}

	return Find(subsets, subsets[value].Parent)
}

func Union(subsets []Subset, x, y int) {
	xroot := Find(subsets, x)
	yroot := Find(subsets, y)

	subsets[xroot].Parent = yroot
}
