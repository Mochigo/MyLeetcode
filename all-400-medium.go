package main

import "math"

func findNthDigit(n int) int {
	base := 9
	l := 1
	for n-base*l > 0 {
		n -= base * l
		base *= 10
		l++
	}
	st := int(math.Pow10(l - 1))
	digitIndex := (n - 1) % l
	offset := (n - 1) / l
	num := st + offset
	return num / int(math.Pow10(l-digitIndex-1)) % 10 // 整除 10 取最低位
}
