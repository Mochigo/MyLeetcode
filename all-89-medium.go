package main

func grayCode1(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
	}

	return ans
}

func grayCode2(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|(1<<(i-1)))
		}
	}
	return ans
}
