package main

import (
	"bytes"
	"strings"
	"testing"
)

type TestCase struct {
	In  string
	Out string
}

var TestCases = []TestCase{
	{
		In: `3
2 2 4
`,
		Out: "4 0 4\n",
	},
	{
		In: `5
3 8 7 5 5
`,
		Out: "2 4 12 2 8\n",
	},
	{
		In: `3
1000000000 1000000000 0
`,
		Out: "0 2000000000 0\n",
	},
}

func TestSolve(t *testing.T) {
	for i, v := range TestCases {
		r := strings.NewReader(v.In)
		w := bytes.NewBufferString("")
		Solve(r, w)
		if w.String() != v.Out {
			t.Errorf("Test case %v is incorrect\n", i+1)
			t.Errorf(" Your Answer: %v\n", w.String())
			t.Errorf("Right Answer: %v\n", v.Out)
		}
	}
}
