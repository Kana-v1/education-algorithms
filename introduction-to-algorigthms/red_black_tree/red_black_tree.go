package redblacktree

import "fmt"

type color byte

const red color = 1
const black color = 0

type node struct {
	color  color // 1 - red  0 - black
	key    int
	left   *node
	right  *node
	parent *node
}

type redBlackTree struct {
	root *node
}

func (tree *redBlackTree) String() string {
	res := ""

	if tree.root != nil {
		tree.root.output("red-black tree:\n", true, &res)
	}

	return res
}

func CreateTree(arr []int) *redBlackTree {
	tree := new(redBlackTree)

	for _, el := range arr {
		tree.Insert(el)
	}

	return tree
}

func (tree *redBlackTree) Insert(v int) {
	x := tree.root

	var newNode *node = &node{key: v, color: red}
	var y *node = nil

	for x != nil { // descent until reach the sentinel
		y = x
		if v < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	if y == nil { // the tree was empty
		tree.root = newNode
	} else {
		if newNode.key < y.key {
			y.left = newNode
		} else {
			y.right = newNode
		}
	}

	tree.insertBalance(newNode)
}

func (tree *redBlackTree) Delete(v int) {
	y := tree.find(v)
	x := new(node)
	origColor := y.color

	if y.left == nil {
		x = y.right
		tree.replace(y, y.right)
	} else {
		if y.right == nil {
			x = y.left
			tree.replace(y, y.left)
		} else {
			min := tree.Min(y.right)
			origColor = min.color
			x = min.right
			if min != y.right {
				tree.replace(min, min.right)
				min.right = y.right
				min.right.parent = min
			} else {
				x.parent = min
				tree.replace(y, min)
				min.left = y.left
				min.left.parent = min
				min.color = y.color
			}
		}
	}

	if origColor == black {
		tree.deleteBalance(x)
	}
}

func (tree *redBlackTree) Min(node *node) *node {
	x := tree.root
	if node != nil {
		x = node
	}

	res := x

	for x != nil {
		res = x
		x = x.left
	}

	return res
}

func (tree *redBlackTree) replace(old, new *node) {
	if old.parent == nil {
		tree.root = new
		new.parent = nil
		return
	}

	if old == old.parent.left {
		old.parent.left = new
	} else {
		old.parent.right = new
	}

	new.parent = old.parent
}

func (tree *redBlackTree) find(v int) *node {
	x := tree.root
	res := new(node)

	for x != nil {
		if x.key == v {
			res.color = x.color
			res.parent = x.parent
			res.key = x.key
		}

		if v < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	return res
}

func (tree *redBlackTree) deleteBalance(node *node) {
	for node != tree.root && node.color == black {
		if node == node.parent.left {
			sibling := node.parent.right

			if sibling.color == red {
				sibling.color = black
				node.parent.color = red
				tree.left_rotate(node.parent)
				sibling = node.parent.right
			}

			if sibling.left.color == black && sibling.right.color == black {
				sibling.color = red
				node = node.parent
			} else {
				if sibling.right.color == black {
					sibling.left.color = black
					sibling.color = red
					tree.right_rotate(sibling)
					sibling = node.parent.right
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.right.color = black
				tree.left_rotate(node.parent)
				node = tree.root
			}
		} else {
			sibling := node.parent.left
			if sibling.color == red {
				sibling.color = black
				node.parent.color = red
				tree.right_rotate(node.parent)
				sibling = node.parent.left
			}

			if sibling.right.color == black && sibling.left.color == black {
				sibling.color = red
				node = node.parent
			} else {
				if sibling.left.color == black {
					sibling.right.color = black
					sibling.color = red
					tree.left_rotate(sibling)
					sibling = node.parent.left
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.left.color = black
				tree.right_rotate(node.parent)
				node = tree.root
			}
		}
	}

	node.color = black
}

func (tree *redBlackTree) insertBalance(node *node) {
	for node != tree.root && node.parent.color == red {
		switch node.parent == node.parent.parent.left {
		case true: // is node's parent a left child?
			uncle := node.parent.parent.right

			if uncle == nil {
				node.parent.color = red
				node.parent.parent.color = red
				tree.left_rotate(node.parent)
				tree.right_rotate(node)
				break
			}

			if uncle.color == red {
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.left_rotate(node)
				}

				node.parent.color = black
				node.parent.parent.color = red
				tree.right_rotate(node.parent.parent)
			}

		case false: // same as before but "right" and "left" rotate
			uncle := node.parent.parent.left

			if uncle == nil {
				node.parent.color = red
				node.parent.parent.color = red
				tree.right_rotate(node.parent)
				tree.left_rotate(node)
				break
			}

			if uncle.color == red {
				node.parent.color = black
				uncle.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.left_rotate(node)
				}

				node.parent.color = black
				node.parent.parent.color = red
				tree.right_rotate(node.parent.parent)
			}

		}
	}

	tree.root.color = black
}

func (tree *redBlackTree) left_rotate(x *node) {
	y := x.right
	x.right = y.left

	if x.right != nil {
		x.right.parent = x
	}

	y.parent = x.parent
	y.left = x
	x.parent = y

	if y.parent == nil {
		tree.root = y
	} else {
		if y.parent.left == x {
			y.parent.left = y
		} else {
			y.parent.right = y
		}
	}
}

func (tree *redBlackTree) right_rotate(y *node) {
	x := y.left
	y.left = x.right

	if y.left != nil {
		y.left.parent = y
	}

	x.parent = y.parent
	x.right = y
	y.parent = x

	if x.parent == nil {
		tree.root = x
	} else {
		if x.parent.right == y {
			x.parent.right = x
		} else {
			x.parent.left = x
		}
	}

}

func (node *node) output(prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix = "|	    "
		} else {
			newPrefix = "		"
		}

		node.right.output(newPrefix, false, str)
	}

	*str += prefix

	if isTail {
		*str += "└──"
	} else {
		*str += "┌──"
	}

	*str = fmt.Sprintf("%s%v\n", *str, node.key)

	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix = "		"
		} else {
			newPrefix = "|	    "
		}

		node.left.output(newPrefix, false, str)
	}

}
