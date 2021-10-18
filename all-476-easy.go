package main

import "fmt"

// 核心思想是，通过将num与相同位数的1相异或，得到结果
// 这里有两种得到相同位数的1，一个是通过循环，找到最小的比num大的2的倍数，通过减一获得
// 一种是通过 x^(x-1)消去最低位的1，直到只剩最高位的1，再将整体 << 1 在 -1 即可获得

func findComplement(num int) int {
	highBit := 0
	for num >= (1 << highBit) {
		highBit++
	}

	return num ^ (1<<highBit - 1)
}

func findComplement2(num int) int {
	mask := num
	for mask&(mask-1) != 0 {
		mask &= mask - 1
	}

	return num ^ (mask<<1 - 1)
}

func main() {
	fmt.Println(findComplement(5))
}
