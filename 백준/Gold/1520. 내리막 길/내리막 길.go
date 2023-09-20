package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 500
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	n, m int
	arr  [MAX][MAX]int
	dp   [MAX][MAX]int
	dy   = [4]int{0, 0, 1, -1}
	dx   = [4]int{1, -1, 0, 0}
)

func solve(y, x int) int {
	if dp[y][x] != -1 {
		return dp[y][x]
	}
	dp[y][x] = 0
	for i := 0; i < 4; i++ {
		py := y + dy[i]
		px := x + dx[i]

		if py < 0 || px < 0 || px >= m || py >= n {
			continue
		}
		if arr[y][x] <= arr[py][px] {
			continue
		}

		dp[y][x] += solve(py, px)
	}

	return dp[y][x]
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &arr[i][j])
			dp[i][j] = -1
		}
	}
	dp[n-1][m-1] = 1
	answer := solve(0, 0)
	fmt.Fprintln(w, answer)
}
