package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r         = bufio.NewReader(os.Stdin)
	w         = bufio.NewWriter(os.Stdout)
	n, answer int
	arr       [505][505]int
	dp        [505][505]int
	dy        = [4]int{0, 0, 1, -1}
	dx        = [4]int{1, -1, 0, 0}
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func move(y, x int) int {
	if dp[y][x] != 0 {
		return dp[y][x]
	}
	dp[y][x] = 1

	for i := 0; i < 4; i++ {
		posX := x + dx[i]
		posY := y + dy[i]

		if posX >= 0 && posY >= 0 && posX < n && posY < n {
			if arr[y][x] < arr[posY][posX] {
				dp[y][x] = max(dp[y][x], move(posY, posX)+1)
			}
		}
	}
	return dp[y][x]
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &arr[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			answer = max(answer, move(i, j))
		}
	}

	fmt.Fprintln(w, answer)
}
