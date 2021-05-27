package main

import (
	"fmt"
	"math/bits"
)

// 这道题可以简化为求两数异或后二进制表示中 1 的个数

// 方法 1 对 异或后的结果 使用Brian Kernighan 算法
// s ^ s-1 可以消去最右侧的 1 从而实现 每次循环都可以消去一个最右侧的1，这时统计循环次数即可知道 s中1的位数
func hammingDistance1(x int, y int) int {
	res := 0

	for s := x ^ y; s > 0; s &= s - 1 {
		res++
	}

	return res
}

// 方法 2 对异或后的结果 与1逐位进行与运算 &，求1的位数，
func hammingDistance2(x int, y int) int {
	res := 0

	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}

	return res
}

// 方法 3 调用编程语言中内置的计算二进制表达中 1 的数量函数。要注意函数的类型需求
func hammingDistance3(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	x, y := 5, 0

	fmt.Println(hammingDistance3(x, y))
}
