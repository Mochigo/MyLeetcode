package main

func distributeCandies(candyType []int) int {
	// 统计有几种类型，然后，比较n/2和类型数哪个大
	n := len(candyType)
	category := make(map[int]bool, n)
	for _, candy := range candyType {
		category[candy] = true
	}

	if len(category) <= n/2 {
		return len(category)
	}
	return n / 2
}
