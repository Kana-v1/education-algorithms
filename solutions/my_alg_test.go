package solutions_test

import (
	"algs/solutions"
	"testing"
)

type TestCase struct {
	expression string
	expected   int
}

func TestAlg(t *testing.T) {
	testcases := []TestCase{
		{
			expression: " -1 - 3 * 4 + 4 /2",
			expected:   -11,
		},
		{
			expression: "-1 - 1 - 1 - 2 * 1",
			expected:   -5,
		},
		{
			expression: "4 + 8 - 2 * 10 - 100",
			expected:   -108,
		},
		{
			expression: "2*3+4",
			expected:   10,
		},
		{
			expression: "  3/  2*  10  * 3-    4+   2",
			expected:   28,
		},
		{
			expression: "1 + 1 + 1",
			expected:   3,
		},
		{
			expression: "    30",
			expected:   30,
		},
		{
			expression: "1",
			expected:   1,
		},
		{
			expression: "3/2*10-4",
			expected:   6,
		},
		{
			expression: "3+5/2 ",
			expected:   5,
		},
		{
			expression: "3+2*2",
			expected:   7,
		},
		{
			expression: " 3/ 2",
			expected:   1,
		},
		{
			expression: " 3+5/2 ",
			expected:   5,
		},
	}

	for i, testcase := range testcases {
		res := solutions.Alg(testcase.expression)

		if res != testcase.expected {
			t.Errorf("testcase %v. Expected: %v, actual: %v", i+1, testcase.expected, res)
		}
	}
}
