package main

import (
	"math"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		// mask=665772=10100010100010101100
		// 2 3 5 7 11 13 17 19
		if 1<<bits.OnesCount(uint(i))&665772 > 0 {
			ans++
		}
	}

	return ans
}

func countPrimeSetBits1(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			ans++
		}
	}

	return ans
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := int(math.Sqrt(float64(x))); i >= 2; i-- {
		if x%i == 0 {
			return false
		}
	}

	return true
}
