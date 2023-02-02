package binarytree

import "fmt"

type BinaryTree struct {
	value int
	right *BinaryTree
	left  *BinaryTree
}

func (tree *BinaryTree) Print() {
	if tree == nil {
		return
	}

	fmt.Printf("%v ", tree.value)
	tree.left.Print()
	tree.right.Print()
}


