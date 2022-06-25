/*
 * @lc app=leetcode id=1791 lang=golang
 *
 * [1791] Find Center of Star Graph
 */

// @lc code=start
package main 

func findCenter(edges [][]int) int {
	nodes := make(map[int]int)

	for _, edge := range edges {
		for _, node := range edge {
			if val, ok := nodes[node]; ok {
				nodes[node] = val + 1
			} else {
				nodes[node] = 1
			}			
		}
	}

	max := -1
	center := 0
	for node, occursTime := range nodes {
		if max < occursTime {
			max = occursTime
			center  = node
		}
	}

	return center
}
// @lc code=end

