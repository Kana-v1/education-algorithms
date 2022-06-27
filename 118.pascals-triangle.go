/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
package main
func generate(numRows int) [][]int {
    res := make([][]int, 0)
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})

	for i := 2; i <= numRows; i++ {
		row := make([]int, 0)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				row = append(row, 1)
				continue
			}

			row = append(row, res[i-2][j-1]+res[i-2][j])
		}

		res = append(res, row)
	}

	return res
}
// @lc code=end

