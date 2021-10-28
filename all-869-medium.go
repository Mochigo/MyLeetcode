package main

import (
	"sort"
	"strconv"
)

func reorderedPowerOf21(n int) bool {
	if n == 1 {
		return true
	}
	// 两大问题，一个是获得重组排序的所有可能，一个是对n判断是否为2的指数
	// 问题一的话，有一种方法是全排列
	digits := []int{}
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	sort.Ints(digits) // 方便去重

	length := len(digits)
	path := []int{}
	used := map[int]bool{}
	cur := 0

	// 剪枝 尾数不是2 4 6 8不可能为 2 的倍数
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == length {
			return isPowerOf2(cur)
		}

		for i, num := range digits {
			if (cur == 0 && num == 0) || used[i] || (i >= 1 && digits[i] == digits[i-1] && !used[i-1]) {
				continue
			}
			path = append(path, num)
			used[i] = true
			cur = cur*10 + num
			if dfs(index + 1) {
				return true
			}
			cur /= 10
			used[i] = false
			path = path[:len(path)-1]
		}

		return false
	}

	return dfs(0)
}

func reorderedPowerOf2(n int) bool {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	m := len(nums)
	vis := make([]bool, m)
	var backtrack func(int, int) bool
	backtrack = func(idx, num int) bool {
		if idx == m {
			return isPowerOfTwo(num)
		}
		for i, ch := range nums {
			// 不能有前导零
			if num == 0 && ch == '0' || vis[i] || i > 0 && !vis[i-1] && ch == nums[i-1] {
				continue
			}
			vis[i] = true
			if backtrack(idx+1, num*10+int(ch-'0')) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return backtrack(0, 0)
}

func reorderedPowerOf23(n int) bool {
	powerOfTwoDigitsSet := map[[10]int]bool{}
	for i := 1; i <= 1e9; i <<= 1 {
		powerOfTwoDigitsSet[countDigits(i)] = true
	}

	return powerOfTwoDigitsSet[countDigits(n)]
}

func countDigits(n int) [10]int {
	var cnt [10]int
	for n != 0 {
		cnt[n%10]++
		n /= 10
	}

	return cnt
}

func isPowerOf2(sum int) bool {
	return sum > 0 && (sum&(sum-1) == 0)
}

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}
