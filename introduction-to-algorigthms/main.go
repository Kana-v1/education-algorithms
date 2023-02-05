package main

import (
	"fmt"
	longestcommonsubsequence "introduction-to-algorithms/longest_common_subsequence"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr := []int{9, 3, 7, 5, 6, 4, 8, 2}
	// arr := []int{1, 7, 4, 3, 5, 10}

	fmt.Println(arr)

	// fmt.Println(mergesort.Merge_sort(arr))
	// quicksort.Quicksort(&arr, 0, len(arr)-1)
	// fmt.Println(arr)

	// fmt.Println(rnd.RandomizedSelect(arr, 0, len(arr)-1, 2))

	str := "abcd"
	substr := "bd"

	fmt.Println()
	fmt.Printf("lcs = %v" ,longestcommonsubsequence.Longest(substr, str))

}
