package main

import "fmt"

func countOrders(n int) int {
	const mod = 1e9 + 7
	ans := 1

	for i := 2; i <= n; i++ {
		// 对于前一个长度为n的连续序列，则有n+1个空位
		// 当Pn+1 Dn+1相连在一起时，则Cn+1 1
		// 当Pn+1 Dn+1不相连在一起时， 则Cn+1 2
		ans = ans * (2*(i-1) + 1 + (2*(i-1)+1)*(2*(i-1)/2)) % mod
	}

	return ans
}

func main() {
	fmt.Println(countOrders(2))
}
