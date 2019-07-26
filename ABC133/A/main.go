package main

import (
	"fmt"
	"io"
	"os"
)

func Solve(stdin io.Reader, stdout io.Writer) {
	var n, a, b int
	_, _ = fmt.Fscanf(stdin, "%d %d %v", &n, &a, &b)
	train := n * a
	if train > b {
		_, _ = fmt.Fprintf(stdout, "%d\n", b)
		return
	}
	_, _ = fmt.Fprintf(stdout, "%d\n", train)
}

func main() {
	Solve(os.Stdin, os.Stdout)
	return
}
