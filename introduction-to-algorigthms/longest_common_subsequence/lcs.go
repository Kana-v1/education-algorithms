package longestcommonsubsequence

import "fmt"

func Longest(subseq, str string) int {
	table := make([][]int, len(subseq)+1)
	for i := range table {
		table[i] = make([]int, len(str)+1)
	}

	for i := 0; i <= len(subseq); i++ {
		for j := 0; j <= len(str); j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
				continue
			}

			if subseq[i-1] == str[j-1] {
				table[i][j] = 1 + table[i-1][j-1]

				fmt.Print(string(subseq[i-1]))
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	fmt.Println()
 
	return table[len(subseq)][len(str)]
}

func max(nums ...int) int {
	max := nums[0]

	for i, el := range nums {
		if max < el {
			max = nums[i]
		}
	}

	return max
}
