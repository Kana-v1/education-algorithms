/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode = nil
	for _, num := range nums {
		root = insert(root, num)
	}

	return root
}

func insert(root *TreeNode, val int) *TreeNode {
	tmp := &TreeNode{
		Val: val,
	}

	if root == nil {
		root = tmp
		return root
	}

	curNode := root

	for curNode != nil && curNode.Val != val {
		if val > curNode.Val {
			if curNode.Right == nil {
				curNode.Right = tmp
				break
			}

			curNode = curNode.Right
		} else {
			if curNode.Left == nil {
				curNode.Left = tmp
				break
			}
			curNode = curNode.Left
		}
	}

	return traversal(root)
}

func traversal(node *TreeNode) *TreeNode {
	if node.Left != nil {
		node.Left = traversal(node.Left)
	}

	if node.Right != nil {
		node.Right = traversal(node.Right)
	}

	balanceFactor := getBalanceFactor(node)
	if balanceFactor > 1 || balanceFactor < -1 {
		switch getImbalanceDirection(node, balanceFactor) {
		case "LL":
			node = rightRotation(node)
		case "LR":
			node = doubleRotation(node, true)
		case "RR":
			node = leftRotation(node)
		case "RL":
			node = doubleRotation(node, false)
		}
	}

	return node
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := 1
	rightHeight := 1

	if node.Left != nil {
		leftHeight += getHeight(node.Left)
	}

	if node.Right != nil {
		rightHeight += getHeight(node.Right)
	}

	if leftHeight > rightHeight {
		return leftHeight
	}

	return rightHeight
}

func getBalanceFactor(node *TreeNode) int {
	return getHeight(node.Left) - getHeight(node.Right)
}

func rightRotation(node *TreeNode) *TreeNode {
	a_node := node
	b_node := node.Left
	b_right_node := b_node.Right

	a_node.Left = b_right_node
	b_node.Right = a_node

	return b_node
}

func leftRotation(node *TreeNode) *TreeNode {
	a_node := node
	b_node := node.Right
	b_left_node := b_node.Left

	a_node.Right = b_left_node
	b_node.Left = a_node

	return b_node
}

func doubleRotation(node *TreeNode, isLr bool) *TreeNode {
	a_node := node

	if isLr {
		b_node := a_node.Left
		c_node := b_node.Right
		c_r_node := c_node.Right
		c_l_node := c_node.Left

		b_node.Right = c_l_node
		a_node.Left = c_r_node
		c_node.Left = b_node
		c_node.Right = a_node

		return c_node
	}

	b_node := a_node.Right
	c_node := b_node.Left
	c_r_node := c_node.Right
	c_l_node := c_node.Left

	b_node.Left = c_r_node
	a_node.Right = c_l_node
	c_node.Right = b_node
	c_node.Left = a_node

	return c_node
}

func getImbalanceDirection(node *TreeNode, balanceFactor int) string {
	if balanceFactor > 1 {
		if getBalanceFactor(node.Left) > 0 {
			return "LL"
		}

		return "LR"
	}

	if balanceFactor < -1 {
		if getBalanceFactor(node.Right) < 0 {
			return "RR"
		}

		return "RL"
	}

	panic("balance factor is invalid")
}

// @lc code=end
