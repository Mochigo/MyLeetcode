package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	if groupSize == 1 {
		return true
	}

	cnt := map[int]int{}
	for _, card := range hand {
		cnt[card]++
	}

	sort.Ints(hand)
	for _, h := range hand {
		for cnt[h] > 0 {
			for i := h; i < h+groupSize; i++ {
				if cnt[i] == 0 {
					return false
				}
				cnt[i]--
			}
		}
	}

	return true
}
