package solutions_test

import (
	"algs/solutions"
	"testing"
)

type TestCase struct {
	board  [][]byte
	word   string
	output bool
}

func TestAlg(t *testing.T) {
	testcases := []TestCase{
		{
			board: [][]byte{
				{'C', 'A', 'A'},
				{'A', 'A', 'A'},
				{'B', 'C', 'D'},
			},
			word:   "AAB",
			output: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "SEE",
			output: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "ABCCED",
			output: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "ABCB",
			output: false,
		},
	}

	for i, testcase := range testcases {
		res := solutions.Alg(testcase.board, testcase.word)

		if res != testcase.output {
			t.Errorf("testcase %v. Expected: %v, actual: %v", i+1, testcase.output, res)
		}
	}
}
