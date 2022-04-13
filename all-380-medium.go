package main

import "math/rand"

type RandomizedSet struct {
	nums []int
	last map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.last[val]; ok {
		return false
	}
	r.nums = append(r.nums, val)
	r.last[val] = len(r.nums) - 1
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.last[val]; !ok {
		return false
	}
	x := r.last[val]
	end := len(r.nums) - 1
	v := r.nums[end]
	r.nums[x], r.nums[end] = r.nums[end], r.nums[x]
	r.last[v] = x
	r.nums = r.nums[:end]
	delete(r.last, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
