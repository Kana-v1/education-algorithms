/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0

	excluded := make(map[int]bool)

	indexes := make([]int, 0)

	for i := range grid {
		indexes = append(indexes, grid[i]...)
	}

	for i := range indexes {
		if _, ok := excluded[i]; !ok && indexes[i] == 1 {
			buf := getElsAround(indexes, i, len(grid[0]), &excluded)

			if buf > res {
				res = buf
			}
		}
	}

	return res
}

func getElsAround(indexes []int, i, rowLen int, exclude *map[int]bool) int {
	(*exclude)[i] = true
	counter := 1

	// right
	if _, ok := (*exclude)[i+1]; !ok && i+1 < len(indexes) && indexes[i+1] == 1 {
		counter += getElsAround(indexes, i+1, rowLen, exclude)
	}

	// left
	if _, ok := (*exclude)[i-1]; !ok && i > 1 && indexes[i-1] == 1 {
		counter += getElsAround(indexes, i-1, rowLen, exclude)
	}

	// down
	if _, ok := (*exclude)[i+rowLen]; !ok && i+rowLen < len(indexes) && indexes[i+rowLen] == 1 {
		counter += getElsAround(indexes, i+rowLen, rowLen, exclude)
	}

	// up
	if _, ok := (*exclude)[i-rowLen]; !ok && i-rowLen >= 0 && indexes[i-rowLen] == 1 {
		counter += getElsAround(indexes, i-rowLen, rowLen, exclude)
	}

	return counter
}

// @lc code=end
