package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 50
)

var (
	r            = bufio.NewReader(os.Stdin)
	w            = bufio.NewWriter(os.Stdout)
	n, m, answer int
	arr          [MAX][MAX]int
	dp           [MAX][MAX]int
	visited      [MAX][MAX]bool
	dy           = [4]int{0, 0, 1, -1}
	dx           = [4]int{1, -1, 0, 0}
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func move(y, x int) int {
	if x < 0 || y < 0 || x >= m || y >= n || arr[y][x] == 0 {
		return 0
	}
	if visited[y][x] {
		fmt.Println("-1")
		os.Exit(0)
	}

	if dp[y][x] != -1 {
		return dp[y][x]
	}

	visited[y][x] = true
	dp[y][x] = 0

	for i := 0; i < 4; i++ {
		posY := y + (arr[y][x] * dy[i])
		posX := x + (arr[y][x] * dx[i])
		dp[y][x] = max(dp[y][x], move(posY, posX)+1)
	}
	visited[y][x] = false
	return dp[y][x]
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(r, &str)
		for j, v := range []rune(str) {
			if v == 'H' {
				arr[i][j] = 0
			} else {
				arr[i][j] = int(v) - '0'
			}
		}
	}
	answer = move(0, 0)

	fmt.Fprintln(w, answer)
}
