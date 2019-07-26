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
		In:  "4 2 9\n",
		Out: "8\n",
	},
	{
		In:  "4 2 7\n",
		Out: "7\n",
	},
	{
		In:  "4 2 8\n",
		Out: "8\n",
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
