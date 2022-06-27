package solutions_test

import (
	"algs/solutions"
	"testing"
)

type TestCase struct {
	output int
	input  [][]int
}

func TestAlg(t *testing.T) {
	input_1 := make([][]int, 0)
	intput_2 := make([][]int, 0)
	intput_3 := make([][]int, 0)
	intput_4 := make([][]int, 0)

	input_1 = append(input_1, []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	)
	intput_2 = append(intput_2, []int{0, 0, 0, 0, 0, 0, 0, 0})

	intput_3 = append(intput_3, []int{1, 1, 1, 1}, []int{1, 1, 1, 1})

	intput_4 = append(intput_4, []int{1, 1, 0, 1, 1}, []int{1, 0, 0, 0, 0}, []int{0, 0, 0, 0, 1}, []int{1, 1, 0, 1, 1})

	testcases := []TestCase{
		{
			output: 3,
			input:  intput_4,
		},
		{
			output: 6,
			input:  input_1,
		},
		{
			output: 8,
			input:  intput_3,
		},
		{
			output: 0,
			input:  intput_2,
		},
	}

	for _, testcase := range testcases {
		res := solutions.Alg(testcase.input)
		if res != testcase.output {
			t.Errorf("should be: %v, actual: %v", testcase.output, res)
		}
	}
}
