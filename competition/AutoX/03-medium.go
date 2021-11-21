package main

import "sort"

type RangeFreqQuery struct {
	appear map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	appear := map[int][]int{}
	// 天然有序
	for i, num := range arr {
		appear[num] = append(appear[num], i)
	}
	return RangeFreqQuery{appear}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if _, has := this.appear[value]; !has {
		return 0
	}

	nums := this.appear[value]
	l := sort.SearchInts(nums, left)
	r := sort.SearchInts(nums, right+1)

	return r - l
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
