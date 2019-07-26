package main

import (
	"fmt"
	"io"
	"math"
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

	count := 0

	for i := 0; i < n-1; i++ {
		src := p[i]
		for j := i + 1; j < n; j++ {
			dst := p[j]
			distSq := distanceSquare(src, dst)
			dist := int(math.Sqrt(float64(distSq)))
			// distが小数だったらintへのキャストで切り捨てられてdist^2 != distSqになるはず！多分！
			if dist*dist == distSq {
				count++
			}
		}
	}
	_, _ = fmt.Fprintf(stdout, "%d\n", count)
}

func distanceSquare(src, dst []int) int {
	dimension := len(src)
	ans := 0
	for i := 0; i < dimension; i++ {
		ans += (src[i] - dst[i]) * (src[i] - dst[i])
	}

	return ans
}

func main() {
	Solve(os.Stdin, os.Stdout)
	return
}
