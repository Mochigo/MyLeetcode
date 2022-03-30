package main

import (
	"container/heap"
	"sort"
)

func busiestServers(k int, arrival []int, load []int) []int {
	available := hi{make([]int, k)}
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}

	busies := hp{}
	requests := make([]int, k)
	maxRequest := 0
	ans := []int{}

	for i := range arrival {
		start, end := arrival[i], arrival[i]+load[i]
		for busies.Len() > 0 && busies[0].finish <= start {
			x := heap.Pop(&busies).(pair)
			heap.Push(&available, i+((x.no-i)%k+k)%k) //为什么能得到与 no 同余。。。
		}

		if available.Len() > 0 {
			x := heap.Pop(&available).(int) % k

			requests[x]++
			if requests[x] > maxRequest {
				maxRequest = requests[x]
				ans = []int{x}
			} else if requests[x] == maxRequest {
				ans = append(ans, x)
			}
			heap.Push(&busies, pair{x, end})
		}
	}

	return ans
}

type hi struct{ sort.IntSlice }

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct {
	no     int
	finish int
}

type hp []pair

func (h hp) Len() int      { return len(h) }
func (h hp) Swap(a, b int) { h[a], h[b] = h[b], h[a] }
func (h hp) Less(a, b int) bool {
	x, y := h[a], h[b]
	return x.finish < y.finish || x.finish == y.finish && x.no < y.no
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
