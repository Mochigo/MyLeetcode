package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	athlete := map[int]int{}
	for i, grade := range score {
		athlete[grade] = i
	}

	sort.Ints(score)
	ans := make([]string, n)
	for i := n - 1; i >= 0; i-- {
		switch i {
		case n - 1:
			ans[athlete[score[i]]] = "Gold Medal"
		case n - 2:
			ans[athlete[score[i]]] = "Silver Medal"
		case n - 3:
			ans[athlete[score[i]]] = "Bronze Medal"
		default:
			ans[athlete[score[i]]] = strconv.Itoa(n - i)
		}
	}

	return ans
}
