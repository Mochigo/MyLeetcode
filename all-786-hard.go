package main

import "container/heap"

type fh []float64

func (h fh) Len() int           { return len(h) }
func (h fh) Less(a, b int) bool { return h[a] <= h[b] }
func (h fh) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h *fh) Push(x interface{}) {
	*h = append(*h, x.(float64))
}
func (h *fh) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MykthSmallestPrimeFraction(arr []int, k int) []int {
	h := &fh{}
	heap.Init(h)
	n := len(arr)
	m := map[float64][]int{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			val := float64(arr[i]) / float64(arr[j])
			heap.Push(h, val)
			m[val] = []int{arr[i], arr[j]}
		}
	}

	for cnt := 1; cnt <= k-1; cnt++ {
		heap.Pop(h)
	}

	return m[heap.Pop(h).(float64)]
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp, n-1)
	for j := 1; j < n; j++ {
		h[j-1] = frac{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(frac)
		if f.i+1 < f.j {
			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type hp []frac

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
