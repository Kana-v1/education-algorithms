package insertion_sort

func insertion_sort(A []int) {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1

		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j = j - 1
		}

		A[j+1] = key
	}
}
