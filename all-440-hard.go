package main

func findKthNumber(n int, k int) int {
	cur := 1    // 标记当前的前缀prefix的字典序
	prefix := 1 // 当前前缀
	for cur < k {
		cnt := getCount(n, prefix)
		if cur+cnt <= k {
			// 说明当前前缀下不存在cur == k，必须要迁移到下一个前缀
			prefix++
			cur += cnt
		} else {
			prefix *= 10
			cur++
		}
	}

	return prefix
}

// 返回当前前缀下的不超过n的所有数字的个数
func getCount(n int, prefix int) int {
	cnt := 0
	next := prefix + 1
	for prefix <= n {
		cnt += min(n+1, next) - prefix
		prefix *= 10
		next *= 10
	}

	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
