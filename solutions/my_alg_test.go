package solutions_test

import (
	"algs/solutions"
	"testing"
)

type TestCase struct {
	triangle [][]int
	expected int
}

func TestAlg(t *testing.T) {
	testcases := []TestCase{
		{
			triangle: [][]int{{-7},{-2,1},{-5,-5,9},{-4,-5,4,4},{-6,-6,2,-1,-5},{3,7,8,-3,7,-9},{-9,-1,-9,6,9,0,7},{-7,0,-6,-8,7,1,-4,9},{-3,2,-6,-9,-7,-6,-9,4,0},{-8,-6,-3,-9,-2,-6,7,-5,0,7},{-9,-1,-2,4,-2,4,4,-1,2,-5,5},{1,1,-6,1,-2,-4,4,-2,6,-6,0,6},{-3,-3,-6,-2,-6,-2,7,-9,-5,-7,-5,5,1}},
			expected: -63,
		},
		{
			triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}},
			expected: -1,
		},
		{
			triangle: [][]int{{2}, {5, 7}, {9, 7, 3}, {10, 1, 6, 2}},
			expected: 14,
		},
		{
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expected: 11,
		},
		{
			triangle: [][]int{{-10}},
			expected: -10,
		},
	}

	for i, testcase := range testcases {
		res := solutions.Alg(testcase.triangle)

		if res != testcase.expected {
			t.Errorf("testcase %v. Expected: %v, actual: %v", i+1, testcase.expected, res)
		}
	}
}
