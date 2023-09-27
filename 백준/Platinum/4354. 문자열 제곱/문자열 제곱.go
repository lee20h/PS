package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 1e6
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
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

func main() {
	defer w.Flush()
	var str string

	for {
		fmt.Fscan(r, &str)
		if str == "." {
			break
		}

		for i := 0; i < len(str); i++ {
			fail[i] = 0
		}

		failFunc(str)

		if fail[len(str)-1] == 0 || fail[len(str)-1]%(len(str)-fail[len(str)-1]) != 0 {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, len(str)/(len(str)-fail[len(str)-1]))
		}
	}
}
