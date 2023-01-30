package randomizedselect

import "math/rand"

func RandomizedSelect(arr []int, floor, ceil, ith int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	j := partition(&arr, floor, ceil)

	if j == ith {
		return arr[j]
	}

	if j < ith {
		return RandomizedSelect(arr, j+1, ceil, ith)
	}

	return RandomizedSelect(arr, floor, j-1, ith)
}

func partition(arr *[]int, floor, ceil int) int {
	if ceil == floor {
		return floor
	}

	rndIdx := rand.Intn(ceil-floor) + floor
	pivot := (*arr)[rndIdx]
	i := floor
	j := ceil

	for i < j {
		for (*arr)[i] <= pivot && i < ceil {
			i++
		}

		for (*arr)[j] > pivot && j > floor {
			j--
		}

		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[rndIdx], (*arr)[j] = (*arr)[j], (*arr)[rndIdx]

	return j
}
