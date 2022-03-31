package main

// 打表法，更适用于第一次初始化后，多次调用范围查询的时候
func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		x := i
		for x != 0 {
			if x%10 == 0 {
				break
			}
			if i%(x%10) != 0 {
				break
			}
			x /= 10
		}

		if x == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
