package main

func getAverages(nums []int, k int) []int {
	n := len(nums)
	c := 2*k + 1

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	if n < c {
		return ans
	}

	sum := 0
	for _, num := range nums[:c] {
		sum += num
	}
	ans[k] = sum / c
	index := k + 1
	for i, r := 0, c; r < n; i, r = i+1, r+1 {
		sum -= nums[i]
		sum += nums[r]
		ans[index] = sum / c
		index++
	}

	return ans
}
