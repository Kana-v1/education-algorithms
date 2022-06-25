package solutions

import (
	"fmt"
	"math"
)

func Alg(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""

	n := int(math.Abs(float64(num)))

	for n/7 != 0 {
		res = fmt.Sprintf("%v%v", n%7, res)

		n /= 7
	}

	res = fmt.Sprintf("%v%v", n, res)

	if num > 0 {
		return res
	}
	
	return fmt.Sprintf("-%v", res)
}
