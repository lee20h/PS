package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func check(n int, L []int) bool {
	if len(L)%n != 0 {
		return false
	}
	for i := 0; i < len(L)/n-1; i++ {
		for j := 0; j < n; j++ {
			if L[j] != L[i*n+n+j] {
				return false
			}
		}
	}
	return true
}

func solve() {
	var t, m, n int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		fmt.Fscan(r, &m, &n)
		L := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &L[i])
		}
		L = append(L, L[0]+m)
		diff := make([]int, n)
		for i := 0; i < n; i++ {
			diff[i] = L[i+1] - L[i]
		}
		for i := 1; i <= n; i++ {
			if check(i, diff) {
				fmt.Fprintln(w, m/(n/i))
				break
			}
		}
	}
}

func main() {
	defer w.Flush()
	solve()
}
