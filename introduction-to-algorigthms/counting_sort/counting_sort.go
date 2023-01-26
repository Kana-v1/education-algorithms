package countingsort

// it's nlgn, however this alg assumes those unput ints are in the small range
func CountringSort(arr []int) []int {
	res := make([]int, len(arr))

	max := arr[0]
	for _, a := range arr {
		if a >= 0 {
			if max < a {
				max = a
			}
		}
	}

	C := make([]int, max+1)

	for i := range arr {
		C[arr[i]]++
	}

	for i := 1; i < len(C); i++ {
		C[i] += C[i-1]
	}

	for i := 0; i < len(arr); i++ {
		el := arr[i]
		index := C[el] - 1

		res[index] = el
		C[el]--
	}

	return res
}
