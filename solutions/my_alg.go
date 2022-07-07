package solutions

import "math"

func Alg(days []int, costs []int) int {
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

		w := 0
		m := 0

		if i >= 7 {
			w = i - 7

			if i >= 30 {
				m = i - 30
			}
		}

		prices[i] = int(math.Min(float64(prices[i]), float64(costs[1]+prices[w])))
		prices[i] = int(math.Min(float64(prices[i]), float64(costs[2]+prices[m])))
	}

	return prices[len(prices)-1]
}
