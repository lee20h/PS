package main

import (
	"bufio"
	"fmt"
	"os"
)

type Seq []int

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func calculatePi(pi []int, str Seq) {
	pi[0] = -1
	j := -1
	for i := 1; i < len(str); i++ {
		for j >= 0 && str[i] != str[j+1] {
			j = pi[j]
		}
		if str[i] == str[j+1] {
			j++
			pi[i] = j
		} else {
			pi[i] = -1
		}
	}
}

func kmp(text, pattern Seq) []int {
	pi := make([]int, len(pattern))
	ans := make([]int, 0)
	if len(pattern) == 0 {
		return ans
	}
	calculatePi(pi, pattern)
	j := -1
	for i := 0; i < len(text); i++ {
		for j >= 0 && text[i] != pattern[j+1] {
			j = pi[j]
		}
		if text[i] == pattern[j+1] {
			j++
			if j+1 == len(pattern) {
				ans = append(ans, i-j)
				j = pi[j]
			}
		}
	}
	return ans
}

func main() {
	defer writer.Flush()
	var N int
	fmt.Fscan(reader, &N)

	v := make([]Seq, N)
	for i := 0; i < N; i++ {
		var c int
		fmt.Fscan(reader, &c)
		u := make(Seq, c)
		for j := range u {
			fmt.Fscan(reader, &u[j])
		}

		var r Seq
		for j := 0; j < c-1; j++ {
			r = append(r, u[j+1]-u[j])
		}
		v[i] = r
	}

	var L int
	fmt.Fscan(reader, &L)
	a := make(Seq, L)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}

	var b Seq
	for i := 0; i < L-1; i++ {
		b = append(b, a[i+1]-a[i])
	}

	ans := make([]int, 0)
	for i := 0; i < N; i++ {
		t := kmp(v[i], b)
		if len(t) > 0 {
			ans = append(ans, i+1)
		}
	}

	if len(ans) == 0 {
		fmt.Fprintln(writer, "-1")
	} else {
		for _, x := range ans {
			fmt.Fprintf(writer, "%d ", x)
		}
		fmt.Fprintln(writer)
	}
}
