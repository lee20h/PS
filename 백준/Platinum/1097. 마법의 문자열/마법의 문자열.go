package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 1000 + 1
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	pi   [MAX]int
	n, k int
)

func kmp(a, b string) int {
	cnt := 0
	for i := 0; i < len(pi); i++ {
		pi[i] = 0
	}
	for i, j := 1, 0; i < len(a); i++ {
		for j > 0 && (j >= len(b) || a[i] != b[j]) {
			j = pi[j-1]
		}
		if a[i] == b[j] {
			pi[i] = j + 1
			j++
		}
	}
	for i, j := 0, 0; i < len(a)-1; i++ {
		for j > 0 && (j >= len(b) || a[i] != b[j]) {
			j = pi[j-1]
		}
		if a[i] == b[j] {
			if j == len(b)-1 {
				cnt++
				j = pi[j]
			} else {
				j++
			}
		}
	}
	return cnt
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n)
	var str []string
	var idx []int
	str = make([]string, n)
	idx = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &str[i])
		idx[i] = i
	}
	fmt.Fscan(r, &k)
	cnt := 0

	for {
		var comb, word string
		for i := 0; i < n; i++ {
			word += str[idx[i]]
		}
		comb = word + word

		if kmp(comb, word) == k {
			cnt++
		}
		if !nextPermutation(idx) {
			break
		}
	}
	fmt.Fprintln(w, cnt)

}

func nextPermutation(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		return false
	}

	j := n - 1
	for nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]

	reverse(nums, i+1, n-1)
	return true
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
