package main

import "fmt"

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	res := make([]bool, len(queries))
	sums := make([]int, n)

	sums[0] = candiesCount[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + candiesCount[i] // 吃完第 i 类糖果前需要吃的糖果数量
	}
	for i, item := range queries {
		ftype, fday, cap := item[0], item[1], item[2]

		currentDay := fday + 1                       // 从第 0 天开始， 因此到第fday天，实际过了 fday+1 天
		maxEat := currentDay * cap                   // 表示第 fday 天 最多能吃的糖果数量
		minEat := currentDay                         // 表示第fday天 最少吃了的糖果数量
		candies := sums[ftype] - candiesCount[ftype] // 表示吃完第i-1类糖果需要吃的糖果数量

		res[i] = maxEat > candies && minEat <= sums[ftype]
	}

	return res
}

func main() {
	candys := []int{46, 5, 47, 48, 43, 34, 15, 26, 11, 25, 41, 47, 15, 25, 16, 50, 32, 42, 32, 21, 36, 34, 50, 45, 46, 15, 46, 38, 50, 12, 3, 26, 26, 16, 23, 1, 4, 48, 47, 32, 47, 16, 33, 23, 38, 2, 19, 50, 6, 19, 29, 3, 27, 12, 6, 22, 33, 28, 7, 10, 12, 8, 13, 24, 21, 38, 43, 26, 35, 18, 34, 3, 14, 48, 50, 34, 38, 4, 50, 26, 5, 35, 11, 2, 35, 9, 11, 31, 36, 20, 21, 37, 18, 34, 34, 10, 21, 8, 5}
	queries := [][]int{
		{85, 54, 42},
	}

	fmt.Println(canEat(candys, queries))

}
