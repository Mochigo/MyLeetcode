package main

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &floatHeap{}
	n := len(classes)
	for _, class := range classes {
		heap.Push(h, &pair{class[0], class[1]})
	}

	for extraStudents > 0 {
		cur := heap.Pop(h).(*pair)
		heap.Push(h, &pair{cur.pass + 1, cur.all + 1})
		extraStudents--
	}

	var sum float64
	for i := 0; i < n; i++ {
		sum += float64((*h)[i].pass) / float64((*h)[i].all)
	}

	return sum / float64(n)
}

type pair struct {
	pass, all int
}

type floatHeap []*pair

func (h *floatHeap) Len() int {
	return len(*h)
}

// 根据+1后的通过率的增长率进行排序
func (h *floatHeap) Less(a, b int) bool {
	x := float64((*h)[a].pass+1)/float64((*h)[a].all+1) - float64((*h)[a].pass)/float64((*h)[a].all)
	y := float64((*h)[b].pass+1)/float64((*h)[b].all+1) - float64((*h)[b].pass)/float64((*h)[b].all)
	return x > y
}

func (h *floatHeap) Swap(a, b int) {
	(*h)[a], (*h)[b] = (*h)[b], (*h)[a]
}
func (h *floatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *floatHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}
