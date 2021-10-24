package main

import (
	"fmt"
	"math"
)

func shoppingOffers1(price []int, special [][]int, needs []int) int {
	// 这里复制一个新的数组
	min := math.MaxInt32
	n := len(price)

	dp := make(map[[6]int]int)
	fmt.Println(dp)

	filterSpecials := [][]int{}
	for _, sp := range special {
		sum := 0
		count := 0
		// num表示sp中item[i]的商品数量，price[i]表示商品item[i]的商品价格
		// sum表示礼包总价值
		for i, num := range sp[:n] {
			sum += num * price[i]
			count += num
		}

		if count > 0 && sum > sp[n] {
			filterSpecials = append(filterSpecials, sp)
		}
	}

	var dfs func([]int, int)
	dfs = func(left []int, cur int) {
		// 遍历每个礼包，减除礼包中各个商品的数量
		isDfs := false
		for _, sp := range filterSpecials {
			val := sp[n]
			tmp := sp[:n]
			flag := true
			for i, item := range tmp {
				left[i] -= item
				if flag && left[i] < 0 {
					flag = false
				}
			}
			// 如果没有过度减除，就dfs
			if flag {
				dfs(left, cur+val)
				isDfs = true
			}

			for i, item := range tmp {
				left[i] += item
			}
		}

		// 如果没有发生遍历, 说明这时候要计算单品了
		if !isDfs {
			for i, num := range left {
				cur += num * price[i]
			}
			if cur < min {
				min = cur
			}
		}
	}
	dfs(needs, 0)

	return min
}

func main() {
	price := []int{2, 5}
	special := [][]int{{3, 0, 5}, {1, 2, 10}}
	needs := []int{3, 2}

	fmt.Println(shoppingOffers1(price, special, needs))
}
