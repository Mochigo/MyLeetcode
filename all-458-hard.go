package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// 我们不需要在乎具体怎么喂毒药
	// 只需要确定有每一位在喂食后有几种状态，以及有多少位
	// 例如当只能喂食一次的时候，每一位只有2个状态，10位数共有2 ^ 10种状态，根据每一位的值可以得知具体是哪种组合导致了相应情况的发生
	// 所以其实是求 m ^ pigs >= buckets
	// pigs log m >= log buckets
	// pigs >= log buckets / log m
	// Ceil 返回 大于或等于x的最小整数
	return int(math.Ceil(math.Log10(float64(buckets)) / math.Log10(float64(minutesToTest/minutesToDie+1))))
}
