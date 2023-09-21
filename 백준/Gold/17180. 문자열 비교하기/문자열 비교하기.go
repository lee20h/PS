package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 301
)

var (
	r             = bufio.NewReader(os.Stdin)
	w             = bufio.NewWriter(os.Stdout)
	n, m          int
	first, second string
	dp            [MAX][MAX]int
)

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func min(x int, y ...int) (m int) {
	m = x
	for _, v := range y {
		if m > v {
			m = v
		}
	}
	return
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n, &m)

	fmt.Fscan(r, &first, &second)
	for i := 1; i <= n; i++ {
		dp[i][0] = 2e10
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = 2e10
	}

	for i, str1_value := range []rune(first) {
		for j, str2_value := range []rune(second) {
			dp[i+1][j+1] = abs(int(str1_value), int(str2_value))
			dp[i+1][j+1] += min(dp[i][j], dp[i+1][j], dp[i][j+1])
		}
	}

	fmt.Println(dp[n][m])
}
