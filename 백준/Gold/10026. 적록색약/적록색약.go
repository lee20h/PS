package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 100
)

var (
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
	visited [MAX][MAX]bool
	color   [MAX]string
	n       int
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

func dfs(y, x int) {
	visited[y][x] = true
	for i := 0; i < 4; i++ {
		px := x + dx[i]
		py := y + dy[i]
		if px < 0 || py < 0 || px >= n || py >= n {
			continue
		}
		if !visited[py][px] && color[py][px] == color[y][x] {
			dfs(py, px)
		}
	}
}

func RGdfs(y, x int) {
	visited[y][x] = true
	for i := 0; i < 4; i++ {
		px := x + dx[i]
		py := y + dy[i]
		if px < 0 || py < 0 || px >= n || py >= n {
			continue
		}

		if !visited[py][px] {
			if (color[y][x] == 'R' || color[y][x] == 'G') && (color[py][px] == 'R' || color[py][px] == 'G') {
				RGdfs(py, px)
			} else if color[py][px] == color[y][x] {
				RGdfs(py, px)
			}
		}
	}
}

func main() {
	defer w.Flush()

	cnt, RGcnt := 0, 0

	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &color[i])
	}
	for i := 0; i < n; i++ {
		for j, _ := range color[i] {
			if !visited[i][j] {
				dfs(i, j)
				cnt++
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}

	for i := 0; i < n; i++ {
		for j, _ := range color[i] {
			if !visited[i][j] {
				RGdfs(i, j)
				RGcnt++
			}
		}
	}

	fmt.Fprintln(w, cnt, RGcnt)
}
