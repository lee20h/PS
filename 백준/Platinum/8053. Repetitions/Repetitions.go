package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	n int
)

func fail(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func kmp(text string, pattern string) bool {
	if len(pattern) > len(text) {
		return false
	}
	n, m := len(text), len(pattern)
	lps := fail(pattern)
	i, j := 0, 0

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			return true
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return false
}

func solve(mid int, words []string) bool {
	refWord := words[0]

	if len(refWord) == 1 {
		for _, word := range words {
			if !kmp(word, refWord) {
				return false
			}
		}
		return true
	}

	for i := 0; i <= len(refWord)-mid; i++ {
		substr := refWord[i : i+mid]
		flag := true
		for j := 1; j < len(words); j++ {
			if !kmp(words[j], substr) {
				flag = false
				break
			}
		}

		if flag {
			return true
		}
	}

	return false
}

func main() {
	defer w.Flush()

	fmt.Fscan(r, &n)

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &words[i])
	}

	left, right, result := 1, len(words[0]), 0

	for left <= right {
		mid := (left + right) / 2

		if solve(mid, words) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(w, result)
}
