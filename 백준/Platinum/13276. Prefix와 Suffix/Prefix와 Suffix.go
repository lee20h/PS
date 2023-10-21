package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type StringIndex struct {
	str   string
	index int
}

func getSA(s string) []int {
	n := len(s)
	saPairs := make([]StringIndex, n)

	for i := 0; i < n; i++ {
		saPairs[i] = StringIndex{s[i:], i}
	}

	sort.Slice(saPairs, func(i, j int) bool {
		return saPairs[i].str < saPairs[j].str
	})

	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = saPairs[i].index
	}

	return sa
}

func getLCP(s string, sa []int) []int {
	n := len(s)
	ra := make([]int, n)
	lcp := make([]int, n)

	for i := 0; i < n; i++ {
		ra[sa[i]] = i
	}

	k := 0
	for i := 0; i < n-1; i++ {
		if k > 0 {
			k--
		}
		j := sa[ra[i]-1]
		for s[i+k] == s[j+k] {
			k++
		}
		lcp[ra[i]] = k
	}

	return lcp
}

func main() {
	defer w.Flush()
	S, _ := r.ReadString('\n')
	A, _ := r.ReadString('\n')
	B, _ := r.ReadString('\n')

	S = strings.TrimSpace(S)
	A = strings.TrimSpace(A)
	B = strings.TrimSpace(B)

	C := make([]bool, len(S))
	var D []int

	p := -1
	for {
		tempIndex := strings.Index(S[p+1:], A)
		if tempIndex == -1 {
			break
		}
		p = tempIndex + p + 1
		C[p] = true
	}

	p = -1
	for {
		tempIndex := strings.Index(S[p+1:], B)
		if tempIndex == -1 {
			break
		}
		p = tempIndex + p + 1
		D = append(D, p+len(B))
	}

	S += "$"
	sa := getSA(S)
	lcp := getLCP(S, sa)

	minLen := len(A)
	if len(B) > minLen {
		minLen = len(B)
	}

	ans := 0
	for i := 1; i < len(S); i++ {
		if C[sa[i]] {
			for _, d := range D {
				if d >= sa[i]+max(lcp[i]+1, minLen) {
					ans++
				}
			}
		}
	}

	fmt.Fprintln(w, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
