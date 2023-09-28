package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 1e6 + 1
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	n    int
	fail [MAX]int
)

func failFunc(str string) {
	for i, j := 1, 0; i < len(str); i++ {
		for j > 0 && str[i] != str[j] {
			j = fail[j-1]
		}
		if str[i] == str[j] {
			j++
			fail[i] = j
		}
	}
}

func kmp(text, word string) int {
	var answer int
	for i, j := 0, 0; i < len(text); i++ {
		for j > 0 && text[i] != word[j] {
			j = fail[j-1]
		}
		if text[i] == word[j] {
			if j == len(word)-1 {
				answer++
				j = fail[j]
			} else {
				j++
			}
		}
	}
	return answer
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n)

	for t := 0; t < n; t++ {
		var word, text string
		fmt.Fscan(r, &word)
		fmt.Fscan(r, &text)
		failFunc(word)
		fmt.Fprintln(w, kmp(text, word))
	}

}
