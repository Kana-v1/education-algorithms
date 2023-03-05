package disjointdatastructure

type node struct {
	value int
	next  *node
	set   *set
}

type set struct {
	head *node
	tail *node
}

type DisjointList struct {
	nodeAddress map[int]*node
}

func (dl *DisjointList) MakeSet(v int) {
	if dl.nodeAddress == nil {
		dl.nodeAddress = make(map[int]*node)
	}

	newSet := &set{
		head: &node{
			value: v,
		},
	}

	dl.nodeAddress[v] = newSet.representative()

	newSet.head.set = newSet
	newSet.tail = newSet.head
}

func (dl *DisjointList) Find(key int) (representative *node) {
	return dl.nodeAddress[key]
}

func (dl *DisjointList) Union(remained, joined *set) {
	cur := joined.representative()

	for cur != nil {
		cur.set = remained
		cur = cur.next
		dl.nodeAddress[cur.value] = cur.set.representative()
	}

	remained.head = joined.tail

	joined = nil
}

func (s *set) representative() *node {
	return s.head
}
