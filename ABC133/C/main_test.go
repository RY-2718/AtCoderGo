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
		In:  `2020 2040\n`,
		Out: "2\n",
	},
	{
		In:  `4 5\n`,
		Out: "20\n",
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
