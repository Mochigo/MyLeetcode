package main

func maximumWealth(accounts [][]int) int {
	ans := 0
	for _, a := range accounts {
		sum := 0
		for _, v := range a {
			sum += v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
