package main

const mod = 1337

func pow(x int, n int) int {
	if n == 0 {
		return 1
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y % mod
	} else {
		return x * y * y % mod
	}
}

func superPow(a int, b []int) int {
	n := len(b)
	ans := 1
	for i := n - 1; i >= 0; i-- {
		ans = ans * pow(a, b[i]) % mod
		// 99 ^ 2345 = 99 ^ 2340 * 99 ^ 5 = (99^10)^234 * (99^5)
		// = ((99^10)^10)^23 * (99^10)^4 * (99^5)
		a = pow(a, 10)
	}

	return ans
}
