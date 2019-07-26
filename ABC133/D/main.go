package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(stdin io.Reader, stdout io.Writer) {
	sc := bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	// a_i = 1/2(x_i + x_(i+1))

	// n = 5のとき
	// sumEvens = [0, a_0, a_0 + a_2, a_0 + a_2 + a_4]
	// sumOdds  = [0, a_1, a_1 + a_3]
	// sum = a_0 + a_1 + ... + a_4 = sumEvens[3] + sumOdds[2] = sumEvens[5/2+1] + sumOdds[5/2]
	sumEvens := make([]int, n/2+2)
	sumOdds := make([]int, n/2+1)
	// それぞれの0番目は番兵
	sumEvens[0] = 0
	sumOdds[0] = 0

	for i := 0; i < n; i++ {
		sc.Scan()
		a, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}

		// 偶奇で分けて累積和を取る
		if i%2 == 0 {
			sumEvens[i/2+1] = sumEvens[i/2] + a
		} else {
			sumOdds[i/2+1] = sumOdds[i/2] + a
		}
	}

	sum := sumEvens[n/2+1] + sumOdds[n/2]

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			others := 2 * (sumOdds[i/2] + sumEvens[n/2+1] - sumEvens[i/2+1])
			_, _ = fmt.Fprintf(stdout, "%d ", sum-others)
		} else if i != n-1 {
			others := 2 * (sumEvens[i/2] + sumOdds[n/2] - sumOdds[i/2])
			_, _ = fmt.Fprintf(stdout, "%d ", sum-others)
		} else {
			// 出力のために場合分け
			// i = n - 1, iは偶数なので n/2 = i/2
			// そのため sumOdds[n/2] - sumOdds[i/2] = 0
			// others := 2 * (sumEvens[i/2] + sumOdds[n/2] - sumOdds[i/2])
			others := 2 * sumEvens[i/2]
			_, _ = fmt.Fprintf(stdout, "%d\n", sum-others)
		}
	}

	return
}

func main() {
	Solve(os.Stdin, os.Stdout)
	return
}
