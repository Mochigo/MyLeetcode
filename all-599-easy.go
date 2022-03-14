package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	cnt := map[string]int{}
	for i, l := range list1 {
		cnt[l] = i
	}
	ans := make([]string, 0)
	min := math.MaxInt32
	for i, l := range list2 {
		if val, ok := cnt[l]; ok {
			if val+i < min {
				ans = []string{l}
				min = val + i
			} else if val+i == min {
				ans = append(ans, l)
			}
		}
	}

	return ans
}
