package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	MAX = 10001
)

var (
	r                          = bufio.NewReader(os.Stdin)
	w                          = bufio.NewWriter(os.Stdout)
	n, m, u, v, cost, ans, cnt int
	parent                     [MAX]int
	q                          *pq
)

type pq []edge

type edge struct {
	cost         int
	node1, node2 int
}

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

func union(node1, node2, cost int) {
	pn1 := find(node1)
	pn2 := find(node2)

	if pn1 == pn2 {
		return
	}
	parent[pn2] = pn1
	ans += cost
	cnt++
}

func initialize() {
	q = &pq{}
	heap.Init(q)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
}

func input(q *pq) {
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &u, &v, &cost)
		heap.Push(q, edge{
			cost:  cost,
			node1: u,
			node2: v,
		})
	}
}

func solve(q *pq) {
	for q.Len() > 0 {
		e := heap.Pop(q).(edge)
		u, v, w := e.node1, e.node2, e.cost

		union(u, v, w)
		if cnt == n-1 {
			break
		}
	}
}

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n, &m)

	initialize()

	input(q)
	solve(q)
	fmt.Fprintln(w, ans)
}
