package mergesort

func Merge_sort(A []int) []int {
	if len(A) == 1 {
		return A
	}

	mid := len(A) / 2
	first := Merge_sort(A[:mid])
	second := Merge_sort(A[mid:])

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))

	i := 0
	j := 0
	k := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}

	for ; i < len(a); i++ {
		c[k] = a[i]
		k++
	}

	for ; j < len(b); j++ {
		c[k] = b[j]
		k++
	}

	return c
}
