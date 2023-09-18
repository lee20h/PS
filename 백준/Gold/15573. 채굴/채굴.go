package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const (
	MAX = 1000
)

type Queue struct {
	v *list.List
}

type pos struct {
	x int
	y int
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return nil
	}

	return q.v.Remove(front)
}

func (q *Queue) IsEmpty() bool {
	if q.v.Len() == 0 {
		return true
	}
	return false
}

var (
	r               = bufio.NewReader(os.Stdin)
	w               = bufio.NewWriter(os.Stdout)
	n, m, k, answer int
	arr             [MAX][MAX]int
	visited         [MAX][MAX]bool
	dy              = [4]int{0, 0, 1, -1}
	dx              = [4]int{1, -1, 0, 0}
)

func solve(strength int) bool {
	stone := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			visited[i][j] = false
		}
	}

	q := NewQueue()

	for i := 0; i < n; i++ {
		if i == 0 {
			for j := 0; j < m; j++ {
				if arr[0][j] <= strength {
					visited[0][j] = true
					q.Push(pos{
						x: j,
						y: 0,
					})
					stone++
				}
			}
		} else {
			if !visited[i][0] && arr[i][0] <= strength {
				visited[i][0] = true
				q.Push(pos{
					x: 0,
					y: i,
				})
				stone++
			}
			if !visited[i][m-1] && arr[i][m-1] <= strength {
				visited[i][m-1] = true
				q.Push(pos{
					x: m - 1,
					y: i,
				})
				stone++
			}
		}
	}
	for !q.IsEmpty() {
		p := q.Pop().(pos)

		for i := 0; i < 4; i++ {
			px := p.x + dx[i]
			py := p.y + dy[i]

			if px < 0 || py < 0 || px >= m || py >= n {
				continue
			}
			if visited[py][px] || arr[py][px] > strength {
				continue
			}

			visited[py][px] = true
			stone++
			q.Push(pos{
				x: px,
				y: py,
			})
		}
	}

	return k <= stone
}

func main() {

	defer w.Flush()
	fmt.Fscan(r, &n, &m, &k)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(r, &arr[i][j])
		}
	}
	low, high := 1, 1000000
	answer = 1000000

	for low <= high {
		mid := (low + high) / 2
		if solve(mid) {
			high = mid - 1
			answer = mid
		} else {
			low = mid + 1
		}
	}
	fmt.Fprintln(w, answer)
}
