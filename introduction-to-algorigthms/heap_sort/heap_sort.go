package heapsort

func Heapsort(array []int) []int {
	build_max_heap(&array)
	size := len(array)

	for i := size - 1; i > 0; i-- {
		tmp := array[i]
		array[i] = array[0]
		array[0] = tmp
		size--
		max_heapify(&array, size, 0)
	}

	return array
}

func max_heapify(tree *[]int, size, i int) {
	left := 2 * i
	right := 2*i + 1
	largest := i

	if left < size && (*tree)[left] > (*tree)[largest] {
		largest = left
	}

	if right < size && (*tree)[right] > (*tree)[largest] {
		largest = right
	}

	if largest != i {
		tmp := (*tree)[i]
		(*tree)[i] = (*tree)[largest]
		(*tree)[largest] = tmp

		max_heapify(tree, size, largest)
	}
}

func build_max_heap(array *[]int) {
	for i := len(*array)/2; i >= 0; i-- {
		max_heapify(array, len(*array), i)
	}
}
