package main

import "math/rand"

type Solution struct {
	orginate []int
	cur      []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.cur, this.orginate)
	return this.cur
}

func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.cur), func(i, j int) {
		this.cur[i], this.cur[j] = this.cur[j], this.cur[i]
	})

	return this.cur
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
