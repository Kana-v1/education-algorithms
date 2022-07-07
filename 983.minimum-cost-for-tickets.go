/*
 * @lc app=leetcode id=983 lang=golang
 *
 * [983] Minimum Cost For Tickets
 */

// @lc code=start
package main

import "math"

func mincostTickets(days []int, costs []int) int {
	isTravelled := make([]bool, days[len(days)-1]+1)

	for _, day := range days {
		isTravelled[day] = true
	}

	prices := make([]int, days[len(days)-1]+1)

	for i := 1; i < len(prices); i++ {
		if !isTravelled[i] {
			prices[i] = prices[i-1]
			continue
		}

		prices[i] = costs[0] + prices[i-1]

		if i >= 7 {
			prices[i] = int(math.Min(float64(prices[i]), float64(costs[1]+prices[i-7])))

			if i >= 30 {
				prices[i] = int(math.Min(float64(prices[i]), float64(costs[2]+prices[i-30])))
			} else {
				prices[i] = int(math.Min(float64(prices[i]), float64(costs[2]+prices[0])))
			}
		} else {
			prices[i] = int(math.Min(float64(prices[i]), float64(costs[1]+prices[0])))
		}
	}

	return prices[len(prices)-1]
}

// @lc code=end
