package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	ans, prev := 10, 9 // 这里是一位数的结果
	for i := 2; i <= n; i++ {
		prev *= 10 - i + 1
		ans += prev
	}

	return ans
}
