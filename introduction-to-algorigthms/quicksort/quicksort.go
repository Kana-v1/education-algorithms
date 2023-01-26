package quicksort

func Quicksort(arr *[]int, floor, ceil int) {
	if len(*arr) > 1 {
		j := partition(arr, floor, ceil)

		if j > floor {
			Quicksort(arr, floor, j-1)
		}

		if j < ceil {
			Quicksort(arr, j+1, ceil)
		}
	}
}

func partition(arr *[]int, floor, ceil int) int {
	pivot := (*arr)[floor]
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

	(*arr)[floor], (*arr)[j] = (*arr)[j], (*arr)[floor]

	return j
}
