package main

import (
	"fmt"
	countingsort "introduction-to-algorithms/counting_sort"
)

func main() {
	arr := []int{9, 3, 7, 5, 3, 6, 4, 3, 8, 2}
	// arr := []int{1, 7, 4, 3, 5, 10}

	fmt.Println(arr)

	// fmt.Println(mergesort.Merge_sort(arr))
	// quicksort.Quicksort(&arr, 0, len(arr)-1)

	fmt.Println(countingsort.CountringSort(arr))
}
