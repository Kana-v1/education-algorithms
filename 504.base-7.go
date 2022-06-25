/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""

	n := int(math.Abs(float64(num)))

	for n/7 != 0 {
		res = fmt.Sprintf("%v%v", n%7, res)

		n /= 7
	}

	res = fmt.Sprintf("%v%v", n, res)

	if num > 0 {
		return res
	}
	return fmt.Sprintf("-%v", res)
}

// @lc code=end
