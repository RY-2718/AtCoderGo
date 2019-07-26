package main

import (
	"fmt"
	"io"
	"os"
)

func Solve(stdin io.Reader, stdout io.Writer) {
	var n, d int
	_, _ = fmt.Fscanf(stdin, "%d %d\n", &n, &d)

	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, d)
		var tmp int
		for j := 0; j < d; j++ {
			_, _ = fmt.Fscan(stdin, &tmp)
			p[i][j] = tmp
		}
	}

	_, _ = fmt.Fprintf(stdout, "\n")
}

func main() {
	Solve(os.Stdin, os.Stdout)
	return
}
