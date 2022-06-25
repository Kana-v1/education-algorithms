/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	numsCopy := make([]int, 0, len(score))
	numsCopy = append(numsCopy, score...)

	sort.Ints(numsCopy)

	buf := make(map[int]string)

	for i := 1; i <= 3; i++ {
		if i > len(numsCopy) {
			break
		}

		medal := ""

		switch i {
		case 1:
			medal = "Gold Medal"
		case 2:
			medal = "Silver Medal"
		case 3:
			medal = "Bronze Medal"
		}

		buf[numsCopy[len(numsCopy)-i]] = medal
	}

	for i := len(numsCopy) - 4; i >= 0; i-- {
		buf[numsCopy[i]] = fmt.Sprint(len(numsCopy) - i)
	}

	output := make([]string, 0, len(score))

	for i := range score {
		output = append(output, buf[score[i]])
	}

	return output

}

// @lc code=end
