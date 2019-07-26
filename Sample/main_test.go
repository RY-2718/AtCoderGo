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
		In: `3 2
1 2
5 5
-2 8
`,
		Out: "1\n",
	},
	{
		In: `3 4
-3 7 8 2
-12 1 10 2
-2 8 9 3
`,
		Out: "2\n",
	},
	{
		In: `5 1
1
2
3
4
5
`,
		Out: "10\n",
	},
}

func TestSolve(t *testing.T) {
	for i, v := range TestCases {
		r := strings.NewReader(v.In)
		w := bytes.NewBufferString("")
		Solve(r, w)
		if w.String() != v.Out {
			t.Errorf("Test case %v is incorrect\n", i+1)
			t.Errorf("Your Answer:  %v\n", w.String())
			t.Errorf("Right Answer: %v\n", v.Out)
		}
	}
}
