package stringmatching

import "fmt"

const prime = 101

func hash(str string, end int) int {
	h := 0
	for i := 0; i <= end; i++ {
		h += int(str[i]) * pow(prime, i)
	}

	return h
}

// calculates x raised to power y
func pow(x, y int) int {
	if y == 0 {
		return 1
	}

	if y%2 == 0 {
		return pow(x*x, y/2)
	}

	return x * pow(x*x, y/2)
}

func recalculateHash(str string, oldIndex, newIndex, oldHash, patternLen int) int {
	newHash := oldHash - int(str[oldIndex])
	newHash /= prime
	newHash += int(str[newIndex]) * pow(prime, patternLen-1)

	return newHash
}

func checkEqual(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func search(text, pattern string) bool {
	m := len(pattern)
	n := len(text)
	patternHash := hash(pattern, m-1)
	textHash := hash(text, m-1)

	for i := 0; i < n-m+1; i++ {
		if patternHash == textHash && checkEqual(text[i:i+m], pattern) {
			return true
		}
		if i < n-m+1 {
			textHash = recalculateHash(text, i, i+m, textHash, m)
		}
	}
	return false
}

func Test() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	fmt.Println(search(text, pattern))
}
