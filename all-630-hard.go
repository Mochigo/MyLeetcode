package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(a, b int) bool {
		return courses[a][1] <= courses[b][1]
	})

	h := &Heap{}
	times := 0 // 上课已用的时间
	for _, course := range courses {
		if t := course[0]; times+t <= course[1] {
			// 如果这门课能够上完，就直接上
			times += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			// 如果上不完，那么就拿这门课与之前用时最长的课进行比较，如果比之前用时最长的课时间短，就替换掉之前用时最长的课，从而减少上课总时长
			times += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

// 重写覆盖
func (h Heap) Less(a, b int) bool { return h.IntSlice[a] > h.IntSlice[b] }
func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *Heap) Pop() interface{} {
	old := h.IntSlice
	x := old[len(old)-1]
	h.IntSlice = old[:len(old)-1]
	return x
}
