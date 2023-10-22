package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const SIZE = 100

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
)

type SuffixArray struct {
	src    string
	len    int
	sa     [SIZE]int
	rk     [SIZE]int
	rkTmp  [SIZE]int
	offset int
}

func (sf *SuffixArray) Cmp(a, b int) bool {
	if sf.rk[a] != sf.rk[b] {
		return sf.rk[a] < sf.rk[b]
	}
	if a+sf.offset >= sf.len || b+sf.offset >= sf.len {
		return a+sf.offset > b+sf.offset // 초과하는 Suffix 가 앞쪽에 오도록
	}
	return sf.rk[a+sf.offset] < sf.rk[b+sf.offset]
}

func (sf *SuffixArray) Init() {
	sf.len = len(sf.src)
	for i := 0; i < sf.len; i++ {
		sf.sa[i] = i
		sf.rk[i] = int(sf.src[i])
	}

	for sf.offset = 1; sf.offset < sf.len; sf.offset <<= 1 {
		sort.Slice(sf.sa[:sf.len], func(i, j int) bool {
			return sf.Cmp(sf.sa[i], sf.sa[j])
		})

		sf.rkTmp[sf.sa[0]] = 0
		for i := 1; i < sf.len; i++ {
			if sf.Cmp(sf.sa[i-1], sf.sa[i]) {
				sf.rkTmp[sf.sa[i]] = sf.rkTmp[sf.sa[i-1]] + 1
			} else {
				sf.rkTmp[sf.sa[i]] = sf.rkTmp[sf.sa[i-1]]
			}
		}
		copy(sf.rk[:], sf.rkTmp[:])
	}
}

func Check(sf *SuffixArray) int {
	if sf.src[sf.sa[0]] > 'a' {
		return 1
	}
	for i := 1; i < sf.len; i++ {
		prev := sf.src[sf.sa[i-1]]
		cur := sf.src[sf.sa[i]]

		if prev < cur-1 {
			return 1
		} else if prev == cur-1 {
			if sf.sa[i]+1 == sf.len {
				continue
			}
			if sf.sa[i-1]+1 == sf.len || sf.rk[sf.sa[i-1]+1] <= sf.rk[sf.sa[i]+1] {
				return 1
			}
		}
	}
	return 0
}

func main() {
	defer w.Flush()
	var sf SuffixArray
	fmt.Fscanln(r, &sf.src)
	sf.src = strings.TrimSpace(sf.src) // 입력 문자열의 공백 제거
	sf.Init()
	fmt.Fprintln(w, Check(&sf))
}
