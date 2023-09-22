package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	MAX = 101
)

var (
	r         = bufio.NewReader(os.Stdin)
	w         = bufio.NewWriter(os.Stdout)
	n, answer int
	parent    [MAX]int
)

type edge struct {
	cost         int
	node1, node2 int
}

type pq []edge

func (h pq) Len() int {
	return len(h)
}

func (h pq) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pq) Push(element interface{}) {
	*h = append(*h, element.(edge))
}

func (h *pq) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

func find(node int) int {
	if node == parent[node] {
		return node
	}
	parent[node] = find(parent[node])
	return parent[node]
}

func union(node1, node2 int) {
	pn1 := find(node1)
	pn2 := find(node2)

	if pn1 < pn2 {
		parent[pn2] = pn1
	} else {
		parent[pn1] = pn2
	}
}

func isConnected() bool {
	for i := 2; i <= n; i++ {
		if find(1) != find(i) {
			return false
		}
	}
	return true
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n)
	sum := 0
	q := &pq{}
	heap.Init(q)
	for i := 1; i <= n; i++ {
		parent[i] = i
		var str string
		fmt.Fscan(r, &str)
		for j, c := range str {
			j = j + 1
			if 'a' <= c && c <= 'z' {
				cost := int(c) - 'a' + 1
				sum += cost
				heap.Push(q, edge{
					cost:  cost,
					node1: i,
					node2: j,
				})
			} else if 'A' <= c && c <= 'Z' {
				cost := int(c) - 'A' + 27
				sum += cost
				heap.Push(q, edge{
					cost:  cost,
					node1: i,
					node2: j,
				})
			}
		}
	}

	e := 0
	for q.Len() > 0 {
		top := heap.Pop(q).(edge)

		if find(top.node1) == find(top.node2) {
			continue
		}
		union(top.node1, top.node2)
		e += top.cost
	}

	if !isConnected() {
		fmt.Fprintln(w, "-1")
	} else {
		fmt.Fprintln(w, sum-e)
	}
}
