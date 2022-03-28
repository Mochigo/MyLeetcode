package main

func hasAlternatingBits1(n int) bool {
	mask := n & 1 // 取最低位的数值
	n >>= 1
	for mask^(n&1) == 1 {
		mask = n & 1 // n & 1 取最低位
		n >>= 1
	}

	return n == 0
}

func hasAlternatingBits2(n int) bool {
	a := n ^ (n >> 1)
	return (a & (a + 1)) == 0
}
