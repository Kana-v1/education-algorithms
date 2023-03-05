package augmentedblacktree

type RBTree struct {
	root *RBTreeNode
}

type RBTreeNode struct {
	left   *RBTreeNode
	right  *RBTreeNode
	parent *RBTreeNode
	key    int
	size   int
}

// select the ith smallest element in the tree
func (node *RBTreeNode) Select(i int) *RBTreeNode {
	r := node.left.size + 1 // rank of x within the subtree rooted at tree

	if i == r {
		return node
	} else {
		if i < r {
			return node.left.Select(i)
		}

		return node.right.Select(i - r)
	}
}

// get the position of node in the linear order of the tree
func (tree *RBTree) Rank(node *RBTreeNode) int {
	r := node.left.size + 1

	for node != tree.root {
		if node == node.parent.right {
			r += node.parent.left.size + 1
		}

		node = node.parent
	}

	return r
}
