package main

import "sort"

type TopVotedCandidate struct {
	times  []int
	leader []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	// 顺序遍历， time作为key
	vote := map[int]int{-1: -1}         // 因为是顺序遍历，所以可以做到记录不同时刻的选票数
	leader := make([]int, len(persons)) // 记录每个时刻的领先者
	max := -1                           // 标记此时票数最多的人，每次和当前被投票的人进行比较
	for i := range times {
		person := persons[i]
		vote[person]++
		if vote[max] <= vote[person] {
			max = person
		}
		leader[i] = max
	}
	return TopVotedCandidate{times, leader}
}

func (this *TopVotedCandidate) Q(t int) int {
	q := sort.SearchInts(this.times, t+1)
	return this.leader[q-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
